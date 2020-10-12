package model

//Account struct
type Account struct {
	AccountID       string       `json:"accountId" bson:"accountId"`
	AccountInfo     *AccountInfo `json:"accountInfo,omitempty" bson:"accountInfo,omitempty"`
	OpeningDate     float64      `json:"openingDate,omitempty" bson:"openingDate,omitempty"`
	StartingBalance float64      `json:"startingBalance,omitempty" bson:"startingBalance,omitempty"`
	IsDeleted       bool         `json:"isDeleted" bson:"isDeleted"`
	LatestBalance   float64      `json:"latestBalance,omitempty" bson:"latestBalance,omitempty"`
}

//AccountInfo struct
type AccountInfo struct {
	AccountGroup string `json:"accountGroup,omitempty" bson:"accountGroup,omitempty"`
	AccountType  string `json:"accountType,omitempty" bson:"accountType,omitempty"`
	AccountName  string `json:"accountName,omitempty" bson:"accountName,omitempty"`
}
