package sectionmodel

type Section struct {
	ID          string
	Description string
	Title       string
	TutorialID  string
}

func NewSection(id, description, title, tutorialID string) *Section {
	return &Section{
		ID:          id,
		Description: description,
		Title:       title,
		TutorialID:  tutorialID,
	}
}
