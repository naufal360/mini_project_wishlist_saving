package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wishlist/dto/payload"
	"wishlist/service/user/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	userTest := payload.Register{
		Name:     "Ahmad Naufal",
		Username: "naufal",
		Email:    "naufal@mail.com",
		Password: "user123naufal",
	}

	testCases := []struct {
		name       string
		body       payload.Register
		buildStubs func(user_service *mocks.MockUserService)
		checkBody  func(recorder *httptest.ResponseRecorder)
	}{
		{
			"created",
			userTest,
			func(user_service *mocks.MockUserService) {
				// hashingUser, _ := util.HashPassword(userTest.Password)
				// arg := model.User{
				// 	Name:     userTest.Name,
				// 	Username: userTest.Username,
				// 	Email:    userTest.Email,
				// 	Password: hashingUser,
				// }
				user_service.EXPECT().
					RegisterUser(gomock.Any()).
					Times(1).
					Return(nil)
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, 200)
				// requireBodyMatchUser(t, recorder.Body, userTest)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			t.Skip()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user_service := mocks.NewMockUserService(ctrl)
			v.buildStubs(user_service)

			// server := newTestServer(t, user_service)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(v.body)
			require.NoError(t, err)

			url := "/signup"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			request.Header.Set(echo.HeaderContentType, "application/json")

			// server.router.ServeHTTP(recorder, request)
			v.checkBody(recorder)
		})
	}
}

// func TestLoginUser(t *testing.T) {
// 	user := payload.Login{
// 		Username: "naufal",
// 		Password: "",
// 	}
// }

// func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user payload.Register) {
// 	data, err := io.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotUser model.User
// 	err = json.Unmarshal(data, &gotUser)

// 	require.NoError(t, err)
// 	hashingUser, err := util.HashPassword(user.Password)

// 	require.NoError(t, err)

// 	require.Equal(t, user.Name, gotUser.Name)
// 	require.Equal(t, user.Username, gotUser.Username)
// 	require.Equal(t, user.Email, gotUser.Email)
// 	require.Equal(t, hashingUser, gotUser.Password)
// }
