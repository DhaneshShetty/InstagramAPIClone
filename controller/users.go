package controller

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-env-ways/models"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() (*mongo.Client, context.Context) {
	godotenv.Load(".env")
	value := os.Getenv("ATLAS_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(value))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}

//New User
func NewUser(w http.ResponseWriter, r *http.Request) {
	var client, ctx = connectDatabase()
	if client == nil {
		return
	}
	defer client.Disconnect(ctx)
	if r.Method == http.MethodPost {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if !emailRegex.MatchString(user.Email) {
			http.Error(w, "Email", http.StatusBadRequest)
			return
		}
		fmt.Print(user.Password)
		regPass := regexp.MustCompile(`^[0-9A-Za-z]{8,20}$`)
		if !regPass.MatchString(user.Password) {
			http.Error(w, "Password should contain atleast 8 characters and max 20", http.StatusBadRequest)
			return
		}
		database := client.Database("InstaDatabase")
		users := database.Collection("users")
		h := sha1.New()
		h.Write([]byte(user.Password))
		sha1_hash := hex.EncodeToString(h.Sum(nil))
		user.Password = sha1_hash
		_, _err := users.InsertOne(ctx, user)
		if _err != nil {
			http.Error(w, _err.Error(), http.StatusForbidden)
			return
		} else {
			w.WriteHeader(http.StatusAccepted)
		}
	} else {
		http.Error(w, "Path doesnt exist", http.StatusBadRequest)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var client, ctx = connectDatabase()
	if client == nil {
		return
	}
	defer client.Disconnect(ctx)
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		usersCollection := client.Database("InstaDatabase").Collection("users")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
		var user models.User
		filter := bson.M{"_id": bson.M{"$eq": objID}}
		err = usersCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			http.Error(w, "UserID doesnt exist", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	} else {
		http.Error(w, "Path doesnt exist", http.StatusBadRequest)
	}
}
