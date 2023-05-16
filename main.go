package main

import (
	"time"

	"github.com/z3orc/dynamic-rpc/internal/env"
	"github.com/z3orc/dynamic-rpc/internal/http/server"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = env.ListenerPort()

func main() {

	//ASCII-banner on launch
	util.Banner("Compass", env.Version, env.Build)
	time.Sleep(10000)

	//Starting HTTP Server
	server := server.New(port)
	server.Start()

}
