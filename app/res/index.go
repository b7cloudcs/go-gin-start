package res

// json2go
// https://www.sojson.com/json/json2go.html

// ErrorMsg
type ErrorMsg struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

// Json Server Req GetList
type JSReqGetList struct {
	Start int    `form:"_start"`
	Limit int    `form:"_limit"`
	Order string `form:"_order"`
	Sort  string `form:"_sort"`
	Q     string `form:"q"`
}
