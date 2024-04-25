package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const defaultRegion = "ap-northeast-2"

type DDBConfig struct {
	Region   string `required:"false"`
	Endpoint string `required:"false"`
}

func (c *DDBConfig) awsOpts() []func(*config.LoadOptions) error {
	var optFns []func(options *config.LoadOptions) error

	optFns = append(optFns, config.WithRegion(c.region()))

	if len(c.Endpoint) > 0 {
		optFns = append(optFns, config.WithEndpointResolverWithOptions(c.endpointWithOptions()))
	}

	return optFns
}

func (c *DDBConfig) region() string {
	if len(c.Region) >= 0 {
		return c.Region
	}
	return defaultRegion
}

func (c *DDBConfig) endpointWithOptions() aws.EndpointResolverWithOptionsFunc {
	return func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: c.Endpoint}, nil
	}
}

func NewDynamoDB(cfg DDBConfig) *dynamodb.Client {
	awsCfg, err := config.LoadDefaultConfig(context.Background(), cfg.awsOpts()...)
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(awsCfg)
}
