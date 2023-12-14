package integration_test

import (
	repo "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/driver/repository/postgresql"
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/driver/usecase"
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/models"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseUpdateDriver() {
	id := 1
	elem := &models.Driver{
		Name:      "asdasd",
		Country:   "USA",
		BirthDate: "2000-01-01T00:00:00Z",
	}

	repo := repo.NewPsqlDriverRepository(suite.db)
	usecase := usecase.NewDriverUsecase(repo)

	old, err := repo.GetDriverById(id)
	suite.Assert().NoError(err)

	err = usecase.Update(id, elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetDriverById(id)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Country, res.Country)
	suite.Assert().Equal(elem.BirthDate, res.BirthDate)

	err = usecase.Update(id, old)
	suite.Assert().NoError(err)
}
