package controller

import (
	"cleango/entity"
	"cleango/errors"
	"cleango/service"
	"encoding/json"
	"net/http"
)

type newController struct{}

var (
	postService service.PostService
)

type PostController interface{
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

//NewPostController dependency injection service service.PostService
func NewPostController(service service.PostService)PostController{
	postService = service
	return &newController{}
}

func (*newController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "Application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*newController) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "Application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message:"Error unmarshaling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message:err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message:"Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
