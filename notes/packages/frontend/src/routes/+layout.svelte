<script lang="ts">
    // AWS Amplifyライブラリから必要なモジュールをインポート
    import {Amplify, Auth} from 'aws-amplify';
    // アプリのスタイルシートをインポート
    import "../app.css";
    // Svelteのナビゲーション関連のモジュールをインポート
    import { goto } from '$app/navigation';
    // 独自の設定ファイルをインポート
    import {config} from "../config";
    // Svelteのマウント関連のモジュールをインポート
    import {onMount} from "svelte";
    // 認証関連の関数をインポート
    import {checkLoginStatus} from "$lib/auth";
    // Svelteのページ関連のストアをインポート
    import { page } from "$app/stores";
    // ローディングコンポーネントをインポート
    import Loading from "$lib/components/Loading.svelte";
    // ストアをインポート
    import { loggedIn } from "$lib/store";

    // AWS Amplifyの設定
    Amplify.configure({
        Auth: {
            // サインインが必須かの設定
            mandatorySignIn: true,
            // Cognito関連の設定
            region: config.cognito.REGION,
            userPoolId: config.cognito.USER_POOL_ID,
            identityPoolId: config.cognito.IDENTITY_POOL_ID,
            userPoolWebClientId: config.cognito.APP_CLIENT_ID,
            // セッションストレージを使用する設定
            storage: global.sessionStorage,
        },
        Storage: {
            // S3の設定
            region: config.s3.REGION,
            bucket: config.s3.BUCKET,
            identityPoolId: config.cognito.IDENTITY_POOL_ID,
        },
        API: {
            // API Gatewayの設定
            endpoints: [
                {
                    name: "notes",
                    endpoint: config.apiGateway.URL,
                    region: config.apiGateway.REGION,
                },
            ],
        },
    });

    // ユーザーを指定のパスにリダイレクトする関数
    function navigateTo(path: string) {
        // goto関数を使用して指定したURLに遷移
        goto(path);
    }

    // ローディング状態とログイン状態の変数を初期化
    let isLoaded = false;

    // コンポーネントがマウントされたときの動作
    onMount(async () => {
        // ログイン状態をチェック
        const isLoggedIn = await checkLoginStatus();
        // ログインしていない場合、ログインページやサインアップページ以外にアクセスしていたら、ログインページにリダイレクト
        if (!isLoggedIn) {
            if ($page.url.pathname !== '/login' && $page.url.pathname !== '/signup') {
                window.location.href = '/login';
            } 
            loggedIn.set(false);
        } else {
            // ログインしている場合、フラグをtrueにセット
            loggedIn.set(true);
        }
        // ローディング完了
        isLoaded = true;
    });

    // ログアウト処理の関数
    function handleLogout() {
        // AmplifyのsignOut関数を使用してログアウト
        Auth.signOut()
            .then(() => {
                // ログアウト成功時、フラグを更新してログインページにリダイレクト
                loggedIn.set(false);
                navigateTo('/login');
            }).catch((err) => {
            // エラーが発生した場合、エラーをコンソールに出力
            console.error(err);
        });
    }
</script>


{#if !isLoaded}
    <!-- isLoaded変数がfalseの場合、アプリケーションがまだ読み込まれていない状態を示します -->
    <Loading />
    <!-- Loadingコンポーネントを表示して、ユーザーにローディング中であることを知らせます -->
{:else}
    <!-- isLoadedがtrueの場合、アプリケーションが読み込み完了していることを示します -->

    <div class="flex flex-col h-screen">
        <!-- 全画面の高さを持つflexboxコンテナを作成 -->

        <header class="p-4 bg-gray-800 text-white flex justify-between items-center">
            <!-- ヘッダー部分。背景色、テキスト色、内側の余白、アイテムの配置などのスタイリングが適用されています -->

            <!-- ロゴエリア -->
            <div>
                <img src="/logo.png" alt="SimpleNote logo" class="h-8" />
                <!-- アプリケーションのロゴを表示 -->
            </div>

            <div class="flex">
                <!-- ナビゲーションボタンのエリア -->

                {#if $loggedIn}
                    <!-- loggedIn変数がtrueの場合、ユーザーがログインしていることを示します -->

                    <!-- ログアウトボタン -->
                    <button on:click={() => handleLogout()}>Log out</button>
                {:else}
                    <!-- ユーザーがログインしていない場合 -->

                    <!-- ログインページへのボタン -->
                    <button class="mr-2" on:click={() => navigateTo('/login')}>Log In</button>

                    <!-- サインアップページへのボタン -->
                    <button on:click={() => navigateTo('/signup')}>Sign Up</button>
                {/if}
            </div>
        </header>

        <main class="flex-grow p-4">
            <!-- メインコンテンツエリア。flex-growクラスにより、利用可能なスペースを最大限に利用します -->

            <slot></slot>
            <!-- Svelteのスロット機能。このコンポーネントを使用する他のコンポーネントがコンテンツを挿入できる部分 -->
        </main>
    </div>
{/if}
