<script lang="ts">
    // このコンポーネントで実行される処理を記述する

    import type { PageData } from './$types';
    import {updateCount} from "$lib";

    // +page.tsで定義されたload関数の返り値がここに入る
    // クリックカウント値がこの中に含まれている(data.count)
    export let data: PageData;

    // ボタンがクリックされた場合の処理
    const onClick = async () => {
        // クリックカウント更新APIの実行
        const cnt = await updateCount();
        // クリックカウント値を更新
        data.count = cnt.count;
    };
</script>

<div class="App">
    {#if data.count}<p>You clicked me {data.count} times.</p>{/if}
    <button on:click={onClick}>Click Me!</button>
</div>

<style>
    .App {
        text-align: center;
    }
    p {
        margin-top: 0;
        font-size: 20px;
    }
    button {
        font-size: 48px;
    }
</style>