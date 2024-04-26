package service_test

import (
	"context"
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

func loadFixtures() (s service.UserService, m *mocks.UserRepository) {
	m = &mocks.UserRepository{}
	s = service.NewUserService(m)
	gin.SetMode(gin.TestMode)
	return s, m
}

func TestRegisterUser(t *testing.T) {
	s, m := loadFixtures()
	user1 := dao.User{
		Id: "1",
		Email: "test@test.com",
		Name: "test",
		RoleID: 1,
	}
	user1Bytes, err := json.Marshal(user1)
	assert.NoError(t, err)
	c, e := gin.CreateTestContext(httptest.NewRequest(http.MethodPost, "/api/user", user1Bytes))
	
    m.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		t.Log("testify/mock과 mockery를 이용한 Mocking. Args: ", args.Get(0))
	}).Return(&dao.User{}, nil)

    // then
    assert.Panics(t, func(){s.RegisterUser(context.TODO().(*gin.Context))})
    // m.AssertCalled(t, "Save", mock.Anything, mock.MatchedBy(userMatcher))
}