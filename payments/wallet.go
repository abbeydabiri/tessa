package payments

import (
	"tessa/database"
)

func DebitWallet(WalletID int64, Amount float64) (payment *Payment) {
	payment = new(Payment)
	payment.Success = true
	return

	formSearch := database.SearchParams{}

	formSearch.ID = WalletID
	wallet := database.Wallets{}
	wallet.GetByID(wallet.ToMap(), &formSearch)
	if wallet.ID == 0 {
		return
	}

	formSearch.ID = wallet.AccountID
	account := database.Accounts{}
	account.GetByID(account.ToMap(), &formSearch)
	if account.ID == 0 {
		return
	}

	if Amount > account.Balance {
		return
	}

	payment.Success = true
	payment.Amount = Amount
	payment.Currency = account.Currency

	account.Balance -= Amount
	accountMap := make(map[string]interface{})
	accountMap["ID"] = account.ID
	accountMap["Balance"] = account.Balance
	account.Update(accountMap)

	//Create Transaction Record
	// Agent Deposit
	// - - - - - - -
	// 		Asset: Debit
	// Liability: Credit
	//
	// Agent Sales
	// - - - - - -
	// 		Asset: Credit
	// Liability: Debit
	//
	//
	// Agent Transfer
	// - - - - - -
	// Liability: Debit
	// Liability: Credit
	//	else do not make payment

	return
}
