package services

import (
	"akshu-21/finance-manager/datamdl"
	"akshu-21/finance-manager/model"
	"errors"

	"github.com/beevik/guid"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

//GetAllCategories : To get all Categories
func GetAllCategories() (interface{}, error) {

	data, err := datamdl.CategoriesDAO.FindData(bson.M{"isDeleted": false})
	if err != nil {
		log.Error("GetAllCategories : error_while_fetching_Categories")
		return nil, errors.New("GetAllCategories : error_while_fetching_Categories")
	}
	return data, nil
}

//AddNewCategory : To add new Category in DB
func AddNewCategory(data *model.Category) error {
	data.CategoryID = guid.New().String()
	data.IsDeleted = false
	err := datamdl.CategoriesDAO.Insert(data)
	if err != nil {
		log.Error("AddNewCategory : error_while_adding_Category")
		return errors.New("AddNewCategory : error_while_adding_Category")
	}
	return nil
}

//UpdateCategory : To add new Category in DB
func UpdateCategory(data *model.Category) error {

	err := datamdl.CategoriesDAO.Update(bson.M{"categoryId": data.CategoryID}, data)
	if err != nil {
		log.Error("UpdateCategory : error_while_updating_Category")
		return errors.New("UpdateCategory : error_while_updating_Category")
	}
	return nil
}

//DeleteCategory : To add new Category in DB
func DeleteCategory(data *model.Category) error {

	err := datamdl.CategoriesDAO.Update(bson.M{"categoryId": data.CategoryID}, bson.M{"isDeleted": true})
	if err != nil {
		log.Error("DeleteCategory : error_while_updating_Category")
		return errors.New("DeleteCategory : error_while_updating_Category")
	}
	return nil
}

//GetCategory : To get all Categorys
func GetCategory(request *model.Category) (interface{}, error) {
	var data interface{}
	response, err := datamdl.CategoriesDAO.FindData(bson.M{"categoryId": request.CategoryID})
	if err != nil {
		log.Error("GetAllCategorys : error_while_fetching_Categorys")
		return nil, errors.New("GetAllCategorys : error_while_fetching_Categorys")
	}
	if len(response) != 0 {
		data = response[0]
	}
	return data, nil
}
