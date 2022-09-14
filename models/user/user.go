package user

import (
	"strings"
	"sybo/clients"
)

// A user
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	GamesPlayed int    `json:"gamesPlayed,omitempty"`
	Score       int    `json:"score,omitempty"`
	Friends     string `json:"friends,omitempty"`
}

type FriendsList struct {
	Friends []string `json:"friends"`
}

type UserLimited struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Friend struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"highscore"`
}

func (u User) ToFriendStruct() Friend {
	return Friend{
		ID:    u.ID,
		Name:  u.Name,
		Score: u.Score,
	}
}

func (u User) ToLimitedStruct() UserLimited {
	return UserLimited{
		ID:   u.ID,
		Name: u.Name,
	}
}

func GetAll() ([]User, error) {

	db, err := clients.GetCon()
	if err != nil {
		return nil, err
	}

	var users []User

	result := db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (u *User) New() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) SaveState() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Model(&u).Updates(map[string]interface{}{"games_played": u.GamesPlayed, "score": u.Score})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) LoadState() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Find(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) UpdateFriends() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Model(&u).Update("friends", u.Friends)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) GetFriends() ([]User, error) {

	db, err := clients.GetCon()
	if err != nil {
		return nil, err
	}

	result := db.Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	var friends []User
	var toFind []string

	toFind = strings.Split(u.Friends, ",")

	result = db.Where("id IN ?", toFind).Find(&friends)

	if result.Error != nil {
		return friends, result.Error
	}

	return friends, nil
}
