<template>
  <div>
    INTENT EDIT
  </div>
</template>

<script>

export default {
  name: 'IntentEdit',
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
    return {
      model: {}
    }
  },
  watch: {
    time: function () {
      console.log('watch time', this.modelProp, this.time)
      this.getModel()
    }
  },
  methods: {
    getModel () {
      console.log('getModel')
    },
    save (e) {
      console.log('save')
    },
    reset () {
      this.model = {}
      // this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/nlu/task/list')
    }
  }
}
</script>
