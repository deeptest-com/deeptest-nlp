<template>
  <div v-show="model.id">
    <div class="header">
      <div v-if="!isEditTitle" class="title">
        <span>{{ model.name }}</span>&nbsp;&nbsp;
        <a @click="editTitle()"><a-icon class="edit-icon" type="edit" /></a>
      </div>
      <div v-if="isEditTitle">
        <a-form layout="inline">
          <a-form-item>
            <a-input v-model="model.name" type="text"></a-input>
          </a-form-item>
          <a-form-item>
            <a @click="saveTitle()"><a-icon type="check" /></a>
          </a-form-item>
        </a-form>
      </div>
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
          <div class="sent-items">
            <div v-for="item in sents" :key="item.id" ref="sentList" class="sent-item">
              <div class="left" :class="{'disabled':item.disabled}">{{ item.text }}</div>
              <div class="right">
                <a-icon @click="editSent(item)" type="edit" class="icon"/> &nbsp;
                <a-icon v-if="!item.disabled" @click="toDisableSent(item)" type="minus" class="icon"/>
                <a-icon v-if="item.disabled" @click="toDisableSent(item)" type="plus" class="icon"/>
                &nbsp;
                <a-icon @click="toDeleteSent(item)" type="delete" class="icon"/>
              </div>
            </div>
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
                v-model="rule.expr"
                @click="ruleInputClick"
                ref="ruleExpr"
                id="ruleExpr"
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
                <a-icon @click="editSent(item)" type="edit" class="icon"/> &nbsp;
                <a-icon v-if="!item.disabled" @click="toDisableSent(item)" type="minus" class="icon"/>
                <a-icon v-if="item.disabled" @click="toDisableSent(item)" type="plus" class="icon"/>
                &nbsp;
                <a-icon @click="toDeleteSent(item)" type="delete" class="icon"/>
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
        {{ ruleSection.slotType }}
        <a-form-model-item prop="value" v-if="ruleSectionType === '_slot_'" :label="$t('form.slot')">
          <a-input v-model="ruleSection.value" />
        </a-form-model-item>
        <a-form-model-item prop="value" v-if="ruleSectionType !== '_slot_'" :label="$t('form.synonym')">
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

import { convertSelectedToSlots, genSent, genSentSlots } from '@/utils/markUtil'
import { getIntent, updateIntent, loadDicts, getSent, saveSent, removeSent, disableSent } from '@/api/manage'

export default {
  name: 'IntentEdit',
  props: {
    modelId: {
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
    return {
      model: {},
      sents: [],
      sent: { content: '' },
      slotEditVisible: false,

      ruleList: [],
      rule: { expr: '' },
      ruleSection: {},
      ruleInputIndex: -1,
      ruleSectionType: '',
      ruleSectionEditVisible: false,

      labelCol: { span: 5 },
      wrapperCol: { span: 15 },

      allSlots: [],
      selectedSlot: {},
      selectedIndex: -1,
      slot: {},
      isEditTitle: false,
      tabKey: 'maintainSent',

      dicts: [],
      rules: {
        slotType: [{ required: true, message: this.$t('valid.slot.type'), trigger: 'change' }],
        value: [{ required: true, message: this.$t('valid.select.dict'), trigger: 'change' }]
      }
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
    resetSent () {
      console.log('resetSent')
      this.rule = {}
      this.$refs.editor.innerHTML = ''
    },

    useDict (dictType) {
      console.log('useDict', dictType)
      this.ruleSectionType = dictType
      this.ruleSectionEditVisible = true

      if (dictType === '_slot_') return

      loadDicts(dictType).then(json => {
        console.log(json)
        this.dicts = json.data
      })
    },
    ruleInputClick () {
      this.ruleInputIndex = this.$refs.ruleExpr.$el.selectionStart
      console.log('this.ruleExpr', this.ruleInputIndex)
    },
    addRuleSection () {
      console.log('addRuleSection')

      let val = this.ruleSection.value
      if (this.ruleSectionType === '_slot_') {
        val = '(' + val + ')'
      } else {
        val = '{' + val + '}'
      }
      this.rule.expr = this.rule.expr.substring(0, this.ruleInputIndex) + val +
        this.rule.expr.substring(this.ruleInputIndex, this.rule.expr.length)

      this.ruleSectionEditVisible = false
      this.ruleSection.value = ''
    },
    cancelRuleSection () {
      console.log('cancelRuleSection')
      this.ruleSectionEditVisible = false
    },
    saveRule () {
      console.log('saveRule')
      this.ruleSectionEditVisible = false
    },
    resetRule () {
      console.log('resetRule')
      this.rule = {}
    },

    toDeleteSent (item) {
      console.log('toDeleteSent')

      removeSent(item).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
    },
    toDisableSent (item) {
      console.log('toDisableSent')

      disableSent(item).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
    },

    editTitle () {
      this.isEditTitle = true
    },
    saveTitle () {
      const data = { id: this.model.id, name: this.model.name }
      updateIntent(data).then(json => {
        this.isEditTitle = false
      })
    },
    tabClick (key) {
      this.tabKey = key
    },
    back () {
      this.$router.push('/nlu/task/list')
    },

    textSelected (event) {
      const mp = convertSelectedToSlots(event.target, document.getElementById('editor'))
      if (!mp.allSlots) return
      // if (!mp.selectedIndex) return
      mp.allSlots.forEach((item, index) => {
        console.log('=' + index + '=', index, item)
      })

      this.allSlots = mp.allSlots
      this.selectedIndex = mp.selectedIndex
      this.selectedSlot = this.allSlots[this.selectedIndex]
      console.log('=curr=', this.selectedSlot)

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
      if (this.slot.slotType === '_slot_') return

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
    }
  }
}
</script>

<style lang="less" scoped>
.header {
  display: flex;
  border-bottom: 1px solid #e9f2fb;
  .title {
    flex: 1;
    font-weight: bolder;
    font-size: 20px;
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
      margin-bottom: 8px;
      line-height: 22px;

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

</style>
