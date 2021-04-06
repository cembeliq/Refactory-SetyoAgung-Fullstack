package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

var baseUrl = "https://jsonplaceholder.typicode.com"

func main() {
	var posts, err = fetchPosts()
	if err != nil {
		fmt.Print("Error!", err.Error())
		return
	}

	fmt.Println(posts)
}

func fetchPosts() ([]Post, error){
	var err error
	var client = &http.Client{}
	var data []Post

	request, err := http.NewRequest("GET", baseUrl+"/posts", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	
	err = json.NewDecoder(response.Body).Decode(&data)

	
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUsers() ([]User, error){
	var err error
	var client = &http.Client{}
	var data []User

	request, err := http.NewRequest("GET", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	
	err = json.NewDecoder(response.Body).Decode(&data)

	
	if err != nil {
		return nil, err
	}

	return data, nil
}

type Post struct {
	UserId	int `json:"userId"`
	Id 	int	`json:"id"`
	Title string	`json:"title"`
	Body string	`json:"body"`
}