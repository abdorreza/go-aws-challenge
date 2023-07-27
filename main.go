package main

import (
	"context"
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
	fmt.Fprintf(w, "<h1 style='color:blue;'>Welcome to the GetPage!</h1>")
	// params := strings.Split(r.URL.Path, "/")
	// fmt.Println(r.URL.Path)
	// fmt.Println(params)

	device, err := db.GetDevice(r.Context(), "/devices/id8")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
	}
	fmt.Println(device)

	fmt.Println("Endpoint Hit: GetPage")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the PostPage!\n\n")
	db.InsertDevice(context.TODO())
	fmt.Println("Endpoint Hit: PostPage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/devices", getPage)
	//http.HandleFunc("/api/devices/wrt", postPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	handleRequests()

	// db.GetDevice()
	// fmt.Println("---------------------------------------------")
	//db.WriteData()

	//lambda.Start(device.Get)

}
