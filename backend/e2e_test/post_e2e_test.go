//go:build e2e

package e2e

import (
	"context"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"io"
	"net/http"
	"testing"
	"time"
)

type RegisterRequest struct {
	Login       string
	Password    string
	Picture     string
	Description string
	Mail        string
}

type PublishRequest struct {
	Content string
	Perms   bool
}

type ViewPost struct {
	Content string `json:"content"`
	Perms   bool   `json:"perms"`
}

type ViewPostsResponse struct {
	Posts []ViewPost `json:"posts"`
}

type AssertBody func(body []byte) error

func customAssertBody() AssertBody {
	return func(bytes []byte) error {
		if len(bytes) == 0 {
			return errors.New("response body is empty")
		}

		var v ViewPostsResponse

		err := json2.Unmarshal(bytes, &v)
		if err != nil {
			return fmt.Errorf("json unmarshal: %w", err)
		}

		if len(v.Posts) != 1 {
			return errors.New("wrong response len")
		}

		if v.Posts[0].Content != "aaa" {
			return errors.New("wrong response content")
		}

		if v.Posts[0].Perms != false {
			return errors.New("wrong response perms")
		}

		return nil
	}
}

func Test_ViewPosts(t *testing.T) {
	tokens := make(map[string][]string)

	cute.NewTestBuilder().
		Title("Просмотр платных и бесплатных постов").
		Tags("posts").
		// Подготавливаем запрос на создание
		CreateStep("Register user1").
		RequestBuilder( // Создаём HTTP-запрос, который будет отправлен
			cute.WithURI("http://localhost:8090/register"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(RegisterRequest{
				Login:       "user1",
				Password:    "password",
				Picture:     "aaa",
				Description: "aaa",
				Mail:        "a@a.a",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusCreated).
		AssertBody(
			json.Present("$.token"),
		).
		NextTest().
		AfterTestExecute(
			func(response *http.Response, errors []error) error {
				b, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}

				temp, err := json.GetValueFromJSON(b, "$.token") // Получаем email из тела ответа
				if err != nil {
					return err
				}

				tokens["User-Token"] = []string{fmt.Sprint(temp)} // Сохраняем email

				return nil
			},
		).
		// Подготавливаем запрос на удаление
		CreateStep("Publish free post").
		RequestBuilder(
			cute.WithURI("http://localhost:8090/posts"),
			cute.WithMethod(http.MethodPost),
			cute.WithHeaders(tokens),
			cute.WithMarshalBody(PublishRequest{
				Content: "aaa",
				Perms:   false,
			}),
		).
		AssertBody(
			json.Equal("$.post.content", "aaa"),
			json.Equal("$.published", true),
		).
		NextTest().
		CreateStep("Publish paid post").
		RequestBuilder(
			cute.WithURI("http://localhost:8090/posts"),
			cute.WithMethod(http.MethodPost),
			cute.WithHeaders(tokens),
			cute.WithMarshalBody(PublishRequest{
				Content: "bbb",
				Perms:   true,
			}),
		).
		AssertBody(
			json.Equal("$.post.content", "bbb"),
			json.Equal("$.published", true),
		).
		NextTest().
		CreateStep("Register user2").
		RequestBuilder( // Создаём HTTP-запрос, который будет отправлен
			cute.WithURI("http://localhost:8090/register"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(RegisterRequest{
				Login:       "user2",
				Password:    "password",
				Picture:     "bbb",
				Description: "bbb",
				Mail:        "b@b.b",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusCreated).
		AssertBody(
			json.Present("$.token"),
		).
		NextTest().
		AfterTestExecute(
			func(response *http.Response, errors []error) error {
				b, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}

				temp, err := json.GetValueFromJSON(b, "$.token") // Получаем email из тела ответа
				if err != nil {
					return err
				}

				tokens["User-Token"] = []string{fmt.Sprint(temp)} // Сохраняем email

				return nil
			},
		).
		CreateStep("View user1 profile posts").
		RequestBuilder( // Создаём HTTP-запрос, который будет отправлен
			cute.WithURI("http://localhost:8090/users/user1/posts"),
			cute.WithMethod(http.MethodGet),
			cute.WithHeaders(tokens),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			cute.AssertBody(customAssertBody()),
		).
		ExecuteTest(context.Background(), t)
}
