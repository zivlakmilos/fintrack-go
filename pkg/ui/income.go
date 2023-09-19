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

type IncomeScreen struct {
	date        string
	income      string
	account     string
	amount      string
	description string
}

func (u *Ui) createIncomeScreen() tview.Primitive {
	incomeAccounts, incomeIds := loadIncomeScreenAccounts(db.AccountTypeIncome, u.db)
	receiveAccounts, receiveIds := loadIncomeScreenAccounts(db.AccountTypeAsset, u.db)

	incomeScreen := IncomeScreen{
		date:   time.Now().Format("2006-01-02"),
		amount: "0.00",
	}
	if len(incomeIds) > 0 {
		incomeScreen.income = incomeIds[0]
	}
	if len(receiveIds) > 0 {
		incomeScreen.account = receiveIds[0]
	}

	form := tview.NewForm().
		AddInputField("Date", incomeScreen.date, 25, nil, func(text string) { incomeScreen.amount = text }).
		AddDropDown("Income", incomeAccounts, 0, func(option string, optionIndex int) {
			if optionIndex >= 0 && optionIndex < len(incomeIds) {
				incomeScreen.income = incomeIds[optionIndex]
			}
		}).
		AddDropDown("Account", receiveAccounts, 0, func(option string, optionIndex int) {
			if optionIndex >= 0 && optionIndex < len(receiveIds) {
				incomeScreen.account = receiveIds[optionIndex]
			}
		}).
		AddInputField("Amount", "0.00", 25, nil, func(text string) { incomeScreen.amount = text }).
		AddTextArea("Description", "", 25, 3, 255, func(text string) { incomeScreen.description = text }).
		AddButton("Save", func() { u.saveIncomeHandler(&incomeScreen) }).
		AddButton("Cancel", func() { u.showPage(0) })

	setFormColors(form)

	form.SetBorder(true)
	form.SetTitle(" New Income ")

	return center(35, 17, form)
}

func loadIncomeScreenAccounts(accountType db.AccountType, con *sql.DB) ([]string, []string) {
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

func (u *Ui) saveIncomeHandler(data *IncomeScreen) {
	amount, err := strconv.ParseFloat(data.amount, 64)
	if err != nil {
		return
	}

	recordId := fmt.Sprintf("%v", uuid.New())

	recordAccount := db.Record{
		AccountId:   data.account,
		RecordId:    recordId,
		Debit:       amount,
		Credit:      0,
		Description: data.description,
		Date:        data.date,
	}

	recordIncome := db.Record{
		AccountId:   data.income,
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
