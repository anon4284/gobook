package entry

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
)

//Entry class includes routes and db connection
type Entry struct {
	router *mux.Router
	svc    *dynamodb.DynamoDB
}

//New create new entry class
func New(router mux.Router, svc dynamodb.DynamoDB) *Entry {
	return &Entry{router: &router, svc: &svc}
}
