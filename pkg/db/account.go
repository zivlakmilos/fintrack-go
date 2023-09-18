package db

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Account struct {
	Id    string
	Code  string
	Title string
	Type  AccountType
}

func CreateAccouont(code string, title string, accType string, db *sql.DB) error {
	id := uuid.New()

	_, err := squirrel.Insert("accounts").
		Columns("id", "code", "title", "type").
		Values(id, code, title, accType).
		RunWith(db).Exec()

	if err != nil {
		return err
	}

	return nil
}

func SelectAccounts(query squirrel.SelectBuilder) ([]Account, error) {
	res, err := query.Columns("id", "code", "title", "type").From("accounts").Query()
	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	for res.Next() {
		acc := Account{}

		err := res.Scan(
			&acc.Id,
			&acc.Code,
			&acc.Title,
			&acc.Type,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	return accounts, nil
}
