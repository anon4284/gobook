package entry

import (
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Entry class includes routes and db connection
type Entry struct {
	svc *dynamodb.DynamoDB
}

//New create new entry class
func New(svc dynamodb.DynamoDB) *Entry {
	return &Entry{svc: &svc}
}

type EntryAddInput struct {
	Name    string
	Email   string
	Content string
}
type EntryAddOutput struct {
	Valid   bool
	Problem string
	Msg     string
}

type EntryGetOutput struct {
	Valid    bool
	Name     string
	Year     string
	DateTime string
	Content  string
}

func validate(entryInput EntryAddInput) *EntryAddOutput {

	if !govalidator.IsEmail(entryInput.Email) {
		return &EntryAddOutput{false, "email", "email must be valid"}
	}
	if len(entryInput.Name) < 3 {
		return &EntryAddOutput{false, "name", "name must be at least 3 characters long"}
	}
	if len(entryInput.Content) < 10 {
		return &EntryAddOutput{false, "name", "Content must be at least 10 characters long"}
	}

	return &EntryAddOutput{true, strconv.Itoa(time.Now().Year()), "Entry added"}

}
