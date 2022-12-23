package main

import(
	"fmt"
	"testgolang/config"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.ConnectDB()
)

func main(){
	fmt.Print("------------------- Init Project --------------------")
	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
}