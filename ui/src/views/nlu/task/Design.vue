<template>
  <div id="design-page">
    <a-modal
      :title="$t('menu.task.design')"
      width="100%"
      dialogClass="full-screen-modal"
      :visible="visible"
      :closable=true
      :footer="null"
      @cancel="cancel"
    >
      <div class="main">
        <div class="left" :style="styl">
          <intent-list
            ref="intentList"
            :models="model.intents"
            @selected="select">
          </intent-list>
        </div>
        <div class="right" :style="styl">
          <intent-edit
            ref="intentEdit"
            :modelId="intentId"
            :visible="intentEditVisible">
          </intent-edit>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script>
import { getTask } from '@/api/manage'
import IntentList from '../intent/List'
import IntentEdit from '../intent/Edit'

export default {
  name: 'TaskDesign',
  components: {
    IntentList, IntentEdit
  },
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    modelProp: {
      type: Object,
      default: () => null
    },
    time: {
      type: Number,
      default: () => 0
    }
  },
  data () {
    const styl = 'height: ' + (document.documentElement.clientHeight - 56) + 'px;'
    return {
      model: {},
      intentId: 0,
      intentEditVisible: false,
      styl: styl
    }
  },
  watch: {
    time: function () {
      console.log('watch time', this.modelProp, this.time)
      this.getModel()
    },
    visible: function () {
      if (this.visible) {
        document.addEventListener('click', this.clearMenu)
      } else {
        document.removeEventListener('click', this.clearMenu)
      }
    }
  },
  mounted () {
  },
  created () {
  },
  methods: {
    getModel () {
      getTask(this.modelProp.id, true).then(json => {
        this.model = json.data
      })
    },
    cancel () {
      console.log('cancel')
      this.$emit('cancel')
    },
    clearMenu () {
      console.log('clear context menu')
    },
    select (intentId) {
      console.log('select', intentId)
      this.intentId = intentId
      this.intentEditVisible = true
    }
  }
}
</script>

<style lang="less" scoped>
.main {
  display: flex;
  .left {
    padding: 8px;
    width: 220px;
    height: 100%;
    border-right: 1px solid #e9f2fb;
  }
  .right {
    flex: 1;
    padding: 8px;
    height: 100%;
  }
}
</style>
