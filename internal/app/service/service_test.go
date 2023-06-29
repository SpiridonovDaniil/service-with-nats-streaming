package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	server "l0/internal/app/http"
	"l0/internal/memory"
	"l0/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {

	testTable := []struct {
		name                   string
		url                    string
		data                   json.RawMessage
		expectedTestStatusCode int
		expectedResponse       string
	}{
		{
			name:                   "create HTTP status 200",
			url:                    "/?key=test",
			data:                   []byte(`{"order_uid":"test"}`),
			expectedTestStatusCode: 200,
			expectedResponse:       `{"order_uid":"test"}`,
		},
		{
			name:                   "create bad request",
			url:                    "/?error=test",
			data:                   []byte(`{"order_uid":"test"}`),
			expectedTestStatusCode: 400,
			expectedResponse:       "[getHandler] search parameters are not specified",
		},
		{
			name:                   "create internal server error",
			url:                    "/?key=test",
			data:                   []byte(`{"order_uid:"test"}`),
			expectedTestStatusCode: 500,
			expectedResponse:       "[getHandler] failed to return JSON answer, error: json: error calling MarshalJSON for type json.RawMessage: invalid character 't' after object key",
		},
		{
			name:                   "not found",
			url:                    "/?key=error",
			data:                   []byte(`{"order_uid":"test"}`),
			expectedTestStatusCode: 404,
			expectedResponse:       "[getHandler] user not found",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var repo repository.Repository
			cashe := memory.New()
			ctx := context.Background()
			cashe.Write(ctx, testCase.data, "test")
			service := New(repo, cashe)
			f := server.NewServer(service)

			req, err := http.NewRequest("GET", testCase.url, strings.NewReader(""))
			req.Header.Add("content-Type", "application/json")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}
