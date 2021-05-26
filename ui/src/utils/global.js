import Vue from 'vue'

const EventBus = new Vue()

export default {
  EventBus,

  ws: {},
  wsEventName: 'wsEvent',
  defaultRoom: 'square',
  setWs: function (conn) {
    this.ws = conn
  }
}
