package jsonUtil

import "encoding/json"

func JsonResp(code int, data interface{}, errMsg string) string {
	resp := make(map[string]interface{})
	resp["status_code"] = code
	resp["error_msg"] = errMsg
	resp["data"] = data

	js, err := json.Marshal(resp)
	if err != nil {
		return err.Error()
	}
	return string(js)
}
