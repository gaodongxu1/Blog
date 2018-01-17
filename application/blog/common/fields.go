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
 package common
 
 const (
	 // Session
	 SessionUserID   = "userId"
	 SessionPhone    = "phone"
	 SessionUserName = "username"
 
	 // DB Name
	 DBUser = "malluser"
	 DBMall = "mall"
	 DBBuss = "business"
 
	 // DB Reserved Value
	 LimitInvalidID  = 0
	 LimitMinPassLen = 6
 
	 // Resp
	 RespKeyStatus = "status"
	 RespKeyType   = "type"
	 RespKeyData   = "data"
 
	 // user常用Key值
	 KeyPhone      = "phone"
	 KeyAvatar     = "avatar"
	 KeyCollection = "collection"
 
	 // ware常用Key值
	 KeyColor  = "color"
	 KeySize   = "size"
	 KeyImage  = "image"
	 KeyDetail = "detail"
 
	 // is_default
	 Undefault = true
	 Default   = false
 
	 // default number
	 DefaultNumber = 1
 
	 // order status  待付款0，待收货1，待评价2，完成交易3，申请退货4，交易关闭5，同意退货6，不同意退货7
	 WaitToPay      = 0
	 WaitToReceive  = 1
	 WaitToEvaluate = 2
	 CloseOrder     = 3
	 RefundWare     = 4
	 CancelOrder    = 5
	 AgreeRefund    = 6
	 DisagreeRefund = 7
 
	 // 激活状态 未激活=1，激活=0
	 Active   = true
	 InActive = false
 
	 // 商品状态 正常态0，下架商品1，活动商品2，推荐商品3
	 NormalWare    = 0
	 RemovedWare   = 1
	 ActivityWare  = 2
	 RecommendWare = 3
 
	 // 用户类别
	 PhoneUser  = true
	 WeChatUser = false
 
	 // 通用数字
	 Zero = 0
 
	 // 默认头像
	 DefaultAvatar = "www.google.com"
 )
 