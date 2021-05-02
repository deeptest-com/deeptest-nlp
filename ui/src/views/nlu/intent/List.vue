<template>
  <div>
    <a-menu
      @click="select"
      style="width: 100%"
      mode="inline"
    >
      <a-menu-item v-for="item in models" :key="item.id">
        {{ item.name }}
      </a-menu-item>
    </a-menu>

  </div>
</template>

<script>

export default {
  name: 'IntentList',
  components: {
  },
  props: {
    models: {
      type: Array,
      required: true
    }
  },
  data () {
    return {
      model: {}
    }
  },
  created () {
  },
  computed: {
  },
  methods: {
    select ({ item, key }) {
      this.$emit('selected', key)
    }
  }
}
</script>
