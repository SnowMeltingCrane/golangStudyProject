package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "test",
		Year:   2019,
		Price:  9,
		Actors: []string{"Tom", "Jerry"},
	}

	// 编码过程 结构体->json
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	} else {
		fmt.Printf("json: %s\n", jsonStr)
	}

	// 解码过程 json->结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json.Unmarshal failed:", err)
		return
	} else {
		fmt.Printf("myMovie: %v\n", myMovie)
	}
}
