package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hola!")
	getRandonJoke("https://icanhazdadjoke.com/")
}

func getRandonJoke(url string) {
	response, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println(response)
}

// type Joke struct {
// 	ID     string `"json:id"`
// 	Joke   string `"json:joke"`
// 	Status int    `"json:status"`
// }
