package db

import (
	"fmt"
	"log"
	"match_me_backend/auth"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

// ! run only once, to initialize demo users.
func InitDemoUsers() bool {
	for i := 0; i < DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		hashedPassword, err := auth.HashPassword(iStr)
		if err != nil {
			log.Println("Error hashing password: ", err)
		}
		err = SaveUser(email, hashedPassword)
		if err != nil {
			log.Println("Error saving user: ", err)
		}
	}
	CreateProfile()
	log.Println("Demo users initialized")
	return true
}

// DO we need to remove the demo users?
func RemoveDemoUsers() {
	for i := 0; i < DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		err := DeleteUser(email)
		if err != nil {
			log.Print(email)
			log.Println(" Error deleting user: ", err)
		}
	}
}

/*
Run only once, to create profiles for demo users.

if you want to run it again  all entries from the tables should be removed first from the following tables.

TRUNCATE TABLE user_interests;
TRUNCATE TABLE users;
TRUNCATE TABLE user_matches;
TRUNCATE TABLE profiles

*/

func CreateProfile() {
	fmt.Println("Creating profiles")
	rand.Seed(uint64(time.Now().UnixNano()))
	birthdate, err := time.Parse("2006-01-02", "1999-01-01")
	var latitude float64
	var longitude float64
	if err != nil {
		log.Println("Error parsing birthdate: ", err)
	}

	for i := 0; i < DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		uuid, err := GetUserUUIDFromUserEmail(email) //

		if err != nil {
			log.Println("Error getting user uuid: ", err)
		}
		latitude = 58.378025 + float64(i)
		longitude = 26.728493 + float64(i)

		SetUsername(uuid, "User"+iStr)
		SetBirthdate(uuid, birthdate) // 1999-01-01 all user have the same birthdate
		SetAbout(uuid, "I am a user "+iStr)
		SetPicturePath(uuid, "default_profile_pic.png")                        // TODO add bot picture no picture or default picture
		SetCity(uuid, "Estonia", "Tartu County", "Tartu", latitude, longitude) // all users are from Tartu random lat and long just stars adding distance to users
		// Add two Genres
		for j := 0; j <= 3; j++ {
			rndNum, err := GenerateRandomNumber(1, 10)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add play style
		for k := 0; k <= 3; k++ {
			rndNum, err := GenerateRandomNumber(11, 16)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add platform
		// rndPlatform, err := GenerateRandomNumber(17, 18)
		err = AddInterestToUser(17, uuid)
		if err != nil {
			log.Printf("Error adding interest to user %s: %v", uuid, err)
		}
		// Add communication
		for m := 0; m <= 2; m++ {
			// rndNum, err := GenerateRandomNumber(22, 23)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(22, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add goals
		for n := 0; n <= 2; n++ {
			rndNum, err := GenerateRandomNumber(27, 31)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add session
		rndSession, err := GenerateRandomNumber(32, 34)
		err = AddInterestToUser(rndSession, uuid)
		if err != nil {
			log.Printf("Error adding interest to user %s: %v", uuid, err)
		}
		// Add vibe

		for n := 0; n <= 3; n++ {
			rndVibe, err := GenerateRandomNumber(35, 41)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndVibe, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add language
		// rndLanguage, err := GenerateRandomNumber(42, 45)
		err = AddInterestToUser(42, uuid)
		if err != nil {
			log.Printf("Error adding Language interest to user %s: %v", uuid, err)
		}

		// err = AddUserMatchForAllExistingUsers(uuid)
		// if err != nil {
		// 	log.Println("Error adding user match for all existing users: ", err)
		// }
	}

}

func GenerateRandomNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("invalid range: min (%d) cannot be greater than max (%d)", min, max)
	}
	return rand.Intn(max-min+1) + min, nil
}
