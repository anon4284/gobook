package entry

import "github.com/aws/aws-sdk-go/service/dynamodb"

//Entry class includes routes and db connection
type Entry struct {
	svc *dynamodb.DynamoDB
}

//New create new entry class
func New(svc dynamodb.DynamoDB) *Entry {
	return &Entry{svc: &svc}
}
