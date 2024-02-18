package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"invoice/common"
	"invoice/common/db"
	"invoice/common/s3"
	"invoice/model"
	"net/http"
	"os"
)

func main() {
	lambda.Start(handlerSave)
}

func handlerSave(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	invoice := model.InvoiceRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &invoice)
	common.LogError("Error when unmarshalling JSON", err)

	parsedDate := common.ValidateAndParseDate(invoice.Date)
	imageLink := uploadImage(invoice.Image)

	saveInvoice(model.InvoiceDocument{
		Date:        parsedDate,
		Image:       imageLink,
		Description: invoice.Description,
	})

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       imageLink,
	}
	return response, nil
}

func saveInvoice(invoice model.InvoiceDocument) {
	conn := db.OpenDbConnection()
	collection := db.GetInvoiceCollection(conn)
	_, err := collection.InsertOne(context.Background(), invoice)
	common.LogError("Error trying to insert invoice", err)
	db.CloseDbConnection(conn)
}

func uploadImage(imageBase64 string) string {
	newSession := s3.GetNewSession()
	imageBytes, err := base64.StdEncoding.DecodeString(imageBase64)
	common.LogError("Error when decoding image base64", err)

	svc := s3manager.NewUploader(newSession)
	res, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String("invoice-" + uuid.New().String() + ".jpg"),
		Body:   bytes.NewReader(imageBytes),
	})
	common.LogError("Error when uploading file", err)

	return res.Location
}
