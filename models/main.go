package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("UserProfile.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Profile{}, &Interest{}, &Family{},
		&Education{}, &Job{}, &Address{}, &Other{},
		&PartnerChoice{}, &PartMartSt{}, &EmailData{}, &VerificationData{}, &Friend{},
		&SentChat{}, &ReceivedChat{}, &Subscription{}, &Album{}, &Contact{},
		&FolderListing{}, &SuccessStory{}, &MediaCoverage{})

	jsonBody := map[string]string{"token": "12931e786a9da517f10c52880c5711eb", "userId": "9798"}
	jsonValue, err := json.Marshal(jsonBody)

	response, err := http.Post("https://www.iitiimshaadi.com/apis/my_profile.json", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	//Populating data in each table
	userTable(db, bodyBytes)
	profileTable(db, bodyBytes)
	familyTable(db, bodyBytes)
	educationTable(db, bodyBytes)
	jobTable(db, bodyBytes)
	addressTable(db, bodyBytes)
	otherTable(db, bodyBytes)
	partnerChoiceTable(db, bodyBytes)
	emailDataTable(db, bodyBytes)
	verificationDataTable(db, bodyBytes)

}
