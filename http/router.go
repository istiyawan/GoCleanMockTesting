package router

import "net/http"

type Router interface{
	GET( URI string, f func(resp http.ResponseWriter, req *http.Request) )
	POST( URI string, f func(resp http.ResponseWriter, req *http.Request) )
	SERVER(port string)
}




