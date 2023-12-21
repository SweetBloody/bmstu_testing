package integration_test

import (
	repo "github.com/SweetBloody/bmstu_testing/backend/internal/pkg/grand_prix/repository/postgresql"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/grand_prix/usecase"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/models"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseUpdateGP() {
	id := 1
	elem := &models.GrandPrix{
		Name:    "asdasd",
		Season:  2022,
		DateNum: 5,
		Month:   "May",
		Place:   "wow",
		TrackId: 1,
	}

	repo := repo.NewPsqlGPRepository(suite.db)
	usecase := usecase.NewGrandPrixUsecase(repo)

	old, err := repo.GetGPById(id)
	suite.Assert().NoError(err)

	err = usecase.Update(id, elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetGPById(id)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Season, res.Season)
	suite.Assert().Equal(elem.DateNum, res.DateNum)
	suite.Assert().Equal(elem.TrackId, res.TrackId)
	suite.Assert().Equal(elem.Place, res.Place)

	err = usecase.Update(id, old)
	suite.Assert().NoError(err)

}
