package filters

import (
	"github.com/astaxie/beego/context"

	. "github.com/gaodongxu1/Blog/application/blog/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(SessionUserID)

		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{RespKeyStatus: ErrLoginRequired}, false, false)
		}
	}
}
