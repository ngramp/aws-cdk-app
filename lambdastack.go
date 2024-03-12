package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MyLambdaStack struct {
	awscdk.Stack
}

func NewMyLambdaStack(scope constructs.Construct, id string, props *awscdk.StackProps) *MyLambdaStack {
	stack := &MyLambdaStack{
		awscdk.NewStack(scope, &id, props),
	}

	// Define Lambda function
	myLambda := awslambda.NewFunction(stack, jsii.String("MyLambdaFunction"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("main"),
		Code:         awslambda.Code_FromAsset(jsii.String("../path/to/your/go/binary"), nil),
		Description:  jsii.String("My Lambda Function"),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
		MemorySize:   jsii.Number(256),
		Environment:  map[string]*string{"KEY": jsii.String("VALUE")},
		FunctionName: jsii.String("MyLambdaFunction"),
	})

	// Create API Gateway
	api := awsapigateway.NewRestApi(stack, jsii.String("MyAPI"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("MyAPI"),
	})

	// Integration between API Gateway and Lambda
	api.Lambda(jsii.String("MyLambdaIntegration"), &awsapigateway.LambdaIntegrationOptions{
		Handler:         myLambda,
		Proxy:           jsii.Bool(true),
		IntegrationName: jsii.String("MyLambdaIntegration"),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewMyLambdaStack(app, "MyLambdaStack", &awscdk.StackProps{
		StackName: jsii.String("MyLambdaStack"),
	})

	app.Synth(nil)
}
