package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"invoice/common"
	"invoice/common/db"
	"invoice/model"
	"time"
)

func main() {
	lambda.Start(handlerSearch)
}

func handlerSearch(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryParam := request.QueryStringParameters
	startDate := common.ValidateAndParseDate(queryParam["startDate"])
	endDate := common.ValidateAndParseDate(queryParam["endDate"])
	foundInvoices := searchInvoice(startDate, endDate)
	invoices, _ := json.Marshal(foundInvoices)
	response := string(invoices)
	if response == "null" {
		response = "[]"
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
	}, nil
}

func searchInvoice(startDate time.Time, endDate time.Time) []model.InvoiceDocument {
	conn := db.OpenDbConnection()
	collection := db.GetInvoiceCollection(conn)
	query := bson.M{"$gte": startDate, "$lte": endDate}
	invoices, err := collection.Find(context.TODO(), bson.M{"date": query})
	common.LogError("Error trying search invoices", err)
	var result []model.InvoiceDocument
	err = invoices.All(context.TODO(), &result)
	common.LogError("Error when decoding invoices", err)
	db.CloseDbConnection(conn)
	return result
}
