// sst/constructs モジュールから必要なクラスをインポート
import { Bucket, StackContext, Table } from "sst/constructs";

// StorageStack 関数をエクスポート
export function StorageStack({ stack }: StackContext) {
    // DynamoDB のテーブルを作成
    // "Notes" という名前のテーブルで、主要なキーとして userId と noteId を持つ
    const table = new Table(stack, "Notes", {
        fields: {
            userId: "string",
            noteId: "string",
        },
        primaryIndex: { partitionKey: "userId", sortKey: "noteId" }, // プライマリインデックスの設定
    });

    // S3 バケットを作成
    // "Uploads" という名前のバケットで、CORS の設定を持つ
    const bucket = new Bucket(stack, "Uploads", {
        cors: [
            {
                maxAge: "1 day",  // キャッシュの有効期限
                allowedOrigins: ["*"],  // 許可されるオリジン
                allowedHeaders: ["*"],  // 許可されるヘッダー
                allowedMethods: ["GET", "PUT", "POST", "DELETE", "HEAD"],  // 許可されるHTTPメソッド
            },
        ],
    });

    // バケットとテーブルの参照を返す
    return {
        bucket,
        table,
    };
}
