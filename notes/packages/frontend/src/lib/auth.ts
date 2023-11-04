// AWS AmplifyのAuthモジュールをインポート
import {Auth} from 'aws-amplify';

// ユーザーのログイン状態を確認する非同期関数を定義
export async function checkLoginStatus(){
    try {
        // 現在認証されているユーザー情報を取得
        const user = await Auth.currentAuthenticatedUser();
        // userオブジェクトがnullでなければtrueを返す（ログイン済みと判断）
        return user != null;
    } catch (err) {
        // 何らかのエラーが発生した場合、falseを返す（未ログインと判断）
        return false;
    }
}
