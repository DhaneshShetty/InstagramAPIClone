package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-env-ways/controller"
	"go-env-ways/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello This is Dhanesh Shetty's Insta Clone API Using Go.Find Dhanesh Shetty on LinkedIn:https://www.linkedin.com/in/dhanesh-shetty/`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/6161108f9621e76fea45394e", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	var user models.User
	er := json.NewDecoder(rr.Body).Decode(&user)
	if er != nil {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
	if user.OID != "6161108f9621e76fea45394e" {
		t.Errorf("handler returned wrong data: got %v",
			rr.Body.String())
	}

}

func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/61611d02852180d63f005091", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	var post models.Post
	er := json.NewDecoder(rr.Body).Decode(&post)
	if er != nil {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
	if post.ID != "61611d02852180d63f005091" {
		t.Errorf("handler returned wrong data: got %v",
			rr.Body.String())
	}

}

func TestNewPost(t *testing.T) {
	buf := new(bytes.Buffer)
	jsonReq := models.Post{
		Caption:  "Coding my way into Appointy",
		ImageUrl: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fcommons.wikimedia.or...",
		UID:      "6161108f9621e76fea45394e"}

	err := json.NewEncoder(buf).Encode(jsonReq)
	if err != nil {
		fmt.Println(err.Error())
	}
	req, err := http.NewRequest("POST", "/posts", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.NewPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func TestNewUser(t *testing.T) {
	buf := new(bytes.Buffer)
	jsonReq := models.User{
		Name:     "Dhanesh Shetty",
		Email:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fcommons.wikimedia.or...",
		Password: "meinKyunBatun"}

	err := json.NewEncoder(buf).Encode(jsonReq)
	if err != nil {
		fmt.Println(err.Error())
	}
	req, err := http.NewRequest("POST", "/users", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.NewUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

}

func TestUsersPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/6161108f9621e76fea45394e", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUserPosts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	var post []models.Post
	er := json.NewDecoder(rr.Body).Decode(&post)
	if er != nil {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
	if post[0].UID != "6161108f9621e76fea45394e" {
		t.Errorf("handler returned wrong data: got %v",
			rr.Body.String())
	}

}

func PasswordEmailError(t *testing.T) {
	buf := new(bytes.Buffer)
	jsonReq := models.User{
		Name:     "Dhanesh Shetty",
		Email:    "dhanesh@dhanesh.com",
		Password: ""}

	err := json.NewEncoder(buf).Encode(jsonReq)
	if err != nil {
		fmt.Println(err.Error())
	}
	req, err := http.NewRequest("POST", "/users", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.NewUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func EmailError(t *testing.T) {
	buf := new(bytes.Buffer)
	jsonReq := models.User{
		Name:     "Dhanesh Shetty",
		Email:    "dhanesh",
		Password: "correctPassword1*"}

	err := json.NewEncoder(buf).Encode(jsonReq)
	if err != nil {
		fmt.Println(err.Error())
	}
	req, err := http.NewRequest("POST", "/users", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.NewUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func UidError(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/6161108f9621e76fe", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUserPosts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
	expected := "UID doesnt exist"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
