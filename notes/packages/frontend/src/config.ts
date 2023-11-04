// 環境変数から公開用の設定をインポート
import {
    PUBLIC_API_ENDPOINT,
    PUBLIC_BUCKET,
    PUBLIC_IDENTITY_POOL_ID,
    PUBLIC_REGION,
    PUBLIC_USER_POOL_CLIENT_ID,
    PUBLIC_USER_POOL_ID
} from "$env/static/public";

// 公開用の設定をオブジェクトとしてエクスポート
export const config = {
    // S3関連の設定
    s3: {
        REGION: PUBLIC_REGION, // リージョンの設定
        BUCKET: PUBLIC_BUCKET, // S3のバケット名
    },
    // API Gateway関連の設定
    apiGateway: {
        REGION: PUBLIC_REGION, // リージョンの設定
        URL: PUBLIC_API_ENDPOINT, // APIのエンドポイント
    },
    // Cognito関連の設定
    cognito: {
        REGION: PUBLIC_REGION, // リージョンの設定
        USER_POOL_ID: PUBLIC_USER_POOL_ID, // ユーザープールID
        APP_CLIENT_ID: PUBLIC_USER_POOL_CLIENT_ID, // アプリクライアントID
        IDENTITY_POOL_ID: PUBLIC_IDENTITY_POOL_ID, // アイデンティティプールID
    },
};
