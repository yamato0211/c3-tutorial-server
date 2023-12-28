package tutorialmodel

type Tutorial struct {
	ID           string
	OwnerID      string
	Description  string
	Score        int
	ThumbnailURL string
	Title        string
}

func NewTutorial(id, ownerID, description, thumbnailURL, title string, score int) *Tutorial {
	return &Tutorial{
		ID:           id,
		OwnerID:      ownerID,
		Description:  description,
		Score:        score,
		ThumbnailURL: thumbnailURL,
		Title:        title,
	}
}
