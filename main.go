package main

import (
	_ "3dprinter-camera/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

