package main

import (
	"RPSBackend/model"
	"RPSBackend/utility"
	"fmt"
)

var user = []model.User{
	{Id: 1, UserName: "John", Score: 5},
	{Id: 2, UserName: "Wick", Score: 6},
}

func DbSetup() {

	utility.Db.Migrator().DropTable(&model.User{})
	utility.Db.Migrator().CreateTable(&model.User{})

	err := utility.Db.Create(&user)
	if err.Error != nil {
		panic(err.Error.Error())
	} else {
		fmt.Println("Dummy table created")
	}

}

func main() {
	utility.DbConnect()
	defer utility.DbClose()
	DbSetup()
}
