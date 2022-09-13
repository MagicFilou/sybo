package user

import (
	"sybo/clients"
)

// A user
type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	GamesPlayed int    `json:"gamesPlayed,omitempty"`
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

func (u *User) SaveState() error {

	db, err := clients.GetCon()
	if err != nil {
		return err
	}

	//result := db.Model(&u).Update("score", u.Score)
	result := db.Model(&u).Updates(map[string]interface{}{"games_played": u.GamesPlayed, "score": u.Score})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
