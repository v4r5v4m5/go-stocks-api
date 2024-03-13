package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/v4r5v4m5/go-stocks-api/router"
)

func main() {
	r := router.Router()

	fmt.Println("[+] Starting go server at port 6969...")

	log.Fatal(http.ListenAndServe(":6969", r))
}
