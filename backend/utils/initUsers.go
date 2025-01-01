package utils

import (
	"fmt"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"strconv"

	"golang.org/x/exp/rand"
)

// ! run only once, to initialize demo users.
func InitDemoUsers() {
	for i := 0; i < db.DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		hashedPassword, err := auth.HashPassword(iStr)
		if err != nil {
			log.Println("Error hashing password: ", err)
		}
		//If I save the user and get all the user then I can just add a new connection to all the existing users.  
		err = db.SaveUser(email, hashedPassword)
		if err != nil {
			log.Println("Error saving user: ", err)
		}
	}
	CreateProfile()
	log.Println("Demo users initialized")
}

func RemoveDemoUsers() {
	for i := 0; i < db.DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		err := db.DeleteUser(email)
		if err != nil {
			log.Print(email)
			log.Println(" Error deleting user: ", err)
		}
	}
}

func CreateProfile() {
	for i := 0; i < db.DEMO_USER_COUNT; i++ {
		rand.Seed(uint64(i))
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		uuid, err := db.GetUserUUIDFromUserEmail(email)
		if err != nil {
			log.Println("Error getting user uuid: ", err)
		}

		// Add two Genres
		for j := 0; j <= 2; j++ {
			rndNum, err := GenerateRandomNumber(1, 10)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			db.AddInterestToUser(rndNum, uuid)
		}
		// Add play style
		for k := 0; k <= 2; k++ {
			rndNum, err := GenerateRandomNumber(11, 16)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			db.AddInterestToUser(rndNum, uuid)
		}
		// Add platform
		rndPlatform, err := GenerateRandomNumber(17, 21)
		db.AddInterestToUser(rndPlatform, uuid)
		// Add communication
		for m := 0; m <= 2; m++ {
			rndNum, err := GenerateRandomNumber(22, 26)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			db.AddInterestToUser(rndNum, uuid)
		}
		// Add goals
		for n := 0; n <= 2; n++ {
			rndNum, err := GenerateRandomNumber(27, 31)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			db.AddInterestToUser(rndNum, uuid)
		}
		// Add session
		rndSession, err := GenerateRandomNumber(32, 34)
		db.AddInterestToUser(rndSession, uuid)
		// Add vibe
		rndVibe, err := GenerateRandomNumber(35, 41)
		db.AddInterestToUser(rndVibe, uuid)
		// Add language
		rndLanguage, err := GenerateRandomNumber(42, 48)
		db.AddInterestToUser(rndLanguage, uuid)
		if err != nil {
			log.Println("Error saving profile: ", err)
		}
		
	}

}

func GenerateRandomNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("invalid range: min (%d) cannot be greater than max (%d)", min, max)
	}
	return rand.Intn(max-min+1) + min, nil
}
