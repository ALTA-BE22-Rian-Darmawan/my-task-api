package data

import (
	"my-task-app/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.DataInterface.
func (u *userQuery) Insert(input user.Core) error {

	userGorm := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Address:  input.Address,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Delete implements user.DataInterface.
func (u *userQuery) Delete(id uint) error {
	tx := u.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *userQuery) Update(id uint, input user.Core) error {
	var userGorm User
	tx := u.db.First(&userGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	userGorm.Name = input.Name
	userGorm.Email = input.Email
	userGorm.Password = input.Password
	userGorm.Phone = input.Phone
	userGorm.Address = input.Address
	tx = u.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements user.DataInterface.
func (u *userQuery) SelectByEmail(email string) (*user.Core, error) {
	// variable penampung datanya
	var userData User
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Phone:     userData.Phone,
		Address:   userData.Address,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &usercore, nil
}

// SelectById implements user.DataInterface.
func (u *userQuery) SelectById(id uint) (*user.Core, error) {
	// variable penampung datanya
	var userData User
	tx := u.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Phone:     userData.Phone,
		Address:   userData.Address,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &usercore, nil
}
