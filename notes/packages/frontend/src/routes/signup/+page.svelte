<script lang="ts">
    // AWS AmplifyのAuthモジュールとSvelteアプリのナビゲーションモジュールをインポート
    import { goto } from '$app/navigation';
    import { Auth } from 'aws-amplify';

    // ユーザーが入力するメールアドレスを格納する変数
    let email: string = '';
    // ユーザーが入力するパスワードを格納する変数
    let password: string = '';
    // ユーザーが入力するサインアップ確認コードを格納する変数
    let code: string = '';
    // エラーメッセージを格納する変数。エラーがない場合はnull
    let errorMessage: string | null = null;
    // ユーザーが登録中かどうかの状態を管理する変数。trueの場合、コード確認ステップが表示される
    let registering: boolean = false;

    // エラーオブジェクトがメッセージプロパティを持っているかどうかを確認するヘルパー関数
    function isErrorWithMessage(error: unknown): error is { message: string } {
        return (error as { message?: string }).message != null;
    }

    // ユーザーをサインアップする関数
    const signUp = async (username: string, password: string) => {
        try {
            // AWS AmplifyのAuthモジュールを使用してユーザーをサインアップ
            const { user } = await Auth.signUp({
                username,
                password,
                autoSignIn: {
                    enabled: true
                }
            });
            // サインアップが成功したら、コード確認画面を表示するフラグをセット
            registering = true;
            return user;
        } catch (error) {
            // エラーが発生した場合、エラーメッセージを表示するためのロジック
            if (isErrorWithMessage(error)) {
                errorMessage = error.message;
            } else {
                errorMessage = 'エラーが発生しました。';
            }
        }
    };

    // サインアップコードの再送をリクエストする関数
    const resend = async (username: string) => {
        try {
            await Auth.resendSignUp(username);
        } catch (error) {
            if (isErrorWithMessage(error)) {
                errorMessage = error.message;
            } else {
                errorMessage = 'エラーが発生しました。';
            }
        }
    };

    // サインアップ時に送信されたコードを確認する関数
    async function confirmSignUp(username: string, code: string) {
        try {
            await Auth.confirmSignUp(username, code);
            // コードの確認が成功したら、ログインページにリダイレクト
            await goto('/login');
        } catch (error) {
            if (isErrorWithMessage(error)) {
                errorMessage = error.message;
            } else {
                errorMessage = 'エラーが発生しました。';
            }
        }
    }

    // サインアップフォームの送信ハンドラ
    function handleSignUp(event: Event) {
        event.preventDefault();
        signUp(email, password);
    }

    // コード確認フォームの送信ハンドラ
    function handleCodeVerification(event: Event) {
        event.preventDefault();
        confirmSignUp(email, code);
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
    <!-- 最大幅を中央揃えし、要素間のスペースを設定 -->
    <div class="max-w-md w-full space-y-8">
        <!-- エラーメッセージがある場合は赤色で表示 -->
        {#if errorMessage}
            <div class="text-red-500 text-center">{errorMessage}</div>
        {/if}
        <div>
            <!-- 登録中であれば「コード確認」と表示、そうでなければ「サインアップ」と表示 -->
            <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
                {#if registering} コード確認 {/if}
                {#if !registering} サインアップ {/if}
            </h2>
        </div>
        <!-- 登録中でない場合、サインアップフォームを表示 -->
        {#if !registering}
            <!-- サインアップフォーム -->
            <form on:submit={handleSignUp} class="mt-8 space-y-6">
                <!-- 入力フィールドのコンテナ -->
                <div class="rounded-md shadow-sm -space-y-px">
                    <!-- Email入力フィールド -->
                    <div>
                        <label for="email-address" class="sr-only">Emailアドレス</label>
                        <input bind:value={email} id="email-address" name="email" type="email" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Emailアドレス">
                    </div>
                    <!-- パスワード入力フィールド -->
                    <div>
                        <label for="password" class="sr-only">パスワード</label>
                        <input bind:value={password} id="password" name="password" type="password" autocomplete="current-password" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="パスワード">
                    </div>
                </div>
                <!-- サインアップボタン -->
                <div>
                    <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                        サインアップ
                    </button>
                </div>
            </form>
        {/if}
        <!-- 登録中の場合、コード確認フォームを表示 -->
        {#if registering}
            <!-- コード確認フォーム -->
            <form on:submit={handleCodeVerification} class="mt-8 space-y-6">
                <!-- コード入力フィールドのコンテナ -->
                <div class="rounded-md shadow-sm -space-y-px">
                    <!-- コード入力フィールド -->
                    <div>
                        <label for="code" class="sr-only">確認コード</label>
                        <input bind:value={code} id="code" name="code" type="text" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="確認コード">
                    </div>
                </div>
                <!-- コード確認ボタンとコード再送ボタンのコンテナ -->
                <div class="flex justify-between">
                    <!-- コード確認ボタン -->
                    <button type="submit" class="group relative flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                        コードを確認
                    </button>
                    <!-- コード再送ボタン -->
                    <button type="button" on:click={() => resend(email)} class="group relative flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                        コード再送
                    </button>
                </div>
            </form>
        {/if}
    </div>
</div>
