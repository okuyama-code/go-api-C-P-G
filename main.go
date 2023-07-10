package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Post struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

var users []User
var posts []Post

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// ユーザーIDを生成
	user.ID = uuid.New().String()

	// ユーザーを追加
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	// 投稿IDを生成
	post.ID = uuid.New().String()

	// 投稿を追加
	posts = append(posts, post)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	// URLパラメータから投稿IDを取得する
	params := mux.Vars(r)
	postID := params["id"]

	// 投稿のスライスをループして、指定された投稿IDと一致する投稿を探す
	for _, post := range posts {
		if post.ID == postID {
			// 一致する投稿が見つかった場合、その投稿をJSON形式でレスポンスとして返す
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	// 一致する投稿が見つからなかった場合、404エラーを返す
	w.WriteHeader(http.StatusNotFound)
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}", GetPostHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
