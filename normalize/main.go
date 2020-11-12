package main

import "errors"

//START OMIT
type UserService struct {
	userVal
}

func (us *UserService) ByID(id uint) (*User, error) {
	return us.userVal.ByID(id)
}

type userVal struct {
	userGorm
}

func (uv *userVal) ByID(id uint) (*User, error) {
	// Validate the ID first...
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}  
	// If the ID is valid, pass it to the userGorm layer
	return uv.userGorm.ByID(id) }
}
//END OMIT

//STARTDB OMIT
type userGorm struct { 
	db *gorm.DB
}
func (ug *userGorm) ByID(id uint) (*User, error) {
	// We DO NOT need to validate data because it is already // validated
	var user User
	err := first(ug.db.Where("id = ?", id), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
//ENDDB OMIT
