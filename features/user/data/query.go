package data

import (
	"errors"
	"mia/my_task_app/features/user"
	"mia/my_task_app/helpers"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: database,
	}
}

// Delete implements user.UserDataInterface.
func (*userQuery) Delete(userId uint) error {
	panic("unimplemented")
}

// Login implements user.UserDataInterface.
func (*userQuery) Login(email string, password string) (dataLogin user.CoreUser, string error, err error) {
	panic("unimplemented")
}

// Update implements user.UserDataInterface.
func (*userQuery) Update(userId uint, updateData user.CoreUser) error {
	panic("unimplemented")
}

// Insert implements user.UserDataInterface.
func (r *userQuery) Insert(inputUser user.CoreUser) (uint, error) {
	NewUser := MapCoreUsertoUser(inputUser)
	NewUser.Password = helpers.HashPassword(inputUser.Password)

	//simpan ke db
	tx := r.db.Create(&NewUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}
	return NewUser.ID, nil
}

// SelectAll implements user.UserDataInterface.
func (r *userQuery) SelectAll() ([]user.CoreUser, error) {
	var userData []User
	tx := r.db.Find(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("failed to read data, row affected is 0")
	}

	//mapping from userData -> CoreUser
	coreUserSlice := ListMapUserToCoreUser(userData)
	return coreUserSlice, nil
}

// SelectByID implements user.UserDataInterface.
func (r *userQuery) Select(userId uint) (user.CoreUser, error) {
	var userData User
	tx := r.db.First(&userData)
	if tx.Error != nil {
		return user.CoreUser{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.CoreUser{}, errors.New("failed to read data, row affected is 0")
	}
	//Mapping User to CoreUser
	coreUser := MapUserToCoreUser(userData)
	return coreUser, nil
}
