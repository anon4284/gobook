package entry

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Init Create database table
func (e *Entry) Init() {
	msg, err := e.createTable()
	if err != nil {
		log.Print(err)
	} else {
		log.Print(msg)
	}
}

//Add adds an entry to the database
func (e *Entry) Add() {

}

//Scan all of the table content
func (e *Entry) Scan() (resp *dynamodb.ScanOutput, err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String("gobook-entries"),
		AttributesToGet: []*string{
			aws.String("Name"),
			aws.String("Content"),
		},
	}
	return e.svc.Scan(params)
}
