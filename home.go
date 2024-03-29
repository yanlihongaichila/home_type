package main

import (
	"flag"
	"fmt"
	"github.com/BarnabyCharles/frame/app"
	config2 "github.com/BarnabyCharles/frame/config"
	"github.com/ghodss/yaml"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"homeRpc/home"
	"homeRpc/internal/config"
	"homeRpc/internal/pkg"
	"homeRpc/internal/server"
	"homeRpc/internal/svc"
	"log"
)

//var configFile = flag.String("f", "etc/home.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config

	//连接配置
	viper, err := config2.InitViper("etc", "./etc/home.yaml", "yaml")
	if err != nil {
		log.Println(err)
		return
	}

	err = app.Init(viper.GetString("Nacos.serverName"),
		viper.GetString("Nacos.group"),
		viper.GetString("Nacos.namespaceId"),
		viper.GetString("Nacos.host"),
		viper.GetInt("Nacos.port"), "mysql")
	if err != nil {
		log.Println(err)
		return
	}

	nacosConfig, err := config2.GetNacosConfig(viper.GetString("Nacos.serverName"), viper.GetString("Nacos.group"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(nacosConfig), &c)
	if err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	//创建表
	err = pkg.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		home.RegisterHomeServer(grpcServer, server.NewHomeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
