package db

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type AccountType string

const (
	AccountTypeAsset     AccountType = "asset"
	AccountTypeLiability AccountType = "liability"
	AccountTypeIncome    AccountType = "income"
	AccountTypeExpense   AccountType = "expense"
)

func Open(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(db *sql.DB) error {
	err := execQuery(db, "CREATE TABLE IF NOT EXISTS accounts (id TEXT PRIMARY KEY, code TEXT, title TEXT, type TEXT)")
	if err != nil {
		return err
	}

	err = execQuery(db, "CREATE TABLE IF NOT EXISTS records (id TEXT PRIMARY KEY, accountId TEXT, credit REAL, debit REAL, description TEXT)")
	if err != nil {
		return err
	}

	err = createAccounts(db)
	if err != nil {
		return err
	}

	return nil
}

func execQuery(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func createAccounts(db *sql.DB) error {
	res, err := squirrel.Select("COUNT(*)").From("accounts").RunWith(db).Query()
	if err != nil {
		return err
	}

	var rows int

	if res.Next() {
		err := res.Scan(&rows)
		if err != nil {
			return err
		}
	}

	err = res.Close()
	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	_, err = squirrel.Insert("accounts").
		Columns("id", "code", "title", "type").
		Values(uuid.New(), "001", "Bank Account", AccountTypeAsset).
		Values(uuid.New(), "101", "Loan", AccountTypeLiability).
		Values(uuid.New(), "201", "Salary", AccountTypeIncome).
		Values(uuid.New(), "202", "Freelance", AccountTypeIncome).
		Values(uuid.New(), "301", "Rent", AccountTypeExpense).
		Values(uuid.New(), "302", "Food", AccountTypeExpense).
		Values(uuid.New(), "303", "Entertainment", AccountTypeExpense).
		RunWith(db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}
