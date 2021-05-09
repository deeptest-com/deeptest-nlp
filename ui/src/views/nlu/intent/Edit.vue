<template>
  <div v-show="model.id">
    <div class="header">
      <div class="title">{{ model.name }}</div>
      <div class="buttons">
        <a-button @click="train()" type="primary">{{ $t('common.train') }}</a-button>
        <a-button @click="test()">{{ $t('common.test') }}</a-button>
      </div>
    </div>

    <div class="edit">
      <div class="edit-title">
        {{ $t('form.sent.edit') }}
      </div>
      <div class="edit-links">
        <a-tag @click="useSynonym" class="tag synonym">{{ $t('form.use.synonym') }}</a-tag>
        <a-tag @click="useLookup" class="tag lookup">{{ $t('form.use.lookup') }}</a-tag>
        <a-tag @click="useRegex" class="tag regex">{{ $t('form.use.regex') }} </a-tag>
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
          <a-button @click="save()">{{ $t('form.save') }}</a-button>
        </div>
      </div>
    </div>

    <div class="sent-list">
      <div class="sent-title">
        {{ $t('form.sent.list') }}
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

    <a-modal
      :title="$t('form.mark')"
      dialogClass="mark-dialog"
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
          </a-radio-group>
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
  </div>
</template>

<script>

import { convertSelectedToSlots, genSent, genSentSlots } from '@/utils/markUtil'
import { getIntent, loadDicts, getSent, saveSent, removeSent, disableSent } from '@/api/manage'

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

      labelCol: { span: 5 },
      wrapperCol: { span: 15 },

      allSlots: [],
      selectedSlot: {},
      selectedIndex: -1,
      slot: {},

      dicts: [],
      rules: {
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
    useRegex () {
      console.log('useRegex')
    },
    useSynonym () {
      console.log('useSynonym')
    },
    useLookup () {
      console.log('useLookup')
    },
    editSent (item) {
      console.log('editSent')
      getSent(item.id).then(json => {
        this.sent = json.data
        this.$refs.editor.innerHTML = this.sent.html
      })
    },
    save () {
      console.log('save')
      this.sent.intentId = this.model.id
      this.sent.html = this.$refs.editor.innerHTML
      this.sent.text = this.$refs.editor.innerText.replace(/\s+/g, '')
      this.sent.slots = genSentSlots(document.getElementById('editor'))

      saveSent(this.sent).then(json => {
        this.sents = json.data
        this.sent = {}
        this.$refs.editor.innerHTML = ''
      })
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

    reset () {
      this.model = {}
    },
    back () {
      this.$router.push('/nlu/task/list')
    },

    textSelected (event) {
      const mp = convertSelectedToSlots(event.target, document.getElementById('editor'))
      if (!mp.selectedIndex) return
      mp.allSlots.forEach((item, index) => {
        console.log('=' + index + '=', index, item)
      })

      this.allSlots = mp.allSlots
      this.selectedIndex = mp.selectedIndex
      this.selectedSlot = this.allSlots[this.selectedIndex]
      console.log('=curr=', this.selectedSlot)

      let id = this.selectedSlot.getAttribute('data-value')
      id = id ? parseInt(id) : ''
      this.slot = { slotType: this.selectedSlot.getAttribute('data-type'), value: id }
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
  margin-top: 12px;
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
      cursor: pointer;
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
        border: 1px solid #e9f2fb;
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

</style>
