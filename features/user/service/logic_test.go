package service

import (
	"mia/my_task_app/features/user"
	"mia/my_task_app/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func TestGetAll(t *testing.T){
//		mocks.UserData:=new(mocks.UserData)
//		returnData:+
//	}
func TestCreate(t *testing.T) {
	mocksUserDataLayer := new(mocks.UserData)

	t.Run("Success Create User", func(t *testing.T) {
		insertData := user.CoreUser{
			ID:          1,
			Name:        "Popol",
			Email:       "Popol@ma.co",
			Password:    "12345",
			PhoneNumber: "0898989382",
		}
		// Expectation on the mock
		mocksUserDataLayer.On("Insert", insertData).Return(uint(1), nil).Once()

		//object service layer dengan mock
		srv := New(mocksUserDataLayer)
		createdUserID, err := srv.Create(insertData)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), createdUserID)
		mocksUserDataLayer.AssertExpectations(t)
	})
}
