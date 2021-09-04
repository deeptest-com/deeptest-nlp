package router

import (
	"fmt"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/utlai/utl/cmd/server/router/handler"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	bizCasbin "github.com/utlai/utl/internal/server/biz/casbin"
	"github.com/utlai/utl/internal/server/biz/jwt"
	"github.com/utlai/utl/internal/server/cfg"
	"github.com/utlai/utl/internal/server/repo"
	"github.com/utlai/utl/internal/server/service"
	"net/http"
)

type Router struct {
	api *iris.Application

	InitService   *serverService.InitService `inject:""`
	JwtService    *jwt.JwtService            `inject:""`
	TokenService  *jwt.TokenService          `inject:""`
	CasbinService *bizCasbin.CasbinService   `inject:""`

	AccountCtrl *handler.AccountCtrl `inject:""`
	FileCtrl    *handler.FileCtrl    `inject:""`

	PermCtrl *handler.PermCtrl `inject:""`
	RoleCtrl *handler.RoleCtrl `inject:""`
	UserCtrl *handler.UserCtrl `inject:""`

	RpcCtrl *handler.RpcCtrl `inject:""`

	ProjectCtrl        *handler.ProjectCtrl        `inject:""`
	NluTaskCtrl        *handler.NluTaskCtrl        `inject:""`
	NluIntentCtrl      *handler.NluIntentCtrl      `inject:""`
	NluRuleCtrl        *handler.NluRuleCtrl        `inject:""`
	NluSentCtrl        *handler.NluSentCtrl        `inject:""`
	NluSlotCtrl        *handler.NluSlotCtrl        `inject:""`
	NluLookupCtrl      *handler.NluLookupCtrl      `inject:""`
	NluLookupItemCtrl  *handler.NluLookupItemCtrl  `inject:""`
	NluSynonymCtrl     *handler.NluSynonymCtrl     `inject:""`
	NluSynonymItemCtrl *handler.NluSynonymItemCtrl `inject:""`
	NluRegexCtrl       *handler.NluRegexCtrl       `inject:""`
	NluRegexItemCtrl   *handler.NluRegexItemCtrl   `inject:""`
	NluDictCtrl        *handler.NluDictCtrl        `inject:""`

	NluRasaCtrl  *handler.NluRasaCtrl  `inject:""`
	NluParseCtrl *handler.NluParseCtrl `inject:""`

	ValidCtrl *handler.ValidCtrl `inject:""`
	WsCtrl    *handler.WsCtrl    `inject:""`

	TokenRepo *repo.TokenRepo `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{}
	router.api = app

	return router
}

func (r *Router) App() {
	iris.LimitRequestBodySize(serverConf.Inst.Options.UploadMaxSize)
	r.api.UseRouter(_httpUtils.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		//if serverConf.Inst.BinData {
		//	app.GetDetail("/", func(ctx iris.Context) { // 首页模块
		//		_ = ctx.View("index.html")
		//	})
		//}

		v1 := app.Party("/v1")
		{
			v1.PartyFunc("/rpc", func(party iris.Party) {
				party.Post("/request", r.RpcCtrl.Request).Name = "转发RPC请求"
			})
			v1.PartyFunc("/client", func(party iris.Party) {
				party.Post("/agent/register", r.RpcCtrl.Request).Name = "转发RPC请求"
			})

			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Post("/login", r.AccountCtrl.UserLogin)

				//登录验证
				admin.Use(r.JwtService.Serve, r.TokenService.Serve, r.CasbinService.Serve)

				admin.Post("/logout", r.AccountCtrl.UserLogout).Name = "退出"
				admin.Get("/expire", r.AccountCtrl.UserExpire).Name = "刷新Token"
				admin.Get("/profile", r.UserCtrl.GetProfile).Name = "个人信息"

				admin.PartyFunc("/projects", func(party iris.Party) {
					party.Get("/", r.ProjectCtrl.List).Name = "项目列表"
					party.Get("/{id:uint}", r.ProjectCtrl.Get).Name = "项目详情"
					party.Post("/", r.ProjectCtrl.Create).Name = "创建项目"
					party.Put("/{id:uint}", r.ProjectCtrl.Update).Name = "更新项目"
					party.Delete("/{id:uint}", r.ProjectCtrl.Delete).Name = "删除项目"

					party.Post("/{id:uint}/setDefault", r.ProjectCtrl.SetDefault).Name = "设置当前项目"
					party.Post("/{id:uint}/disable", r.ProjectCtrl.Disable).Name = "禁用/启动项目"

					party.Get("/listForSelect", r.ProjectCtrl.ListForSelect).Name = "查询下拉框用项目列表"
				})

				admin.PartyFunc("/rasa/compile/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluRasaCtrl.Compile).Name = "编译项目"
				})
				admin.PartyFunc("/rasa/training/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluRasaCtrl.Training).Name = "训练项目"
				})
				admin.PartyFunc("/rasa/start/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluRasaCtrl.Start).Name = "启动服务"
				})
				admin.PartyFunc("/rasa/stop/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluRasaCtrl.Stop).Name = "停止服务"
				})
				admin.PartyFunc("/rasa/reloadRes/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluRasaCtrl.ReloadRes).Name = "刷新资源"
				})
				admin.PartyFunc("/rasa/parse/{id:uint}", func(party iris.Party) {
					party.Post("/", r.NluParseCtrl.Parse).Name = "解析说法"
				})

				admin.PartyFunc("/tasks", func(party iris.Party) {
					party.Get("/", r.NluTaskCtrl.List).Name = "任务列表"
					party.Get("/{id:uint}", r.NluTaskCtrl.Get).Name = "任务详情"
					party.Post("/", r.NluTaskCtrl.Create).Name = "创建任务"
					party.Put("/{id:uint}", r.NluTaskCtrl.Update).Name = "更新任务"
					party.Delete("/{id:uint}", r.NluTaskCtrl.Delete).Name = "删除任务"
					party.Post("/{id:uint}/disable", r.NluTaskCtrl.Disable).Name = "禁用/启动任务"
				})

				admin.PartyFunc("/intents", func(party iris.Party) {
					party.Get("/", r.NluIntentCtrl.List).Name = "意图列表"
					party.Get("/{id:uint}", r.NluIntentCtrl.Get).Name = "意图详情"
					party.Post("/", r.NluIntentCtrl.Create).Name = "创建意图"
					party.Put("/{id:uint}", r.NluIntentCtrl.Update).Name = "更新意图"
					party.Delete("/{id:uint}", r.NluIntentCtrl.Delete).Name = "删除意图"
					party.Post("/{id:uint}/disable", r.NluIntentCtrl.Disable).Name = "禁用/启动意图"
					party.Post("/move", r.NluIntentCtrl.Move).Name = "移动意图"
				})
				admin.PartyFunc("/rules", func(party iris.Party) {
					party.Get("/", r.NluRuleCtrl.List).Name = "规则列表"
					party.Get("/{id:uint}", r.NluRuleCtrl.Get).Name = "规则详情"
					party.Post("/", r.NluRuleCtrl.Create).Name = "创建规则"
					party.Put("/{id:uint}", r.NluRuleCtrl.Update).Name = "更新规则"
					party.Delete("/{id:uint}", r.NluRuleCtrl.Delete).Name = "删除规则"
					party.Post("/{id:uint}/disable", r.NluRuleCtrl.Disable).Name = "禁用/启动规则"
				})
				admin.PartyFunc("/sents", func(party iris.Party) {
					party.Get("/", r.NluSentCtrl.List).Name = "说法列表"
					party.Get("/{id:uint}", r.NluSentCtrl.Get).Name = "说法详情"
					party.Post("/", r.NluSentCtrl.Create).Name = "创建说法"
					party.Put("/{id:uint}", r.NluSentCtrl.Update).Name = "更新说法"
					party.Delete("/{id:uint}", r.NluSentCtrl.Delete).Name = "删除说法"
					party.Post("/{id:uint}/disable", r.NluSentCtrl.Disable).Name = "禁用/启动句子"
				})
				admin.PartyFunc("/slots", func(party iris.Party) {
					party.Get("/", r.NluSlotCtrl.List).Name = "语义槽列表"
					party.Get("/{id:uint}", r.NluSlotCtrl.Get).Name = "语义槽详情"
					party.Post("/", r.NluSlotCtrl.Create).Name = "创建语义槽"
					party.Put("/{id:uint}", r.NluSlotCtrl.Update).Name = "更新语义槽"
					party.Delete("/{id:uint}", r.NluSlotCtrl.Delete).Name = "删除语义槽"
					party.Post("/{id:uint}/disable", r.NluSlotCtrl.Disable).Name = "禁用/启动语义槽"
				})

				admin.PartyFunc("/synonyms", func(party iris.Party) {
					party.Get("/", r.NluSynonymCtrl.List).Name = "同义词列表"
					party.Get("/{id:uint}", r.NluSynonymCtrl.Get).Name = "同义词详情"
					party.Post("/", r.NluSynonymCtrl.Create).Name = "创建同义词"
					party.Put("/{id:uint}", r.NluSynonymCtrl.Update).Name = "更新同义词"
					party.Delete("/{id:uint}", r.NluSynonymCtrl.Delete).Name = "删除同义词"
					party.Post("/{id:uint}/disable", r.NluSynonymCtrl.Disable).Name = "禁用/启动同义词"
				})
				admin.PartyFunc("/synonymItems", func(party iris.Party) {
					party.Get("/", r.NluSynonymItemCtrl.List).Name = "同义词项列表"
					party.Get("/{id:uint}", r.NluSynonymItemCtrl.Get).Name = "同义词项详情"
					party.Post("/", r.NluSynonymItemCtrl.Create).Name = "创建同义词项"
					party.Put("/{id:uint}", r.NluSynonymItemCtrl.Update).Name = "更新同义词项"
					party.Delete("/{id:uint}", r.NluSynonymItemCtrl.Delete).Name = "删除同义词项"
					party.Post("/{id:uint}/disable", r.NluSynonymItemCtrl.Disable).Name = "禁用/启动同义词项"
					party.Post("/batchRemove", r.NluSynonymItemCtrl.BatchRemove).Name = "批量删除同义词项"
				})

				admin.PartyFunc("/lookups", func(party iris.Party) {
					party.Get("/", r.NluLookupCtrl.List).Name = "同类词列表"
					party.Get("/{id:uint}", r.NluLookupCtrl.Get).Name = "同类词详情"
					party.Post("/", r.NluLookupCtrl.Create).Name = "创建同类词"
					party.Put("/{id:uint}", r.NluLookupCtrl.Update).Name = "更新同类词"
					party.Delete("/{id:uint}", r.NluLookupCtrl.Delete).Name = "删除同类词"
					party.Post("/{id:uint}/disable", r.NluLookupCtrl.Disable).Name = "禁用/启动同类词"
				})
				admin.PartyFunc("/lookupItems", func(party iris.Party) {
					party.Get("/", r.NluLookupItemCtrl.List).Name = "同类词项列表"
					party.Get("/{id:uint}", r.NluLookupItemCtrl.Get).Name = "同类词项详情"
					party.Post("/", r.NluLookupItemCtrl.Create).Name = "创建同类词项"
					party.Put("/{id:uint}", r.NluLookupItemCtrl.Update).Name = "更新同类词项"
					party.Delete("/{id:uint}", r.NluLookupItemCtrl.Delete).Name = "删除同类词项"
					party.Post("/{id:uint}/disable", r.NluLookupItemCtrl.Disable).Name = "禁用/启动同类词项"
					party.Post("/batchRemove", r.NluLookupItemCtrl.BatchRemove).Name = "批量删除同类词项"
				})

				admin.PartyFunc("/regexes", func(party iris.Party) {
					party.Get("/", r.NluRegexCtrl.List).Name = "正则表达式列表"
					party.Get("/{id:uint}", r.NluRegexCtrl.Get).Name = "正则表达式详情"
					party.Post("/", r.NluRegexCtrl.Create).Name = "创建正则表达式"
					party.Put("/{id:uint}", r.NluRegexCtrl.Update).Name = "更新正则表达式"
					party.Delete("/{id:uint}", r.NluRegexCtrl.Delete).Name = "删除正则表达式"
					party.Post("/{id:uint}/disable", r.NluRegexCtrl.Disable).Name = "禁用/启动正则表达式"
				})
				admin.PartyFunc("/regexItems", func(party iris.Party) {
					party.Get("/", r.NluRegexItemCtrl.List).Name = "正则表达式项列表"
					party.Get("/{id:uint}", r.NluRegexItemCtrl.Get).Name = "正则表达式项详情"
					party.Post("/", r.NluRegexItemCtrl.Create).Name = "创建正则表达式项"
					party.Put("/{id:uint}", r.NluRegexItemCtrl.Update).Name = "更新正则表达式项"
					party.Delete("/{id:uint}", r.NluRegexItemCtrl.Delete).Name = "删除正则表达式项"
					party.Post("/{id:uint}/disable", r.NluRegexItemCtrl.Disable).Name = "禁用/启动正则表达式项"
					party.Post("/batchRemove", r.NluRegexItemCtrl.BatchRemove).Name = "批量删除正则表达式项"
				})

				admin.PartyFunc("/dicts", func(party iris.Party) {
					party.Get("/", r.NluDictCtrl.List).Name = "词典列表"
				})
				admin.PartyFunc("/valid", func(party iris.Party) {
					party.Post("/", r.ValidCtrl.Valid).Name = "表单验证"
				})

				admin.PartyFunc("/users", func(party iris.Party) {
					party.Get("/", r.UserCtrl.GetAllUsers).Name = "用户列表"
					party.Get("/{id:uint}", r.UserCtrl.GetUser).Name = "用户详情"
					party.Post("/", r.UserCtrl.CreateUser).Name = "创建用户"
					party.Put("/{id:uint}", r.UserCtrl.UpdateUser).Name = "编辑用户"
					party.Delete("/{id:uint}", r.UserCtrl.DeleteUser).Name = "删除用户"
				})
				admin.PartyFunc("/roles", func(party iris.Party) {
					party.Get("/", r.RoleCtrl.GetAllRoles).Name = "角色列表"
					party.Get("/{id:uint}", r.RoleCtrl.GetRole).Name = "角色详情"
					party.Post("/", r.RoleCtrl.CreateRole).Name = "创建角色"
					party.Put("/{id:uint}", r.RoleCtrl.UpdateRole).Name = "编辑角色"
					party.Delete("/{id:uint}", r.RoleCtrl.DeleteRole).Name = "删除角色"
				})
				admin.PartyFunc("/permissions", func(party iris.Party) {
					party.Get("/", r.PermCtrl.GetAllPermissions).Name = "权限列表"
					party.Get("/{id:uint}", r.PermCtrl.GetPermission).Name = "权限详情"
					party.Post("/", r.PermCtrl.CreatePermission).Name = "创建权限"
					party.Put("/{id:uint}", r.PermCtrl.UpdatePermission).Name = "编辑权限"
					party.Delete("/{id:uint}", r.PermCtrl.DeletePermission).Name = "删除权限"
				})
			})
		}

		// enable websocket
		websocketAPI := r.api.Party("/api/v1/ws")
		m := mvc.New(websocketAPI)
		m.Register(
			&prefixedLogger{prefix: "DEV"},
		)
		m.HandleWebsocket(handler.NewWsCtrl())

		websocketServer := websocket.New(
			gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), m)
		websocketAPI.Get("/", websocket.Handler(websocketServer))
	}
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}
