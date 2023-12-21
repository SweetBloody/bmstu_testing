package integration_test

import (
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/models"
	repo "github.com/SweetBloody/bmstu_testing/backend/internal/pkg/team/repository/postgresql"
	"github.com/SweetBloody/bmstu_testing/backend/internal/pkg/team/usecase"
)

func (suite *UsecaseRepositoryTestSuite) TestUsecaseCreateTeam() {
	elem := &models.Team{
		Name:    "asdasd",
		Country: "USA",
		Base:    "New",
	}

	repo := repo.NewPsqlTeamRepository(suite.db)
	usecase := usecase.NewTeamUsecase(repo)

	id, err := usecase.Create(elem)

	suite.Assert().NoError(err)
	elem.ID = id

	res, err := repo.GetTeamById(elem.ID)
	suite.Assert().NoError(err)

	suite.Assert().Equal(elem.ID, res.ID)
	suite.Assert().Equal(elem.Name, res.Name)
	suite.Assert().Equal(elem.Country, res.Country)
	suite.Assert().Equal(elem.Base, res.Base)
}
