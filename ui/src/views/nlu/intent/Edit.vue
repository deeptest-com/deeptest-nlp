<template>
  <div v-show="model.id">
    <div class="header">
      <span class="label">{{ $t('form.code') }}</span>
      <span v-if="!isEditCode" class="code">
        <span>{{ model.code }}</span>&nbsp;&nbsp;
        <a @click="editCode()"><a-icon class="edit-icon" type="edit" /></a>
      </span>
      <span v-if="isEditCode">
        <a-form layout="inline">
          <a-form-item>
            <a-input v-model="model.code" type="text"></a-input>
          </a-form-item>
          <a-form-item>
            <a @click="saveCode()"><a-icon type="check" /></a>
          </a-form-item>
        </a-form>
      </span>

      <span class="label">{{ $t('form.name') }}</span>
      <span v-if="!isEditTitle" class="title">
        <span>{{ model.name }}</span>&nbsp;&nbsp;
        <a @click="editTitle()"><a-icon class="edit-icon" type="edit" /></a>
      </span>
      <span v-if="isEditTitle">
        <a-form layout="inline">
          <a-form-item>
            <a-input v-model="model.name" type="text"></a-input>
          </a-form-item>
          <a-form-item>
            <a @click="saveTitle()"><a-icon type="check" /></a>
          </a-form-item>
        </a-form>
      </span>

      <div class="buttons"></div>
    </div>

    <a-tabs
      :activeKey="tabKey"
      @change="tabClick"
    >
      <a-tab-pane key="maintainSent" :tab="$t('form.maintain.nlu.sent')">
        <div class="edit">
          <div class="edit-links">
            <a-tag class="tag synonym">{{ $t('form.synonym') }}</a-tag>
            <a-tag class="tag lookup">{{ $t('form.lookup') }}</a-tag>
            <a-tag class="tag regex">{{ $t('form.regex') }} </a-tag>
            <a-tag class="tag _slot_">{{ $t('form.slot') }} </a-tag>
          </div>
          <div class="edit-inputs">
            <div class="left">
              <div
                id="editor"
                contenteditable="true"
                @mouseup="textSelected"
                class="editor"
                ref="editor"
                oncontextmenu="return false;"
                spellcheck="false">
              </div>
              <div class="tips">
                <a-icon type="info-circle" class="icon"/>
                <span>{{ $t('form.select.to.mark') }}</span>
              </div>
            </div>
            <div class="right">
              <a-button @click="saveSent()">{{ $t('form.save') }}</a-button>
              <a-button @click="resetSent()">{{ $t('form.reset') }}</a-button>
            </div>
          </div>
        </div>

        <div class="sent-list">
          <div class="sent-title">
            {{ $t('form.list') }}
          </div>
          <div class="sent-items" :style="styl">
            <draggable v-model="sents" group="sents" @change="resort" @start="drag" @end="drop">
              <div
                v-for="(item, index) in sents"
                @mouseover="mouseover(index)"
                :class="{'mouse-over': index == mouseLocation}"
                :key="item.id"
                ref="sentList"
                class="sent-item">
                <div @click="editSent(item)" class="no link">{{ index + 1 }}</div>
                <div @click="editSent(item)" class="left link" :class="{'disabled':item.disabled}">
                  <span v-html="item.html"></span>
                </div>
                <div class="right">
                  <a-icon @click="editSent(item)" type="edit" class="icon"/> &nbsp;
                  <a-icon v-if="!item.disabled" @click="disableSent(item)" type="minus" class="icon"/>
                  <a-icon v-if="item.disabled" @click="disableSent(item)" type="plus" class="icon"/>

                  <a-icon @click="deleteSent(item)" type="delete" class="icon"/>
                  <a-icon @click.stop="clickMove" type="drag" class="icon" style="cursor: move;" />
                </div>
              </div>
            </draggable>
          </div>
        </div>
      </a-tab-pane>

      <a-tab-pane key="maintainRule" :tab="$t('form.maintain.nlu.rule')">
        <div class="edit">
          <div class="edit-links">
            <a-tag @click="useDict('synonym')" class="tag btn synonym">{{ $t('form.synonym') }}</a-tag>
            <a-tag @click="useDict('lookup')" class="tag btn lookup">{{ $t('form.lookup') }}</a-tag>
            <a-tag @click="useDict('regex')" class="tag btn regex">{{ $t('form.regex') }} </a-tag>
            <a-tag @click="useDict('_slot_')" class="tag btn _slot_">{{ $t('form.slot') }} </a-tag>
          </div>
          <div class="edit-inputs">
            <div class="left">
              <a-input
                v-model="rule.text"
                @click="ruleInputClick"
                ref="ruleText"
                id="ruleText"
                type="text"
                style="height: 40px;"></a-input>
            </div>
            <div class="right">
              <a-button @click="saveRule()">{{ $t('form.save') }}</a-button>
              <a-button @click="resetRule()">{{ $t('form.reset') }}</a-button>
            </div>
          </div>
        </div>
        <div class="sent-list">
          <div class="sent-title">
            {{ $t('form.list') }}
          </div>
          <div class="sent-items">
            <div v-for="item in ruleList" :key="item.id" ref="ruleList" class="sent-item">
              <div class="left" :class="{'disabled':item.disabled}">{{ item.text }}</div>
              <div class="right">
                <a-icon @click="editRule(item)" type="edit" class="icon"/> &nbsp;
                <a-icon v-if="!item.disabled" @click="disableRule(item)" type="minus" class="icon"/>
                <a-icon v-if="item.disabled" @click="disableRule(item)" type="plus" class="icon"/>
                &nbsp;
                <a-icon @click="deleteRule(item)" type="delete" class="icon"/>
              </div>
            </div>
          </div>
        </div>
      </a-tab-pane>
    </a-tabs>

    <a-modal
      :title="$t('form.mark')"
      dialogClass="mark-dialog"
      :width="650"
      :visible="slotEditVisible"
      @cancel="() => cancelSlot()"
      @ok="() => saveSlot()"
    >
      <a-form-model ref="form" :model="slot" :rules="rules" :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-model-item prop="slotType" :label="$t('form.slot.type')">
          <a-radio-group v-model="slot.slotType" @change="slotTypeChanged">
            <a-radio value="synonym">{{ $t('form.synonym') }}</a-radio>
            <a-radio value="lookup">{{ $t('form.lookup') }}</a-radio>
            <a-radio value="regex">{{ $t('form.regex') }}</a-radio>
            <a-radio value="_slot_">{{ $t('form.slot') }}</a-radio>
          </a-radio-group>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="slot.slotType === '_slot_'" :label="$t('form.slot')">
          <a-input v-model="slot.value" />
          <a-select v-model="slot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.name">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="slot.slotType === 'synonym'" :label="$t('form.synonym')">
          <a-select v-model="slot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="slot.slotType === 'lookup'" :label="$t('form.lookup')">
          <a-select v-model="slot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="slot.slotType === 'regex'" :label="$t('form.regex')">
          <a-select v-model="slot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

      </a-form-model>
    </a-modal>

    <a-modal
      :title="$t('form.pls.select')"
      dialogClass="mark-dialog"
      :width="650"
      :visible="ruleSectionEditVisible"
      @cancel="() => cancelRuleSection()"
      @ok="() => addRuleSection()"
    >
      <a-form-model ref="form" :model="ruleSection" :rules="rules" :label-col="labelCol" :wrapper-col="wrapperCol">
        <template v-if="ruleSection.type === '_slot_'">
          <a-form-model-item prop="value" :label="$t('form.name')">
            <a-input v-model="ruleSection.value" />
          </a-form-model-item>
          <a-form-model-item prop="value" :label="$t('form.placeholder')">
            <a-input v-model="ruleSection.placeholder" />
          </a-form-model-item>
        </template>

        <a-form-model-item prop="value" v-if="ruleSection.type !== '_slot_'" :label="$t('form.name')">
          <a-select v-model="ruleSection.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.name">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>

