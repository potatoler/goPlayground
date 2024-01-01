package banking

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func Remittance(db *sql.DB, clerk User, src, dst, amount, password string) error {
	sql := fmt.Sprintf("EXEC MakeRemittance '%s', '%s', %s, '%s', '%s'", src, dst, amount, password, clerk.Account())
	res, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer res.Close()
	var status int
	for res.Next() {
		res.Scan(&status)
	}
	if status == 1 {
		return errors.New("Failed")
	} else {
		return nil
	}
}

func Deposit(db *sql.DB, clerk User, account, amount, password string) error {
	sql := fmt.Sprintf("EXEC DepositWithdraw '%s', '%s', %s, '%s'", account, password, amount, clerk.Account())
	res, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer res.Close()
	var status int
	for res.Next() {
		res.Scan(&status)
	}
	if status == 1 {
		return errors.New("Failed")
	} else {
		return nil
	}
}

func Withdraw(db *sql.DB, clerk User, account, amount, password string) error {
	sql := fmt.Sprintf("EXEC DepositWithdraw '%s', '%s', -%s, '%s'", account, password, amount, clerk.Account())
	res, err := db.Query(sql)
	if err != nil {
		return err
	}
	var status int
	for res.Next() {
		res.Scan(&status)
	}
	if status == 1 {
		return errors.New("Failed")
	} else {
		return nil
	}
}

func Register(db *sql.DB, clerk User, name, id, address, telephone, password, comfirm string) (string, error) {
	if password != comfirm {
		return "", errors.New("password mismatch")
	}
	sql := fmt.Sprintf("EXEC OpenAccount '%s', '%s', '%s', '%s', '%s'", name, id, address, telephone, password)
	res, err := db.Query(sql)
	if err != nil {
		return "", err
	}
	var newCard string
	for res.Next() {
		res.Scan(&newCard)
	}
	return newCard, nil
}
