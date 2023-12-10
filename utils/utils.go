package utils

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type Attachment struct {
	Name string
	Data []byte
}

func GenerateRandomSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func ValidateExpiration(expiration int64) bool {
	return expiration > time.Now().Unix()
}

func GetTokenData(token *jwt.Token) jwt.MapClaims {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorln("error in parsing token ")
		return nil
	}
	return claims
}

func GetExpDuration(exp int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(exp)).Unix()
}

func RetryWithBackoff(fn func() (interface{}, error), maxAttempts int) (interface{}, error) {
	const initialDelay = time.Millisecond * 100
	const maxDelay = time.Second * 5
	retryCount := 0
	delay := initialDelay
	for retryCount < maxAttempts {
		resp, err := fn()
		if err == nil {
			return resp, nil
		}
		fmt.Printf("Attempt %d failed: %s\n", retryCount+1, err)
		if delay < maxDelay {
			delay *= 2
		}
		time.Sleep(delay)
		retryCount++
	}

	return nil, errors.New("all attempts failed")
}

func ConvertSliceToPtr(sl []string) []*string {
	strPtrs := make([]*string, 0)
	for _, str := range sl {
		strPtrs = append(strPtrs, aws.String(str))
	}
	return strPtrs
}

func CreateEmailTemplate(body string, attachments []Attachment) *bytes.Buffer {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	textPart := make(textproto.MIMEHeader)
	textPart.Set("Content-Type", "text/plain; charset=UTF-8")
	textPart.Set("Content-Transfer-Encoding", "quoted-printable")
	textPart.Set("Content-Disposition", `inline`)
	writer.CreatePart(textPart)
	buf.Write([]byte(body))
	for _, attachment := range attachments {
		part := make(textproto.MIMEHeader)
		part.Set("Content-Type", "application/octet-stream")
		part.Set("Content-Disposition", `attachment; filename="`+attachment.Name+`"`)
		attachmentPart, _ := writer.CreatePart(part)
		attachmentPart.Write(attachment.Data)
	}
	writer.Close()
	return buf
}
