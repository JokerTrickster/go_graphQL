package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"main/src/common"
	"strings"
)

func AwsGetParam(path string) (string, error) {
	ctx := context.TODO()
	// get ssm param
	param, err := AwsClientSsm.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: common.GetBoolPointer(true),
	})
	if err != nil {
		return "", err
	}
	return aws.ToString(param.Parameter.Value), nil
}

func AwsGetParams(paths []string) ([]string, error) {
	ctx := context.TODO()
	// get ssm param
	params, err := AwsClientSsm.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          paths,
		WithDecryption: common.GetBoolPointer(true),
	})
	if err != nil {
		return nil, err
	}
	result := make([]string, len(paths))
	for i, path := range paths {
		val := ""
		for _, parameter := range params.Parameters {
			if strings.Contains(aws.ToString(parameter.ARN), path) {
				val = aws.ToString(parameter.Value)
				break
			}
		}
		result[i] = val
	}
	return result, nil
}
