package common

import "fmt"

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误:%s",
}


// GetMsg get error information based on Code
func GetMsg(code int, err ...interface{}) string {
	var msg = MsgFlags[ERROR]
	msg, ok := MsgFlags[code]
	if ok {
		msg = MsgFlags[code]
	}
	if len(err) > 0 {
		return fmt.Sprintf(msg, err...)
	}
	return msg
}