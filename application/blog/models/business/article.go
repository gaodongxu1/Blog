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
 *     Initial: 2018/01/15        Gao Dongxu
 */


 package models
 
  import (
	"fmt"
	_"fmt"
	 "time"
	 "github.com/astaxie/beego/orm"
 )
 type ArticleServce struct{}
 
 var AS *ArticleServce
 
 type Article struct {
	 Id       int         `orm:"column(Id)"`
	 Title    string      `orm:"column(Title)"`
	 Content  string      `orm:"column(Content)"`
	 Status   int         `orm:"column(Status)"`
	 PublishedTime  string `orm:"column(PublishedTime)"`
 }
 
 // 注册模型
 func init() {
	 orm.RegisterModel(new(Article))
 }
 // 添加文章
 func (insert *ArticleServce) Insert(title string, content string) error {
	 o := orm.NewOrm()
 
	 sql := "INSERT INTO Blog.article(title,content,PublishedTime,status)VALUES(?,?,?,0)"
	 values := []interface{}{title, content,time.Now().Format("2006-01-02 15:04:05")}
	 raw := o.Raw(sql, values)
	 _, err := raw.Exec()
	 return err
 }

 func(read *ArticleServce) Read(title string) error {
	o := orm.NewOrm()
	var a []Article
	sql :="SELECT * FROM Blog.article where Status = 1"
	_,err := o.Raw(sql).QueryRows(&a)
	fmt.Printf("asd %v\n",a)
	return err
 }