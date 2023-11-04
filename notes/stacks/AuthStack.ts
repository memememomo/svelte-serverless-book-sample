// 必要なクラスや関数を sst/constructs モジュール、同じディレクトリの他のスタックモジュール、aws-cdk-lib/aws-iam からインポート
import {Cognito, StackContext, use} from "sst/constructs";
import {ApiStack} from "./ApiStack";
import {StorageStack} from "./StorageStack";
import * as iam from "aws-cdk-lib/aws-iam";

// AuthStack 関数をエクスポート
export function AuthStack({ stack, app }: StackContext) {
    // ApiStack および StorageStack からリソースを取得
    const { api } = use(ApiStack);
    const { bucket } = use(StorageStack);

    // Cognito 認証を作成
    const auth = new Cognito(stack, "Auth", {
        login: ["email"],  // email を使ったログインを設定
    });

    // 認証されたユーザーに権限を付与
    auth.attachPermissionsForAuthUsers(stack, [
        api,  // APIへのアクセス権限
        // S3バケット内の特定のパスに対する権限を設定
        new iam.PolicyStatement({
            actions: ["s3:*"],  // S3 に対するすべてのアクションを許可
            effect: iam.Effect.ALLOW,  // 許可するエフェクト
            resources: [
                bucket.bucketArn + "/private/${cognito-identity.amazonaws.com:sub}/*",  // 特定のリソースパス
            ],
        }),
    ]);

    // スタックの出力として、リージョンとCognito関連の情報を追加
    stack.addOutputs({
        Region: app.region,
        UserPoolId: auth.userPoolId,
        UserPoolClientId: auth.userPoolClientId,
        IdentityPoolId: auth.cognitoIdentityPoolId,
    });

    // auth オブジェクトを戻り値として返す
    return {
        auth,
    };
}
