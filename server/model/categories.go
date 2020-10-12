package model

//Category struct
type Category struct {
	CategoryID    string `json:"categoryId" bson:"categoryId"`
	CategoryGroup string `json:"categoryGroup,omitempty" bson:"categoryGroup,omitempty"`
	CategoryType  string `json:"categoryType,omitempty" bson:"categoryType,omitempty"`
	CategoryName  string `json:"categoryName,omitempty" bson:"categoryName,omitempty"`
	IsDeleted     bool   `json:"isDeleted,omitempty" bson:"isDeleted,omitempty"`
}
