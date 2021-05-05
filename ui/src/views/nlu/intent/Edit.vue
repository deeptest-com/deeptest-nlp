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
        <a-tag @click="useRegex" class="tag">{{ $t('form.use.regex') }} </a-tag>
        <a-tag @click="useSynonym" class="tag">{{ $t('form.use.synonym') }}</a-tag>
        <a-tag @click="useLookup" class="tag">{{ $t('form.use.lookup') }}</a-tag>
      </div>
      <div class="edit-inputs">
        <div class="left">
          <div
            id="editor"
            contenteditable="true"
            @mouseup="textSelected"
            class="editor"
            ref="editor"
            spellcheck="false">
          </div>
          <div class="tips">
            <a-icon type="info-circle" class="icon"/>
            <span>{{ $t('form.select.to.mark') }}</span>
          </div>
        </div>
        <div class="right">
          <a-button @click="add()">{{ $t('form.save') }}</a-button>
        </div>
      </div>
    </div>

    <div class="sent-list">
      <div class="sent-title">
        {{ $t('form.sent.list') }}
      </div>
      <div class="sent-items">
        <div v-for="item in sents" :key="item.id" ref="sentList" class="sent-item">
          <div class="left">{{ item.content }}</div>
          <div class="right">
            <a-icon @click="edit(item)" type="edit" class="icon"/>
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
      <a-form-model ref="form" :model="selectedSlot" :rules="rules" :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-model-item prop="slotType" :label="$t('form.slot.type')">
          <a-radio-group v-model="selectedSlot.slotType" @change="slotTypeChanged">
            <a-radio value="synonym">{{ $t('form.synonym') }}</a-radio>
            <a-radio value="lookup">{{ $t('form.lookup') }}</a-radio>
            <a-radio value="regex">{{ $t('form.regex') }}</a-radio>
          </a-radio-group>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="selectedSlot.slotType === 'synonym'" :label="$t('form.synonym')">
          <a-select v-model="selectedSlot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="selectedSlot.slotType === 'lookup'" :label="$t('form.lookup')">
          <a-select v-model="selectedSlot.value">
            <a-select-option v-for="(item, index) in dicts" :key="index" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item prop="value" v-if="selectedSlot.slotType === 'regex'" :label="$t('form.regex')">
          <a-select v-model="selectedSlot.value">
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

import { convertSelectedToSlots, genSent } from '@/utils/markUtil'
import { getIntent, loadDicts } from '@/api/manage'

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
        if (json.code === 200) {
          this.model = json.data
          this.sents = this.model.sents

          setTimeout(() => {
            this.$refs.editor.innerHTML = '<span id="1">abc</span>123<span id="3">xyz</span>'
          }, 500)
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

    add () {
      console.log('add')
      let index = -1
      for (let i = 0; i < this.sents.length; i++) {
        if (this.sents[i].id === this.sent.id) {
          index = i
          break
        }
      }

      const content = this.$refs.sent.innerHTML
      if (index > -1) {
        const item = this.sents[index]
        item.content = content
        this.sents.splice(index, 1, item)
      } else {
        const item = { content: content }
        this.sents.push(item)
      }
      this.$refs.sent.innerHTML = '<span></span>'
      this.sent = {}
    },
    edit (item) {
      console.log('edit')
      this.sent = item
      this.$refs.sent.innerHTML = this.sent.content
    },
    save (e) {
      console.log('save')
    },
    reset () {
      this.model = {}
    },
    back () {
      this.$router.push('/nlu/task/list')
    },

    textSelected (event) {
      const mp = convertSelectedToSlots(event.target, document.getElementById('editor'))
      this.allSlots = mp.all
      this.selectedSlot = mp.selected
      console.log('=1=', this.allSlots, this.selectedSlot)

      if (!this.selectedSlot) return

      this.slotEditVisible = true
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

        const sentContent = genSent(this.allItems, this.selectedItems, this.slot)
        console.log('saveSlot', sentContent)

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
        color: #52c41a;
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

<style lang="less">
.mark-dialog {
  top: 230px;
  .ant-modal-header {
    padding: 10px 24px !important;
  }
  .ant-modal-footer {
    border-top: 0 !important;
    .ant-btn {
      margin-bottom: 0px;
    }
  }
}
</style>
