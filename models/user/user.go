package user

import (
	"fmt"
	"strings"

	"sybo/clients"
	cfg "sybo/configs"
	"sybo/models"

	"gorm.io/gorm"
)

// User: A user for the game !
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	GamesPlayed int    `json:"gamesPlayed"`
	Score       int    `json:"score"`
	Friends     string `json:"friends"`
}

// FriendsList: convenience struct to parse the update friend body request.
type FriendsList struct {
	Friends []string `json:"friends"`
}

// UserLimited: struct to return the data of users without all the info
type UserLimited struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Friend: struct to return the data of frtiends with a different score key. If highscore and score would be harmonized User struct could be used instead
type Friend struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"highscore"`
}

// ToFriendStruct: convenience method to convert a User to a Friend
func (u User) ToFriendStruct() Friend {
	return Friend{
		ID:    u.ID,
		Name:  u.Name,
		Score: u.Score,
	}
}

// ToLimitedStruct: convenience method to convert a User to a UserLimited
func (u User) ToLimitedStruct() UserLimited {
	return UserLimited{
		ID:   u.ID,
		Name: u.Name,
	}
}

// Get: get all the users with the given params
func Get(wds []models.WhereData) ([]User, error) {

	db, err := clients.GetCon()
	if err != nil {
		return nil, err
	}

	var result *gorm.DB
	var users []User

	//Build the where query to find users according to the where params
	query, args := models.BuildWHereQuery(wds)

	//This also means that if no parameters are given the get all is still valid
	if len(query) > 0 {
		result = db.Where(query, args...).Find(&users)
	} else {
		result = db.Find(&users)
	}

	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

// New: add a new user
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

// SaveState: Update a user with current state
func (u *User) SaveState() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Model(&u).Updates(map[string]interface{}{"games_played": u.GamesPlayed, "score": u.Score})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No users with that user ID")
	}

	return nil
}

// LoadState: Load the state of a user
func (u *User) LoadState() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Find(&u)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No users with that user ID")
	}

	return nil
}

// UpdateFriends: Update a user's friends list
func (u *User) UpdateFriends() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	result := db.Model(&u).Update("friends", u.Friends)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No users with that user ID")
	}

	return nil
}

// GetFriends: Get a user's friends list
func (u *User) GetFriends() ([]User, error) {

	//Get db con
	db, err := clients.GetCon()
	if err != nil {
		return nil, err
	}

	//First get the user iteself
	result := db.Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No users with that user ID")
	}

	//Then get all the friends in the list
	var friends []User
	var toFind []string

	toFind = strings.Split(u.Friends, cfg.SEPARATOR)

	result = db.Where("id IN ?", toFind).Find(&friends)

	if result.Error != nil {
		return friends, result.Error
	}

	return friends, nil
}
