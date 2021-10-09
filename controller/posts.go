package controller

import (
	"encoding/json"
	"go-env-ways/models"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	var client, ctx = connectDatabase()
	if client == nil {
		return
	}
	defer client.Disconnect(ctx)
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
		query := r.URL.Query()
		limit, e := strconv.Atoi(query.Get("limit"))
		page, ef := strconv.Atoi(query.Get("page"))
		if e != nil {
			limit = -1
		}
		if ef != nil {
			page = 0
		}
		offset := (page - 1) * limit
		database := client.Database("InstaDatabase")
		postsCollection := database.Collection("posts")
		usersCollection := database.Collection("users")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filterUID := bson.M{"_id": bson.M{"$eq": objID}}
		var user models.User
		uidErr := usersCollection.FindOne(ctx, filterUID).Decode(&user)
		if uidErr != nil {
			http.Error(w, "UID doesnt exist", http.StatusInternalServerError)
			return
		}
		var post []models.Post
		var skip int64 = int64(offset)
		var limt int64 = int64(limit)
		var opts options.FindOptions = options.FindOptions{}
		if limit > 0 && skip >= 0 {
			opts = options.FindOptions{
				Limit: &limt,
				Skip:  &skip,
			}
		}

		filter := bson.M{"uid": bson.M{"$eq": id}}
		cur, err := postsCollection.Find(ctx, filter, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if e := cur.All(ctx, &post); e != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	} else {
		http.Error(w, "Path doesnt exist", http.StatusBadRequest)
	}
}

//Get Particular Post
func GetPost(w http.ResponseWriter, r *http.Request) {
	var client, ctx = connectDatabase()
	if client == nil {
		return
	}
	defer client.Disconnect(ctx)
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/posts/")
		database := client.Database("InstaDatabase")
		postsCollection := database.Collection("posts")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
		var post models.Post
		filter := bson.M{"_id": bson.M{"$eq": objID}}
		err = postsCollection.FindOne(ctx, filter).Decode(&post)
		if err != nil {
			http.Error(w, "Post Doesnt Exist", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)

	} else {
		http.Error(w, "Path doesnt exist", http.StatusBadRequest)
	}
}

//New Post
func NewPost(w http.ResponseWriter, r *http.Request) {
	var client, ctx = connectDatabase()
	if client == nil { // this will throw an error

		return
	}
	defer client.Disconnect(ctx)
	if r.Method == http.MethodPost {
		var post models.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if (post.Caption == "") || (post.ImageUrl == "") {
			http.Error(w, "Caption,ImageUrl cant be empty", http.StatusBadRequest)
			return
		}
		database := client.Database("InstaDatabase")
		usersCollection := database.Collection("users")
		objID, err := primitive.ObjectIDFromHex(post.UID)
		if err != nil {
			return
		}
		filter := bson.M{"_id": bson.M{"$eq": objID}}
		var user models.User
		uidErr := usersCollection.FindOne(ctx, filter).Decode(&user)
		if uidErr != nil {
			http.Error(w, "UID doesnt exist", http.StatusInternalServerError)
			return
		}
		postsCollection := database.Collection("posts")
		_, _err := postsCollection.InsertOne(ctx, post)
		if _err != nil {
			http.Error(w, _err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusAccepted)
		}
	} else {
		http.Error(w, "Path doesnt exist", http.StatusBadRequest)
	}
}
