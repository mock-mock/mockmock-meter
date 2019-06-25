/*
 TODO:write something...
*/
package getusers

// User is mock-mock User Object
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []User {
	users := []User{
		{ID: 1, Name: "test1"},
		{ID: 2, Name: "test2"}}
	return users
}
