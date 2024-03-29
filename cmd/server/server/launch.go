package server

import (
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/utlai/utl/cmd/server/router"
	_db "github.com/utlai/utl/internal/pkg/db"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	bizCasbin "github.com/utlai/utl/internal/server/biz/casbin"
	"github.com/utlai/utl/internal/server/biz/jwt"
	"github.com/utlai/utl/internal/server/biz/redis"
	"github.com/utlai/utl/internal/server/conf"
	serverCron "github.com/utlai/utl/internal/server/cron"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverRes "github.com/utlai/utl/res/server"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12/context"
)

func Launch() {
	serverConf.Init()
	_db.InitDB("server")

	irisServer := NewServer(nil)
	if irisServer == nil {
		panic("Http 初始化失败")
	}
	irisServer.App.Logger().SetLevel(serverConf.Inst.LogLevel)

	if _commonUtils.IsRelease() {
		irisServer.App.HandleDir("/",
			&assetfs.AssetFS{Asset: serverRes.Asset, AssetDir: serverRes.AssetDir, AssetInfo: serverRes.AssetInfo,
				Prefix: "ui/dist"}, iris.DirOptions{
				IndexName: "index.html",
				Compress:  false,
			})
	}

	router := router.NewRouter(irisServer.App)
	injectObj(router)
	router.InitService.Init()
	router.App()

	if serverConf.Inst.Redis.Enable {
		redisUtils.InitRedisCluster(serverConf.GetRedisUris(), serverConf.Inst.Redis.Pwd)
	}

	iris.RegisterOnInterrupt(func() {
		defer _db.GetInst().Close()
	})

	if _commonUtils.IsPortInUse(serverConf.Inst.Port) {
		panic(fmt.Sprintf("端口 %d 已被使用", serverConf.Inst.Port))
	}

	// start the service
	err := irisServer.Serve()
	if err != nil {
		panic(err)
	}
}

func injectObj(router *router.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},

		// repo

		// middleware
		&inject.Object{Value: bizCasbin.NewEnforcer()},
		&inject.Object{Value: jwt.NewJwtService()},

		// service
		&inject.Object{Value: serverService.NewSeeder()},
		&inject.Object{Value: serverCron.NewServerCron()},

		// controller

		// router
		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}

type Server struct {
	App       *iris.Application
	AssetFile http.FileSystem
}

func NewServer(assetFile http.FileSystem) *Server {
	app := iris.Default()
	return &Server{
		App:       app,
		AssetFile: assetFile,
	}
}

func (s *Server) Serve() error {
	if serverConf.Inst.Https {
		host := fmt.Sprintf("%s:%d", serverConf.Inst.Host, 443)
		if err := s.App.Run(iris.TLS(host, serverConf.Inst.CertPath, serverConf.Inst.CertKey)); err != nil {
			return err
		}
	} else {
		if err := s.App.Run(
			iris.Addr(fmt.Sprintf("%s:%d", serverConf.Inst.Host, serverConf.Inst.Port)),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			iris.WithTimeFormat(time.RFC3339),
		); err != nil {
			return err
		}
	}

	return nil
}

type PathName struct {
	Name   string
	Path   string
	Method string
}

// 获取路由信息
func (s *Server) GetRoutes() []*model.Permission {
	var rrs []*model.Permission
	names := getPathNames(s.App.GetRoutesReadOnly())
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("路由权限集合：%v", names))
		fmt.Println(fmt.Sprintf("Iris App ：%v", s.App))
	}
	for _, pathName := range names {
		if !isPermRoute(pathName.Name) {
			rr := &model.Permission{Name: pathName.Path, DisplayName: pathName.Name, Description: pathName.Name, Act: pathName.Method}
			rrs = append(rrs, rr)
		}
	}
	return rrs
}

func getPathNames(routeReadOnly []context.RouteReadOnly) []*PathName {
	var pns []*PathName
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("routeReadOnly：%v", routeReadOnly))
	}
	for _, s := range routeReadOnly {
		pn := &PathName{
			Name:   s.Name(),
			Path:   s.Path(),
			Method: s.Method(),
		}
		pns = append(pns, pn)
	}

	return pns
}

// 过滤非必要权限
func isPermRoute(name string) bool {
	exceptRouteName := []string{"OPTIONS", "GET", "POST", "HEAD", "PUT", "PATCH", "payload"}
	for _, er := range exceptRouteName {
		if strings.Contains(name, er) {
			return true
		}
	}
	return false
}
