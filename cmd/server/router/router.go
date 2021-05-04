package router

import (
	"fmt"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/utlai/utl/cmd/server/router/handler"
	"github.com/utlai/utl/internal/server/biz/middleware"
	middlewareUtils "github.com/utlai/utl/internal/server/biz/middleware/misc"
	"github.com/utlai/utl/internal/server/cfg"
	"github.com/utlai/utl/internal/server/repo"
	"github.com/utlai/utl/internal/server/service"
	"net/http"
)

type Router struct {
	api *iris.Application

	InitService   *service.InitService      `inject:""`
	JwtService    *middleware.JwtService    `inject:""`
	TokenService  *middleware.TokenService  `inject:""`
	CasbinService *middleware.CasbinService `inject:""`

	AccountCtrl *handler.AccountCtrl `inject:""`
	FileCtrl    *handler.FileCtrl    `inject:""`

	InitCtrl *handler.InitCtrl `inject:""`
	PermCtrl *handler.PermCtrl `inject:""`
	RoleCtrl *handler.RoleCtrl `inject:""`
	UserCtrl *handler.UserCtrl `inject:""`

	RpcCtrl *handler.RpcCtrl `inject:""`

	ProjectCtrl    *handler.ProjectCtrl    `inject:""`
	NluTaskCtrl    *handler.NluTaskCtrl    `inject:""`
	NluIntentCtrl  *handler.NluIntentCtrl  `inject:""`
	NluSentCtrl    *handler.NluSentCtrl    `inject:""`
	NluSlotCtrl    *handler.NluSlotCtrl    `inject:""`
	NluLookupCtrl  *handler.NluLookupCtrl  `inject:""`
	NluSynonymCtrl *handler.NluSynonymCtrl `inject:""`
	NluRegexCtrl   *handler.NluRegexCtrl   `inject:""`
	NluDictCtrl    *handler.NluDictCtrl    `inject:""`

	WsCtrl *handler.WsCtrl `inject:""`

	TokenRepo *repo.TokenRepo `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{}
	router.api = app

	return router
}

func (r *Router) App() {
	iris.LimitRequestBodySize(serverConf.Config.Options.UploadMaxSize)
	r.api.UseRouter(middlewareUtils.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		if serverConf.Config.BinData {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		v1 := app.Party("/v1")
		{
			v1.PartyFunc("/rpc", func(party iris.Party) {
				party.Post("/request", r.RpcCtrl.Request).Name = "转发RPC请求"
			})

			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Get("/init", r.InitCtrl.InitData)
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
					party.Post("/{id:uint}/disable", r.NluIntentCtrl.Disable).Name = "禁用/启动项目"
				})
				admin.PartyFunc("/sents", func(party iris.Party) {
					party.Get("/", r.NluSentCtrl.List).Name = "句子列表"
					party.Get("/{id:uint}", r.NluSentCtrl.Get).Name = "句子详情"
					party.Post("/", r.NluSentCtrl.Create).Name = "创建句子"
					party.Put("/{id:uint}", r.NluSentCtrl.Update).Name = "更新句子"
					party.Delete("/{id:uint}", r.NluSentCtrl.Delete).Name = "删除句子"
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

				admin.PartyFunc("/lookups", func(party iris.Party) {
					party.Get("/", r.NluLookupCtrl.List).Name = "词表列表"
					party.Get("/{id:uint}", r.NluLookupCtrl.Get).Name = "词表详情"
					party.Post("/", r.NluLookupCtrl.Create).Name = "创建词表"
					party.Put("/{id:uint}", r.NluLookupCtrl.Update).Name = "更新词表"
					party.Delete("/{id:uint}", r.NluLookupCtrl.Delete).Name = "删除词表"
					party.Post("/{id:uint}/disable", r.NluLookupCtrl.Disable).Name = "禁用/启动词表"
				})
				admin.PartyFunc("/lookupItems", func(party iris.Party) {
					party.Get("/", r.NluLookupCtrl.List).Name = "词表项列表"
					party.Get("/{id:uint}", r.NluLookupCtrl.Get).Name = "词表项详情"
					party.Post("/", r.NluLookupCtrl.Create).Name = "创建项词表"
					party.Put("/{id:uint}", r.NluLookupCtrl.Update).Name = "更新词表项"
					party.Delete("/{id:uint}", r.NluLookupCtrl.Delete).Name = "删除词表项"
					party.Post("/{id:uint}/disable", r.NluLookupCtrl.Disable).Name = "禁用/启动词表项"
					party.Post("/batchRemove", r.NluLookupCtrl.BatchRemove).Name = "批量删除词表项"
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
					party.Get("/", r.NluSynonymCtrl.List).Name = "词表项列表"
					party.Get("/{id:uint}", r.NluSynonymCtrl.Get).Name = "词表项详情"
					party.Post("/", r.NluSynonymCtrl.Create).Name = "创建项词表"
					party.Put("/{id:uint}", r.NluSynonymCtrl.Update).Name = "更新词表项"
					party.Delete("/{id:uint}", r.NluSynonymCtrl.Delete).Name = "删除词表项"
					party.Post("/{id:uint}/disable", r.NluSynonymCtrl.Disable).Name = "禁用/启动词表项"
					party.Post("/batchRemove", r.NluSynonymCtrl.BatchRemove).Name = "批量删除词表项"
				})

				admin.PartyFunc("/dicts", func(party iris.Party) {
					party.Get("/", r.NluDictCtrl.List).Name = "词典列表"
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

		websocketAPI := r.api.Party("/api/v1/ws")
		m := mvc.New(websocketAPI)
		m.Register(
			&prefixedLogger{prefix: "DEV"},
		)
		m.HandleWebsocket(handler.NewWsCtrl())
		websocketServer := websocket.New(
			gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}),
			m)
		websocketAPI.Get("/", websocket.Handler(websocketServer))
	}
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}
