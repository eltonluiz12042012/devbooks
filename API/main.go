package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API is running!!!")
	r := router.Gerar()

	Log.Fatal(http.ListenAndServe(":5000", r))
}
