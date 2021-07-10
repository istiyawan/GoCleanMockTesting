package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type chiRouter struct{}

var(
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router{
	return &chiRouter{}
}

func (*chiRouter) GET( URI string, f func(resp http.ResponseWriter, req *http.Request) ){
 chiDispatcher.Get(URI,f)
}
func (*chiRouter) POST( URI string, f func(resp http.ResponseWriter, req *http.Request) ){
	chiDispatcher.Post(URI,f)
}
func (*chiRouter) SERVER(port string){
	fmt.Printf("Chi Http server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
