package main

import (
	"api/src/config"
	"api/src/router"
	//"encoding/base64"
	"fmt"
	"log"
	//"math/rand"
	"net/http"
)

//Func que gera a chave secret_key jwt
/*func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}*/

func main() {
	config.Carregar()
	fmt.Printf("API is running at port:%d", config.Porta)
	r := router.Gerar()

	fmt.Println(config.Secretkey)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
