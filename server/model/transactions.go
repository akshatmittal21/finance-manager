package model

//Transaction struct
type Transaction struct {
	TransactionID              string                  `json:"transactionId" bson:"transactionId"`
	TransactionName            string                  `json:"transactionName,omitempty" bson:"transactionName,omitempty"`
	TransactionDate            float64                 `json:"transactionDate,omitempty" bson:"transactionDate,omitempty"`
	Amount                     float64                 `json:"amount,omitempty" bson:"amount,omitempty"`
	Tags                       []string                `json:"tags,omitempty" bson:"tags,omitempty"`
	TransactionType            string                  `json:"transactionType,omitempty" bson:"transactionType,omitempty"`
	IsAmountReceived           bool                    `json:"isAmountReceived" bson:"isAmountReceived"`
	IsDeleted                  bool                    `json:"isDeleted" bson:"isDeleted"`
	ReceiptPath                string                  `json:"receiptPath,omitempty" bson:"receiptPath,omitempty"`
	TransactionAccountInfo     *TransactionAccountInfo `json:"accountInfo,omitempty" bson:"accountInfo,omitempty"`
	TransferredToAccountInfo   *TransactionAccountInfo `json:"transferredToAccountInfo,omitempty" bson:"transferredToAccountInfo,omitempty"`
	TransferredFromAccountInfo *TransactionAccountInfo `json:"transferredFromAccountInfo,omitempty" bson:"transferredFromAccountInfo,omitempty"`
	CategoryInfo               *CategoryInfo           `json:"categoryInfo,omitempty" bson:"categoryInfo,omitempty"`
}

// TransactionAccountInfo struct
type TransactionAccountInfo struct {
	AccountID   string `json:"accountId,omitempty" bson:"accountId,omitempty"`
	AccountType string `json:"accountType,omitempty" bson:"accountType,omitempty"`
	AccountName string `json:"accountName,omitempty" bson:"accountName,omitempty"`
}

// CategoryInfo struct
type CategoryInfo struct {
	CategoryID   string `json:"categoryId,omitempty" bson:"categoryId,omitempty"`
	CategoryName string `json:"categoryName,omitempty" bson:"categoryName,omitempty"`
}
