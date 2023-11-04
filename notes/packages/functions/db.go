package memo

import (
	"github.com/pkg/errors"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/guregu/dynamo"
)

// 環境変数からテーブル名を取得。環境変数名はSSTで定義されている命名規則に従う
// https://docs.sst.dev/resource-binding#how-it-works
func tableName() string {
	return os.Getenv("SST_Table_tableName_Notes")
}

// connectはDynamoDBに接続するための関数。
func connect() dynamo.Table {
	sess := session.Must(session.NewSession())                                // AWS SDKを使用して新しいセッションを作成
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")}) // 指定されたリージョンでDynamoDBに接続
	return db.Table(tableName())                                              // テーブルにアクセスするためのインスタンスを返す
}

// Note はメモを表す構造体。
type Note struct {
	UserID     string    `dynamo:"userId" json:"userId"`         // ユーザのID
	NoteID     string    `dynamo:"noteId" json:"noteId"`         // メモのID
	Content    string    `dynamo:"content" json:"content"`       // メモの内容
	Attachment string    `dynamo:"attachment" json:"attachment"` // 添付ファイルや画像のURLなど
	CreatedAt  time.Time `dynamo:"createdAt" json:"createdAt"`   // メモの作成日時
}

// CreateNote は新しいNoteをDynamoDBに保存する関数。
func CreateNote(newNote *Note) error {
	table := connect()
	// Putメソッドで新しいNoteをテーブルに保存
	err := table.Put(newNote).Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateNoteByNoteID はNoteIDに基づき、既存のNoteをDynamoDBに更新する関数。
func UpdateNoteByNoteID(newNote *Note) error {
	table := connect()
	// Putメソッドで既存のNoteを上書きして更新
	err := table.Put(newNote).Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteNoteByNoteID は指定されたNoteIDのNoteをDynamoDBから削除する関数。
func DeleteNoteByNoteID(noteID string, userID string) error {
	table := connect()
	// Deleteメソッドで指定されたuserIDとnoteIDに一致するNoteを削除
	err := table.
		Delete("userId", userID).
		Range("noteId", noteID).
		Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetNoteByNoteID は指定されたNoteIDのNoteをDynamoDBから取得する関数。
func GetNoteByNoteID(noteID string, userID string) (*Note, error) {
	table := connect()
	var note Note
	// Getメソッドで指定されたuserIDとnoteIDに一致するNoteを取得
	err := table.
		Get("userId", userID).
		Range("noteId", dynamo.Equal, noteID).
		One(&note)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &note, nil
}

// ListNotes は指定されたUserIDに関連するすべてのNoteをDynamoDBから取得する関数。
func ListNotes(userID string) ([]*Note, error) {
	table := connect()
	// Getメソッドで指定されたuserIDに関連するすべてのNoteを取得
	var notes []*Note
	err := table.
		Get("userId", userID).All(&notes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return notes, nil
}
