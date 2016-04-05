package routers

import (
	"github.com/welemon/coveralls-test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
