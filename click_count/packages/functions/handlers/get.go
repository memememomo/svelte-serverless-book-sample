package main

import (
	"counter"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(counter.HandlerGet)
}
