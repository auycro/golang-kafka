package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	connection_string := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		"root",
		"example",
		"mysql",
		"3306",
		"micro",
	)

	dial := mysql.Open(connection_string)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", db.Config.Logger)

	//userRepo := domain.NewUserRepository(db)
	//users, _ := userRepo.FindAll()
	//for _, v := range users {
	//	fmt.Printf("%#v,%#v\n", v.FirstName, v.LastName)
	//}
}
