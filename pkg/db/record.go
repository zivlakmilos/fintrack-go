package db

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Record struct {
	Id          string
	AccountId   string
	RecordId    string
	Debit       float64
	Credit      float64
	Description string
	Date        string
}

func CreateRecord(record Record, db *sql.DB) error {
	record.Id = fmt.Sprintf("%v", uuid.New())

	_, err := squirrel.Insert("records").Columns("id", "accountId", "recordId", "debit", "credit", "description", "date").
		Values(record.Id, record.AccountId, record.RecordId, record.Debit, record.Credit, record.Description, record.Date).
		RunWith(db).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func SelectRecords(query squirrel.SelectBuilder) ([]Record, error) {
	res, err := query.Columns("id", "accountId", "recordId", "debit", "credit", "description", "date").From("records").Query()
	if err != nil {
		return nil, err
	}

	records := []Record{}

	for res.Next() {
		record := Record{}

		err := res.Scan(
			&record.Id,
			&record.AccountId,
			&record.RecordId,
			&record.Debit,
			&record.Credit,
			&record.Description,
			&record.Date,
		)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
