<script>
    // svelteライブラリからonMount関数をインポート
    import { onMount } from 'svelte';
    // $lib/authからcheckLoginStatus関数をインポート（ログイン状態をチェックする）
    import { checkLoginStatus } from '$lib/auth';

    // onMountは、コンポーネントがマウントされたときに実行されるコールバック関数を受け取る
    onMount(async () => {
        // ログイン状態をチェック
        const isLoggedIn = await checkLoginStatus();
        // ログインしていない場合、リダイレクトしない
        if (!isLoggedIn) {
            return;
        }
        // '/notes'ページへリダイレクト
        window.location.href = '/notes';
    });
</script>
