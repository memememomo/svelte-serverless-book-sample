// vitest モジュールからテスト関数をインポート
import { it } from 'vitest';

// sst/project モジュールからプロジェクト初期化関数をインポート
import { initProject } from "sst/project";

// sst/constructs モジュールから必要なクラスや関数をインポート
import { App, getStack } from "sst/constructs";

// 同じディレクトリの StorageStack モジュールから StorageStack クラスをインポート
import { StorageStack } from "../stacks/StorageStack";

// aws-cdk-lib/assertions モジュールからテンプレートアサーションクラスをインポート
import { Template } from "aws-cdk-lib/assertions";


// "Test StorageStack" というテストを定義
it("Test StorageStack", async () => {
    // プロジェクトを初期化
    await initProject({});

    // 新しい App インスタンスを作成し、その中で StorageStack をスタックとして追加
    const app = new App({ mode: "deploy"});
    app.stack(StorageStack);

    // StorageStack から CloudFormation スタックのテンプレートを取得
    const template = Template.fromStack(getStack(StorageStack));

    // アサーションを使用して、生成されたテンプレートに特定のリソースとプロパティが含まれていることを確認
    template.hasResourceProperties("AWS::DynamoDB::Table", {
        BillingMode: "PAY_PER_REQUEST",
    });
});
