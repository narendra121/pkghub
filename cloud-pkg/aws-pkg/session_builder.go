package awspkg

import "github.com/aws/aws-sdk-go/aws/session"

type AwsCfg struct {
	ClientID     string
	ClientSecret string
	Region       string
	AwsSession   *session.Session
}

type AwsCfgBuilder struct {
	awsCfg AwsCfg
}

func NewAwsCfgBuilder() *AwsCfgBuilder {
	return &AwsCfgBuilder{awsCfg: AwsCfg{}}
}

func (acb *AwsCfgBuilder) AddClientID(clientID string) *AwsCfgBuilder {
	acb.awsCfg.ClientID = clientID
	return acb
}

func (acb *AwsCfgBuilder) AddClientSecret(clientSecret string) *AwsCfgBuilder {
	acb.awsCfg.ClientSecret = clientSecret
	return acb
}

func (acb *AwsCfgBuilder) AddRegion(region string) *AwsCfgBuilder {
	acb.awsCfg.Region = region
	return acb
}

func (acb *AwsCfgBuilder) Build() *AwsCfg {
	return &acb.awsCfg
}
