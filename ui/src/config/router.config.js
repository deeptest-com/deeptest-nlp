// eslint-disable-next-line
import { UserLayout, BasicLayout, BlankLayout } from '@/layouts'
import { bxAnaalyse } from '@/core/icons'

const RouteView = {
  name: 'RouteView',
  render: (h) => h('router-view')
}

export const asyncRouterMap = [

  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/platform/project/list',
    children: [
      {
        path: '/platform',
        name: 'platform',
        component: RouteView,
        meta: { title: 'menu.platform', keepAlive: true, icon: bxAnaalyse, permission: ['nlu'] },
        hideChildrenInMenu: false,
        children: [
          {
            path: '/platform/dashboard',
            name: 'platform-dashboard',
            component: () => import('@/views/platform/Dashboard'),
            meta: { title: 'menu.platform.dashboard', keepAlive: true, permission: ['nlu'] }
          },
          {
            path: '/platform/project',
            name: 'project',
            redirect: '/platform/project/list',
            component: RouteView,
            meta: { title: 'menu.project', keepAlive: true, permission: ['project'] },
            hideChildrenInMenu: true,
            children: [
              {
                path: '/platform/project/list',
                name: 'project-list',
                component: () => import('@/views/platform/project/List'),
                meta: { title: 'menu.project.list', keepAlive: true, permission: ['project'] }
              },
              {
                path: '/platform/project/:id/edit',
                name: 'project-edit',
                component: () => import('@/views/platform/project/Edit'),
                meta: { title: 'menu.project.edit', keepAlive: true, permission: ['project'] }
              }
            ]
          }
        ]
      },
      {
        path: '/nlu',
        name: 'nlu',
        component: RouteView,
        meta: { title: 'menu.nlu', keepAlive: true, icon: bxAnaalyse, permission: [ 'nlu' ] },
        hideChildrenInMenu: false,
        children: [
          {
            path: '/nlu/intent',
            name: 'intent',
            redirect: '/nlu/intent/list',
            component: RouteView,
            meta: { title: 'menu.intent', keepAlive: true, permission: [ 'intent' ] },
            hideChildrenInMenu: true,
            children: [
              {
                path: '/nlu/intent/list',
                name: 'intent-list',
                component: () => import('@/views/nlu/intent/List'),
                meta: { title: 'menu.intent.list', keepAlive: true, permission: [ 'intent' ] }
              },
              {
                path: '/nlu/intent/edit',
                name: 'intent-edit',
                component: () => import('@/views/nlu/intent/Edit'),
                meta: { title: 'menu.intent.edit', keepAlive: true, permission: [ 'intent' ] }
              }
            ]
          },
          {
            path: '/nlu/synonym',
            name: 'synonym',
            redirect: '/nlu/synonym/list',
            component: RouteView,
            meta: { title: 'menu.synonym', keepAlive: true, permission: [ 'synonym' ] },
            hideChildrenInMenu: true,
            children: [
              {
                path: '/nlu/synonym/list',
                name: 'synonym-list',
                component: () => import('@/views/nlu/synonym/List'),
                meta: { title: 'menu.synonym.list', keepAlive: true, permission: [ 'synonym' ] }
              },
              {
                path: '/nlu/synonym/edit',
                name: 'synonym-edit',
                component: () => import('@/views/nlu/synonym/Edit'),
                meta: { title: 'menu.synonym.edit', keepAlive: true, permission: [ 'synonym' ] }
              }
            ]
          },
          {
            path: '/nlu/lookup',
            name: 'lookup',
            redirect: '/nlu/lookup/list',
            component: RouteView,
            meta: { title: 'menu.lookup', keepAlive: true, permission: [ 'lookup' ] },
            hideChildrenInMenu: true,
            children: [
              {
                path: '/nlu/lookup/list',
                name: 'lookup-list',
                component: () => import('@/views/nlu/lookup/List'),
                meta: { title: 'menu.intent.list', keepAlive: true, permission: [ 'lookup' ] }
              },
              {
                path: '/nlu/lookup/edit',
                name: 'lookup-edit',
                component: () => import('@/views/nlu/lookup/Edit'),
                meta: { title: 'menu.intent.edit', keepAlive: true, permission: [ 'lookup' ] }
              }
            ]
          }
        ]
      },
      {
        path: '/sys',
        name: 'sys',
        redirect: '/sys/settings',
        component: RouteView,
        meta: { title: 'menu.sys.admin', keepAlive: true, icon: bxAnaalyse, permission: ['sys'] },
        hideChildrenInMenu: false,
        children: [
          {
            path: '/sys/settings',
            name: 'sys-settings',
            component: () => import('@/views/sys/Settings'),
            meta: { title: 'menu.sys.settings', keepAlive: true, permission: ['sys'] }
          }
        ]
      },

      // dashboard
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/workplace',
        component: RouteView,
        meta: { title: 'menu.dashboard', icon: bxAnaalyse, permission: [ 'dashboard' ] },
        children: [
          {
            path: '/dashboard/analysis/:pageNo([1-9]\\d*)?',
            name: 'Analysis',
            component: () => import('@/views/dashboard/Analysis'),
            meta: { title: 'menu.dashboard.analysis', permission: [ 'dashboard' ] }
          },
          // 外部链接
          {
            path: 'https://www.baidu.com/',
            name: 'Monitor',
            meta: { title: 'menu.dashboard.monitor', target: '_blank' }
          },
          {
            path: '/dashboard/workplace',
            name: 'Workplace',
            component: () => import('@/views/dashboard/Workplace'),
            meta: { title: 'menu.dashboard.workplace', keepAlive: true, permission: [ 'dashboard' ] }
          }
        ]
      },

      // forms
      {
        path: '/form',
        redirect: '/form/base-form',
        component: RouteView,
        meta: { title: '表单页', icon: 'form', permission: [ 'form' ] },
        children: [
          {
            path: '/form/base-form',
            name: 'BaseForm',
            component: () => import('@/views/form/basicForm'),
            meta: { title: '基础表单', keepAlive: true, permission: [ 'form' ] }
          },
          {
            path: '/form/step-form',
            name: 'StepForm',
            component: () => import('@/views/form/stepForm/StepForm'),
            meta: { title: '分步表单', keepAlive: true, permission: [ 'form' ] }
          },
          {
            path: '/form/advanced-form',
            name: 'AdvanceForm',
            component: () => import('@/views/form/advancedForm/AdvancedForm'),
            meta: { title: '高级表单', keepAlive: true, permission: [ 'form' ] }
          }
        ]
      },

      // list
      {
        path: '/list',
        name: 'list',
        component: RouteView,
        redirect: '/list/table-list',
        meta: { title: '列表页', icon: 'table', permission: [ 'table' ] },
        children: [
          {
            path: '/list/table-list/:pageNo([1-9]\\d*)?',
            name: 'TableListWrapper',
            hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
            component: () => import('@/views/list/TableList'),
            meta: { title: '查询表格', keepAlive: true, permission: [ 'table' ] }
          },
          {
            path: '/list/basic-list',
            name: 'BasicList',
            component: () => import('@/views/list/BasicList'),
            meta: { title: '标准列表', keepAlive: true, permission: [ 'table' ] }
          },
          {
            path: '/list/card',
            name: 'CardList',
            component: () => import('@/views/list/CardList'),
            meta: { title: '卡片列表', keepAlive: true, permission: [ 'table' ] }
          },
          {
            path: '/list/search',
            name: 'SearchList',
            component: () => import('@/views/list/search/SearchLayout'),
            redirect: '/list/search/article',
            meta: { title: '搜索列表', keepAlive: true, permission: [ 'table' ] },
            children: [
              {
                path: '/list/search/article',
                name: 'SearchArticles',
                component: () => import('../views/list/search/Article'),
                meta: { title: '搜索列表（文章）', permission: [ 'table' ] }
              },
              {
                path: '/list/search/project',
                name: 'SearchProjects',
                component: () => import('../views/list/search/Projects'),
                meta: { title: '搜索列表（项目）', permission: [ 'table' ] }
              },
              {
                path: '/list/search/application',
                name: 'SearchApplications',
                component: () => import('../views/list/search/Applications'),
                meta: { title: '搜索列表（应用）', permission: [ 'table' ] }
              }
            ]
          }
        ]
      },

      // profile
      {
        path: '/profile',
        name: 'profile',
        component: RouteView,
        redirect: '/profile/basic',
        meta: { title: '详情页', icon: 'profile', permission: [ 'profile' ] },
        children: [
          {
            path: '/profile/basic',
            name: 'ProfileBasic',
            component: () => import('@/views/profile/basic'),
            meta: { title: '基础详情页', permission: [ 'profile' ] }
          },
          {
            path: '/profile/advanced',
            name: 'ProfileAdvanced',
            component: () => import('@/views/profile/advanced/Advanced'),
            meta: { title: '高级详情页', permission: [ 'profile' ] }
          }
        ]
      },

      // result
      {
        path: '/result',
        name: 'result',
        component: RouteView,
        redirect: '/result/success',
        meta: { title: '结果页', icon: 'check-circle-o', permission: [ 'result' ] },
        children: [
          {
            path: '/result/success',
            name: 'ResultSuccess',
            component: () => import(/* webpackChunkName: "result" */ '@/views/result/Success'),
            meta: { title: '成功', keepAlive: false, hiddenHeaderContent: true, permission: [ 'result' ] }
          },
          {
            path: '/result/fail',
            name: 'ResultFail',
            component: () => import(/* webpackChunkName: "result" */ '@/views/result/Error'),
            meta: { title: '失败', keepAlive: false, hiddenHeaderContent: true, permission: [ 'result' ] }
          }
        ]
      },

      // Exception
      {
        path: '/exception',
        name: 'exception',
        component: RouteView,
        redirect: '/exception/403',
        meta: { title: '异常页', icon: 'warning', permission: [ 'exception' ] },
        children: [
          {
            path: '/exception/403',
            name: 'Exception403',
            component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/403'),
            meta: { title: '403', permission: [ 'exception' ] }
          },
          {
            path: '/exception/404',
            name: 'Exception404',
            component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404'),
            meta: { title: '404', permission: [ 'exception' ] }
          },
          {
            path: '/exception/500',
            name: 'Exception500',
            component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/500'),
            meta: { title: '500', permission: [ 'exception' ] }
          }
        ]
      },

      // account
      {
        path: '/account',
        component: RouteView,
        redirect: '/account/center',
        name: 'account',
        meta: { title: '个人页', icon: 'user', keepAlive: true, permission: [ 'user' ] },
        children: [
          {
            path: '/account/center',
            name: 'center',
            component: () => import('@/views/account/center'),
            meta: { title: '个人中心', keepAlive: true, permission: [ 'user' ] }
          },
          {
            path: '/account/settings',
            name: 'settings',
            component: () => import('@/views/account/settings/Index'),
            meta: { title: '个人设置', hideHeader: true, permission: [ 'user' ] },
            redirect: '/account/settings/base',
            hideChildrenInMenu: true,
            children: [
              {
                path: '/account/settings/base',
                name: 'BaseSettings',
                component: () => import('@/views/account/settings/BaseSetting'),
                meta: { title: '基本设置', hidden: true, permission: [ 'user' ] }
              },
              {
                path: '/account/settings/security',
                name: 'SecuritySettings',
                component: () => import('@/views/account/settings/Security'),
                meta: { title: '安全设置', hidden: true, keepAlive: true, permission: [ 'user' ] }
              },
              {
                path: '/account/settings/custom',
                name: 'CustomSettings',
                component: () => import('@/views/account/settings/Custom'),
                meta: { title: '个性化设置', hidden: true, keepAlive: true, permission: [ 'user' ] }
              },
              {
                path: '/account/settings/binding',
                name: 'BindingSettings',
                component: () => import('@/views/account/settings/Binding'),
                meta: { title: '账户绑定', hidden: true, keepAlive: true, permission: [ 'user' ] }
              },
              {
                path: '/account/settings/notification',
                name: 'NotificationSettings',
                component: () => import('@/views/account/settings/Notification'),
                meta: { title: '新消息通知', hidden: true, keepAlive: true, permission: [ 'user' ] }
              }
            ]
          }
        ]
      }

      // other
      /*
      {
        path: '/other',
        name: 'otherPage',
        component: PageView,
        meta: { title: '其他组件', icon: 'slack', permission: [ 'dashboard' ] },
        redirect: '/other/icon-selector',
        children: [
          {
            path: '/other/icon-selector',
            name: 'TestIconSelect',
            component: () => import('@/views/other/IconSelectorView'),
            meta: { title: 'IconSelector', icon: 'tool', keepAlive: true, permission: [ 'dashboard' ] }
          },
          {
            path: '/other/list',
            component: RouteView,
            meta: { title: '业务布局', icon: 'layout', permission: [ 'support' ] },
            redirect: '/other/list/tree-list',
            children: [
              {
                path: '/other/list/tree-list',
                name: 'TreeList',
                component: () => import('@/views/other/TreeList'),
                meta: { title: '树目录表格', keepAlive: true }
              },
              {
                path: '/other/list/edit-table',
                name: 'EditList',
                component: () => import('@/views/other/TableInnerEditList'),
                meta: { title: '内联编辑表格', keepAlive: true }
              },
              {
                path: '/other/list/user-list',
                name: 'UserList',
                component: () => import('@/views/other/UserList'),
                meta: { title: '用户列表', keepAlive: true }
              },
              {
                path: '/other/list/role-list',
                name: 'RoleList',
                component: () => import('@/views/other/RoleList'),
                meta: { title: '角色列表', keepAlive: true }
              },
              {
                path: '/other/list/system-role',
                name: 'SystemRole',
                component: () => import('@/views/role/RoleList'),
                meta: { title: '角色列表2', keepAlive: true }
              },
              {
                path: '/other/list/permission-list',
                name: 'PermissionList',
                component: () => import('@/views/other/PermissionList'),
                meta: { title: '权限列表', keepAlive: true }
              }
            ]
          }
        ]
      }
      */
    ]
  },
  {
    path: '*', redirect: '/404', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  }

]
