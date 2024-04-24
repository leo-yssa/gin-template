package main

import (
	"fmt"

	"github.com/leo-yssa/gin-template/backend/configs"
	"github.com/leo-yssa/gin-template/backend/routers"
)

func main() {
	configs.InitConfig()
	r := routers.NewRouter()
	port := configs.Config.Server.Port
	r.Run(fmt.Sprintf(":%s", port));
}