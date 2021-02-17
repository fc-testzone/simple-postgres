package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Login  string
	Passwd string
}

var db *gorm.DB

func ConnectDB(ip string, port int, user string, passwd string, base string) error {
	var err error
	db, err = gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, passwd, ip, port, base)), &gorm.Config{})
	return err
}

func InsertUsers() error {
	var users []User

	// Find user by login
	var err = db.Find(&users, &User{Login: "serg"}).Error
	if err != nil {
		return err
	}

	// Insert new user
	if len(users) == 0 {
		var err = db.Create(&User{Login: "serg", Passwd: "123"}).Error
		if err != nil {
			return err
		}
		fmt.Println("User \"serg\" not found. Create new user")
	} else {
		fmt.Println("User \"serg\" exists. User not created")
	}

	//
	// Second user
	//

	// Find user by login
	err = db.Find(&users, &User{Login: "diman"}).Error
	if err != nil {
		return err
	}

	// Insert new user
	if len(users) == 0 {
		var err = db.Create(&User{Login: "diman", Passwd: "999999999"}).Error
		if err != nil {
			return err
		}
		fmt.Println("User \"diman\" not found. Create new user")
	} else {
		fmt.Println("User \"diman\" exists. User not created")
	}

	return nil
}

func UpdateUsers() error {
	return db.Model(&User{}).Where(&User{Login: "serg"}).Update("passwd", "00000000").Error
}

func DeleteUsers() error {
	return db.Delete(&User{}, &User{Login: "serg"}).Error
}

func SelectUsers() error {
	var out []User

	var err = db.Find(&out, &User{Login: "serg"}).Error

	for _, u := range out {
		fmt.Println(u.Login + " " + u.Passwd)
	}
	fmt.Println("=====================")

	return err
}

func SelectUsers2() error {
	var users []User

	var err = db.Find(&users, &User{Login: "alex", Passwd: "456"}).Error

	for _, u := range users {
		fmt.Println(u.Login + " " + u.Passwd)
	}
	fmt.Println("=====================")

	return err
}

func SelectAllUsers() error {
	var users []User

	var err = db.Find(&users).Error

	for _, u := range users {
		fmt.Println(u.Login + " " + u.Passwd)
	}
	fmt.Println("=====================")

	return err
}

func main() {
	var err = ConnectDB("127.0.0.1", 5432, "postgres", "123456", "futcity")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = InsertUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SelectAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SelectUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SelectUsers2()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = UpdateUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SelectAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = DeleteUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SelectAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
