package main

//START OMIT
func hash(s string) int {
	return len(s) % 5
}

hash("jon")=3 and hash("password")=3   *BAD* // HL

//END OMIT

//STARTPW OMIT
password
abc123
abcd1234
aksdfj
//ENDPW OMIT

//STARTRAINBOW OMIT
fake-hash-value-for-password
fake-hash-value-for-abc123
fake-hash-value-for-abcd1234
fake-hash-value-for-aksdfj
//ENDRAINBOW OMIT

//STARTUSR OMIT
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"` // HL
	PasswordHash string `gorm:"not null"` // HL
}

func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}

//ENDUSR OMIT
