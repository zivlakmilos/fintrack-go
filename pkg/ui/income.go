package ui

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rivo/tview"
	"github.com/zivlakmilos/fintrack-go/pkg/db"
)

func (u *Ui) createIncomeScreen() tview.Primitive {
	incomeAccounts := loadIncomeScreenAccounts(db.AccountTypeIncome, u.db)
	receiveAccounts := loadIncomeScreenAccounts(db.AccountTypeAsset, u.db)

	form := tview.NewForm().
		AddDropDown("Income", incomeAccounts, 0, nil).
		AddDropDown("Account", receiveAccounts, 0, nil).
		AddInputField("Amount", "0.00", 25, nil, nil).
		AddButton("Save", nil).
		AddButton("Cancel", func() { u.showPage(0) })

	setFormColors(form)

	form.SetBorder(true)
	form.SetTitle(" New Income ")

	return center(35, 12, form)
}

func loadIncomeScreenAccounts(accountType db.AccountType, con *sql.DB) []string {
	accounts, err := db.SelectAccounts(squirrel.Select().Where("type=?", accountType).OrderBy("code").RunWith(con))
	if err != nil {
		return nil
	}

	items := []string{}
	for _, acc := range accounts {
		items = append(items, fmt.Sprintf("%s - %s", acc.Code, acc.Title))
	}

	return items
}
