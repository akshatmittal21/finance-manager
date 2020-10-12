package services

import (
	"akshu-21/finance-manager/datamdl"
	"akshu-21/finance-manager/model"
	"encoding/json"
	"errors"
	"github.com/beevik/guid"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//AddNewTransaction : To add new Transaction in DB
func AddNewTransaction(data *model.Transaction) (string, error) {
	data.TransactionID = guid.New().String()
	data.IsDeleted = false
	data.TransactionDate = float64(time.Now().Unix())

	err := datamdl.TransactionsDAO.Insert(data)
	if err != nil {
		log.Error("AddNewTransaction : error_while_adding_Transaction")
		return "", errors.New("AddNewTransaction : error_while_adding_Transaction")
	}
	err = updateLatestBalance(data.TransactionAccountInfo.AccountID, data.Amount, data.TransactionType, data.IsAmountReceived)
	if err != nil {
		log.Error("AddNewTransaction : error_while_updating_latest_balance")
		return "", errors.New("AddNewTransaction : error_while_updating_latest_balance")
	}

	return data.TransactionID, nil
}

//UpdateTransaction : To add new Transaction in DB
func UpdateTransaction(data *model.Transaction) error {

	transaction := model.Transaction{}
	if data.Amount != 0 {
		transactionDetails, err := GetTransaction(data)
		if err != nil {
			log.Error("UpdateTransaction: unable_to_fetch_Transaction")
			return errors.New("UpdateTransaction: unable_to_fetch_Transaction")
		}
		byteArray, err := json.Marshal(transactionDetails)
		if err != nil {
			log.Error("updateLatestBalance : error_while_marshalling")
			return errors.New("updateLatestBalance : error_while_marshalling")
		}
		json.Unmarshal(byteArray, &transaction)
	}
	var err error
	// Rollback previous amount
	if data.Amount != 0 {
		err = updateLatestBalance(transaction.TransactionAccountInfo.AccountID, transaction.Amount, "INCOME", true)
		if err != nil {
			log.Error("UpdateTransaction : error_while_updating_latest_balance")
			return errors.New("UpdateTransaction : error_while_updating_latest_balance")
		}
	}
	// Update transaction
	err = datamdl.TransactionsDAO.Update(bson.M{"transactionId": data.TransactionID}, data)
	if err != nil {
		log.Error("UpdateTransaction : error_while_updating_Transaction")
		return errors.New("UpdateTransaction : error_while_updating_Transaction")
	}
	// Update new amount
	if data.Amount != 0 {
		err = updateLatestBalance(transaction.TransactionAccountInfo.AccountID, data.Amount, data.TransactionType, data.IsAmountReceived)
		if err != nil {
			log.Error("UpdateTransaction : error_while_updating_latest_balance")
			return errors.New("UpdateTransaction : error_while_updating_latest_balance")
		}
	}
	return nil
}

//DeleteTransaction : To add new Transaction in DB
func DeleteTransaction(data *model.Transaction) error {

	err := datamdl.TransactionsDAO.Update(bson.M{"transactionId": data.TransactionID}, bson.M{"isDeleted": true})
	if err != nil {
		log.Error("DeleteTransaction : error_while_updating_Transaction")
		return errors.New("DeleteTransaction : error_while_updating_Transaction")
	}
	transactionDetails, err := GetTransaction(data)
	if err != nil {
		log.Error("DeleteTransaction: unable_to_fetch_Transaction")
		return errors.New("DeleteTransaction: unable_to_fetch_Transaction")
	}
	byteArray, err := json.Marshal(transactionDetails)
	if err != nil {
		log.Error("DeleteTransaction : error_while_marshalling")
		return errors.New("DeleteTransaction : error_while_marshalling")
	}
	transaction := model.Transaction{}
	json.Unmarshal(byteArray, &transaction)
	err = updateLatestBalance(transaction.TransactionAccountInfo.AccountID, transaction.Amount, "INCOME", true)
	if err != nil {
		log.Error("DeleteTransaction : error_while_updating_latest_balance")
		return errors.New("DeleteTransaction : error_while_updating_latest_balance")
	}
	return nil
}

//GetAllTransactions : To get all Transactions
func GetAllTransactions() (interface{}, error) {

	data, err := datamdl.TransactionsDAO.FindData(bson.M{"isDeleted": false})
	if err != nil {
		log.Error("GetAllTransactions : error_while_fetching_Transactions")
		return nil, errors.New("GetAllTransactions : error_while_fetching_Transactions")
	}
	return data, nil
}

//GetTransaction : To get all Transactions
func GetTransaction(request *model.Transaction) (interface{}, error) {
	var data interface{}
	response, err := datamdl.TransactionsDAO.FindData(bson.M{"transactionId": request.TransactionID})
	if err != nil {
		log.Error("GetAllTransactions : error_while_fetching_Transactions")
		return nil, errors.New("GetAllTransactions : error_while_fetching_Transactions")
	}
	if len(response) != 0 {
		data = response[0]
	}
	return data, nil
}

//updateLatestBalance : to update balance for account
func updateLatestBalance(accountID string, amount float64, transactionType string, isAmountRecieved bool) error {
	var accountInfo interface{}
	response, err := datamdl.AccountsDAO.FindData(bson.M{"accountId": accountID})
	if err != nil {
		log.Error("updateLatestBalance : error_while_fetching_account")
		return errors.New("updateLatestBalance : error_while_fetching_account")
	}
	if len(response) != 0 {
		accountInfo = response[0]
	}
	byteArray, err := json.Marshal(accountInfo)
	if err != nil {
		log.Error("updateLatestBalance : error_while_marshalling")
		return errors.New("updateLatestBalance : error_while_marshalling")
	}
	account := model.Account{}
	json.Unmarshal(byteArray, &account)
	latestBalance := account.LatestBalance
	if transactionType == "EXPENSE" {
		latestBalance = account.LatestBalance - amount
	} else if transactionType == "INCOME" {
		latestBalance = account.LatestBalance + amount
	} else if transactionType == "TRANSFER" {
		if isAmountRecieved {
			latestBalance = account.LatestBalance + amount
		} else {
			latestBalance = account.LatestBalance - amount
		}
	}
	err = datamdl.AccountsDAO.Update(bson.M{"accountId": accountID}, bson.M{"latestBalance": latestBalance})
	if err != nil {
		log.Error("UpdateAccount : error_while_updating_account")
		return errors.New("UpdateAccount : error_while_updating_account")
	}
	return nil
}
