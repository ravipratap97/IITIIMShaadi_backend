package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	//User registration table
	User struct {
		gorm.Model
		Username string `json:"username"`
		Fullname string `json:"name"`
		Email    string `json:"email"`
		//Password string    `json:"password"`
		Gender string `json:"gender"`
		Mobile string `json:"mobile_no"`
		//DOB    time.Time.Date() `json:"birth_date"`
	}
	//Profile of user
	Profile struct {
		ProfileFor string `json:"profile_created_for"`
		//Image
		Religion   string `json:"religion"`
		Caste      string `json:"caste"`
		MotherTng  string `json:"mother_tounge"`
		AltMobile  string `json:"alternate_no"`
		MaritalSts string `json:"marital_status"`
		Interest   string `json:"interest"`
		Diet       string `json:"diet"`
		Height     string `json:"height"`
		Drink      string `json:"drink"`
		Smoke      string `json:"smoke"`
		HealthIss  string `json:"health_issue"`
		AboutMe    string `json:"about_me"`
	}
	//Family details of user
	Family struct {
		FatherName string `json:"father_name"`
		FatherOccp string `json:"father_occupation"`
		MotherName string `json:"mother_name"`
		MotherOccp string `json:"mother_occupation"`
		Brother    string `json:"brother"`
		Sister     string `json:"sister"`
	}
	//Education details of user
	Education struct {
		HighestQual string `json:"highest_education"`
		PostGradClg string `json:"post_graduation_college"`
		PostGradYr  string `json:"post_graduation_year"`
		PostGrad    string `json:"post_graduation"`
		GradClg     string `json:"graduation_college"`
		GradYr      string `json:"graduation_year"`
		Grad        string `json:"graduation"`
		SchoolYr    string `json:"schooling_year"`
		School      string `json:"schooling"`
	}
	//Job details of user
	Job struct {
		Company  string `json:"name_of_company"`
		JobTitle string `json:"working_as"`
		JobLoc   string `json:"job_location"`
		LinkedIn string `json:"linked_in"`
		IncomeAn string `json:"annual_income"`
	}
	//Address details of user
	Address struct {
		PermAddr  string `json:"pemanent_address"`
		PermCity  string `json:"permanent_city"`
		permState string `json:"permanent_state"`
		PermCntry string `json:"permanent_country"`
		PermZipCd string `json:"permanent_zipcode"`
		CurrAddr  string `json:"current_address"`
		CurrCity  string `json:"current_city"`
		CurrState string `json:"current_state"`
		CurrCntry string `json:"current_country"`
		CurrZipCd string `json:"current_zipcode"`
	}
	//Other details regarding user
	Other struct {
		MPrivacy int8   `json:"mprivacy"`
		BrideChe string `json:"choice_of_bride"`
		GroomChe string `json:"choice_of_groom"`
		PrflCmpl string `json:'profile_complition"`
	}
	//User preferred PartnerChoice
	PartnerChoice struct {
		PartnerChe string   `json:"choice_of_partner"`
		PartnerEdu string   `json:"prefered_partner_education"`
		PartnerRlg string   `json:"prefered_partner_religion"`
		PartnerCst string   `json:"prefered_partner_caste"`
		PartnerCtr string   `json:"prefered_partner_country"`
		PartMinAge string   `json:"prefered_partner_min_age"`
		PartMaxAge string   `json:"prefered_partner_max_age"`
		PartHtMax  string   `json:"prefered_partner_height_max"`
		PartHtMin  string   `json:"prefered_partner_height_min"`
		PartMartSt []string `json:"prefered_partner_marital_status"`
	}
	//EmailData of user
	EmailData struct {
		EmailIdNo   int16     `json:"id"`
		UserNick    string    `json:"user_nickname"`
		EmailStTm   time.Time `json:"email_sent_time"`
		DocName     string    `json:"document_name"`
		DocVerified string    `json:"document_verified"`
		UserUrl     string    `json:"user_url"`
		//CreatedDate
		FbProfileId  string `json:"facebook_profileid"`
		FbStatus     string `json:"facebook_status"`
		GgProfileId  string `json:"google_profileid"`
		GgStatus     string `json:"google_status"`
		UserActKey   string `json:"user_activation_key"`
		AccStatus    string `json:"account_status"`
		SendMsg      string `json:"send_message"`
		SendReq      string `json:"send_request"`
		Shortlistd   string `json:"shortlisted"`
		Favourite    string `json:"favourate"`
		DispName     string `json:"display_name"`
		PrflEmlSent  string `json:"profile_email_sent"`
		EmailSentSt  string `json:"emailsentstatus"`
		EmailSentSt1 string `json:"emailsentstatus1"`
		DeactEmail   string `json:"deactivatemail"`
		Pemail       string `json:"pemail"`
	}
	//VerificationData for user
	VerificationData struct {
		EmailStatus   string `json:"emailStatus"`
		MobileStatus  string `json:"mobileStatus"`
		BiodataStatus string `json:"biodata_status"`
		IdProofStatus string `json:"identity_proof_verified"`
		//DocVerified  string 	`json:"document_verified"`
	}
)

var db *gorm.DB

func main() {
	db, err := gorm.Open(sqlite.Open("UserProfile.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	jsonBody := map[string]string{"token": "12931e786a9da517f10c52880c5711eb", "userId": "9798"}
	jsonValue, err := json.Marshal(jsonBody)

	response, err := http.Post("https://www.iitiimshaadi.com/apis/my_profile.json", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var userData map[string]interface{}
	json.Unmarshal(bodyBytes, &userData)
	user := userData["basicData"].(map[string]interface{})

	db.Create(&User{Username: user["username"].(string),
		Fullname: user["name"].(string),
		Email:    userData["emailData"].(map[string]interface{})["email"].(string),
		Gender:   user["gender"].(string),
		Mobile:   user["mobile_no"].(string),
		//DOB:      user["birth_data"].(),
	})

	//userTable(response)

}

/*
func userTable(response *http.Response) {
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var userData map[string]interface{}
	json.Unmarshal(bodyBytes, &userData)
	user := userData["basicData"].(map[string]interface{})

	db.Create(&User{Username: user["username"].(string),
		Fullname: user["name"].(string),
		Email:    userData["emailData"].(map[string]interface{})["email"].(string),
		Gender:   user["gender"].(string),
		Mobile:   user["mobile_no"].(string),
		//DOB:      user["birth_data"].(string),
	})
}
*/
