package api

import (
	"github.com/gaodongxu1/Blog/application/blog/controllers/business"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/insert", &business.ArticleControllers{}, "post:Insert")
	beego.Router("/read", &business.ArticleControllers{}, "post:ReadTitle")
	beego.Router("/update", &business.ArticleControllers{}, "post:Update")
}
