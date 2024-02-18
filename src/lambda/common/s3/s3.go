package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"invoice/common"
	"os"
)

func GetNewSession() *session.Session {
	if os.Getenv("ENVIRONMENT") == "local" {
		s3Url := os.Getenv("URL_S3_LOCAL")
		sess, err := session.NewSession(&aws.Config{
			Region:           aws.String("us-east-1"),
			Credentials:      credentials.NewStaticCredentials("test", "test", ""),
			S3ForcePathStyle: aws.Bool(true),
			Endpoint:         aws.String(s3Url),
		})
		common.LogError("Error when creating new session (local)", err)
		return sess
	}
	sess, err := session.NewSession(&aws.Config{})
	common.LogError("Error when creating new session", err)
	return sess
}
