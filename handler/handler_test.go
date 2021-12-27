package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"inert/model"
	"inert/repository"
	"inert/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	name           string
	method         string
	input          string
	exceptedValue  string
	exceptedStatus int
	mockArg        interface{}
	mockErr        error
}

var (
	// parsedTime, _ = time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z")
	repo        = repository.InitStorage(nil)
	s           = service.InitService(repo)
	hr          = InitInertHandler(s)
	listAddUser = []Dummy{
		{
			name:           "ok",
			method:         http.MethodPost,
			input:          `{"id":0,"name":"Bekzh","age":28,"participants":{"10":"adam@mail.kz"}}`,
			exceptedValue:  `1` + "\n",
			exceptedStatus: 200,
			mockArg:        1,
			mockErr:        nil,
		},
		{
			name:           "bad request",
			method:         http.MethodPost,
			input:          `{"ids":1`,
			exceptedValue:  `"unexpected end of JSON input"` + "\n",
			exceptedStatus: 400,
			mockArg:        0,
			mockErr:        nil,
		},
		{
			name:           "not find participiant id",
			method:         http.MethodPost,
			input:          `{"id":0,"name":"Bekzh","age":28,"participants":{"0":"molly@mail.kz"}}`,
			exceptedValue:  `"not find participiant by id"` + "\n",
			exceptedStatus: 500,
			mockArg:        0,
			mockErr:        errors.New("not find participiant by id"),
		},
	}
)

func TestAddUser(t *testing.T) {
	as := assert.New(t)
	//	hr := Handler{}

	mock := new(service.MockService)

	for _, test := range listAddUser {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, "/add", bytes.NewBuffer([]byte(test.input))) //]byte(test.input)
			if err != nil {
				t.Errorf("%v", err)
			}
			resRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(hr.AddUser)

			var obj model.Inert
			err = json.Unmarshal([]byte(test.input), &obj)
			if err != nil {
				// t.Errorf("%v", err)
				log.Println(err)
			}
			log.Print(obj, "test side obj")
			mock.On("AddUser", obj).Return(test.mockArg, test.mockErr)
			hr.Service = mock
			log.Print("prep mock func", test.mockArg.(int))
			handler.ServeHTTP(resRecorder, req)

			as.Equal(test.exceptedValue, resRecorder.Body.String())
			as.Equal(test.exceptedStatus, resRecorder.Result().StatusCode)
		})
	}
}
