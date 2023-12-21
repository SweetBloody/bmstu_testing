package integration_test

import (
	repo "github.com/SweetBloody/bmstu_testing/backend/internal/pkg/driver/repository/postgresql"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/driver/usecase"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/models"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseCreateDriver() {
	elem := &models.Driver{
		Name:      "asdasd",
		Country:   "USA",
		BirthDate: "2000-01-01T00:00:00Z",
	}

	repo := repo.NewPsqlDriverRepository(suite.db)
	usecase := usecase.NewDriverUsecase(repo)

	id, err := usecase.Create(elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetDriverById(elem.ID)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.ID, res.ID)
	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Country, res.Country)
	suite.Assert().Equal(elem.BirthDate, res.BirthDate)
}
