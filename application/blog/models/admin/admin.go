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
 package business
 
 import (
	 "github.com/astaxie/beego/orm"
 
	 . "github.com/gaodongxu1/Blog/application/blog/common"
	 . "github.com/gaodongxu1/Blog/application/blog/utility"
 )
 
 // 对外接口
 type AdminServiceProvider struct {
 }
 
 var AdminServer *AdminServiceProvider
 
 // 管理员
 type Manager struct {
	 ID   uint64   `orm:"column(id);pk"`
	 Name string   `orm:"column(name)"  json:"name"`
	 Pass string   `orm:"column(pass)"  json:"pass"`
 }
 
 // 添加管理员用户
 func (this *AdminServiceProvider) Create(user Manager) error {
	 o := orm.NewOrm()
 
	 // 哈希加密
	 hash, err := GenerateHash(user.Pass)
 
	 if err != nil {
		 return err
	 }
	 password := string(hash)
 
	 sql := "INSERT INTO Blog.admin(name,password) VALUES (?,?)"
	 values := []interface{}{user.Name, password}
	 raw := o.Raw(sql, values)
	 _, err = raw.Exec()
 
	 return err
 }
 
 // 管理员用户登录
 func (this *AdminServiceProvider) Login(name string, password string) (bool, error) {
	 o := orm.NewOrm()
	 var pass string
 
	 err := o.Raw("SELECT password FROM  Blog.admin WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&pass)
 
	 if err != nil { 
		 return false, err
	 } else if !CompareHash([]byte(pass), password) {
		 return false, nil
	 }
 
	 return true, nil
 }
 
 // 根据用户名修改密码
 func (this *AdminServiceProvider) ChangePass(name string, newpassword string) error {
	 o := orm.NewOrm()
 
	 hash, err := GenerateHash(newpassword)
 
	 if err != nil {
		 return err
	 } else {
		 password := string(hash)
 
		 sql := "UPDATE Blog.admin SET password=? WHERE name=? LIMIT 1"
		 values := []interface{}{password, name}
		 raw := o.Raw(sql, values)
		 _, err := raw.Exec()
 
		 return err
	 }
 }
 
 // 禁用用户账号
 func (this *AdminServiceProvider) Inactive(name string) error {
	 o := orm.NewOrm()
 
	 sql := "UPDATE malluser.user SET active=? WHERE name=? LIMIT 1"
	 values := []interface{}{InActive, name}
	 raw := o.Raw(sql, values)
	 _, err := raw.Exec()
 
	 return err
 }
 
 // 激活用户账号
 func (this *AdminServiceProvider) Active(name string) error {
	 o := orm.NewOrm()
 
	 sql := "UPDATE malluser.user SET active=? WHERE name=? LIMIT 1"
	 values := []interface{}{Active, name}
	 raw := o.Raw(sql, values)
	 _, err := raw.Exec()
 
	 return err
 }
 
 // 根据用户名查找用户是否存在
 func (this *AdminServiceProvider) IsUserExist(name string) error {
	 o := orm.NewOrm()
	 var id uint64
 
	 err := o.Raw("SELECT id FROM malluser.user WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&id)
 
	 return err
 }
 