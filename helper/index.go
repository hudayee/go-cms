package helper

func ToWeb(code int, msg string, data interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	m["data"] = data
	return
}
