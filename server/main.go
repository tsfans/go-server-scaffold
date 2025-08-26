package main

import (
	"github.com/tsfans/go/framework"
	"github.com/tsfans/go/server/controller"
)

func main() {
	framework.LoadRoute(controller.InitProbeRoute)
	framework.LoadServerRoute(controller.InitAllServerRoute)
	framework.Run()
}
