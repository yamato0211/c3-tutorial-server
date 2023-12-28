package entity

import "github.com/yamato0211/c3-tutorial-server/pkg/domain/value"

type User struct {
	ID          string
	Name        value.UserName
	Description string
	Score       value.UserScore
	IconURL     string
	FirebaseID  string
}

func NewUser(id, name, description, iconURL, firebaseID string, score int) (*User, error) {
	userName := value.NewUserName(name)
	if err := userName.Validate(); err != nil {
		return nil, err
	}

	userScore := value.NewScore(score)
	if err := userScore.Validate(); err != nil {
		return nil, err
	}
	return &User{
		ID:          id,
		Name:        userName,
		Description: description,
		Score:       userScore,
		IconURL:     iconURL,
		FirebaseID:  firebaseID,
	}, nil
}
