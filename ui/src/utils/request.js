import axios from 'axios'
// import store from '@/store'
import storage from 'store'
import router from '../router'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'

const request = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL,
  timeout: 5000
})

const errorHandler = (error) => {
  console.log('===errorHandler===', error)
  if (error.response) {
    const data = error.response.data
    // 从 localstorage 获取 token
    // const token = storage.get(ACCESS_TOKEN)
    if (error.response.status === 403) {
      notification.error({
        message: 'Forbidden',
        description: data.message
      })
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      notification.error({
        message: 'Unauthorized',
        description: 'Authorization verification failed'
      })
      // if (token) {
      //   store.dispatch('Logout').then(() => {
      //     setTimeout(() => {
      //       window.location.reload()
      //     }, 1500)
      //   })
      // }

      router.push({ path: '/user/login' })
    }
  }
  return Promise.reject(error)
}

// request interceptor
request.interceptors.request.use(config => {
  const method = ('' + config.method).toLowerCase()
  if (config.params) {
    if (method === 'get') {
      let queryParams = ''
      let i = 0
      for (const key in config.params) {
        queryParams += key + '=' + config.params[key]
        if (i < config.params - 1) queryParams += '&'
        i++
      }
      console.log(queryParams)
      if (config.url.indexOf('?') < 0) {
        config.url += '?' + queryParams
      } else {
        config.url += '&' + queryParams
      }
    } else {
      for (const key in config.params) {
        config.params[key] = '' + config.params[key]
      }
    }
  }
  console.log('===Request===', config)

  const jwtToken = storage.get(ACCESS_TOKEN)
  if (jwtToken) {
    config.headers['Authorization'] = 'Bearer ' + jwtToken
    console.log('add token in request header', jwtToken)
  }

  return config
}, errorHandler)

// response interceptor
request.interceptors.response.use((response) => {
  console.log('===Response===', response)

  if (response.data.code === 401) {
    if (router.currentRoute.path !== '/user/login') {
      router.push({ path: '/user/login' })
    }

    return
  }

  return response.data
}, errorHandler)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, request)
  }
}

export default request

export {
  installer as VueAxios,
  request as axios
}
