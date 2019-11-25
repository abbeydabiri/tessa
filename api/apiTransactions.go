package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/justinas/alice"

	"golang.org/x/crypto/sha3"
	// "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"tessa/blockchain"
	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerTransactions(middlewares alice.Chain, router *Router) {
	router.Get("/api/transactions", middlewares.ThenFunc(apiTransactionsGet))
	router.Post("/api/transactions", middlewares.ThenFunc(apiTransactionsPost))
	router.Post("/api/transactions/search", middlewares.ThenFunc(apiTransactionsSearch))
	go apiJobCheckTransactions()
}

func apiTransactionsGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
		table.GetByID(table.ToMap(), formSearch)
		tableMap := table.ToMap()

		Token := database.Tokens{}
		config.Get().Postgres.Get(&Token, "select * from tokens where id = $1 limit 1", table.TokenID)
		tableMap["Token"] = Token

		objAccountToken := database.AccountTokens{}
		sqlAccountToken := "select * from accounttokens where tokenid = $1 and accountid = $2 limit 1"
		config.Get().Postgres.Get(&objAccountToken, sqlAccountToken, table.TokenID, table.AccountID)
		tableMap["AccountToken"] = objAccountToken

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTransactionsPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
		table.FillStruct(tableMap)

		message.Code = http.StatusInternalServerError
		if table.Title == "" {
			message.Message += "Title is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		var userID, walletID uint64
		if claims := utils.VerifyJWT(httpRes, httpReq); claims != nil {
			if claims["ID"] != nil {
				userID = uint64(claims["ID"].(float64))
			}

			if claims["WalletID"] != nil {
				walletID = uint64(claims["WalletID"].(float64))
			}
		}

		account := database.Accounts{}
		if table.WalletID == 0 || table.AccountID == 0 {
			config.Get().Postgres.Get(&account, "select * from accounts where userid = $1 and walletid = $2 limit 1", userID, walletID)

			if account.ID > uint64(0) {
				table.WalletID = walletID
				table.AccountID = account.ID
			}
		}

		if table.WalletID == 0 {
			message.Message += "Wallet ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.AccountID == 0 {
			message.Message += "Account ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.TokenID == 0 {
			message.Message += "Token ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		// amount := big.NewInt(0)
		// if amount.Cmp(table.Amount) == 0 {
		if table.Amount == 0 {
			message.Message += "Amount is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//Check if From Address is mine then it is a debit
		if table.FromAddress == "" {
			message.Message += "From Address is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//Check if To Address is mine then it is a credit
		if table.ToAddress == "" {
			message.Message += "To Address is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//From and To Address cannot be the same

		if table.ID == 0 {
			txHash, err := apiTransactionBroadcast(table)
			if err != nil {
				message.Message += "Transaction Broadcast Failed " + err.Error()
				json.NewEncoder(httpRes).Encode(message)
				return
			}
			tableMap["Reference"] = txHash
		}

		if table.ID == 0 {
			tableMap["Workflow"] = "pending"
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}

		// //Update Balance
		// sqlDebit := "select sum(amount) from transactions where fromaddress = $1 and tokenid = $2 and workflow = 'success'"
		// sqlCredit := "select sum(amount) from transactions where toaddress = $1 and tokenid = $2 and workflow = 'success'"
		// sqlUpdateBalance := "update accounttokens set balance = $1 where accountid = (select id from accounts where address = $2) and tokenid = $3"

		// fromDebit := float64(0)
		// config.Get().Postgres.Get(&fromDebit, sqlDebit, table.FromAddress, table.TokenID)

		// fromCredit := float64(0)
		// config.Get().Postgres.Get(&fromCredit, sqlCredit, table.FromAddress, table.TokenID)

		// fromBalance := fromCredit - fromDebit
		// if _, err := config.Get().Postgres.Exec(sqlUpdateBalance, fromBalance, table.FromAddress, table.TokenID); err != nil {
		// 	log.Println(err.Error())
		// }

		// toDebit := float64(0)
		// config.Get().Postgres.Get(&toDebit, sqlDebit, table.ToAddress, table.TokenID)

		// toCredit := float64(0)
		// config.Get().Postgres.Get(&toCredit, sqlCredit, table.ToAddress, table.TokenID)

		// toBalance := toCredit - toDebit
		// if _, err := config.Get().Postgres.Exec(sqlUpdateBalance, toBalance, table.ToAddress, table.TokenID); err != nil {
		// 	log.Println(err.Error())
		// }
		// //Update Balance

		message.Body = table.ID
		message.Code = http.StatusOK
		message.Message = "Transaction Broadcasted"
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTransactionsSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}

		sqlExtra := ""
		if formSearch.Filter["ToAddress"] != "" && formSearch.Filter["FromAddress"] != "" {
			sqlExtra = fmt.Sprintf("(toaddress = '%s' or fromaddress = '%s') and", formSearch.Filter["ToAddress"], formSearch.Filter["FromAddress"])
			delete(formSearch.Filter, "ToAddress")
			delete(formSearch.Filter, "FromAddress")
		}

		var searchList []interface{}
		searchResults := table.SearchExtra(table.ToMap(), formSearch, sqlExtra)
		for _, result := range searchResults {
			tableMap := result.ToMap()

			Token := database.Tokens{}
			config.Get().Postgres.Get(&Token, "select * from tokens where id = $1 limit 1", result.TokenID)
			tableMap["Token"] = Token

			objAccountToken := database.AccountTokens{}
			sqlAccountToken := "select * from accounttokens where tokenid = $1 and accountid = $2 limit 1"
			config.Get().Postgres.Get(&objAccountToken, sqlAccountToken, result.TokenID, result.AccountID)
			tableMap["AccountToken"] = objAccountToken

			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTransactionBroadcast(transaction database.Transactions) (txHash string, err error) {

	tokenAddressString := ""
	sqlAddress := "select address from tokens where id = $1 limit 1"
	config.Get().Postgres.Get(&tokenAddressString, sqlAddress, transaction.TokenID)

	if tokenAddressString == "" {
		err = errors.New("Token Address is required")
		log.Println(err.Error())
		return
	}

	if transaction.FromAddress == "" {
		err = errors.New("From Address is required")
		log.Println(err.Error())
		return
	}

	if transaction.ToAddress == "" {
		err = errors.New("To Address is required")
		log.Println(err.Error())
		return
	}

	if transaction.Amount <= float64(0.00) {
		err = errors.New("Amount is required")
		log.Println(err.Error())
		return
	}

	if blockchain.Client == nil {
		blockchain.EthClientDial(blockchain.InfuraNetwork)
	}

	ownerPrivateKey, ownerAddress := blockchain.EthGenerateKey(config.Get().Mnemonic, 1)
	nonce, errx := blockchain.Client.PendingNonceAt(context.Background(), ownerAddress)
	if errx != nil {
		err = errx
		log.Println(err.Error())
		return
	}

	value := big.NewInt(0)
	gasPrice, errx := blockchain.Client.SuggestGasPrice(context.Background())
	if errx != nil {
		err = errx
		log.Println(err.Error())
		return
	}

	var data []byte
	amount := new(big.Int)
	toAddress := common.HexToAddress(transaction.ToAddress)
	fromAddress := common.HexToAddress(transaction.FromAddress)
	tokenAddress := common.HexToAddress(tokenAddressString)

	fnSignatureString := ""
	if strings.ContainsAny(strings.ToLower(transaction.Code), "mint") {
		fnSignatureString = "mintTokens(address,uint256)"
		amount = new(big.Int).SetUint64(uint64(transaction.Amount))
	} else {
		fnSignatureString = "transferFrom(address,address,uint256)"
		amount = new(big.Int).SetUint64(uint64(transaction.Amount * float64(blockchain.WEI)))
	}

	paddedFromAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	fnSignature := []byte(fnSignatureString)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(fnSignature)
	methodID := hash.Sum(nil)[:4]

	if strings.ContainsAny(strings.ToLower(transaction.Code), "mint") {
		data = append(data, methodID...)
		data = append(data, paddedToAddress...)
		data = append(data, paddedAmount...)
	} else {
		data = append(data, methodID...)
		data = append(data, paddedFromAddress...)
		data = append(data, paddedToAddress...)
		data = append(data, paddedAmount...)
	}

	gasLimit := blockchain.ETHGasLimit // in units

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := blockchain.Client.NetworkID(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ownerPrivateKey)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = blockchain.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println(err.Error())
		return
	}

	txHash = signedTx.Hash().Hex()
	return
}
