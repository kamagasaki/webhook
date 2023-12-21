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

// FetchUserData retrieves user data from the db_bot table based on the phone number
func FetchUserData(phoneNumber string) (string, string, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHostname, DBPortNum, DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	var nama, email string
	err = db.QueryRow("SELECT nama, email FROM db_bot WHERE phone_number = ?", phoneNumber).Scan(&nama, &email)
	if err != nil {
		return "", "", err
	}

	return nama, email, nil
}
