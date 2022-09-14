package user

import (
	"strings"

	cfg "sybo/configs"
	"sybo/models"
	"sybo/models/user"

	"github.com/google/uuid"
)

//GetAll: fetch all the users.
//TODO add some query params
func Get(wds []models.WhereData) ([]user.UserLimited, error) {

	//Get all the users
	users, err := user.Get(wds)
	if err != nil {

		return nil, err
	}

	//Limit the data available, aka remove score and others
	var usersFormated []user.UserLimited

	for _, u := range users {

		usersFormated = append(usersFormated, u.ToLimitedStruct())
	}

	return usersFormated, err
}

//New: add a new user
func New(u *user.User) error {

	//Make sure we have a uuid generated
	//TODO could check if there is a uuid provided and use it but it would be prone to errors
	u.ID = uuid.New().String()

	return u.New()
}

//SaveState: Save current state for the user
func SaveState(u *user.User) error {

	return u.SaveState()
}

//LoadState: Load current state for the user.
func LoadState(u *user.User) error {

	return u.LoadState()
}

//UpdateFriends: update the list of friend of a user. Here is replaces all the current list.
func UpdateFriends(friends user.FriendsList, u *user.User) error {

	//Concat all the friends in a comma separated list
	u.Friends = strings.Join(friends.Friends, cfg.SEPARATOR)

	return u.UpdateFriends()
}

//GetFriends: get all the friends from a user.
func GetFriends(u *user.User) ([]user.Friend, error) {

	//Get all the friends
	friends, err := u.GetFriends()
	if err != nil {

		return nil, err
	}

	//Convert from the user to the friend format (essentiall for the highscore). Could be extended in the future
	var friendsFormated []user.Friend

	for _, f := range friends {

		friendsFormated = append(friendsFormated, f.ToFriendStruct())
	}

	return friendsFormated, err
}
