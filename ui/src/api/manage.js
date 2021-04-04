import request from '@/utils/request'

const prefix = '/v1/admin'

const api = {
  profile: `${prefix}/profile`,
  projects: `${prefix}/projects`,

  user: `${prefix}/user`,
  role: `${prefix}/role`,
  service: `${prefix}/service`,
  permission: `${prefix}/permission`,
  permissionNoPager: `${prefix}/permission/no-pager`,
  orgTree: `${prefix}/org/tree`
}

export const WsApi = 'ws://127.0.0.1:8085/api/v1/ws'

export function requestSuccess (code) {
  return code === 200
}

export function getProfile (parameter) {
  return request({
    url: api.profile,
    method: 'get',
    data: parameter
  })
}

export function listProject (params) {
  return request({
    url: api.projects,
    method: 'get',
    params: params
  })
}
export function getProject (id) {
  return request({
    url: api.projects + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveProject (model) {
  console.log('---', model.id)
  return request({
    url: api.projects,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}

export function getUserList (parameter) {
  return request({
    url: api.user,
    method: 'get',
    params: parameter
  })
}

export function getRoleList (parameter) {
  return request({
    url: api.role,
    method: 'get',
    params: parameter
  })
}

export function getServiceList (parameter) {
  return request({
    url: api.service,
    method: 'get',
    params: parameter
  })
}

export function getPermissions (parameter) {
  return request({
    url: api.permissionNoPager,
    method: 'get',
    params: parameter
  })
}

// id == 0 add     post
// id != 0 update  put
export function saveService (parameter) {
  return request({
    url: api.service,
    method: parameter.id === 0 ? 'post' : 'put',
    data: parameter
  })
}
