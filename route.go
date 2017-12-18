package route

import "net/http"

const (
	//CONNECT CONNECT请求
	CONNECT = "CONNECT"
	//DELETE DELETE请求
	DELETE = "DELETE"
	//GET GET请求
	GET = "GET"
	//HEAD HEAD请求
	HEAD = "HEAD"
	//OPTIONS OPTIONS请求
	OPTIONS = "OPTIONS"
	//PATCH PATCH请求
	PATCH = "PATCH"
	//POST POST请求
	POST = "POST"
	//PUT PUT请求
	PUT = "PUT"
	//TRACE TRACE请求
	TRACE = "TRACE"
)

//Route 路由定义以及封装
type Route struct {
}

//HTTPServer http请求处理
func HTTPServer(w *http.ResponseWriter, r *http.Request) {

}
