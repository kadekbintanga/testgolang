package main

import(
	"fmt"
	"testgolang/config"
	"gorm.io/gorm"
	"testgolang/app/routers"
)

var(
	db *gorm.DB = config.ConnectDB()
)

func main(){
	fmt.Print("------------------- Init Project --------------------")
	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.InitRouter()
}