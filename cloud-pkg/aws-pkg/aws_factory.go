package awspkg

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/narendra121/pkghub/utils"
	log "github.com/sirupsen/logrus"
)

type AwsFactory interface {
	DownloadObjectFromS3(dest *os.File, buketName, objPath string) error
	UploadObjectToS3(src *os.File, bucketName, objPath string) (string, error)
	DeleteObjectFromS3(src *os.File, bucketName, objPath string) error
	SendEmail(senderEmail string, toEmails, ccEmails, bccEmails []string, subject, body string, attachments []utils.Attachment) error
}

func NewAwsFactory(cfg *AwsCfg) AwsFactory {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(cfg.ClientID, cfg.ClientSecret, ""),
	})
	if err != nil {
		log.Error("error in creating aws session ", err)
		return nil
	}
	cfg.AwsSession = sess
	return cfg
}
