package e2e

import (
	"context"
	"github.com/ozontech/cute"
	"net/http"
	"testing"
	"time"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PublishRequest struct {
	GpName string `json:"gp_name"`
}

func Test_Drivers(t *testing.T) {
	cute.NewTestBuilder().
		Title("1 этап").
		Tags("driver").
		// Подготавливаем запрос на создание
		CreateStep("1 этап").
		RequestBuilder( // Создаём HTTP-запрос, который будет отправлен
			cute.WithURI("http://localhost:8080/auth/login"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(RegisterRequest{
				Login:    "admin",
				Password: "admin",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		NextTest().
		CreateStep("Get drivers").
		RequestBuilder(
			cute.WithURI("http://localhost:8080/api/drivers"),
			cute.WithMethod(http.MethodGet),
		).
		NextTest().
		CreateStep("change name").
		RequestBuilder(
			cute.WithURI("http://localhost:8080/api/grandprix/100/name"),
			cute.WithMethod(http.MethodPatch),
			cute.WithMarshalBody(PublishRequest{
				GpName: "bbb",
			}),
		).
		NextTest().
		CreateStep("Get name").
		RequestBuilder( // Создаём HTTP-запрос, который будет отправлен
			cute.WithURI("http://localhost:8080/api/grandprix/100"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		ExecuteTest(context.Background(), t)
}
