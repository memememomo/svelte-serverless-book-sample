package counter

import (
	"os"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/guregu/dynamo"
)

// 環境変数からテーブル名を取得。環境変数名はSSTで定義されている命名規則に従う
// https://docs.sst.dev/resource-binding#how-it-works
func tableName() string {
	return os.Getenv("SST_Table_tableName_counter")
}

// DynamoDBに接続する処理
func connect() dynamo.Table {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	return db.Table(tableName())
}

// Clicks DynamoDBのデータ構造を表す構造体
type Clicks struct {
	Counter string `dynamo:"counter"` // パーティションキー
	Count   uint64 `dynamo:"count"`   // カウント値
}

// GetCount カウント値を取得する
func GetCount() (uint64, error) {
	// DynamoDBに接続
	table := connect()

	// パーティションキーが"clicks"のデータを取得
	var result Clicks
	err := table.Get("counter", "clicks").One(&result)
	if err != nil {
		if !errors.Is(err, dynamo.ErrNotFound) {
			return 0, err
		}
		result.Count = 0
	}

	// カウント値を返す
	return result.Count, nil
}

// UpdateCount カウント値を更新する
func UpdateCount() error {
	// DynamoDBに接続
	table := connect()

	// カウント値を取得
	count, err := GetCount()
	if err != nil {
		return err
	}

	// カウント値を更新
	err = table.Put(Clicks{Counter: "clicks", Count: count + 1}).Run()
	if err != nil {
		return err
	}

	return nil
}
