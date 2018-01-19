/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/01/17        Gao Dongxu
 */
 package admin
 
 import (
	"fmt"
	 "encoding/json"
 
	 "github.com/astaxie/beego"
	 "github.com/astaxie/beego/orm"
 
	 . "github.com/gaodongxu1/Blog/application/blog/common"
	  "github.com/gaodongxu1/Blog/application/blog/logger"
	 . "github.com/gaodongxu1/Blog/application/blog/models/admin"
 )
 
 type AdminController struct {
	 beego.Controller
 }
 
 // 添加管理员账号
 func (this *AdminController) Create() {
	 var admin Manager

	fmt.Println(12345)

	 err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)

	 fmt.Println("ssssss:", this.Ctx.Input.RequestBody, "error", err)
	 if err != nil {
		 logger.GlobalLogReporter.Error(err)//日志
 
		 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrInvalidParam}
	 } else {
		 fmt.Println("admin:", admin)
		 logger.GlobalLogReporter.Info("admin:", admin)
		 err := AdminServer.Create(admin)
 
		 if err != nil {
			 logger.GlobalLogReporter.Error(err)
 
			 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrMysqlQuery}
		 } else {
 
			 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrSucceed}
		 }
	 }
 
	 this.ServeJSON()
 }
 
 // 管理员登录
 func (this *AdminController) Login() {
	 var admin Manager
 
	 err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
 
	 if err != nil {
		 logger.GlobalLogReporter.Error(err)
 
		 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrInvalidParam}
	 } else {
		 flag, err := AdminServer.Login(admin.Name, admin.Pass)
 
		 if err != nil {
			 if err == orm.ErrNoRows {
				 logger.GlobalLogReporter.Debug("Invalid name!")
 
				 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrInvalidUser}
			 } else {
				 logger.GlobalLogReporter.Error(err)
 
				 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrMysqlQuery}
			 }
		 } else {
			 if !flag {
				 logger.GlobalLogReporter.Debug("Wrong Pass!")
 
				 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrWrongPass}
			 } else {
				 this.SetSession(SessionUserID, admin.Name)
 
				 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrSucceed}
			 }
		 }
	 }
 
	 this.ServeJSON()
 }
 
 // 管理员修改自己的密码
 func (this *AdminController) ChangePass() {
	 var admin struct {
		 Oldpass string
		 Newpass string
	 }
 
	 err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
 
	 if err != nil {
		 logger.GlobalLogReporter.Error(err)
 
		 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrInvalidParam}
	 } else {
		 name := this.GetSession(SessionUserName)
		 flag, err := AdminServer.Login(name.(string), admin.Oldpass)
 
		 if err != nil {
			 logger.GlobalLogReporter.Error(err)
 
			 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrMysqlQuery}
		 } else {
			 if !flag {
				 logger.GlobalLogReporter.Debug("Wrong Pass!")
 
				 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrWrongPass}
			 } else {
				 err := AdminServer.ChangePass(name.(string), admin.Newpass)
 
				 if err != nil {
					 logger.GlobalLogReporter.Error(err)
 
					 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrMysqlQuery}
				 } else {
 
					 this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrSucceed}
				 }
			 }
		 }
	 }
 
	 this.ServeJSON()
 }
 
 