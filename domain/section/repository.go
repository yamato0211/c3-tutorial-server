package sectionrepository

import sectionmodel "github.com/yamato0211/c3-tutorial-server/domain/section/model"

type SectionRepository interface {
	GetSection(id string) (*sectionmodel.Section, error)
	GetSections(tutorialID string) ([]*sectionmodel.Section, error)
	UpdateSection(section *sectionmodel.Section) (*sectionmodel.Section, error)
	DeleteSection(id string) error
}
