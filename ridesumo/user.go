package ridesumo

import (
	"sessionauth"
)

type User struct {
	uniqueId string
	Name     string
	Email    string

	authenticated bool
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.uniqueId
}

// GetById will populate a user object from a database model with
// a matching id.
// func (u *User) GetById(id interface{}) error {
func (u *User) GetById(id interface{}) error {
	// log.Println("GetById called for user: " + id.(string))
	// may want to speed this up by storing the user in the datastore and only refreshing on command
	// context := appengine.NewContext(req)
	// var err error
	// u.data, err = nxtcutils.GetUser(context, id.(string), "shippingAddresses", "paymentCards", "orders", "scripts", "scripts.product")
	// if err != nil {
	// 	return err
	// }

	// u.Username = u.data["username"].(string)
	// u.UniqId = u.data["objectId"].(string)

	// err := dbmap.SelectOne(u, "SELECT * FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return err
	// }

	return nil
}
