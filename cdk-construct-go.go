package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type CdkConstructGoStackProps struct {
	awscdk.StackProps
}

func NewCdkConstructGoStack(scope constructs.Construct, id string, props *CdkConstructGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	myFunction := `function handler(event) {
		var request = event.request;
		var uri = request.uri;
	
		// Check whether the URI is missing a file name.
		if (uri.endsWith('/')) {
			request.uri += 'index.html';
		}
		// Check whether the URI is missing a file extension.
		else if (!uri.includes('.')) {
			request.uri += '/index.html';
		}
		return request;
	}`

	awscloudfront.NewFunction(stack, jsii.String("CFFuncUpdateSubdirPathCDK"), &awscloudfront.FunctionProps{
		FunctionName: jsii.String("update-subdir-path-cdk"),
		Code:         awscloudfront.FunctionCode(awscloudfront.FunctionCode_FromInline(jsii.String(myFunction))),
		Comment:      jsii.String("Path rewrite CloudFront Function deployed with CDK"),
	},
	)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkConstructGoStack(app, "CdkConstructGoStack", &CdkConstructGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
