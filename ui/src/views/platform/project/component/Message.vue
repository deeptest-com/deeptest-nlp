<template>
  <div>
    <div v-if="msg.type!=='question'" class="message a">
      <div class="avatar">
        <a-avatar slot="avatar" icon="android" class="my-avatar-icon" :style="{ fontSize: '23px' }" />
      </div>
      <div class="box a">
        <div class="content">
          <span v-if="msg.type==='welcome'">{{ $t('msg.testing.can.not.understand') }}</span>
          <span v-if="msg.type!=='welcome'">{{ msg.content }}</span>
        </div>
        <div class="action">
          <a @click="view('result')">{{ $t('common.view.result') }}</a>
          &nbsp;&nbsp;&nbsp;  | &nbsp;&nbsp;&nbsp;
          <a @click="view('json')">{{ $t('common.view.json') }}</a>
          &nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;
          <a @click="view('')">{{ $t('common.view.nothing') }}</a>
        </div>
      </div>
    </div>

    <div v-if="msg.type==='question'" class="message q">
      <div class="avatar">
        <a-avatar slot="avatar" icon="user" class="my-avatar-icon" :style="{ fontSize: '23px' }" />
      </div>
      <div class="box q">
        <div class="content">
          <span v-if="msg.type!=='welcome'">{{ msg.content }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ChatMessage',
  props: {
    msg: {}
  },
  data () {
    return {
      typeMap: {}
    }
  },
  watch: {
  },
  mounted () {
  },
  created () {
    this.typeMap = {
      true: {
        type: 'processing',
        text: this.$t('status.enable')
      },
      false: {
        type: 'default',
        text: this.$t('status.disable')
      }
    }
  },
  filters: {
    typeFilter (type, typeMap) {
      return typeMap[type].type
    }
  },
  methods: {
    view (mode) {
      this.$emit('view', mode)
    }
  }
}
</script>

<style lang="less" scoped>

.message {
  position: relative;
  display: flex;
  margin-bottom: 10px;

  .avatar {
    width: 30px;
    padding: 15px 0;
  }
  &.q {
    flex-direction:row-reverse;
    .avatar {
      width: 20px;
      padding: 10px 0;
    }
  }
  .box{
    position: relative;
    padding: 10px;
    min-width: 220px;
    top:0px;
    height: 60px;
    -moz-border-radius: 12px;
    -webkit-border-radius: 12px;
    border-radius: 12px;

    &.a {
      left: 28px;
      background: #f2f4f5;
      &:before {
        position: absolute;
        content: "";
        width: 0;
        height: 0;
        right: 100%;
        top: 18px;
        border-top: 13px solid transparent;
        border-bottom: 13px solid transparent;
        border-right: 26px solid #f2f4f5;
      }
    }
    &.q {
      height: 45px;
      right: 25px;
      background: #f2f4f5;
      &:after {
        position: absolute;
        content: "";
        width: 0;
        height: 0;
        right: -22px;
        top: 10px;
        border-top: 13px solid transparent;
        border-bottom: 13px solid transparent;
        border-left: 26px solid #f2f4f5;
      }
    }

    .content {
      color: rgba(0, 0, 0, 0.45);
      text-align: left;
    }
    .action {
      font-size: 12px;
      line-height: 23px;
    }
  }
}

</style>
