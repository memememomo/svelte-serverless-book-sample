// sst/constructs モジュールおよび同じディレクトリ内の他のスタックモジュールから必要なクラスや関数をインポート
import {StackContext, SvelteKitSite, use} from "sst/constructs";
import {ApiStack} from "./ApiStack";
import {AuthStack} from "./AuthStack";
import {StorageStack} from "./StorageStack";

// FrontendStack 関数をエクスポート
export function FrontendStack({ stack, app }: StackContext) {
    // ApiStack、AuthStack、StorageStack のそれぞれからリソースを取得
    const { api } = use(ApiStack);
    const { auth } = use(AuthStack);
    const { bucket } = use(StorageStack);

    // SvelteKit を使用したフロントエンドサイトを作成
    const site = new SvelteKitSite(stack, "SvelteSite", {
        path: "packages/frontend",  // ソースコードのパス
        environment: {  // 環境変数の設定
            PUBLIC_API_ENDPOINT: api.url,  // API のエンドポイント
            PUBLIC_REGION: app.region,  // リージョン情報
            PUBLIC_BUCKET: bucket.bucketName,  // S3 バケットの名前
            PUBLIC_USER_POOL_ID: auth.userPoolId,  // Cognito ユーザープールID
            PUBLIC_USER_POOL_CLIENT_ID: auth.userPoolClientId,  // Cognito ユーザープールクライアントID
            PUBLIC_IDENTITY_POOL_ID: auth.cognitoIdentityPoolId || "",  // Cognito アイデンティティプールID
        },
    });

    // サイトのURLをスタックの出力として追加
    stack.addOutputs({
        SiteUrl: site.url,
    });
}
