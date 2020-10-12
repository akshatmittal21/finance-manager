package services

import (
	"akshu-21/finance-manager/datamdl"
	"akshu-21/finance-manager/model"

	"errors"
	"time"

	"github.com/beevik/guid"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

//AddNewAccount : To add new account in DB
func AddNewAccount(data *model.Account) error {
	data.AccountID = guid.New().String()
	data.IsDeleted = false
	data.OpeningDate = float64(time.Now().Unix())
	err := datamdl.AccountsDAO.Insert(data)
	if err != nil {
		log.Error("AddNewAccount : error_while_adding_account")
		return errors.New("AddNewAccount : error_while_adding_account")
	}
	return nil
}

//UpdateAccount : To add new account in DB
func UpdateAccount(data *model.Account) error {

	err := datamdl.AccountsDAO.Update(bson.M{"accountId": data.AccountID}, data)
	if err != nil {
		log.Error("UpdateAccount : error_while_updating_account")
		return errors.New("UpdateAccount : error_while_updating_account")
	}
	return nil
}

//DeleteAccount : To add new account in DB
func DeleteAccount(data *model.Account) error {

	err := datamdl.AccountsDAO.Update(bson.M{"accountId": data.AccountID}, bson.M{"isDeleted": true})
	if err != nil {
		log.Error("DeleteAccount : error_while_updating_account")
		return errors.New("DeleteAccount : error_while_updating_account")
	}
	return nil
}

//GetAllAccounts : To get all accounts
func GetAllAccounts() (interface{}, error) {

	data, err := datamdl.AccountsDAO.FindData(bson.M{"isDeleted": false})
	if err != nil {
		log.Error("GetAllAccounts : error_while_fetching_accounts")
		return nil, errors.New("GetAllAccounts : error_while_fetching_accounts")
	}
	return data, nil
}

//GetAccount : To get all accounts
func GetAccount(request *model.Account) (interface{}, error) {
	var data interface{}
	response, err := datamdl.AccountsDAO.FindData(bson.M{"accountId": request.AccountID})
	if err != nil {
		log.Error("GetAllAccounts : error_while_fetching_accounts")
		return nil, errors.New("GetAllAccounts : error_while_fetching_accounts")
	}
	if len(response) != 0 {
		data = response[0]
	}
	return data, nil
}
