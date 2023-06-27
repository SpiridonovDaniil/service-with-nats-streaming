package service

//func TestService_Create(t *testing.T) {
//	type mockBehavior func(s *mock_subscription.Mockservice, expectedError error)
//	testTable := []struct {
//		name                   string
//		inputBody              string
//		mockBehavior           mockBehavior
//		expectedTestStatusCode int
//		expectedError          error
//		expectedResponse       string
//	}{
//		{
//			name:      "create HTTP status 201",
//			inputBody: `{"name" : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
//			mockBehavior: func(s *mock_subscription.Mockservice, expectedError error) {
//				s.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedError)
//			},
//			expectedTestStatusCode: 201,
//			expectedResponse:       "",
//		},
//		{
//			name:                   "create bad request",
//			inputBody:              `{"name" : "", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
//			expectedTestStatusCode: 400,
//			expectedResponse:       "[createHandler] bad request, name, phone number, reason and adding user fields must be filled in",
//		},
//		{
//			name:                   "create bad request, syntax error",
//			inputBody:              `{"name : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
//			expectedTestStatusCode: 400,
//			expectedResponse:       "[createHandler] failed to parse request, error: invalid character 'm' after object key",
//		},
//		{
//			name:      "create internal server error",
//			inputBody: `{"name" : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
//			mockBehavior: func(s *mock_subscription.Mockservice, expectedError error) {
//				s.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedError)
//			},
//			expectedTestStatusCode: 500,
//			expectedError:          fmt.Errorf("[create] error adding a user to the blacklist"),
//			expectedResponse:       "[createHandler] [create] error adding a user to the blacklist",
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			service := mock_subscription.NewMockservice(c)
//			if testCase.name == "create HTTP status 201" {
//				testCase.mockBehavior(service, testCase.expectedError)
//			}
//			if testCase.name == "create internal server error" {
//				testCase.mockBehavior(service, testCase.expectedError)
//			}
//
//			f := server.NewServer(service)
//			req, err := http.NewRequest("POST", "/", strings.NewReader(testCase.inputBody))
//			req.Header.Add("content-Type", "application/json")
//			assert.NoError(t, err)
//
//			resp, err := f.Test(req)
//			assert.NoError(t, err)
//
//			body, err := io.ReadAll(resp.Body)
//			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
//			assert.Equal(t, testCase.expectedResponse, string(body))
//		})
//	}
//}
