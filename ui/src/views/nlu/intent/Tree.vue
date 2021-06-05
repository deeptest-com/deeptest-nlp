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
    />
    <div v-if="treeNode" :style="this.tmpStyle" class="tree-context-menu">
      <a-menu @click="menuClick" mode="inline" class="menu">
        <a-menu-item key="addNeighbor" v-if="!isRoot">
          <a-icon type="plus" />{{ $t('menu.create.intent') }}
        </a-menu-item>
        <a-menu-item key="addChild" v-if="isRoot">
          <a-icon type="plus" />{{ $t('menu.create.intent') }}
        </a-menu-item>
        <a-menu-item key="remove" v-if="!isRoot">
          <a-icon type="delete" />{{ $t('menu.remove.intent') }}
        </a-menu-item>
      </a-menu>
    </div>

    <a-modal
      title="确认删除"
      :width="400"
      :visible="removeVisible"
      okText="确认"
      cancelText="取消"
      @ok="removeNode"
      @cancel="cancelRemove">
      <div>确认删除选中节点？</div>
    </a-modal>
  </div>
</template>

<script>

import { listIntent, createIntent, removeIntent, moveIntent, wrapperIntents } from '@/api/manage'

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
  },
  mounted () {
  },
  computed: {
    isRoot () {
      console.log('isRoot', this.treeNode)
      return this.treeNode.id === 0
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
        title: node._props.title
        // parentID: node._props.dataRef.parentID || null
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
      console.log(info, info.dragNode.eventKey, info.node.eventKey, info.dropPosition)

      moveIntent(info.dragNode.eventKey, info.node.eventKey, info.dropPosition).then(res => {
        this.treeData = [res.data]

        this.selectedKeys = [res.modelId] // select
        this.$emit('selected', res.modelId)
      })
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
      this.treeData = wrapperIntents(json.data, this.$t('menu.intent'))
      this.getOpenKeys(this.treeData[0])

      // this.selectedKeys = [json.model.id] // select
    },
    removeNode () {
      console.log('removeNode', this.targetModel)
      this.removeVisible = false
      removeIntent(this.targetModel).then(json => {
        console.log('removeIntent', json)
        this.removeCallback(json)
      })
    },
    removeCallback (json) {
      this.treeData = [json.data]
    },
    cancelRemove (e) {
      e.preventDefault()
      this.removeVisible = false
    },
    getOpenKeys (node) {
      if (!node) return

      this.openKeys.push(node.id)
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
