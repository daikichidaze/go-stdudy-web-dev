package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content`
	Author  string `json:"author`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// jsonData, err := ioutil.ReadAll(jsonFile)
	// if err != nil {
	// 	fmt.Println("Error reading JSON data:", err)
	// 	return
	// }

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		}
		fmt.Println(post)
	}
}
