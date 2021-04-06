package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

var baseUrl = "https://jsonplaceholder.typicode.com"

func main() {

	posts, err := fetchPosts()

	if err != nil {
		fmt.Print("Error!", err.Error())
		return
	}

	users, err := fetchUsers()
	if err != nil {
		fmt.Print("Error!", err.Error())
		return
	}

	for _, post := range posts {
		for _, user := range users {
			if user.ID == post.UserID {
				post.User = user
			}
		}
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)


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
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`    
	Title  string `json:"title"` 
	Body   string `json:"body"`  
	User   User   `json:"user"`  
}

type User struct {
	ID       int64   `json:"id"`      
	Name     string  `json:"name"`    
	Username string  `json:"username"`
	Email    string  `json:"email"`   
	Address  Address `json:"address"` 
	Phone    string  `json:"phone"`   
	Website  string  `json:"website"` 
	Company  Company `json:"company"` 
}

type Address struct {
	Street  string `json:"street"` 
	Suite   string `json:"suite"`  
	City    string `json:"city"`   
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`    
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`       
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`         
}
