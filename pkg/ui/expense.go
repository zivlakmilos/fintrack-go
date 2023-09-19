package ui

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	"github.com/zivlakmilos/fintrack-go/pkg/db"
)

type ExpenseScreen struct {
	date        string
	expense     string
	account     string
	amount      string
	description string
}

func (u *Ui) createExpenseScreen() tview.Primitive {
	expenseAccounts, expenseIds := loadIncomeScreenAccounts(db.AccountTypeIncome, u.db)
	receiveAccounts, receiveIds := loadIncomeScreenAccounts(db.AccountTypeAsset, u.db)

	expenseScreen := ExpenseScreen{
		date:   time.Now().Format("2006-01-02"),
		amount: "0.00",
	}
	if len(expenseIds) > 0 {
		expenseScreen.expense = expenseIds[0]
	}
	if len(receiveIds) > 0 {
		expenseScreen.account = receiveIds[0]
	}

	form := tview.NewForm().
		AddInputField("Date", expenseScreen.date, 25, nil, func(text string) { expenseScreen.amount = text }).
		AddDropDown("Expense", expenseAccounts, 0, func(option string, optionIndex int) {
			if optionIndex >= 0 && optionIndex < len(expenseIds) {
				expenseScreen.expense = expenseIds[optionIndex]
			}
		}).
		AddDropDown("Account", receiveAccounts, 0, func(option string, optionIndex int) {
			if optionIndex >= 0 && optionIndex < len(receiveIds) {
				expenseScreen.account = receiveIds[optionIndex]
			}
		}).
		AddInputField("Amount", "0.00", 25, nil, func(text string) { expenseScreen.amount = text }).
		AddTextArea("Description", "", 25, 3, 255, func(text string) { expenseScreen.description = text }).
		AddButton("Save", func() { u.saveExpenseHandler(&expenseScreen) }).
		AddButton("Cancel", func() { u.showPage(0) })

	setFormColors(form)

	form.SetBorder(true)
	form.SetTitle(" New Expense ")

	return center(35, 17, form)
}

func loadExpenseScreenAccounts(accountType db.AccountType, con *sql.DB) ([]string, []string) {
	accounts, err := db.SelectAccounts(squirrel.Select().Where("type=?", accountType).OrderBy("code").RunWith(con))
	if err != nil {
		return nil, nil
	}

	items := []string{}
	ids := []string{}

	for _, acc := range accounts {
		items = append(items, fmt.Sprintf("%s - %s", acc.Code, acc.Title))
		ids = append(ids, acc.Id)
	}

	return items, ids
}

func (u *Ui) saveExpenseHandler(data *ExpenseScreen) {
	amount, err := strconv.ParseFloat(data.amount, 64)
	if err != nil {
		return
	}

	recordId := fmt.Sprintf("%v", uuid.New())

	recordIncome := db.Record{
		AccountId:   data.expense,
		RecordId:    recordId,
		Debit:       amount,
		Credit:      0,
		Description: data.description,
		Date:        data.date,
	}

	recordAccount := db.Record{
		AccountId:   data.account,
		RecordId:    recordId,
		Debit:       0,
		Credit:      amount,
		Description: data.description,
		Date:        data.date,
	}

	err = db.CreateRecord(recordAccount, u.db)
	if err != nil {
		return
	}

	err = db.CreateRecord(recordIncome, u.db)
	if err != nil {
		return
	}

	u.showPage(0)
}