import draggable from 'vuedraggable'
import { convertSelectedToSlots, genSent, genSentSlots } from '@/utils/markUtil'
import {
  getIntent, updateIntent, loadDicts, getSent, saveSent, removeSent, disableSent,
  getRule, saveRule, removeRule, disableRule, resortSent
} from '@/api/manage'

export default {
  name: 'IntentEdit',
  components: {
    draggable
  },
  props: {
    modelId: {
      type: Number,
      default: () => 0
    },
    projectId: {
      type: Number,
      default: () => 0
    }
  },
  mounted () {
    console.log('mounted')
  },
  destroyed () {
    console.log('destroyed')
  },
  data () {
    const styl = 'overflow-y: auto; height: ' + (document.documentElement.clientHeight - 330) + 'px;'
    return {
      model: {},
      sents: [],
      sent: { content: '' },
      slotEditVisible: false,

      ruleList: [],
      rule: { text: '' },
      ruleSection: {},
      ruleInputIndex: -1,
      ruleSectionEditVisible: false,

      labelCol: { span: 5 },
      wrapperCol: { span: 15 },

      allSlots: [],
      selectedSlot: {},
      selectedIndex: -1,
      slot: {},
      isEditCode: false,
      isEditTitle: false,
      tabKey: 'maintainSent',
      mouseLocation: -1,
      isDrag: false,

      dicts: [],
      rules: {
        slotType: [{ required: true, message: this.$t('valid.slot.type'), trigger: 'change' }],
        value: [{ required: true, message: this.$t('valid.select.dict'), trigger: 'change' }]
      },
      styl: styl
    }
  },
  watch: {
    modelId: function () {
      if (this.modelId <= 0) return
      console.log('watch modelId', this.modelId)
      this.getModel()
    }
  },
  methods: {
    getModel () {
      console.log('getModel')
      getIntent(this.modelId).then(json => {
        console.log(json)
        this.sent = {}
        this.$refs.editor.innerHTML = ''

        if (json.code === 200) {
          this.model = json.data
          this.sents = this.model.sents
          this.ruleList = this.model.rules
        }
      })
    },

    editSent (item) {
      console.log('editSent')
      getSent(item.id).then(json => {
        this.sent = json.data
        this.$refs.editor.innerHTML = this.sent.html
      })
    },
    saveSent () {
      const text = this.$refs.editor.innerText.trim()
      console.log('saveSent', text)
      if (text === '') return

      this.sent.intentId = this.model.id
      this.sent.html = this.$refs.editor.innerHTML
      this.sent.text = text.replace(/\s+/g, '')
      this.sent.slots = genSentSlots(document.getElementById('editor'))

      saveSent(this.sent).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
    },
    deleteSent (item) {
      console.log('deleteSent')

      removeSent(item).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
    },
    disableSent (item) {
      console.log('disableSent')

      disableSent(item).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
    },
    clickMove (e) {
      console.log('clickMove')
    },
    resetSent () {
      console.log('resetSent')
      this.rule = { text: '' }
      this.$refs.editor.innerHTML = ''
    },

    editRule (item) {
      console.log('editRule')
      getRule(item.id).then(json => {
        this.rule = json.data
      })
    },
    saveRule () {
      this.rule.intentId = this.model.id
      this.rule.projectId = this.projectId
      console.log('saveRule', this.rule)

      saveRule(this.rule).then(json => {
        console.log(json)
        this.ruleList = json.data
        this.rule = { text: '' }
      })
    },
    deleteRule (item) {
      console.log('toDeleteRule')

      removeRule(item).then(json => {
        this.ruleList = json.data
        this.rule = { text: '' }
      })
    },
    disableRule (item) {
      console.log('toDisableRule')

      disableRule(item).then(json => {
        this.ruleList = json.data
        this.rule = { text: '' }
      })
    },
    resetRule () {
      console.log('resetRule')
      this.rule = { text: '' }
    },

    useDict (dictType) {
      console.log('useDict', dictType)
      this.ruleSection.type = dictType
      this.ruleSectionEditVisible = true

      loadDicts(dictType).then(json => {
        console.log(json)
        this.dicts = json.data
      })
    },
    ruleInputClick () {
      this.ruleInputIndex = this.$refs.ruleText.$el.selectionStart
      console.log('this.ruleText', this.ruleInputIndex)
    },
    addRuleSection () {
      console.log('addRuleSection')

      let val = this.ruleSection.value + ':' + this.ruleSection.placeholder
      if (this.ruleSection.type === '_slot_') {
        val = '(' + val + ')'
      } else {
        val = '{' + val + '}'
      }
      this.rule.text = this.rule.text.substring(0, this.ruleInputIndex) + val +
        this.rule.text.substring(this.ruleInputIndex, this.rule.text.length)

      this.ruleSectionEditVisible = false
      this.ruleSection = {}
      this.ruleInputIndex = this.rule.text.length
    },
    cancelRuleSection () {
      console.log('cancelRuleSection')
      this.ruleSectionEditVisible = false
    },

    editTitle () {
      this.isEditTitle = true
    },
    saveTitle () {
      const data = { id: this.model.id, name: this.model.name }
      updateIntent(data).then(json => {
        this.isEditTitle = false

        this.$global.EventBus.$emit(this.$global.intentUpdateEvent, { 'id': this.model.id })
      })
    },
    editCode () {
      this.isEditCode = true
    },
    saveCode () {
      const data = { id: this.model.id, code: this.model.code }
      updateIntent(data).then(json => {
        this.isEditCode = false

        this.$global.EventBus.$emit(this.$global.intentUpdateEvent, { 'id': this.model.id })
      })
    },

    tabClick (key) {
      this.tabKey = key
    },
    back () {
      this.$router.push('/nlu/task/list')
    },

    textSelected (event) {
      console.log('==== textSelected')
      const mp = convertSelectedToSlots(event.target, document.getElementById('editor'))

      if (!mp.allSlots) return

      mp.allSlots.forEach((item, index) => {
        console.log('--' + index + '--', item)
      })

      this.allSlots = mp.allSlots
      this.selectedIndex = mp.selectedIndex
      this.selectedSlot = this.allSlots[this.selectedIndex]
      console.log('--curr--', this.selectedSlot)

      const dataType = this.selectedSlot.getAttribute('data-type')
      let dataValue = this.selectedSlot.getAttribute('data-value')
      if (dataType !== '_slot_') {
        dataValue = dataValue ? parseInt(dataValue) : ''
      }
      this.slot = { slotType: dataType === 'null' ? '' : dataType, value: dataValue }
      this.slotEditVisible = true

      this.slotTypeChanged()
    },
    slotTypeChanged () {
      loadDicts(this.slot.slotType).then(json => {
        console.log(json)
        this.dicts = json.data
      })
    },
    saveSlot () {
      console.log('saveSlot', this.slot)

      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate', valid)
          return false
        }

        const sentContent = genSent(this.allSlots, this.selectedIndex, this.slot)
        console.log('saveSlot', sentContent)
        this.$refs.editor.innerHTML = sentContent
        this.sent.html = sentContent

        this.slotEditVisible = false
        window.getSelection().removeAllRanges()
      })
    },
    cancelSlot () {
      console.log('cancelSlot')
      this.slotEditVisible = false
      window.getSelection().removeAllRanges()
    },
    mouseover (index) {
      console.log('mouseover')
      this.mouseLocation = index
    },
    resort (evt) {
      console.log('resort', evt.moved.newIndex, evt.moved.oldIndex, evt.moved.element)

      const src = evt.moved.element
      const target = this.sents[evt.moved.newIndex + 1]

      resortSent(src.id, target.id, this.model.id).then(json => {
        console.log('resort', json)
      })

      return true
    },
    drag (index) {
      console.log('drag')
      this.isDrag = true
    },
    drop (index) {
      console.log('drop')
      this.isDrag = false
    }
  }
}
</script>

