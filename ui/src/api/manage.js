import request from '@/utils/request'

const prefix = '/v1/admin'

const api = {
  profile: `${prefix}/profile`,
  projects: `${prefix}/projects`,
  lookups: `${prefix}/lookups`,
  synonyms: `${prefix}/synonyms`,

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
  return request({
    url: !model.id ? api.projects : api.projects + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function setDefaultProject (model) {
  return request({
    url: api.projects + '/' + model.id + '/setDefault',
    method: 'post',
    params: {}
  })
}
export function disableProject (model) {
  return request({
    url: api.projects + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeProject (model) {
  return request({
    url: api.projects + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

export function listLookup (params) {
  return request({
    url: api.lookups,
    method: 'get',
    params: params
  })
}
export function getLookup (id) {
  return request({
    url: api.lookups + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveLookup (model) {
  return request({
    url: !model.id ? api.lookups : api.lookups + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableLookup (model) {
  return request({
    url: api.lookups + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeLookup (model) {
  return request({
    url: api.lookups + '/' + model.id,
    method: 'delete',
    params: {}
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
