package memo

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

// HandlePostBody はHTTP POSTリクエストから受け取るJSONデータの構造を定義した構造体。
// この構造体はメモの内容と添付ファイルを受け取るために使用されます。
type HandlePostBody struct {
	Content    string `json:"content"`    // メモの内容
	Attachment string `json:"attachment"` // 添付ファイルや画像のURLなど
}

// commonHeadersはCORS(Cross-Origin Resource Sharing)対応のためのHTTPヘッダーを返す関数。
// この関数はAPIのレスポンスに必要なヘッダーを設定するために使用されます。
func commonHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",    // 全てのオリジンからのリクエストを許可
		"Access-Control-Allow-Credentials": "true", // クレデンシャル(例: クッキー)を含むリクエストを許可
	}
}

// Response500 は、サーバー内部エラーが発生した場合のHTTPレスポンスを生成します。
func Response500(err error, logger *slog.Logger) (events.APIGatewayV2HTTPResponse, error) {
	logger.Error("Internal Server Error", Trace(err))
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 500,
		Body:       "{\"error\": \"Internal Server Error\"}",
		Headers:    commonHeaders(),
	}, nil
}

// Response200 は、成功した場合のHTTPレスポンスを生成します。
func Response200(body string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    commonHeaders(),
	}, nil
}

// Response403 は、クライアントがリソースへのアクセス権を持っていない場合のHTTPレスポンスを生成します。
func Response403() (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 403,
		Body:       "{\"error\": \"Forbidden\"}",
		Headers:    commonHeaders(),
	}, nil
}

// IsAuthorized は、リクエストが認証されているかどうかを判定します。
func IsAuthorized(req events.APIGatewayV2HTTPRequest) bool {
	return req.RequestContext.Authorizer != nil && req.RequestContext.Authorizer.IAM != nil
}

// UserID は、リクエストからユーザーIDを取得します。
func UserID(req events.APIGatewayV2HTTPRequest) string {
	return req.RequestContext.Authorizer.IAM.CognitoIdentity.IdentityID
}

// HandleDelete は、ノートを削除するリクエストを処理します。
func HandleDelete(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// ロガーのインスタンスを取得
	logger, err := GetLogger()
	if err != nil {
		return Response500(err, logger)
	}

	// リクエストが認証されていない場合は403レスポンスを返す
	if !IsAuthorized(req) {
		return Response403()
	}

	// パスパラメータからノートIDを取得
	noteID := req.PathParameters["id"]

	// 指定されたノートIDとユーザーIDでノートを削除
	err = DeleteNoteByNoteID(noteID, UserID(req))
	if err != nil {
		return Response500(err, logger)
	}

	// 削除が成功した場合のレスポンスを返す
	return Response200("")
}

// HandleUpdate はノートの更新リクエストを処理します。
func HandleUpdate(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// ロガーのインスタンスを取得
	logger, err := GetLogger()
	if err != nil {
		return Response500(err, logger)
	}

	// 認証の確認
	if !IsAuthorized(req) {
		return Response403()
	}

	// パスパラメータからノートIDを取得
	noteID := req.PathParameters["id"]
	userID := UserID(req)

	// リクエストボディをデコード
	var body HandlePostBody
	err = json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return Response500(err, logger)
	}

	// 現在のノート情報を取得
	note, err := GetNoteByNoteID(noteID, userID)
	if err != nil {
		return Response500(err, logger)
	}

	// 更新するノートの情報を作成
	note.Content = body.Content // リクエストから取得したノートの内容

	// 付属情報が存在する場合は、それを設定
	if body.Attachment != "" {
		note.Attachment = body.Attachment
	}

	// ノートの更新
	err = UpdateNoteByNoteID(note)
	if err != nil {
		return Response500(err, logger)
	}

	return Response200("")
}

// HandleList はユーザーのノート一覧を取得するリクエストを処理します。
func HandleList(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// ロガーのインスタンスを取得
	logger, err := GetLogger()
	if err != nil {
		return Response500(err, logger)
	}

	// 認証の確認
	if !IsAuthorized(req) {
		return Response403()
	}

	// ノートの一覧を取得
	notes, err := ListNotes(UserID(req))
	if err != nil {
		return Response500(err, logger)
	}

	// ノートの一覧をJSON形式に変換
	resBody, err := json.Marshal(notes)
	if err != nil {
		return Response500(err, logger)
	}

	return Response200(string(resBody))
}

// HandleGet は特定のノートを取得するリクエストを処理します。
func HandleGet(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// ロガーのインスタンスを取得
	logger, err := GetLogger()
	if err != nil {
		return Response500(err, logger)
	}

	// 認証の確認
	if !IsAuthorized(req) {
		return Response403()
	}

	// パスパラメータからノートIDを取得
	noteID := req.PathParameters["id"]

	// 指定されたノートIDとユーザーIDでノートを取得
	note, err := GetNoteByNoteID(noteID, UserID(req))
	if err != nil {
		return Response500(err, logger)
	}

	// ノート情報をJSON形式に変換
	resBody, err := json.Marshal(note)
	if err != nil {
		return Response500(err, logger)
	}

	return Response200(string(resBody))
}

// HandlePost は新しいノートを作成するリクエストを処理します。
func HandlePost(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// ロガーのインスタンスを取得
	logger, err := GetLogger()
	if err != nil {
		return Response500(err, logger)
	}

	// リクエスト情報をロガーに出力
	logger.With(req).Info("HandlePost")

	// 認証の確認
	if !IsAuthorized(req) {
		return Response403()
	}

	// リクエストボディをデコード
	var body HandlePostBody
	err = json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return Response500(err, logger)
	}

	// UUIDを生成
	guuid, err := uuid.NewRandom()
	if err != nil {
		return Response500(err, logger)
	}

	// 新しいノートの情報を作成
	newNote := &Note{
		UserID:     UserID(req),     // 認証されたユーザのID
		NoteID:     guuid.String(),  // 生成されたUUIDをノートIDとして設定
		Content:    body.Content,    // リクエストから取得したノートの内容
		Attachment: body.Attachment, // リクエストから取得した添付情報
		CreatedAt:  time.Now(),      // 現在のタイムスタンプを作成日時として設定
	}

	// ノートをデータベースに作成
	err = CreateNote(newNote)
	if err != nil {
		return Response500(err, logger)
	}

	// 新しいノートの情報をJSON形式に変換
	resBody, err := json.Marshal(newNote)
	if err != nil {
		return Response500(err, logger)
	}

	return Response200(string(resBody))
}
