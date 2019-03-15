package main

import (
	_ "github.com/3xxx/flowtest/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 在开发环境下，您可以使用以下指令来开启查询调试模式
	// orm.Debug = true
	beego.Run()
}
