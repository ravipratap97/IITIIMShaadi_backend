package datamodels

import (
	"fmt"

	"gorm.io/gorm"
)

func DataFriends(db *gorm.DB) {
	// var users User

	// db.Where("username = ?", "tejaulli").Find(&users)
	// fmt.Println(users)

	var friends []Friend
	// db.Where("user_id = ?", "9798").Find(&friends)
	db.Find(&friends, "user_id = ?", "9798")

	fmt.Println(friends)
	// fmt.Println(friends)


}
