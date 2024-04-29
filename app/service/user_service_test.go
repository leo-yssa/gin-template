package service_test

import (
	"bytes"
	"encoding/json"
	"gin-api/app/domain/dao"
	"gin-api/app/repository/mocks"
	"gin-api/app/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func loadFixtures() (s service.UserService, u *mocks.UserRepository, r *mocks.RoleRepository) {
	u = &mocks.UserRepository{}
	r = &mocks.RoleRepository{}
	s = service.NewUserService(u)
	gin.SetMode(gin.TestMode)
	return s, u, r
}

func TestRegisterUser(t *testing.T) {
	s, u, r := loadFixtures()
	r.On("Save", &dao.Role{
		ID: 1,
		Role: "sys",
	}).Return(&dao.Role{})
	user1 := &dao.User{
		Id: "1",
		Email: "test@test.com",
		Name: "test",
		RoleID: 1,
	}
	u.On("Save", mock.Anything).Return(dao.User{}, nil)

	user1Bytes, err := json.Marshal(user1)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, err = http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(user1Bytes))
	c.Request.Header.Set("Content-Type", "application/json")
	assert.NotPanics(t, func(){s.RegisterUser(c)})
	assert.Equal(t, http.StatusAccepted, w.Code)
}