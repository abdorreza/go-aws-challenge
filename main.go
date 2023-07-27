package main

import (
	"fmt"

	"github.com/abdorreza/go-aws-challenge/db"
)

func main() {
	db.GetOneData()
	fmt.Println("---------------------------------------------")
	db.GetAllData()

	//lambda.Start(getAllData)

}
