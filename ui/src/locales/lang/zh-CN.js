import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

const locale = {
  'message': '-',
  'menu.home': '主页',

  'menu.platform': '智能平台',
  'menu.platform.dashboard': '仪表盘',
  'menu.sys.admin': '系统管理',
  'menu.sys.settings': '全局设置',

  'menu.dashboard': '仪表盘',
  'menu.dashboard.analysis': '分析页',
  'menu.dashboard.monitor': '监控页',
  'menu.dashboard.workplace': '工作台',

  'menu.nlu': '自然语言理解',
  'menu.project': '项目',
  'menu.project.list': '项目列表',
  'menu.project.edit': '项目编辑',
  'menu.intent': '意图',
  'menu.intent.list': '意图列表',
  'menu.intent.edit': '意图编辑',
  'menu.synonym': '同义词',
  'menu.synonym.list': '同义词列表',
  'menu.synonym.edit': '同义词编辑',
  'menu.lookup': '词表',
  'menu.lookup.list': '词表列表',
  'menu.lookup.edit': '词表编辑',

  'common.notify': '通知',

  'form.create': '新建',
  'form.edit': '编辑',
  'form.remove': '删除',
  'form.disable': '禁用',
  'form.enable': '启用',
  'form.back': '返回',
  'form.save': '保存',
  'form.submit': '提交',
  'form.reset': '重置',
  'form.cancel': '取消',
  'form.search': '查询',
  'form.collapse': '收缩',
  'form.expand': '展开',
  'form.ok': '确认',
  'form.confirmToRemove': '确认删除？',

  'form.all': '所有',
  'form.name': '名称',
  'form.status': '状态',

  'status.enable': '启用',
  'status.disable': '禁用',

  'msg.warn': '提醒',
  'msg.canNotDisableDefaultProject': '不能禁用默认的项目。',
  'msg.canNotDeleteDefaultProject': '不能删除默认的项目。'
}

export default {
  ...components,
  ...locale
}
