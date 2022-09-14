package user

import (
	"strings"

	"sybo/models/user"

	"github.com/google/uuid"
)

func New(u *user.User) error {

	u.ID = uuid.New().String()

	return u.New()
}

func SaveState(u *user.User) error {

	return u.SaveState()
}

func LoadState(u *user.User) error {

	return u.LoadState()
}

func UpdateFriends(friends user.FriendsList, u *user.User) error {

	u.Friends = strings.Join(friends.Friends, ",")

	return u.UpdateFriends()

}

func GetFriends(u *user.User) ([]user.Friend, error) {

	friends, err := u.GetFriends()

	if err != nil {

		return nil, err
	}

	var friendsFormated []user.Friend

	for _, f := range friends {

		friendsFormated = append(friendsFormated, f.ToFriendStruct())
	}

	return friendsFormated, err
}
