package integration_test

import (
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/models"
	repo "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/track/repository/postgresql"
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/track/usecase"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseCreateTrack() {
	elem := &models.Track{
		Name:    "asdasd",
		Country: "USA",
		Town:    "New",
	}

	repo := repo.NewPsqlTrackRepository(suite.db)
	usecase := usecase.NewTrackUsecase(repo)

	id, err := usecase.Create(elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetTrackById(elem.ID)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.ID, res.ID)
	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Country, res.Country)
	suite.Assert().Equal(elem.Town, res.Town)
}
