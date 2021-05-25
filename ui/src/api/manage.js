import request from '@/utils/request'

const prefix = '/v1/admin'

const api = {
  profile: `${prefix}/profile`,
  projects: `${prefix}/projects`,
  convert: `${prefix}/convert`,
  training: `${prefix}/training`,
  tasks: `${prefix}/tasks`,
  intents: `${prefix}/intents`,
  sents: `${prefix}/sents`,
  slots: `${prefix}/slots`,
  lookups: `${prefix}/lookups`,
  lookupItems: `${prefix}/lookupItems`,
  synonyms: `${prefix}/synonyms`,
  synonymItems: `${prefix}/synonymItems`,
  regexes: `${prefix}/regexes`,
  regexItems: `${prefix}/regexItems`,
  dicts: `${prefix}/dicts`,
  valid: `${prefix}/valid`,

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

// 项目
export function listProject (params) {
  return request({
    url: api.projects,
    method: 'get',
    params: params
  })
}
export function listForSelect (params) {
  return request({
    url: api.projects + '/listForSelect',
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
export function convert (model) {
  return request({
    url: api.convert + '/' + model.id,
    method: 'post',
    params: { }
  })
}
export function training (model) {
  return request({
    url: api.training + '/' + model.id,
    method: 'post',
    params: { }
  })
}

// 任务
export function listTask (params) {
  return request({
    url: api.tasks,
    method: 'get',
    params: params
  })
}
export function getTask (id, withIntents) {
  return request({
    url: api.tasks + '/' + id,
    method: 'get',
    params: { withIntents: withIntents }
  })
}
export function saveTask (model) {
  return request({
    url: !model.id ? api.tasks : api.tasks + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableTask (model) {
  return request({
    url: api.tasks + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeTask (model) {
  return request({
    url: api.tasks + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

// 意图
export function listIntent (params) {
  return request({
    url: api.intents,
    method: 'get',
    params: params
  })
}
export function getIntent (id) {
  return request({
    url: api.intents + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveIntent (model) {
  return request({
    url: !model.id ? api.intents : api.intents + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableIntent (model) {
  return request({
    url: api.intents + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeIntent (model) {
  return request({
    url: api.intents + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

export function listSent (params) {
  return request({
    url: api.sents,
    method: 'get',
    params: params
  })
}
export function getSent (id) {
  return request({
    url: api.sents + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveSent (model) {
  return request({
    url: !model.id ? api.sents : api.sents + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableSent (model) {
  return request({
    url: api.sents + '/' + model.id + '/disable',
    method: 'post',
    params: { intentId: model.intentId }
  })
}
export function removeSent (model) {
  return request({
    url: api.sents + '/' + model.id,
    method: 'delete',
    params: { intentId: model.intentId }
  })
}

export function listSlot (params) {
  return request({
    url: api.slots,
    method: 'get',
    params: params
  })
}
export function getSlot (id) {
  return request({
    url: api.slots + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveSlot (model) {
  return request({
    url: !model.id ? api.slots : api.slots + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableSlot (model) {
  return request({
    url: api.slots + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeSlot (model) {
  return request({
    url: api.slots + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

// 同义词表
export function listSynonym (params) {
  return request({
    url: api.synonyms,
    method: 'get',
    params: params
  })
}
export function getSynonym (id) {
  return request({
    url: api.synonyms + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveSynonym (model) {
  return request({
    url: !model.id ? api.synonyms : api.synonyms + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableSynonym (model) {
  return request({
    url: api.synonyms + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeSynonym (model) {
  return request({
    url: api.synonyms + '/' + model.id,
    method: 'delete',
    params: {}
  })
}
// 同义词表项
export function listSynonymItem (params) {
  return request({
    url: api.synonymItems,
    method: 'get',
    params: params
  })
}
export function getSynonymItem (id) {
  return request({
    url: api.synonymItems + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveSynonymItem (model) {
  return request({
    url: !model.id ? api.synonymItems : api.synonymItems + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableSynonymItem (model) {
  return request({
    url: api.synonymItems + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeSynonymItem (model) {
  return request({
    url: api.synonymItems + '/' + model.id,
    method: 'delete',
    params: {}
  })
}
export function batchRemoveSynonymItem (data) {
  return request({
    url: api.synonymItems + '/batchRemove',
    method: 'post',
    data: data
  })
}

// 同类词表
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
// 同类词表项
export function listLookupItem (params) {
  return request({
    url: api.lookupItems,
    method: 'get',
    params: params
  })
}
export function getLookupItem (id) {
  return request({
    url: api.lookupItems + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveLookupItem (model) {
  return request({
    url: !model.id ? api.lookupItems : api.lookupItems + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableLookupItem (model) {
  return request({
    url: api.lookupItems + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeLookupItem (model) {
  return request({
    url: api.lookupItems + '/' + model.id,
    method: 'delete',
    params: {}
  })
}
export function batchRemoveLookupItem (data) {
  return request({
    url: api.lookupItems + '/batchRemove',
    method: 'post',
    data: data
  })
}

// 正则表达式
export function listRegex (params) {
  return request({
    url: api.regexes,
    method: 'get',
    params: params
  })
}
export function getRegex (id) {
  return request({
    url: api.regexes + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveRegex (model) {
  return request({
    url: !model.id ? api.regexes : api.regexes + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableRegex (model) {
  return request({
    url: api.regexes + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeRegex (model) {
  return request({
    url: api.regexes + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

// 正则表达式项
export function listRegexItem (params) {
  return request({
    url: api.regexItems,
    method: 'get',
    params: params
  })
}
export function getRegexItem (id) {
  return request({
    url: api.regexItems + '/' + id,
    method: 'get',
    params: {}
  })
}
export function saveRegexItem (model) {
  return request({
    url: !model.id ? api.regexItems : api.regexItems + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableRegexItem (model) {
  return request({
    url: api.regexItems + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeRegexItem (model) {
  return request({
    url: api.regexItems + '/' + model.id,
    method: 'delete',
    params: {}
  })
}
export function batchRemoveRegexItem (data) {
  return request({
    url: api.regexItems + '/batchRemove',
    method: 'post',
    data: data
  })
}

export function loadDicts (type) {
  return request({
    url: api.dicts,
    method: 'get',
    params: { type: type }
  })
}
export function validProjectPath (value) {
  return request({
    url: api.valid,
    method: 'post',
    data: { method: 'validProjectPath', value: value }
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
