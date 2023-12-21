package integration_test

import (
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/models"
	repo "github.com/SweetBloody/bmstu_testing/backend/internal/pkg/track/repository/postgresql"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/track/usecase"
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
