package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	authHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/auth/delivery/http"
	driverHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/driver/delivery/http"
	grandPrixHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/grand_prix/delivery/http"
	qualHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/qual_result/delivery/http"
	raceHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/race_result/delivery/http"
	teamHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/team/delivery/http"
	trackHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/track/delivery/http"
	userHandler "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/user/delivery/http"

	driverRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/driver/repository/postgresql"
	grandPrixRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/grand_prix/repository/postgresql"
	qualRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/qual_result/repository/postgresql"
	raceRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/race_result/repository/postgresql"
	teamRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/team/repository/postgresql"
	trackRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/track/repository/postgresql"
	userRepository "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/user/repository/postgresql"

	driverUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/driver/usecase"
	grandPrixUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/grand_prix/usecase"
	qualUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/qual_result/usecase"
	raceUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/race_result/usecase"
	teamUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/team/usecase"
	trackUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/track/usecase"
	userUsecase "git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/pkg/user/usecase"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"git.iu7.bmstu.ru/kaa20u554/testing/backend/internal/app/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// @title FormulOne Web-Server
// @version 1.0
// @description API Server for F1 Grand-Prix info
// @termsOfService  http://swagger.io/terms/

// @host localhost:5259
// @BasePath /
func main() {
	params := fmt.Sprintf("user=postgresql dbname=postgresql password=postgresql host=%s port=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driverRepo := driverRepository.NewPsqlDriverRepository(db)
	teamRepo := teamRepository.NewPsqlTeamRepository(db)
	trackRepo := trackRepository.NewPsqlTrackRepository(db)
	gpRepo := grandPrixRepository.NewPsqlGPRepository(db)
	raceRepo := raceRepository.NewPsqlRaceResultRepository(db)
	qualRepo := qualRepository.NewPsqlQualResultRepository(db)
	userRepo := userRepository.NewPsqlUserRepository(db)

	driverUcase := driverUsecase.NewDriverUsecase(driverRepo)
	teamUcase := teamUsecase.NewTeamUsecase(teamRepo)
	trackUcase := trackUsecase.NewTrackUsecase(trackRepo)
	gpUcase := grandPrixUsecase.NewGrandPrixUsecase(gpRepo)
	raceUcase := raceUsecase.NewRaceResultUsecase(raceRepo)
	qualUcase := qualUsecase.NewQualResultUsecase(qualRepo)
	userUcase := userUsecase.NewUserUsecase(userRepo)

	m := mux.NewRouter()

	driverHandler.NewDriverHandler(m, driverUcase, raceUcase)
	teamHandler.NewTeamHandler(m, teamUcase)
	trackHandler.NewTrackHandler(m, trackUcase)
	grandPrixHandler.NewDriverHandler(m, gpUcase, raceUcase, qualUcase)
	raceHandler.NewRaceResultHandler(m, raceUcase)
	qualHandler.NewQualResultHandler(m, qualUcase)
	authHandler.NewAuthHandler(m, userUcase)
	userHandler.NewUserHandler(m, userUcase)

	m.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
	//m.HandleFunc("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")
	mMiddleware := middleware.LogMiddleware(m)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", handlers.CORS()(mMiddleware))
}
