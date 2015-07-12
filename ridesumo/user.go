package ridesumo

import (
	"appengine/datastore"
	"log"
	"middleware"
)

const (
	DatastoreNameUser string = "user"
)

type User struct {
	Email string

	FirstName, LastName  string
	City, State, Country string
	Gender               string
	StravaAccessToken    string

	// should probably store last update on this
	// HACK: VERY tight coupling here... what if Strava goes away?
	// datastore can't take ClubSummary type for some reason... maybe because it's an array?
	// StravaUserRecord strava.AthleteDetailed

	// Strava does it like this
	// "access_token":"a9cf6599f14fb7a81eb7e9731b60909acc0bc671",
	// "token_type":"Bearer",
	// "athlete":{
	// 	"id":323132,
	// 	"resource_state":3,
	// 	"firstname":"Mark",
	// 	"lastname":"Blackwell",
	// 	"profile_medium":"https://dgalywyr863hv.cloudfront.net/pictures/athletes/323132/2690845/1/medium.jpg",
	// 	"profile":"https://dgalywyr863hv.cloudfront.net/pictures/athletes/323132/2690845/1/large.jpg",
	// 	"city":"Sydney",
	// 	"state":"New South Wales",
	// 	"country":"Australia",
	// 	"sex":"M",
	// 	"friend":null,
	// 	"follower":null,
	// 	"premium":true,
	// 	"created_at":"2012-03-14T04:59:12Z",
	// 	"updated_at":"2015-06-22T22:16:45Z",
	// 	"badge_type_id":1,
	// 	"follower_count":32,
	// 	"friend_count":36,
	// 	"mutual_friend_count":0,
	// 	"athlete_type":0,
	// 	"date_preference":"%d/%m/%Y",
	// 	"measurement_preference":"meters",
	// 	"email":"mblackwell8@gmail.com",
	// 	"ftp":310,
	// 	"weight":74.5,
	// 	"clubs":[],
	// 	"bikes":[],
	// 	"shoes":[]
	// }

	// Only exported fields (beginning with an upper case letter) will be saved to the datastore
	// uniqueId      int64
	// an internal cache of the datastore.Key
	datastoreKey *datastore.Key

	authenticated bool
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() middleware.User {
	return &User{}
}

func NewUser(email, firstName, lastName string) *User {
	newUser := &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	return newUser
}

func (u *User) Put() error {
	var err error
	if u.datastoreKey == nil {
		newKey := datastore.NewIncompleteKey(middleware.AppEngineContext, DatastoreNameUser, nil)
		u.datastoreKey, err = datastore.Put(middleware.AppEngineContext, newKey, u)
	} else {
		_, err = datastore.Put(middleware.AppEngineContext, u.datastoreKey, u)
	}

	return err
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
	return u.datastoreKey.IntID()
}

// GetById will populate a user object from a database model with
// a matching id.
// func (u *User) GetById(id interface{}) error {
func (u *User) GetById(id interface{}) error {
	log.Printf("looking up user ID %d", id)
	u.datastoreKey = datastore.NewKey(middleware.AppEngineContext, DatastoreNameUser, "", id.(int64), nil)
	err := datastore.Get(middleware.AppEngineContext, u.datastoreKey, u)
	if err != nil {
		log.Panicf("ERROR looking up user ID %s", err.Error())
	}

	return err
}
