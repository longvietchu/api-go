package tests

import (
	"api-go/configs"
	"api-go/dto"
	"api-go/service"
	"api-go/tests/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWalkCoursesCrud(t *testing.T) {
	requestCreate := helpers.Request{
		Method: http.MethodPost,
		Url:    "/courses",
	}
	requestGet := helpers.Request{
		Method: http.MethodGet,
		Url:    "/courses",
	}
	handlerFuncCreate := func(s *configs.Server, c echo.Context) error {
		return service.Create(c)
	}
	handlerFuncGet := func(s *configs.Server, c echo.Context) error {
		return service.GetAll(c)
	}

	cases := []helpers.TestCase {
		{
			"Create course success",
			requestCreate,
			dto.CreateCourseRequest{
				Title:       "title",
				Description: "description",
				Creator:     "creator",
				Level:       "level",
				URL:         "url",
				Language:    "language",
				Commitment:  "commitment",
				Rating:      "rating",
			},
			handlerFuncCreate,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 201,
				BodyPart:   "Course successfully created",
			},
		},
		{
			"Create course with empty title",
			requestCreate,
			dto.CreateCourseRequest{
				Title:       "",
				Description: "description",
				Creator:     "creator",
				Level:       "level",
				URL:         "url",
				Language:    "language",
				Commitment:  "commitment",
				Rating:      "rating",
			},
			handlerFuncCreate,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Required fields are empty",
			},
		},
		{
			"Get courses success",
			requestGet,
			"",
			handlerFuncGet,
			&helpers.QueryMock{
				Query: `SELECT * FROM "courses"  WHERE `,
				Reply: helpers.MockReply{{"id": 1, "title": "title", "description": "description"}},
			},
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "[{\"title\":\"title\",\"description\":\"description\",\"id\":1}]",
			},
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
			}
		})
	}
}