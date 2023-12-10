package awspkg

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
)

func (cfg *AwsCfg) DownloadObjectFromS3(dest *os.File, buketName, objPath string) error {

	downloader := s3manager.NewDownloader(cfg.AwsSession)
	if _, err := downloader.Download(dest, &s3.GetObjectInput{
		Bucket: aws.String(buketName),
		Key:    aws.String(objPath),
	}); err != nil {
		log.Errorln("unable to download the object from s3 ", err)
		return err
	}
	return nil
}

func (cfg *AwsCfg) UploadObjectToS3(src *os.File, bucketName, objPath string) (string, error) {
	uploader := s3manager.NewUploader(cfg.AwsSession)
	uploadResp, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
		Body:   src,
	})
	if err != nil {
		log.Errorln("error in uploading file to s3 ", err)
		return "", err
	}
	return uploadResp.Location, nil
}

func (cfg *AwsCfg) DeleteObjectFromS3(src *os.File, bucketName, objPath string) error {
	s3Sess := s3.New(cfg.AwsSession)
	if _, err := s3Sess.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: aws.String(objPath)}); err != nil {
		log.Errorln("error in deleting object from s3 ", err)
		return err
	}
	return nil
}
