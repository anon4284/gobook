package entry

import (
	"log"
	"strconv"
	"time"

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
func (e *Entry) Add(email string, name string, content string) (*dynamodb.PutItemOutput, error) {
	t := time.Now().String()
	year := time.Now().Year()
	params := &dynamodb.PutItemInput{
		TableName: aws.String("gobook-entries"),
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"Email": {
				Exists: aws.Bool(false),
			},
		},
		Item: map[string]*dynamodb.AttributeValue{
			"Email": &dynamodb.AttributeValue{
				S: aws.String(email),
			},
			"Year": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(year)),
			},
			"DateTime": &dynamodb.AttributeValue{
				S: aws.String(t),
			},
			"Name": &dynamodb.AttributeValue{
				S: aws.String(name),
			},
			"Content": &dynamodb.AttributeValue{
				S: aws.String(content),
			},
		},
	}
	return e.svc.PutItem(params)
}

//Scan all of the table content
func (e *Entry) Scan() (resp *dynamodb.ScanOutput, err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String("gobook-entries"),
		AttributesToGet: []*string{
			aws.String("Name"),
			aws.String("DateTime"),
			aws.String("Content"),
		},
	}
	return e.svc.Scan(params)
}
