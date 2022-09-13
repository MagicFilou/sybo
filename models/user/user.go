package user

import (
	"sybo/clients"
)

// A user
type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	GamesPlayed int    `json:"games_played,omitempty"`
	Score       int    `json:"score,omitempty"`
	Friends     string `json:"friends,omitempty"`
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
