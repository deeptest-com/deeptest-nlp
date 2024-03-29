<template>
  <div>
    <a-tree
      ref="fieldTree"
      class="draggable-tree"
      :show-line="true"
      :show-icon="false"
      :expandedKeys.sync="openKeys"
      :selectedKeys.sync="selectedKeys"
      :tree-data="treeData"
      :replaceFields="replaceFields"
      @select="onSelect"
      @rightClick="onRightClick"
      :draggable="true"
      @dragenter="onDragEnter"
      @drop="onDrop"
      @load="load"
    />
    <div v-if="treeNode" :style="this.tmpStyle" class="tree-context-menu">
      <a-menu @click="menuClick" mode="inline" class="menu">
        <template v-if="!isRoot">
          <a-menu-item key="addNeighbor" >
            <a-icon type="plus" />{{ $t('menu.create.intent') }}
          </a-menu-item>

          <a-menu-item key="disable" v-if="isDisabled">
            <a-icon type="reload" />{{ $t('menu.enable.intent') }}
          </a-menu-item>
          <a-menu-item key="disable" v-if="!isDisabled">
            <a-icon type="pause" />{{ $t('menu.disable.intent') }}
          </a-menu-item>

          <a-menu-item key="remove">
            <a-icon type="delete" />{{ $t('menu.remove.intent') }}
          </a-menu-item>
        </template>

        <a-menu-item key="addChild" v-if="isRoot">
          <a-icon type="plus" />{{ $t('menu.create.intent') }}
        </a-menu-item>
      </a-menu>
    </div>

    <a-modal
      :title="$t('form.pls.select')"
      :width="400"
      :visible="removeVisible"
      :okText="$t('form.ok')"
      :cancelText="$t('form.cancel')"
      @ok="removeNode"
      @cancel="cancelRemove">
      <div>{{ $t('form.confirm.to.remove') }}</div>
    </a-modal>
  </div>
</template>

<script>

import { listIntent, createIntent, disableIntent, removeIntent, moveIntent, wrapperIntents } from '@/api/manage'

export default {
  name: 'IntentTree',
  components: {
  },
  props: {
    taskId: {
      type: Number,
      default: () => 0
    },
    time: {
      type: Number,
      default: () => 0
    }
  },
  data () {
    return {
      treeData: [],
      treeNode: null,
      openKeys: [],
      selectedKeys: [],
      targetId: 0,
      removeVisible: false,
      replaceFields: { title: 'name', key: 'id', value: 'id', children: 'children' }
    }
  },
  watch: {
    time: function () {
      console.log('watch time in tree', this.treeData, this.time)

      this.loadData()
      this.getOpenKeys(this.treeData[0])
    }
  },
  created () {
    this.$global.EventBus.$on(this.$global.intentUpdateEvent, (json) => {
      console.log('IntentUpdateEvent in tree page', json)
      this.loadData()
    })
  },
  destroyed () {
    this.$global.EventBus.$destroy()
  },
  mounted () {
  },
  computed: {
    isRoot () {
      console.log('isRoot', this.treeNode)
      return this.treeNode.id === 0
    },
    isDisabled () {
      console.log('isDisabled', this.treeNode)
      return this.treeNode.disabled
    }
  },
  methods: {
    loadData () {
      listIntent(this.taskId).then(json => {
        console.log('listIntent', json)
        this.updateCallback(json)
      })
    },
    onSelect (selectedKeys, e) { // selectedKeys, e:{selected: bool, selectedNodes, node, event}
      console.log('onSelect', selectedKeys, e.selectedNodes, e.node, e.node.eventKey)
      if (selectedKeys.length === 0) {
        selectedKeys[0] = e.node.eventKey // keep selected
      }

      if (e.node.eventKey !== 0) {
        this.$emit('selected', e.node.eventKey)
      }
    },
    onRightClick ({ event, node }) {
      event.preventDefault()
      console.log('onRightClick', node)

      const y = event.currentTarget.getBoundingClientRect().top
      const x = event.currentTarget.getBoundingClientRect().right

      this.treeNode = {
        pageX: x,
        pageY: y,
        id: node._props.eventKey,
        title: node._props.title,
        disabled: node.disabled
      }

      this.tmpStyle = {
        position: 'fixed',
        maxHeight: 40,
        textAlign: 'center',
        left: `${x + 10 - 0}px`,
        top: `${y + 6 - 0}px`
        // display: 'flex',
        // flexDirection: 'row'
      }
    },
    menuClick (e) {
      console.log('menuClick', e, this.treeNode)

      this.targetId = this.treeNode.id
      if (e.key === 'addNeighbor') {
        this.addNeighbor()
      } else if (e.key === 'addChild') {
        this.addChild()
      } else if (e.key === 'disable') {
        this.disableNode()
      } else if (e.key === 'remove') {
        this.removeVisible = true
      }
      this.clearMenu()
    },
    clearMenu () {
      console.log('clearMenu')
      this.treeNode = null
    },
    onDragEnter (info) {
      console.log(info)
      // expandedKeys 需要受控时设置
      this.expandedKeys = info.expandedKeys
    },
    onDrop (info) {
      console.log('onDrop', info, info.dropToGap, info.dragNode.eventKey, info.node.eventKey, info.dropPosition)
      if (info.node.eventKey === 0) return

      const pos = info.dropPosition - Number(info.node.pos.split('-')[2])
      moveIntent(info.dragNode.eventKey, info.node.eventKey, pos, this.taskId).then(json => {
        console.log('moveIntent', json)
        this.updateCallback(json)
      })
    },
    load (loadedKeys, event, node) {
      console.log('filterTreeNode', node)
    },
    addChild () {
      console.log('addChild', this.targetId)

      createIntent(this.taskId, this.targetId, 'child', this.$t('menu.intent.new')).then(json => {
        console.log('createIntent', json)
        this.updateCallback(json)
      })
    },
    addNeighbor () {
      console.log('addNeighbor', this.targetId)

      createIntent(this.taskId, this.targetId, 'neighbor', this.$t('menu.intent.new')).then(json => {
        console.log('createIntent', json)
        this.updateCallback(json)
      })
    },
    updateCallback (json) {
      console.log(json.data)

      this.treeData = wrapperIntents(json.data.models, this.$t('menu.intent'))
      this.getOpenKeys(this.treeData[0])

      if (json.data && json.data.model) {
        this.selectedKeys = [json.data.model.id]
        this.$emit('selected', json.data.model.id)
      }
    },
    disableNode () {
      console.log('disableNode', this.targetId)
      disableIntent(this.targetId, this.taskId).then(json => {
        console.log('disableIntent', json)
        this.updateCallback(json)
      })
    },
    removeNode () {
      console.log('removeNode', this.targetId)
      this.removeVisible = false
      removeIntent(this.targetId, this.taskId).then(json => {
        console.log('removeIntent', json)
        this.updateCallback(json)
      })
    },
    cancelRemove (e) {
      e.preventDefault()
      this.removeVisible = false
    },
    getOpenKeys (node) {
      if (!node) return

      this.openKeys.push(node.id)
      if (node.disabled) node.disabled = true

      if (node.children) {
        node.children.forEach((item) => {
          this.getOpenKeys(item)
        })
      }
    }
  }
}
</script>

<style lang="less" scoped>
.tree-context-menu {
  z-index: 9;
  .ant-tree-node-content-wrapper {
    display: block !important;
  }
  .menu {
    border: 1px solid #ebedf0;
    background: #f0f2f5;
    .ant-menu-item {
      padding-left: 12px !important;
      height: 22px;
      line-height: 21px;
      text-align: left;
    }
  }
}
</style>
