package datamodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/gorm"
)

func FriendsDataTable(db *gorm.DB) {
	// jsonBody := map[string]string{"token": "d53f61981a50b6e5baad02eec136db6c", "userId": "83441"}
	// jsonValue, err := json.Marshal(jsonBody)

	// response, err := http.Post("https://www.iitiimshaadi.com/apis/all_friends_category.json", "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer response.Body.Close()
	// bodyBytes, _ := ioutil.ReadAll(response.Body)

	// var allFriends map[string]interface{}
	// json.Unmarshal(bodyBytes, &allFriends)
	// friends := allFriends["allFriend"]
	// fmt.Println(friends)


	// User Detail
	userFile := "my_profile.json"
	_, UserFileErr := os.Stat(userFile)
	if os.IsNotExist(UserFileErr) {
		fmt.Printf("file [%s] does not exist", userFile)
	}

	UserByteValue, _ := ioutil.ReadFile(userFile)
	var userData map[string]interface{}
	err := json.Unmarshal([]byte(UserByteValue), &userData)
	if err != nil {
		log.Fatal(err)
	}

	// Freind Detail
	friendFile := "all_friends_category.json"
	_, frinedFileErr := os.Stat(friendFile)
	if os.IsNotExist(frinedFileErr) {
		fmt.Printf("file [%s] does not exist", friendFile)
	}

	friendByteValue, _ := ioutil.ReadFile(friendFile)
	var friendList map[string]interface{}
	err = json.Unmarshal([]byte(friendByteValue), &friendList)
	if err != nil {
		log.Fatal(err)
	}

	allFriend := friendList["allFriend"].([]interface{})
	shorlistedFriendCount := friendList["shorlistedFriendCount"]
	allRequestFriendCount := friendList["allRequestFriendCount"]
	allInvitedFriendCount := friendList["allInvitedFriendCount"]

	var shortListed bool
	var requestFrnd bool
	var inviteedFrnd bool

	if int(shorlistedFriendCount.(float64)) > 0 {
		shortListed = true
	} else {
		shortListed = false
	}

	if int(allRequestFriendCount.(float64)) > 0 {

		requestFrnd = true
	} else {
		requestFrnd = false
	}

	if int(allInvitedFriendCount.(float64)) > 0 {
		inviteedFrnd = true
	} else {
		inviteedFrnd = false
	}

	for i := 0; i < len(allFriend); i++ {

		db.Create(&Friend{
			FriendId:     int(allFriend[i].(map[string]interface{})["id"].(float64)),
			ShortListed:  shortListed,
			RequestFrnd:  requestFrnd,
			InviteedFrnd: inviteedFrnd,
			UserID:       int(userData["emailData"].(map[string]interface{})["id"].(float64)),
		})
	}

}
