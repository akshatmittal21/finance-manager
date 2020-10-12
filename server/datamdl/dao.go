package datamdl

var (
	//AccountsDAO - For DB Operation
	AccountsDAO *MongoDAO
	//CategoriesDAO - For DB Operation
	CategoriesDAO *MongoDAO
	//TransactionsDAO - For DB Operation
	TransactionsDAO *MongoDAO
)

//InitDAO - Associate DAOs with respective collections
//Create new DAO here for new collection
func InitDAO() {
	AccountsDAO = GetMongoDAO("accounts")
	CategoriesDAO = GetMongoDAO("categories")
	TransactionsDAO = GetMongoDAO("transactions")
}
