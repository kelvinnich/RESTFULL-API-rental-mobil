package repository

import (
	"log"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	"restfull-api-rental-mobil/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type UserConnection struct{
	connected *gorm.DB
}

func ConnectCustomerRepository(db *gorm.DB) entityrepository.UserRepository{
	return &UserConnection{
		connected: db,
	}
}

func(db *UserConnection) InsertUser(user models.User) models.User{
	user.Password = hashPassword([]byte(user.Password))
	db.connected.Exec("INSERT INTO users(nik, name, email, password, no_wa) VALUES(?,?,?,?,?)", user.NIK, user.Name, user.Email, user.Password, user.NoWa)
	return user
}

func(db *UserConnection) UpdateUser(user models.User) models.User{
	if user.Password != "" {
		user.Password = hashPassword([]byte(user.Password))
	}else {
		temp := models.User{}
		db.connected.Find(&temp, user.Password)
		user.Password = temp.Password
	}

	db.connected.Exec("UPDATE users SET name = ?, email = ?, password = ?, no_wa = ? WHERE nik = ?", user.Name, user.Email, user.Password, user.NoWa, user.NIK)
	return user
}

func(db *UserConnection) VerifyCredintial(email string, password string) interface{}{
	var user models.User
	res := db.connected.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func(db *UserConnection) IsDuplicateWA(wa uint64) (tx *gorm.DB){
	var user models.User
	return db.connected.Where("wa = ?", user.NoWa).Take(&user)
}

func(db *UserConnection) FindByWA(wa uint64) models.User{
	var user models.User
	db.connected.Where("wa = ?", user.NoWa).Take(&user)
	return user
}

func(db *UserConnection) ProfileUser(nik string) models.User{
	var user models.User
	db.connected.Find(&user, nik)
	return user
}

func hashPassword(password []byte) string{
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}

	return string(hash)
}