package main

import (
	"fmt"
	"go-env-ways/controller"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello This is Dhanesh Shetty's Insta Clone API Using Go.Find Dhanesh Shetty on LinkedIn:https://www.linkedin.com/in/dhanesh-shetty/")
}

func handleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/users", controller.NewUser)
	http.HandleFunc("/posts", controller.NewPost)
	http.HandleFunc("/posts/", controller.GetPost)
	http.HandleFunc("/users/", controller.GetUser)
	http.HandleFunc("/posts/users/", controller.GetUserPosts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
