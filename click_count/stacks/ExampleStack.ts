import {Api, StackContext, Table, SvelteKitSite} from "sst/constructs";


export function ExampleStack({ stack }: StackContext) {
    // DynamoDB テーブルを作成
    const table = new Table(stack, "counter", {
        fields: {
            counter: "string",
        },
        primaryIndex: {partitionKey: "counter"}, // パーティションキーを設定
    });

    // API GatewayエンドポイントとLambdaを作成
    const api = new Api(stack, "api", {
        // 全Lambda関数の共通設定
        defaults: {
            function: {
                bind: [table], // DynamoDBテーブルをバインド
                runtime: "go", // ランタイムをGoに設定
            },
        },
        // APIエンドポイントの設定
        routes: {
            "GET /hello": "packages/functions/handlers/hello.go",
            "GET /count": "packages/functions/handlers/get.go", // 追加: カウント取得API
            "POST /count": "packages/functions/handlers/update.go", // 追加: カウント更新API
        },
    });

    // SvelteKitのデプロイ設定
    const site = new SvelteKitSite(stack, "SvelteSite", {
        path: "packages/frontend", // デプロイするディレクトリを指定
        environment: {
           PUBLIC_API_ENDPOINT: api.url, // APIエンドポイントを環境変数に設定
        },
    });

    // 値の出力
    stack.addOutputs({
        SiteUrl: site.url, // フロントエンドのURL
        ApiEndpoint: api.url, // APIエンドポイントのURL
    });
};
