/*
 TODO:write something...
*/
package getusers

import (
	"github.com/jinzhu/gorm"

	"log"
	"time"

	_ "github.com/lib/pq"
)

// User is mock-mock User Object
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// Users is array of user
type Users struct {
	Users []User `json:"users"`
}

// GetTest is return from mock data
func GetTest() []User {
	users := []User{
		{ID: 1, Name: "test1"},
		{ID: 2, Name: "test2"}}
	return users
}

// GetFromDB is return from DB data
func GetFromDB() []User {
	users := getUserFromDB()
	return users
}

func getUserFromDB() []User {
	db, err := gorm.Open("postgres", "db connect info")
	defer db.Close()
	checkError(err)
	var users []User
	db.Find(&users)
	return users
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
