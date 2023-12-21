package integration_test

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
)

type UsecaseRepositoryTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (suite *UsecaseRepositoryTestSuite) SetupSuite() {
	flag.Parse()
	params := fmt.Sprintf("user=postgresql dbname=postgresql password=postgresql host=%s port=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	suite.db = db
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UsecaseRepositoryTestSuite))
}
