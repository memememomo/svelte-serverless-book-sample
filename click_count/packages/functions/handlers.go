package counter

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response レスポンスのBodyに入れるデータ構造を定義する
type Response struct {
	Count uint64 `json:"count"` // クリックカウント値
}

// HandlerGet クリックカウント値を返すLambdaハンドラ
func HandlerGet(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// クリックカウント値を取得
	count, err := GetCount()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// クリックカウント値を含んだJSONレスポンス
	body, err := json.Marshal(Response{
		Count: count,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// API Gateway用のレスポンス
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

// HandlerUpdate クリックカウント値を更新するLambdaハンドラ
func HandlerUpdate(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// クリックカウント値を更新
	err := UpdateCount()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// クリックカウント値を取得
	count, err := GetCount()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// クリックカウント値を含んだJSONレスポンス
	body, err := json.Marshal(Response{
		Count: count,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// API Gateway用のレスポンス
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}
