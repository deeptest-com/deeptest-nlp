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

  'menu.platform': '平台',
  'menu.platform.dashboard': '仪表盘',
  'menu.settings': '设置',
  'menu.settings.sys': '系统设置',
  'menu.settings.account': '账号设置',

  'menu.nlu.full': '自然语言理解',
  'menu.nlu': 'NLU',
  'menu.project': '项目',
  'menu.project.list': '项目列表',
  'menu.project.view': '项目查看',
  'menu.project.test': '项目测试',
  'menu.project.edit': '项目编辑',
  'menu.task': '任务',
  'menu.task.list': '任务列表',
  'menu.task.edit': '任务编辑',
  'menu.task.design': '任务设计',
  'menu.intent': '意图',
  'menu.intent.list': '意图列表',
  'menu.intent.edit': '意图编辑',
  'menu.sent': '句子',
  'menu.sent.list': '句子列表',
  'menu.sent.edit': '句子编辑',
  'menu.slot': '语义槽',
  'menu.slot.list': '意语义槽列表',
  'menu.slot.edit': '语义槽编辑',
  'menu.synonym': '同义词',
  'menu.synonym.list': '同义词表',
  'menu.synonym.edit': '编辑同义词表',
  'menu.synonym.items': '同义词项',
  'menu.synonym.create.item': '创建同义词项',
  'menu.synonym.edit.item': '编辑同义词项',
  'menu.lookup': '同类词',
  'menu.lookup.list': '同类词表',
  'menu.lookup.edit': '编辑同类词表',
  'menu.lookup.items': '同类词项',
  'menu.lookup.create.item': '创建同类词项',
  'menu.lookup.edit.item': '编辑同类词项',
  'menu.regex': '正则表达式',
  'menu.regex.list': '正则表达式列表',
  'menu.regex.edit': '编辑正则表达式',
  'menu.regex.items': '正则表达式项',
  'menu.regex.create.item': '创建正则表达式项',
  'menu.regex.edit.item': '编辑正则表达式项',

  'common.login': '登录',
  'common.logout': '登出',
  'form.view': '查看',
  'form.create': '新建',
  'form.edit': '编辑',
  'form.design': '设计',
  'form.maintain': '维护',
  'form.remove': '删除',
  'form.disable': '禁用',
  'form.enable': '启用',
  'form.compile': '编译',
  'form.training': '训练',
  'form.test': '测试',
  'form.back': '返回',
  'form.save': '保存',
  'form.submit': '提交',
  'form.send': '发送',
  'form.reset': '重置',
  'form.cancel': '取消',
  'form.search': '查询',
  'form.collapse': '收缩',
  'form.expand': '展开',
  'form.ok': '确认',
  'form.confirm.to.remove': '确认删除？',

  'form.all': '所有',
  'form.no': '序号',
  'form.code': '编码',
  'form.name': '名称',
  'form.path': '路径',
  'form.content': '内容',
  'form.status': '状态',
  'form.desc': '描述',
  'form.createdBy': '创建人',
  'form.createdAt': '创建时间',
  'form.updatedAt': '更新时间',
  'form.is.default': '是否默认',
  'form.opt': '操作',
  'form.opt.log': '操作日志',

  'form.sent.edit': '编辑说法',
  'form.sent.list': '说法列表',
  'form.slot.type': '插槽类型',
  'form.regex': '正则表达式',
  'form.synonym': '同义词',
  'form.lookup': '词表',
  'form.use.regex': '使用正则表达式',
  'form.use.synonym': '使用同义词',
  'form.use.lookup': '使用词表',
  'form.mark': '标记',
  'form.select.to.mark': '请选择文本进行标记，选中内容中包括的原有语义槽将被移除。',
  'form.input.sent': '请输入您想说的话',
  'form.nav.history': '按键盘的↑ ↓键，切换历史说法.',

  'status.enable': '启用',
  'status.disable': '禁用',

  'valid.required.code': '请输入编码',
  'valid.format.code': '编码必须以字母开头，且只可包含字母和数字。',
  'valid.required.name': '请输入名称',
  'valid.required.project': '请选择项目',
  'valid.required.path': '请输入路径',
  'valid.project.path': '不是一个合法的项目路径',
  'valid.required.content': '请输入内容',
  'valid.select.dict': '请选择词典。',

  'common.status': '状态',
  'common.info': '消息',
  'common.tips': '提示',
  'common.confirm': '确认',
  'common.notify': '通知',
  'common.create': '新建',
  'common.back': '返回',
  'common.training': '训练',
  'common.test': '测试',
  'common.compile': '编译',
  'common.start_service': '启动服务',
  'common.start_training': '训练开始',
  'common.end_training': '训练结束',
  'common.view.result': '显示结果',
  'common.view.json': '显示JSON',
  'common.view.nothing': '关闭显示',

  'msg.warn': '提醒',
  'msg.confirm.to.logout': '确认退出？',
  'msg.forbidden': '确认退出？',
  'msg.unauthorized': '未授权的',
  'msg.auth.fail': '授权失败',
  'msg.compile.success': '成功编译项目。',
  'msg.training.start': '成功发起训练。',
  'msg.service.start': '成功启动服务。',

  'msg.testing.welcome': '请问有什么可以帮您？',
  'msg.testing.pardon': '我听不懂您说了什么，请再说一遍。',

  'app.setting.pagestyle': '页面演示设置',
  'app.setting.pagestyle.light': '淡色',
  'app.setting.pagestyle.dark': '深色',
  'app.setting.pagestyle.realdark': '深黑',
  'app.setting.themecolor': '主题色彩',
  'app.setting.navigationmode': '导航模式',
  'app.setting.content-width': '内容宽度',
  'app.setting.fixedheader': '固定头部',
  'app.setting.fixedsidebar': '固定菜单栏',
  'app.setting.sidemenu': '左侧菜单栏布局',
  'app.setting.topmenu': '顶部菜单布局',
  'app.setting.content-width.fixed': '固定',
  'app.setting.content-width.fluid': '流式',
  'app.setting.othersettings': '其他设置',
  'app.setting.weakmode': '弱模式',
  'app.setting.copy': '复制设置',
  'app.setting.loading': '加载主题',
  'app.setting.copyinfo': '拷贝成功，请替换src/models/setting.js里的默认设置。',
  'app.setting.production.hint': '设置面板仅显示在开发模式中，请手工进行编辑。'
}

export default {
  ...components,
  ...locale
}
