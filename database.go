package webhook

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func isPhoneNumberRegistered(phoneNumber string) (bool, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHostname, DBPortNum, DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return false, err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE phone_number = ?)", phoneNumber).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
