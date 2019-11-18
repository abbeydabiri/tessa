package api

import (
	"context"
	"errors"

	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/justinas/alice"

	"math/big"
	"tessa/config"
	"tessa/database"

	"tessa/blockchain"
	"tessa/smarttoken"
)

func apiHandlerTokens(middlewares alice.Chain, router *Router) {
	router.Get("/api/tokens", middlewares.ThenFunc(apiTokensGet))
	router.Post("/api/tokens", middlewares.ThenFunc(apiTokensPost))
	router.Post("/api/tokens/search", middlewares.ThenFunc(apiTokensSearch))
}

func apiTokensGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTokensPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
		table.FillStruct(tableMap)

		//Check if Token name or symbol is already in use
		if table.ID == 0 {
			numTokens := 0
			sqlCheck := "select count(id) from tokens where lower(title) = lower($1) or lower(symbol) = lower($2)"
			err := config.Get().Postgres.Get(&numTokens, sqlCheck, table.Title, table.Symbol)
			if err != nil {
				log.Println(err.Error())
			}
			if numTokens > 0 {
				message.Message += "Token Name or Symbol already exists!!! \n"
				message.Code = http.StatusInternalServerError
				json.NewEncoder(httpRes).Encode(message)
				return
			}
		}
		//Check if Token name or symbol is already in use

		message.Code = http.StatusInternalServerError
		if table.Company == "" {
			message.Message += "Company is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.RC == "" {
			message.Message += "Company RC is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Title == "" {
			message.Message += "Token Title is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Symbol == "" {
			message.Message += "Token Symbol is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Project == "" {
			message.Message += "Token Project is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.ProjectCost < 1.00 {
			message.Message += "Token Project Cost is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.MaxTotalSupply < 1 {
			message.Message += "Token Total Supply is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.ID == 0 {

			tableMap["Code"] = ""
			tableMap["Address"] = ""
			deployedToken, err := apiTokenDeploy(table.Symbol, table.Title, table.MaxTotalSupply, table.Seed)
			if err == nil {
				tableMap["Code"] = deployedToken["transaction"]
				tableMap["Address"] = deployedToken["address"]
			}

			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Code = http.StatusOK
		message.Message = "Token Created and Smart Contract Deployed to " + tableMap["Address"].(string)
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTokensSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}
		var searchList []interface{}
		searchResults := table.Search(table.ToMap(), formSearch)
		for _, result := range searchResults {
			tableMap := result.ToMap()
			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTokenDeploy(Symbol, Name string, maxtotalsupply, seed float64) (token map[string]string, err error) {

	token = make(map[string]string)
	if Symbol == "" {
		err = errors.New("Symbol is required")
		return
	}

	if Name == "" {
		err = errors.New("Name is required")
		return
	}

	if maxtotalsupply <= float64(0) {
		err = errors.New("Max Total Supply is required")
		return
	}

	Seed := new(big.Int).SetFloat64(seed)
	MaxTotalSupply := new(big.Int).SetFloat64(maxtotalsupply)

	if blockchain.Client == nil {
		blockchain.EthClientDial(blockchain.InfuraNetwork)
	}

	privateKey, fromAddress := blockchain.EthGenerateKey(config.Get().Mnemonic, 1)
	nonce, errx := blockchain.Client.PendingNonceAt(context.Background(), fromAddress)
	if errx != nil {
		err = errx
		return
	}

	gasPriceTemp, errx := blockchain.Client.SuggestGasPrice(context.Background())
	if errx != nil {
		err = errx
		return
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6009850) // in units
	gasPrice := big.NewInt(1)
	auth.GasPrice = gasPrice.Mul(gasPrice, gasPriceTemp)

	// // address, tx, instance, err := smartcontracts.DeploySmartcontracts(auth, client, Symbol, Name, MaxTotalSupply, Seed)
	address, tx, _, err := smarttoken.DeploySmartToken(auth, blockchain.Client, Symbol, Name, MaxTotalSupply, Seed)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	token["address"] = address.Hex()
	token["transaction"] = tx.Hash().Hex()

	return
}
