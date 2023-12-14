package integration_test

import (
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/models"
	repo "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/team/repository/postgresql"
	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/team/usecase"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseUpdateTeam() {
	id := 1
	elem := &models.Team{
		Name:    "asdasd",
		Country: "USA",
		Base:    "New",
	}

	repo := repo.NewPsqlTeamRepository(suite.db)
	usecase := usecase.NewTeamUsecase(repo)

	old, err := repo.GetTeamById(id)
	suite.Assert().NoError(err)

	err = usecase.Update(id, elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetTeamById(id)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Country, res.Country)
	suite.Assert().Equal(elem.Base, res.Base)

	err = usecase.Update(id, old)
	suite.Assert().NoError(err)
}
