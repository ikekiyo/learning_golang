package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{}
	for i := 0; i < 10; i++ {
		title := "Hello_%d"
		articles = append(
			articles,
			Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"},
		)
		fmt.Println("Endpoint Hit: returnAllArticles")
		json.NewEncoder(w).Encode(articles)
	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleRequest()
}
