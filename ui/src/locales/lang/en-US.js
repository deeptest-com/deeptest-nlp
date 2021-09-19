import antdEnUS from 'ant-design-vue/es/locale-provider/en_US'
import momentEU from 'moment/locale/eu'

const components = {
  antLocale: antdEnUS,
  momentName: 'eu',
  momentLocale: momentEU
}

const locale = {
  'message': '-',
  'menu.home': 'Home',

  'menu.platform': 'Platform',
  'menu.platform.dashboard': 'Dashboard',
  'menu.settings': 'Settings',
  'menu.settings.sys': 'System Settings',
  'menu.settings.account': 'Account Settings',

  'menu.nlu.full': 'Natural Language Understanding',
  'menu.nlu': 'NLU',
  'menu.project': 'Project',
  'menu.project.list': 'Project List',
  'menu.project.view': 'Project View',
  'menu.project.test': 'Project Test',
  'menu.project.edit': 'Project Edit',
  'menu.task': 'Task',
  'menu.task.list': 'Task List',
  'menu.task.edit': 'Task Edit',
  'menu.task.design': 'Task Design',
  'menu.intent': 'Intent',
  'menu.intent.list': 'Intent List',
  'menu.intent.edit': 'Intent Edit',
  'menu.sent': 'Sentence',
  'menu.sent.list': 'Sentence List',
  'menu.sent.edit': 'Sentence Edit',
  'menu.slot': 'Slot',
  'menu.slot.list': 'Slot List',
  'menu.slot.edit': 'Slot Edit',

  'menu.placeholder': 'Placeholder',
  'menu.placeholder.list': 'Placeholder List',
  'menu.placeholder.edit': 'Placeholder Edit',

  'menu.synonym': 'Synonym',
  'menu.synonym.list': 'Synonym List',
  'menu.synonym.edit': 'Synonym Edit',
  'menu.synonym.items': 'Synonym Items',
  'menu.synonym.create.item': 'Create Synonym',
  'menu.synonym.edit.item': 'Edit Synonym',
  'menu.lookup': 'Lookup',
  'menu.lookup.list': 'Lookup List',
  'menu.lookup.edit': 'Lookup Edit',
  'menu.lookup.items': 'Lookup Items',
  'menu.lookup.create.item': 'Crete Lookup Item',
  'menu.lookup.edit.item': 'Edit Lookup Item',
  'menu.regex': 'Regex',
  'menu.regex.list': 'Regex List',
  'menu.regex.edit': 'Regex Edit',
  'menu.regex.items': 'Regex Items',
  'menu.regex.create.item': 'Crete Regex Item',
  'menu.regex.edit.item': 'Edit Regex Item',
  'menu.create.intent': 'Create Intent',
  'menu.enable.intent': 'Enable Intent',
  'menu.disable.intent': 'Disable Intent',
  'menu.remove.intent': 'Remove Intent',
  'menu.intent.new': 'New Intent',

  'form.view': 'View',
  'form.create': 'Create',
  'form.edit': 'Edit',
  'form.list': 'List',
  'form.design': 'Design',
  'form.maintain': 'Maintain',
  'form.remove': 'Delete',
  'form.disable': 'Disable',
  'form.enable': 'Enable',
  'form.compile': 'Compile',
  'form.training': 'Training',
  'form.test': 'Test',
  'form.back': 'Back',
  'form.save': 'Save',
  'form.submit': 'Submit',
  'form.send': 'Send',
  'form.reset': 'Reset',
  'form.confirm': 'Confirm',
  'form.cancel': 'Cancel',
  'form.search': 'Search',
  'form.collapse': 'Collapse',
  'form.expand': 'Expand',

  'form.pls.select': 'Please Select',
  'form.ok': 'Ok',
  'form.confirm.to.remove': 'Confirm to delete?',

  'form.all': 'All',
  'form.no': 'NO',
  'form.code': 'Code',
  'form.name': 'Name',
  'form.placeholder': 'Placeholder',
  'form.path': 'Path',
  'form.content': 'Content',
  'form.status': 'Status',
  'form.desc': 'Description',
  'form.createdBy': 'Created By',
  'form.createdAt': 'Created Time',
  'form.updatedAt': 'Updated Time',
  'form.set.default': 'Set Default',
  'form.is.default': 'Is Default',
  'form.opt': 'Operation',
  'form.opt.log': 'Operation Log',

  'form.slot.type': 'Slot Type',
  'form.regex': 'Regular Expression',
  'form.synonym': 'Synonym',
  'form.lookup': 'Lookup',
  'form.slot': 'Slot',
  'form.mark': 'Mark',
  'form.select.to.mark': 'Please select text to mark, the original slots in selected area will be removed.',
  'form.input.sent': 'Please input your question',
  'form.nav.history': 'Use ↑ ↓ to navigate through history.',
  'form.maintain.nlu.sent': 'Sentence',
  'form.maintain.nlu.rule': 'Rule',

  'form.exec.selenium.on.agent': 'Exec command on agent',

  'status.enable': 'Enable',
  'status.disable': 'Disable',

  'valid.required.code': 'Please input code',
  'valid.format.code': 'Code must start with a letter, and contains only letters and numbers.',
  'valid.required.name': 'Please input name.',
  'valid.required.project': 'Please select project.',
  'valid.required.content': 'Please input content.',
  'valid.required.path': 'Path not exist.',
  'valid.select.dict': 'Please select dictionary.',
  'valid.slot.type': 'Please select the type.',
  'valid.project.path': 'Not a valid project path.',
  'valid.dict.code.unique': 'Code must be unique.',

  'common.status': '状态',
  'common.login': 'Login',
  'common.logout': 'Logout',
  'common.info': 'Info',
  'common.tips': 'Tips',
  'common.confirm': 'Confirmation',
  'common.notify': 'Notification',
  'common.create': 'Create',
  'common.back': 'Back',
  'common.training': 'Training',
  'common.test': 'Test',
  'common.compile': 'Compile',
  'common.start_service': 'Start Service',
  'common.stop_service': 'Stop Service',
  'common.reload.pattern': 'Reload Resource',
  'common.start_training': 'Training Start',
  'common.end_training': 'Training End',
  'common.view.result': 'View Result',
  'common.view.json': 'View JSON',
  'common.view.nothing': 'View Nothing',
  'common.nlu.confidence': 'Confidence',
  'common.nlu.slots': 'Slot',
  'common.nlu.exec.result': 'Exec Result',

  'status.start.training': 'Training Start',
  'status.end.training': 'Training End',
  'status.start.service': 'Service Start',
  'status.stop.service': 'Service End',
  'msg.warn': 'Warning',
  'msg.confirm.to.logout': 'Do you really log-out.',
  'msg.forbidden': 'Forbidden',
  'msg.unauthorized': 'Unauthorized',
  'msg.auth.fail': 'Authorization verification failed',
  'msg.compile.success': 'Compile successfully.',
  'msg.training.start': 'Training started.',
  'msg.service.start': 'Service started.',
  'msg.service.stop': 'Service stopped.',
  'msg.service.not.start': 'Service not started',
  'msg.rasa.request.failed': 'Rasa request failed, response %s.',

  'msg.testing.welcome': 'Can I help you?',
  'msg.testing.pardon': 'Could you please say it again?',

  'app.setting.pagestyle': 'Page style setting',
  'app.setting.pagestyle.light': 'Light style',
  'app.setting.pagestyle.dark': 'Dark style',
  'app.setting.pagestyle.realdark': 'RealDark style',
  'app.setting.themecolor': 'Theme Color',
  'app.setting.navigationmode': 'Navigation Mode',
  'app.setting.content-width': 'Content Width',
  'app.setting.fixedheader': 'Fixed Header',
  'app.setting.fixedsidebar': 'Fixed Sidebar',
  'app.setting.sidemenu': 'Side Menu Layout',
  'app.setting.topmenu': 'Top Menu Layout',
  'app.setting.content-width.fixed': 'Fixed',
  'app.setting.content-width.fluid': 'Fluid',
  'app.setting.othersettings': 'Other Settings',
  'app.setting.weakmode': 'Weak Mode',
  'app.setting.copy': 'Copy Setting',
  'app.setting.loading': 'Loading theme',
  'app.setting.copyinfo': 'copy success，please replace defaultSettings in src/models/setting.js',
  'app.setting.production.hint': 'Setting panel shows in development environment only, please manually modify'
}

export default {
  ...components,
  ...locale
}
