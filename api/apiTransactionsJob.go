package api

import (
	"log"
	"time"
	"context"

	"github.com/ethereum/go-ethereum/common"

	"tessa/config"
	"tessa/database"
	"tessa/blockchain"
)

func apiJobCheckTransactions() {
	for {
		transactionList := []database.Transactions{}
		sqlSearch := "select * from transactions where workflow = 'pending' and reference != '' and fromaddress != '' and toaddress != '' limit 30"
		if err := config.Get().Postgres.Select(&transactionList, sqlSearch); err != nil {
			log.Println(err.Error())
		}

		//Loop through all eligible Users Transaction,
		for _, transaction := range transactionList {

			//Check if transaction is failed or successfull
			if blockchain.Client == nil {
				blockchain.EthClientDial(blockchain.InfuraNetwork)
			}

			txHash := common.HexToHash(transaction.Reference)
			if txReceipt, err := blockchain.Client.TransactionReceipt(context.Background(), txHash); err == nil {
				switch txReceipt.Status {
				case 0:
					transaction.Workflow = "fail"
				case 1:
					transaction.Workflow = "success"
				}	
				transaction.Update(transaction.ToMap())
				apiTransactionsUpdateBalance(transaction.ToAddress, transaction.FromAddress, transaction.TokenID)
			}
		}
		<-time.Tick(time.Second * 10)
	}
}

func apiTransactionsUpdateBalance(toAddress, fromAddress string, tokenid uint64) {
	if toAddress != "" && fromAddress != "" && tokenid != 0 {
		//Update Balance
		sqlDebit := "select sum(amount) from transactions where fromaddress = $1 and tokenid = $2 and workflow = 'success'"
		sqlCredit := "select sum(amount) from transactions where toaddress = $1 and tokenid = $2 and workflow = 'success'"
		sqlUpdateBalance := "update accounttokens set balance = $1 where accountid = (select id from accounts where address = $2) and tokenid = $3"
	
		fromDebit := float64(0)
		config.Get().Postgres.Get(&fromDebit, sqlDebit, fromAddress, tokenid)
	
		fromCredit := float64(0)
		config.Get().Postgres.Get(&fromCredit, sqlCredit, fromAddress, tokenid)
	
		fromBalance := fromCredit - fromDebit
		if _, err := config.Get().Postgres.Exec(sqlUpdateBalance, fromBalance, fromAddress, tokenid); err != nil {
			log.Println(err.Error())
		}
	
		toDebit := float64(0)
		config.Get().Postgres.Get(&toDebit, sqlDebit, toAddress, tokenid)
	
		toCredit := float64(0)
		config.Get().Postgres.Get(&toCredit, sqlCredit, toAddress, tokenid)
	
		toBalance := toCredit - toDebit
		if _, err := config.Get().Postgres.Exec(sqlUpdateBalance, toBalance, toAddress, tokenid); err != nil {
			log.Println(err.Error())
		}
		//Update Balance
	}
}
