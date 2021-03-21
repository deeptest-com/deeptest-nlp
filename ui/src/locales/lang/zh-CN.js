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
  'menu.platform.project': '项目管理',
  'menu.project.list': '项目列表',
  'menu.sys.admin': '系统管理',
  'menu.sys.settings': '全局设置',

  'menu.dashboard': '仪表盘',
  'menu.dashboard.analysis': '分析页',
  'menu.dashboard.monitor': '监控页',
  'menu.dashboard.workplace': '工作台',

  'common.notify': '通知',

  'action.create': '创建',
  'action.edit': '编辑',
  'action.back': '返回',
  'form.save': '保存',
  'form.reset': '重置',
  'form.cancel': '取消'
}

export default {
  ...components,
  ...locale
}
