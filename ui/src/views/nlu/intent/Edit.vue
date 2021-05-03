<template>
  <div v-if="model.id">
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
          <div contenteditable="true" class="editor" ref="editor" spellcheck="false"></div>
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

  </div>
</template>

<script>

import { getIntent } from '@/api/manage'

export default {
  name: 'IntentEdit',
  props: {
    modelId: {
      type: Number,
      default: () => 0
    }
  },
  mounted () {
    document.addEventListener('mouseup', this.textSelected)
  },
  destroyed () {
    document.removeEventListener('mouseup', this.textSelected)
  },
  data () {
    return {
      model: {},
      sents: [],
      sent: { content: '' }
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
            this.$refs.editor.innerHTML = '<span id="1">aaa</span><span id="2">123</span><span id="3">bbb</span>'
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
    getParentSpanNodeIfNeeded (target) {
      if (target.parentNode && target.parentNode.nodeName.toLowerCase() === 'span') {
        target = target.parentNode
      }
      return target
    },
    textSelected (event) {
      let target = event.target
      console.log(target)
      target = this.getParentSpanNodeIfNeeded(target)

      if (target === this.$refs.editor || target.contains(this.$refs.editor) ||
        target.parentNode === this.$refs.editor || target.parentNode.contains(this.$refs.editor)) {
        const slt = window.getSelection()
        console.log('anchorNode', slt.anchorNode)
        console.log('anchorOffset', slt.anchorOffset)

        console.log('focusNode', slt.focusNode)
        console.log('focusOffset', slt.focusOffset)

        const range = window.getSelection().getRangeAt(0)
        console.log('range', range)
        const startContainer = this.getParentSpanNodeIfNeeded(range.startContainer)
        const endContainer = this.getParentSpanNodeIfNeeded(range.endContainer)
        console.log('startContainer', startContainer)
        console.log('endContainer', endContainer)
        console.log('is same', startContainer === endContainer)

        const items = []
        let item = startContainer
        while (true) {
          item = this.getParentSpanNodeIfNeeded(item)

          let tp = item.nodeName.toLowerCase()
          tp = tp.replace('#', '')
          let html = ''
          if (tp === 'span') {
            html = item.outerHTML
          } else {
            html = item.wholeText
          }

          console.log(tp, html)
          items.push({ type: tp, content: html })

          if (item.id === endContainer.id) {
            break
          }

          if (item.nextSibling) {
            item = item.nextSibling
          } else {
            item = item.parentNode.nextSibling
          }
          if (!item) {
            break
          }
        }

        const startText = startContainer.textContent
        const endText = endContainer.textContent

        const startLeft = slt.anchorOffset
        let startRight = startText.length
        let endLeft = 0
        const endRight = slt.focusOffset

        if (startContainer === endContainer) {
          startRight = endRight
          endLeft = startLeft
        }

        items[0].selected = startText.substr(startLeft, startRight - startLeft)
        console.log('start', items[0].selected, startLeft, startRight - startLeft)

        items[items.length - 1].selected = endText.substr(endLeft, endRight - endLeft)
        console.log('end', items[items.length - 1].selected, endLeft, endRight - endLeft)

        console.log(items)

        // const commonAncestorContainer = range.commonAncestorContainer
        // const firstChild = commonAncestorContainer.firstChild
        // const nodes = commonAncestorContainer.childNodes
        // let foundFirst = false
        // nodes.forEach((item, index) => {
        //   if (item === firstChild) foundFirst = true
        //
        //   const tp = item.nodeName.toLowerCase()
        //   let html = ''
        //   if (tp === 'span') html = item.outerHTML
        //   else html = item.wholeText
        //   if (foundFirst) {
        //     console.log(index, tp, html)
        //   }
        // })
      }
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
