package main

import (
	"cleango/controller"
	router "cleango/http"
	"cleango/service"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
	"cleango/repository"
)

var(
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewChiRouter()//independent from http framework
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "server up")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVER(port)
}