<style lang="less" scoped>
.header {
  display: flex;
  border-bottom: 1px solid #e9f2fb;
  .label {
    display: inline-block;
    margin-right: 6px;
    font-weight: bolder;
    font-size: 20px;
    line-height: 20px;
  }
  .code {
    width: 160px;
    font-size: 20px;
    line-height: 20px;
  }
  .title {
    flex: 1;
    font-size: 20px;
    line-height: 20px;
  }
  .buttons {
    width: 160px;
  }
}

.edit {
  .edit-title {
    font-weight: bolder;
    font-size: 18px;
  }
  .edit-links {
    text-line: 20px;
    margin: 6px 0;
    .tag {
      margin: 4px 8px 4px 0px;
      line-height: 26px;
      &.btn {
        cursor: pointer;
      }
    }
  }
  .edit-inputs {
    display: flex;
    .left {
      flex: 1;
      .editor {
        padding: 0px 6px;
        height: 40px;
        line-height: 38px;
        font-size: 22px;
        border: 1px solid #d9d9d9;
        outline: none;
      }
      .tips {
        padding: 2px 6px;
        color: rgba(0, 0, 0, 0.45);
        .icon {
          display: inline-block;
          padding-right: 5px;
        }
      }
    }
    .right {
      width: 160px;
      button {
        height: 40px;
      }
    }
    .icon {
      cursor: pointer;
    }
  }
}

.sent-list {
  margin-top: 16px;
  .sent-title {
    margin-bottom: 6px;
    font-weight: bolder;
    font-size: 18px;
  }
  .sent-items {
    .sent-item {
      display: flex;
      padding: 8px;
      line-height: 30px;
      .no {
        width: 50px;
      }
      .left {
        flex: 1;
        border-bottom: 1px solid #e9f2fb;
      }
      .right {
        width: 160px;
        .icon {
          padding: 3px 5px;
          font-size: 18px;
          cursor: pointer;
        }
      }
    }
  }
}

.edit-icon {
  color: gray;
}
.mouse-over {
  background-color: #f5f5f5;
}

</style>
