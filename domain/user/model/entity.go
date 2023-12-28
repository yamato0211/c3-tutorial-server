package usermodel

type User struct {
	ID          string
	Name        UserName
	Description string
	Score       int
	IconURL     string
	FirebaseID  string
}

func NewUser(id, name, description, iconURL, firebaseID string) (*User, error) {
	userName := NewUserName(name)
	if err := userName.Validate(); err != nil {
		return nil, err
	}
	return &User{
		ID:          id,
		Name:        userName,
		Description: description,
		Score:       0,
		IconURL:     iconURL,
		FirebaseID:  firebaseID,
	}, nil
}

func (u *User) SetName(name string) error {
	n := NewUserName(name)
	if err := n.Validate(); err != nil {
		return err
	}

	u.Name = n
	return nil
}
