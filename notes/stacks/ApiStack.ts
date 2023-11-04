// 必要なクラスや関数を sst/constructs モジュールおよび同じディレクトリの StorageStack モジュールからインポート
import { Api, StackContext, use, Config } from "sst/constructs";
import { StorageStack } from "./StorageStack";

// ApiStack 関数をエクスポート
export function ApiStack({ stack }: StackContext) {
    // StorageStack からテーブルリソースを取得
    const { table } = use(StorageStack);

    // HOGE_SECRET_KEY を秘密としてコンフィグから取得
    const HOGE_SECRET_KEY = new Config.Secret(stack, "HOGE_SECRET_KEY");

    // APIの設定と作成
    const api = new Api(stack, "Api", {
        cors: true,  // CORSを有効にする
        defaults: {
            function: {
                bind: [table, HOGE_SECRET_KEY],  // ファンクションにテーブルと秘密鍵をバインド
                runtime: "go",  // 実行環境としてGoを指定
            },
            authorizer: "iam",  // IAMを使ってAPIの認可を行う
        },
        routes: {
            // 各エンドポイントとハンドラーのマッピング
            "POST /notes": "packages/functions/handlers/create.go",
            "GET /notes": "packages/functions/handlers/list.go",
            "GET /notes/{id}": "packages/functions/handlers/get.go",
            "PUT /notes/{id}": "packages/functions/handlers/update.go",
            "DELETE /notes/{id}": "packages/functions/handlers/delete.go",
        },
    });

    // スタックの出力としてAPIエンドポイントのURLを追加
    stack.addOutputs({
        ApiEndpoint: api.url,
    });

    // api オブジェクトを戻り値として返す
    return {
        api,
    };
}
