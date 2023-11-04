// Package memo パッケージはログの生成やエラートレースのユーティリティを提供します。
package memo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

// getLogger は与えられたログハンドラを使用して新しいslog.Loggerを生成します。
// このロガーは一意のUUIDをIDとして持ちます。
func getLogger(h slog.Handler) (*slog.Logger, error) {
	// 新しいUUIDを生成
	id, err := uuid.NewRandom()
	if err != nil {
		// UUIDの生成中にエラーが発生した場合は、エラースタックと共に返します。
		return nil, errors.WithStack(err)
	}

	// 新しいロガーを作成し、UUIDをIDとして設定
	logger := slog.New(h).With(slog.Attr{
		Key:   "__id__",
		Value: slog.StringValue(id.String()),
	})
	return logger, nil
}

// handleOptions はslogハンドラのオプションを返します。
func handleOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true, // ソースコードの位置情報をログに追加するオプション
	}
}

// GetLogger はJSON形式のログをstdoutに出力する新しいslog.Loggerを生成して返します。
func GetLogger() (*slog.Logger, error) {
	// JSON形式のハンドラを使用して、新しいロガーを取得
	return getLogger(slog.NewJSONHandler(os.Stdout, handleOptions()))
}

// Trace はエラーの詳細なトレース情報をslog.Attrとして返します。
func Trace(err error) slog.Attr {
	// エラートレース情報を文字列として返す
	return slog.String("trace", fmt.Sprintf("%+v", err))
}
