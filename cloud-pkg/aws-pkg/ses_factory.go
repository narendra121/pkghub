package awspkg

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/narendra121/pkghub/utils"
	log "github.com/sirupsen/logrus"
)

func (a *AwsCfg) SendEmail(senderEmail string, toEmails, ccEmails, bccEmails []string, subject, body string, attachments []utils.Attachment) error {
	svc := ses.New(a.AwsSession)
	emailTemplate := utils.CreateEmailTemplate(body, attachments)
	_, err := svc.SendRawEmail(&ses.SendRawEmailInput{
		Destinations: aws.StringSlice(append(append(toEmails, ccEmails...), bccEmails...)),
		Source:       aws.String(senderEmail),
		RawMessage:   &ses.RawMessage{Data: emailTemplate.Bytes()},
	})
	if err != nil {
		log.Errorln("error in sending email ", err)
		return err
	}
	return nil
}
