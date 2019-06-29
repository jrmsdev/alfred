package routers

import (
	"github.com/astaxie/beego"
	"github.com/jrmsdev/alfred/internal/server/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
