/*
 * MIT License
 *
 * Copyright (c) 2018 Gao Dongxu.
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
 package common

 const (
	ErrSucceed       = 0  // 成功
	ErrInvalidParam  = 1  // 参数错误
	ErrReqtAticle    = 1  // 获取文章失败
	ErrInvalidPerm   = 2  // 权限错误
	ErrInvalidUser   = 3  // 用户不存在
	ErrInactiveUser  = 4  // 用户被禁用
	ErrOperationPass = 5  // 操作密码错误
	ErrUserExists    = 6  // 用户已存在
	ErrArticleExists = 7  // 每日一文已存在
	ErrNoArea        = 8  // 区域经理没有绑定区域
	ErrWrongCode     = 9  // 验证码错误
	ErrWrongMsg      = 10 // 消息发送失败
	ErrWrongPass     = 11 // 密码错误
	ErrPhoneExists   = 12 // 手机号已经注册
	ErrPhoneNumber   = 13 // 手机号码错误
	ErrNoConsignee   = 14 // 收件人未填写
	ErrNoAddress     = 15 // 收货地址未填写
	ErrNotMatch      = 16 // 请核对用户名和密码
	ErrInvalidPass   = 17 // 密码不符合格式
	ErrNoData        = 18 // 未查询到数据

	ErrSession       = 400 // Session 错误
	ErrLoginRequired = 401 // 未登录
	ErrActive        = 402 // 用户未激活

	ErrMysqlQuery = 500 // MySQL 错误

)
