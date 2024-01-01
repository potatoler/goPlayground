package banking

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Client struct {
	Id        string
	Name      string
	Address   string
	Telephone string
}

type User struct {
	client_id string
	number    string
	name      string
	isAdmin   bool
}

func (u *User) ID() string {
	return u.client_id
}

func (u *User) Account() string {
	return u.number
}

func (u *User) Name() string {
	return u.name
}

func (u *User) IsAdmin() bool {
	return u.isAdmin
}

func Login(db *sql.DB, number string, password string) (User, error) {
	var err error
	adminFlag := false
	//Query if account exist
	sql := fmt.Sprintf("select card_number, passwd, client_id, blocked from Account where card_number='%s'", number)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	var accNumber, accPasswd, accClient string
	var blocked int
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&accNumber, &accPasswd, &accClient, &blocked)
	}
	if accNumber == "" {
		adminFlag = true
		sql := fmt.Sprintf("select serial, passwd from Clerk where serial='%s'", number)
		rows, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&accNumber, &accPasswd)
		}
	}
	if accNumber == "" {
		err = errors.New("login: account not found")
		return User{}, err
	} else if blocked == -1 {
		err = errors.New("login: account blocked")
		return User{}, err
	} else if accPasswd != password {
		if !adminFlag {
			sql := fmt.Sprintf("UPDATE Account SET blocked = blocked + 1 WHERE card_number = '%s'", number)
			rows, err := db.Query(sql)
			defer rows.Close()
			if err != nil {
				panic(err.Error())
			}
		}
		err = errors.New("login: wrong password")
		return User{}, err
	} else {
		var user User
		user.isAdmin = adminFlag
		user.number = accNumber
		if !user.isAdmin {
			user.client_id = accClient
			sql := fmt.Sprintf("SELECT name FROM Client WHERE id = '%s'", user.client_id)
			rows, err := db.Query(sql)
			if err != nil {
				panic(err.Error())
			}
			for rows.Next() {
				rows.Scan(&user.name)
			}
			tailNumber := user.Account()
			tailNumber = "..." + tailNumber[len(tailNumber)-4:]
			fmt.Printf("Welcome! %s(%s)\n", user.name, tailNumber)
			return user, err
		} else {
			sql := fmt.Sprintf("SELECT name FROM Clerk WHERE serial = '%s'", user.Account())
			rows, er := db.Query(sql)
			if er != nil {
				panic(err.Error())
			}
			for rows.Next() {
				rows.Scan(&user.name)
			}
			fmt.Printf("Ready to work, %s(%s)\n", user.name, user.Account())
			return user, err
		}
	}
}

func GetUserTitleString(db *sql.DB, user User) string {
	var userTitle string
	if !user.isAdmin {
		sql := fmt.Sprintf("SELECT name FROM Client WHERE id = '%s'", user.client_id)
		rows, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			rows.Scan(&user.name)
		}
		tailNumber := user.Account()
		tailNumber = "..." + tailNumber[len(tailNumber)-4:]
		userTitle = fmt.Sprintf("%s(%s)", user.name, tailNumber)
	} else {
		sql := fmt.Sprintf("SELECT name FROM Clerk WHERE serial = '%s'", user.Account())
		rows, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			rows.Scan(&user.name)
		}
		userTitle = fmt.Sprintf("%s(%s)", user.name, user.Account())
	}
	return userTitle
}

func GetUserTitleStringPair(db *sql.DB, user User) (string, string) {
	if !user.isAdmin {
		sql := fmt.Sprintf("SELECT name FROM Client WHERE id = '%s'", user.client_id)
		rows, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			rows.Scan(&user.name)
		}
		tailNumber := user.Account()
		tailNumber = "..." + tailNumber[len(tailNumber)-4:]
		return user.name, tailNumber
	} else {
		sql := fmt.Sprintf("SELECT name FROM Clerk WHERE serial = '%s'", user.Account())
		rows, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			rows.Scan(&user.name)
		}
		return user.name, user.Account()
	}
}

func Unblock(db *sql.DB, account, password, id string) error {
	sql := fmt.Sprintf("SELECT passwd, client_id FROM Account WHERE card_number = '%s'", account)
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	var realPassword, realID string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&realPassword, &realID)
	}
	if realPassword == password && realID == id {
		sql = fmt.Sprintf("UPDATE Account SET blocked = blocked + 1 WHERE card_number = '%s'", account)
		_, err = db.Query(sql)
		return err
	} else {
		return errors.New("Info not match")
	}
}
