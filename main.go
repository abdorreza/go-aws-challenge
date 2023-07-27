package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdorreza/go-aws-challenge/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GetPage!")
	db.ReadData()
	fmt.Println("Endpoint Hit: GetPage")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the PostPage!")
	db.ReadData()
	fmt.Println("Endpoint Hit: PostPage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/devices", getPage)
	http.HandleFunc("/api/devices/{id}", postPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	//handleRequests()

	db.ReadData()
	fmt.Println("---------------------------------------------")
	db.GetAllData()
	fmt.Println("---------------------------------------------")
	//db.WriteData()

	//lambda.Start(device.Get)

}
