/*
 * @Author: weihua hu
 * @Date: 2024-12-16 00:48:47
 * @LastEditTime: 2024-12-16 15:04:15
 * @LastEditors: weihua hu
 * @Description:
 */
package result

type Result struct {
	Code int         `json:"code"` // 编码：1成功，0和其它数字为失败
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 数据
}

// Success 返回成功结果
func Success(data ...interface{}) Result {
	if len(data) > 0 {
		return Result{
			Code: 1,
			Msg:  "success",
			Data: data[0],
		}
	}
	return Result{
		Code: 1,
		Msg:  "success",
		Data: nil,
	}
}

// Error 返回错误结果
func Error(msg string) Result {
	return Result{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
}
