package tutorialrepository

import tutorialmodel "github.com/yamato0211/c3-tutorial-server/domain/tutorial/model"

type TutorialRepository interface {
	GetTutorial(userID string) ([]*tutorialmodel.Tutorial, error)
	CreateTutorial(tutorial *tutorialmodel.Tutorial) (*tutorialmodel.Tutorial, error)
	UpdateTutorial(tutorial *tutorialmodel.Tutorial) (*tutorialmodel.Tutorial, error)
}
