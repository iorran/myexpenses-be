package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"invoice/common"
	"invoice/common/db"
)

func main() {
	lambda.Start(handlerDelete)
}

func handlerDelete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	conn := db.OpenDbConnection()
	collection := db.GetInvoiceCollection(conn)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Wrong ID format",
		}, nil
	}
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idHex})
	common.LogError("Error when trying to delete invoice", err)

	if res.DeletedCount == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Invoice does not exist",
		}, nil
	}

	db.CloseDbConnection(conn)
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}
