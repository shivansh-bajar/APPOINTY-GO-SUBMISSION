package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"usr-id"`
	Name     string `json:"name"`
	Email    string `json:"mail"`
	Password string `json:"pass"`
}
type Post struct {
	Id      string `json:"usr-id"`
	Name    string `json:"name"`
	Caption string `json:"cap"`
	Photo   string `json:"url"`
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "%v\n", r.PostForm)
		name := r.FormValue("name")
		email := r.FormValue("mail")
		password := r.FormValue("pass")

		md5HashInBytes := md5.Sum([]byte(password))
		md5HashInString := hex.EncodeToString(md5HashInBytes[:])

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", email)
		fmt.Fprintf(w, "Password = %s\n", password)

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://shivanshbajar:123@cluster0.helmj.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("insta").Collection("user")
		user := bson.D{{Key: "Name", Value: name},
			{Key: "Email", Value: email}, {Key: "Password", Value: md5HashInString}}

		res, insertErr := collection.InsertOne(ctx, user)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)
	} else {
		http.ServeFile(w, r, "userform.html")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		id := r.FormValue("id")

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://shivanshbajar:123@cluster0.helmj.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("insta").Collection("user")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("Invalid id")
		}

		var user bson.M
		if err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "USER: %s", user)
	}
}

func addPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "%v\n", r.PostForm)
		name := r.FormValue("name")
		caption := r.FormValue("cap")
		url := r.FormValue("url")
		timestamp := time.Now()

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Caption = %s\n", caption)
		fmt.Fprintf(w, "Photo Url = %s\n", url)
		fmt.Fprintf(w, "Timestamp = %s\n", timestamp)

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://shivanshbajar:123@cluster0.helmj.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("insta").Collection("posts")

		post := bson.D{{Key: "Name", Value: name},
			{Key: "Caption", Value: caption}, {Key: "Post URL", Value: url}, {Key: "TimeStamp", Value: timestamp}}

		res, insertErr := collection.InsertOne(ctx, post)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)
	} else {
		http.ServeFile(w, r, "postform.html")
	}
}
func getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		postid := r.FormValue("postid")

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://shivanshbajar:123@cluster0.helmj.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("insta").Collection("posts")

		objectId, err := primitive.ObjectIDFromHex(postid)
		if err != nil {
			log.Println("Invalid id")
		}

		var post bson.M
		if err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&post); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "POST: %s", post)

	}
}
func getPostList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		key := r.FormValue("id")
		//fmt.Fprintf(w, key)

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://shivanshbajar:123@cluster0.helmj.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("insta").Collection("posts")

		/*
			objectId, err := primitive.ObjectIDFromHex(key)
			if err != nil {
				log.Println("Invalid id")
			}
		*/
		postslist, err := collection.Find(ctx, bson.M{"Name": key})
		if err != nil {
			log.Fatal(err)
		}
		var listofpost []bson.M
		if err = postslist.All(ctx, &listofpost); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "All posts by user: %s", listofpost)

	}
}

func handleRequests() {

	http.HandleFunc("/users", addUsers)
	http.HandleFunc("/posts", addPosts)
	http.HandleFunc("/users/", getUser)
	http.HandleFunc("/posts/", getPost)
	http.HandleFunc("/posts/users/", getPostList)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequests()
}
