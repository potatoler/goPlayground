package banking

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetClientProfile(db *sql.DB, user User) (Client, error) {
	var cli Client
	sql := fmt.Sprintf("SELECT * FROM Client WHERE id = '%s'", user.client_id)
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return Client{}, err
	}
	for rows.Next() {
		rows.Scan(&cli.Id, &cli.Name, &cli.Address, &cli.Telephone)
	}
	return cli, err
}

func Balance(db *sql.DB, user User) (float64, error) {
	var err error
	var balance float64
	sql := fmt.Sprintf("SELECT balance FROM Account WHERE card_number = '%s'", user.number)
	rows, err := db.Query(sql)
	if err != nil {
		return balance, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&balance)
	}
	return balance, err
}

func ChangeProfile(db *sql.DB, id, telephone, address string) error {
	sql := fmt.Sprintf("UPDATE Client SET telephone = '%s', address = '%s' WHERE id = '%s'", telephone, address, id)
	rows, err := db.Query(sql)
	defer rows.Close()
	return err
}
