<template>
  <div id="design-page">
    <a-modal
      :title="$t('menu.task.design')"
      width="100%"
      dialogClass="full-screen-modal"
      :visible="visible"
      :closable="true"
      :footer="null"
      @cancel="cancel"
    >
      <div class="main">
        <div class="left" :style="styl">
          <intent-tree
            ref="intentTree"
            :taskId="model.id"
            :time="time2"
            @selected="select">
          </intent-tree>
        </div>
        <div class="right">
          <intent-edit
            v-if="intentEditVisible"
            ref="intentEdit"
            :modelId="intentId"
            :projectId="model.projectId">
          </intent-edit>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script>
import { getTask } from '@/api/manage'
import IntentTree from '../intent/Tree'
import IntentEdit from '../intent/Edit'

export default {
  name: 'TaskDesign',
  components: {
    IntentTree, IntentEdit
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
    const styl = 'overflow-y: auto; height: ' + (document.documentElement.clientHeight - 56) + 'px;'
    return {
      model: {},
      intentId: 0,
      intentEditVisible: false,
      styl: styl,
      time2: 0
    }
  },
  watch: {
    time: function () {
      console.log('watch time', this.modelProp, this.time)
      this.getModel()
    },
    visible: function () {
      console.log('watch visible', this.visible)

      if (this.visible) {
        this.intentEditVisible = true
        document.addEventListener('click', this.clearMenu)
      } else {
        this.intentEditVisible = false
        this.intentId = 0
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
        console.log('getTask', this.model)

        this.time2 = Date.now()
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
      console.log('select', intentId, this.intentEditVisible)
      this.intentId = intentId
    }
  }
}
</script>

<style lang="less" scoped>
.main {
  display: flex;
  .left {
    padding: 8px;
    width: 280px;
    height: 100%;
    border-right: 1px solid #e9f2fb;
  }
  .right {
    flex: 1;
    padding: 8px 15px;
    height: 100%;
  }
}
</style>
