package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var AwsClientSsm *ssm.Client

func InitAws() error {
	var awsConfig aws.Config
	var err error
	awsConfig, err = AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion("us-east-1"))
	if err != nil {
		return err
	}
	AwsClientSsm = ssm.NewFromConfig(awsConfig)
	return nil
}
