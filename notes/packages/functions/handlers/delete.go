package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"memo"
)

func main() {
	lambda.Start(memo.HandleDelete)
}
