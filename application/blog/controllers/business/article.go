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
 *     Initial: 2018/01/16        Gao Dongxu
 */
 package business
 
 import (
	"fmt"
	_"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/gaodongxu1/Blog/application/blog/common"
	"github.com/gaodongxu1/Blog/application/blog/models/business"
 )
 
 type ArticleControllers struct {
	 beego.Controller
 }
 
 // 添加类别
 func (this *ArticleControllers) Insert() {
	 var article struct {
		 Title    string
		 Content  string
	 }
 
	 err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
 
	 if err != nil {
		 this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	 } else {
		err = models.AS.Insert(article.Title, article.Content)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	 }
 
	 this.ServeJSON()
 }
 // 根据标题查找文章
 func(this *ArticleControllers) ReadTitle() {
	 var Messages struct {
		 Title string `json:"title"`
	 }
	 err := json.Unmarshal(this.Ctx.Input.RequestBody, &Messages)

	 // error错误信息，如果error为空说明没错，不为空说明有错。
	 if err != nil {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: "Title error"}
	} else {
		messages := models.AS.Read(Messages.Title)
		if err != nil {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyData: messages}
			fmt.Printf("显示结果: %v\n",Messages.Title)
		}
	}
	this.ServeJSON()
 }
 // 根据ID更改文章标题，内容，状态
 func(this *ArticleControllers) Update() {
	 var article struct {
		 Id        int      `json:"id"`
		 Title     string   `json:"title"`
		 Content   string   `json:"content"`
		 Status    int      `json:"status"`
	 }
	 err :=json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	 if err  !=nil {
		 this.Data["json"] = map[string]interface{}{common.RespKeyStatus: "param error"}
	 } else{
		err = models.AS.Update(article.Title, article.Content, article.Status,article.Id)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		 fmt.Printf("更新结果: %v\n",err)
	 }
	 this.ServeJSON()
 }
