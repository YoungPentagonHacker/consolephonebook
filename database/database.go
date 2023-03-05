package database

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type PersonDto struct {
	Name         string
	PhoneNumbers []string
}

const path = "database/db.json"

func init() {
	if _, err := os.ReadFile(path); err != nil {
		os.Create(path)
	}
}

func DeleteUser(name string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	users := []PersonDto{}
	isExist := false
	json.Unmarshal(file, &users)
	res := []PersonDto{}
	for _, x := range users {
		if x.Name != name {
			res = append(res, x)
		} else {
			isExist = true
		}
	}
	if isExist == false {
		return errors.New("Удаляемого пользователя не существует")
	}
	data, err := json.MarshalIndent(res, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, data, os.ModePerm)
	return nil
}

func AddNumber(name string, numbers []string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	isExist := false
	users := []PersonDto{}
	json.Unmarshal(file, &users)
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			users[i].PhoneNumbers = append(users[i].PhoneNumbers, numbers...)
			isExist = true
		}
	}
	if isExist == false {
		users = append(users, PersonDto{name, numbers})
	}
	data, err := json.MarshalIndent(users, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, data, os.ModePerm)
}

func GetUsers() []PersonDto {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	users := []PersonDto{}

	json.Unmarshal(file, &users)
	return users
}

func AddUser(name string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	users := []PersonDto{}
	json.Unmarshal(file, &users)
	for _, x := range users {
		if x.Name == name {
			return errors.New("Такой человек уже записан")
		}
	}
	users = append(users, PersonDto{name, []string{}})

	data, err := json.MarshalIndent(users, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, data, os.ModePerm)
	return nil
}
