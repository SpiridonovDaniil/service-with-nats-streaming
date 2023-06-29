package memory

import (
	"context"
	"encoding/json"
	"testing"

	"l0/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestCashe_Write(t *testing.T) {
	testTable := []struct {
		name             string
		data             json.RawMessage
		id               string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "true",
			data:             []byte(`{"order_uid":"test"}`),
			id:               "test",
			expectedResponse: "{\"order_uid\":\"test\"}",
			expectedError:    nil,
		},
		{
			name:             "false",
			data:             []byte(`{"order_uid":"test"}`),
			id:               "test",
			expectedResponse: "{\"order_uid\":\"test\"}",
			expectedError:    models.ErrAlreadyInTheDB,
		},
	}
	ctx := context.Background()
	cashe := New()
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name == "false" {
				_ = cashe.Write(ctx, testCase.data, testCase.id)
			}
			err := cashe.Write(ctx, testCase.data, testCase.id)
			result, errRead := cashe.Read(ctx, testCase.id)
			if errRead != nil {
				t.Errorf("ошибка в методе чтения из структуры Cashe")
				t.Fail()
			}

			assert.Equal(t, testCase.expectedResponse, string(result))
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestCashe_Read(t *testing.T) {
	testTable := []struct {
		name             string
		data             json.RawMessage
		id               string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "true",
			data:             []byte(`{"order_uid":"test_true"}`),
			id:               "test_true",
			expectedResponse: "{\"order_uid\":\"test_true\"}",
			expectedError:    nil,
		},
		{
			name:             "false",
			data:             []byte(`{"order_uid":"test_false"}`),
			id:               "test_false",
			expectedResponse: "",
			expectedError:    models.ErrNotFound,
		},
	}
	ctx := context.Background()
	cashe := New()
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.name == "true" {
				errWrite := cashe.Write(ctx, testCase.data, testCase.id)
				if errWrite != nil {
					t.Errorf("ошибка в методе записи в структуру Cashe")
					t.Fail()
				}
			}

			result, err := cashe.Read(ctx, testCase.id)

			assert.Equal(t, testCase.expectedResponse, string(result))
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
