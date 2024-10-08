package main

import (
	"flag"
	"fmt"
	conf2 "greet/config"
	"greet/internal/config"
	"greet/internal/handler"
	"greet/internal/svc"
	"greet/webs"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func init() {
	conf2.DBInit()
	conf2.LogInIt()
}

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)
	//todo:: 完善
	// grpc.GrpcTptodbInit()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	http.HandleFunc("/ws", webs.WebSocketHandler)
	log.Println("Server started on port 8080")
	go http.ListenAndServe(":8080", nil)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
