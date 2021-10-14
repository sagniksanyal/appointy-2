package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"crypto/md5"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID       string `bson:"Id"`
	Username string `bson:"Name"`
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

type Post struct {
	PID     string `bson:"Post_id"`
	ID      string `bson:"Id"`
	Caption string `bson:"Caption"`
	Image   string `bson:"Image_URL"`
}

func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //Parse url parameters passed, then parse the response packet for the POST body (request body)
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "This is the home page!") // write data to response
}

func generateMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("email:", r.Form["email"])
		fmt.Println("password:", r.Form["password"])
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mishratest:12345678M@mishradb.nxo5m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database("myStorageDB").Collection("users_collection")
		pass := generateMD5Hash(r.Form["password"][0])
		docs := []interface{}{
			bson.D{{"Id", r.Form["id"][0]}, {"Name", r.Form["username"][0]}, {"Email", r.Form["email"][0]}, {"Password", pass}},
		}
		res, insertErr := collection.InsertMany(ctx, docs)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)

		cur, currErr := collection.Find(ctx, bson.D{})
		if currErr != nil {
			panic(currErr)
		}
		defer cur.Close(ctx)

		var users []User
		if err = cur.All(ctx, &users); err != nil {
			panic(err)
		}
		fmt.Println(users)
		fmt.Fprint(w, "Succesfully sent data to mongoDB")
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DF")
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	x := r.URL.Path
	x = x[7:]
	fmt.Println(x)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mishratest:12345678M@mishradb.nxo5m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("myStorageDB").Collection("users_collection")

	filter := bson.D{{"Id", x}}
	var result User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, result)

}

func newpost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("post.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("Post ID:", r.Form["pid"])
		fmt.Println("User id:", r.Form["id"])
		fmt.Println("caption:", r.Form["caption"])
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mishratest:12345678M@mishradb.nxo5m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database("myStorageDB").Collection("posts_collection")
		docs := bson.D{{"Post_id", r.Form["pid"][0]}, {"Id", r.Form["id"][0]}, {"Caption", r.Form["caption"][0]}, {"Image_URL", r.Form["url"][0]}}
		res, insertErr := collection.InsertOne(ctx, docs)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)

		cur, currErr := collection.Find(ctx, bson.D{})
		if currErr != nil {
			panic(currErr)
		}
		defer cur.Close(ctx)

		var posts []Post
		if err = cur.All(ctx, &posts); err != nil {
			panic(err)
		}
		fmt.Println(posts)
		fmt.Fprint(w, "Succesfully sent data to mongoDB")
	}
}

func postbyid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DF")
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	x := r.URL.Path
	x = x[7:]
	fmt.Println(x)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mishratest:12345678M@mishradb.nxo5m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("myStorageDB").Collection("posts_collection")

	filter := bson.D{{"Post_id", x}}
	var result Post
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, result)
}

func postallbyid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DF")
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	x := r.URL.Path
	x = x[13:]
	fmt.Println(x)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mishratest:12345678M@mishradb.nxo5m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("myStorageDB").Collection("posts_collection")

	filter := bson.D{{"Id", x}}
	findOptions := options.Find()
	var results []Post

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	fmt.Fprintln(w, results)
}

func main() {
	http.HandleFunc("/", home) // setting router rule
	http.HandleFunc("/users", login)
	http.HandleFunc("/users/", users)
	http.HandleFunc("/posts", newpost)
	http.HandleFunc("/posts/", postbyid)
	http.HandleFunc("/posts/users/", postallbyid)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
