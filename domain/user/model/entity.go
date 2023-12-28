package usermodel

type User struct {
	ID          string
	Name        string
	Description string
	Score       int
	IconURL     string
	FirebaseID  string
}

func NewUser(id, name, description, iconURL, firebaseID string) (*User, error) {
	return &User{
		ID:          id,
		Name:        name,
		Description: description,
		Score:       0,
		IconURL:     iconURL,
		FirebaseID:  firebaseID,
	}, nil
}
