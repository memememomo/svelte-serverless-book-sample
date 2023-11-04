<script lang="ts">
    // Svelteアプリケーションの内部ナビゲーションモジュールからgoto関数をインポート
    import { goto } from '$app/navigation';
    // AWS AmplifyのAuthモジュールをインポート
    import { Auth } from 'aws-amplify';
    // ストアをインポート
    import { loggedIn } from '$lib/store';

    // ユーザーが入力するemailとpasswordの変数を初期化
    let email: string = '';
    let password: string = '';
    // エラーメッセージを表示するための変数を初期化
    let errorMessage: string | null = null;

    // ログイン処理を行う非同期関数
    async function handleLogin(event: Event): Promise<void> {
        // デフォルトのフォーム送信を阻止
        event.preventDefault();
        // エラーメッセージを初期化
        errorMessage = null;

        try {
            // AWS AmplifyのAuthモジュールを使用してログイン処理を実行
            await Auth.signIn(email, password);
            // ログインが成功したらストアのloggedIn変数をtrueに設定
            loggedIn.set(true);
            // ログインが成功したらホームページにリダイレクト
            await goto('/');
        } catch (error) {
            // ログインに失敗した場合、エラーメッセージを設定
            errorMessage = 'ログイン中にエラーが発生しました。';
            // エラー内容をコンソールに出力
            console.error(error);
        }
    }
</script>

<style>
    .container {
        display: flex;           /* フレックスボックスを有効にする */
        flex-direction: column;  /* 子要素を垂直方向に配置 */
        align-items: center;     /* 子要素を水平方向の中央に配置 */
        height: 100vh;           /* ビューポートの高さを100%に設定 */
        max-width: 800px;        /* 画面の幅に応じて最大幅を設定 */
        margin: 0 auto;          /* 水平方向に中央寄せするためのマージン設定 */
        padding: 1rem;           /* コンテンツの周りに余白を追加 */
    }
</style>

<div class="container">
    <!-- 最大幅を中程度に設定し、全ての子要素の間隔を設定 -->
    <div class="max-w-md w-full space-y-8">
        <!-- エラーメッセージがある場合、そのエラーメッセージを赤色で表示 -->
        {#if errorMessage}
            <div class="text-red-500 text-center">{errorMessage}</div>
        {/if}
        <!-- ヘッダーセクション：ログインのタイトルを表示 -->
        <div>
            <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
                ログイン
            </h2>
        </div>
        <!-- フォームセクション：ユーザーが情報を入力する部分 -->
        <form on:submit={handleLogin} class="mt-8 space-y-6">
            <!-- メールアドレスとパスワードの入力フィールド -->
            <div class="rounded-md shadow-sm -space-y-px">
                <!-- メールアドレス入力フィールド -->
                <div>
                    <label for="email-address" class="sr-only">Emailアドレス</label>
                    <input bind:value={email} id="email-address" name="email" type="email" autocomplete="email" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Emailアドレス">
                </div>
                <!-- パスワード入力フィールド -->
                <div>
                    <label for="password" class="sr-only">パスワード</label>
                    <input bind:value={password} id="password" name="password" type="password" autocomplete="current-password" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="パスワード">
                </div>
            </div>
            <!-- ログインボタン -->
            <div>
                <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                    ログイン
                </button>
            </div>
        </form>
    </div>
</div>
