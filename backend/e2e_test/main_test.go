//go:build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/app/http"
	"os"
	"testing"
	"text/template"
	"time"
)

//func setupTestDatabase() (testcontainers.Container, string, int, error) {
//	containerReq := testcontainers.ContainerRequest{
//		Image:        "postgres:latest",
//		ExposedPorts: []string{"5432/tcp"},
//		WaitingFor:   wait.ForListeningPort("5432/tcp"),
//		Env: map[string]string{
//			"POSTGRES_DB":       "testdb",
//			"POSTGRES_PASSWORD": "postgres",
//			"POSTGRES_USER":     "postgres",
//		},
//	}
//
//	dbContainer, _ := testcontainers.GenericContainer(
//		context.Background(),
//		testcontainers.GenericContainerRequest{
//			ContainerRequest: containerReq,
//			Started:          true,
//		})
//
//	host, _ := dbContainer.Host(context.Background())
//	port, _ := dbContainer.MappedPort(context.Background(), "5432")
//
//	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Int())
//	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, "", 0, fmt.Errorf("gorm open: %w", err)
//	}
//
//	sqlDB, err := pureDB.DB()
//	if err != nil {
//		return nil, "", 0, fmt.Errorf("get db: %w", err)
//	}
//
//	if err = goose.Up(sqlDB, "../../deployments/migrations/migrations_postgres"); err != nil {
//		return nil, "", 0, fmt.Errorf("up migrations: %w", err)
//	}
//
//	return dbContainer, host, port.Int(), nil
//}

//type dbConfig struct {
//	Host string
//	Port int
//}

func TestMain(m *testing.M) {
	tmplFile := flag.String("tmpl", "../../configs/test_config_template.yml", "tmpl file name")
	cfgFile := flag.String("cfg", "../../configs/test_config.yml", "config file name")

	for i := 0; i < 10; i++ {
		fmt.Println()
	}

	flag.Parse()
	rc := 0

	defer func() {
		os.Exit(rc)
	}()

	dbContainer, host, port, err := setupTestDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "setup test database: %s", err)
		rc = 1
		return
	}
	defer dbContainer.Terminate(context.Background())

	tmpl, err := template.ParseFiles(*tmplFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse template: %s", err)
		rc = 1
		return
	}

	cfg, err := os.Create(*cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create cfg file: %s", err)
		rc = 1
		return
	}

	v := dbConfig{
		Host: host,
		Port: port,
	}

	err = tmpl.Execute(cfg, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "execute tmpl: %s", err)
		rc = 1
		return
	}

	err = cfg.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "close file: %s", err)
		rc = 1
		return
	}

	a := http.New()

	err = a.Init(*cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %s", err)
		rc = 1
		return
	}

	go func() {
		err := a.Run(context.Background())
		fmt.Println(err)
	}()

	t := time.NewTimer(2 * time.Second)

	select {
	case <-a.Ready():
		rc = m.Run()
	case <-t.C:
		fmt.Println("timeout")
		rc = 1
	}
}
