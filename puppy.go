// main.go
package main

import (
	"fmt"
	"os"

	"github.com/pclubiitk/puppy-love/db"
	"github.com/pclubiitk/puppy-love/router"
	"github.com/pclubiitk/puppy-love/utils"
	"github.com/pclubiitk/puppy-love/config"

	"github.com/kataras/iris"
)

func executeFirst(ctx *iris.Context) {
	fmt.Println(string(ctx.Path()[:]))
	ctx.Next()
}

func main() {
	config.CfgInit()

	sessionDb := db.RedisSession()

	mongoDb, error := db.MongoConnect()
	if error != nil {
		fmt.Print("[Error] Could not connect to MongoDB")
		fmt.Print("[Error] " + config.CfgMgoUrl)
		fmt.Print(os.Environ())
		os.Exit(1)
	}

	utils.Randinit()

	iris.UseSessionDB(sessionDb)
	iris.Config.Gzip = true
	iris.UseFunc(executeFirst)

	router.PuppyRoute(mongoDb)

	iris.Listen(config.CfgAddr)
}
