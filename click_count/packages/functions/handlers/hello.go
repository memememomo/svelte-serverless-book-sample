package main
import (
        "context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// 処理を行うハンドラー
func handler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// API Gatewayのレスポンスとして、ステータスコード200と文字列を返す
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}

// main関数
func main() {
	// ハンドラーを登録して実行
	lambda.Start(handler)
}
