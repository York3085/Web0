package dao

import (
	"os"
)

var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func AddUser(username, password string) {
	database[username] = password
}

func Save0(data0 string) error {
	file, err := os.Create(data0)
	if err != nil {
		return err
	}
	defer file.Close()

	for username, password := range database {
		line := username + ":" + password + "\n"
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}
