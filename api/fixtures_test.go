package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vbogretsov/sgmock"
)

const apiKey = "xxx"

var (
	msg1 = sgmock.Message{
		From: sgmock.Address{
			Email: "sender@mail.com",
		},
		Personalizations: []sgmock.Personalizations{
			{
				Subject: "test2",
				To: []sgmock.Address{
					{
						Email: "to1@mail.com",
					},
				},
			},
		},
		Content: sgmock.Content{
			Type:  "text/plain",
			Value: "Hello!",
		},
	}
	msg2 = sgmock.Message{
		From: sgmock.Address{
			Email: "sender@mail.com",
		},
		Personalizations: []sgmock.Personalizations{
			{
				Subject: "test2",
				To: []sgmock.Address{
					{
						Email: "to1@mail.com",
					},
					{
						Email: "to2@mail.com",
					},
				},
			},
		},
		Content: sgmock.Content{
			Type:  "text/html",
			Value: "<h>Hello!<h/>",
		},
	}
)

type httpError struct {
	code int
}

func (e httpError) Error() string {
	return strconv.Itoa(e.code)
}

type client struct {
	e *echo.Echo
}

func (c client) send(token string, msg sgmock.Message) error {
	buf, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	req := httptest.NewRequest(echo.POST, "/v3/mail/send", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {

	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	rec := httptest.NewRecorder()

	c.e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		return httpError{code: rec.Code}
	}

	return nil
}

func (c client) list() ([]sgmock.Message, error) {
	req := httptest.NewRequest(echo.GET, "/ctl/list", nil)
	req.Header.Set("Accept", "application/json")
	rec := httptest.NewRecorder()

	c.e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		return nil, httpError{code: rec.Code}
	}

	var res []sgmock.Message
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c client) clear() error {
	req := httptest.NewRequest(echo.POST, "/ctl/clear", nil)
	rec := httptest.NewRecorder()

	c.e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		return httpError{code: rec.Code}
	}

	return nil
}
