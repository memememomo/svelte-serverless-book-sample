<script>
    // svelte/transition モジュールから fade トランジションをインポート
    import { fade } from 'svelte/transition';

    // オプションでスタイルやアニメーションをカスタマイズできます
</script>

<!-- fadeトランジションを使用して表示するローディング画面 -->
<div in:fade class="loading">
    <!-- スピンアニメーションを持つローディングアイコン -->
    <div class="spinner"></div>
    Loading... <!-- テキスト表示 -->
</div>

<style>
    /* ローディング画面のスタイリング */
    .loading {
        position: fixed; /* 画面の固定位置に配置 */
        top: 0; /* 上端からの距離 */
        left: 0; /* 左端からの距離 */
        width: 100vw; /* 画面の全幅を使用 */
        height: 100vh; /* 画面の全高さを使用 */
        background: rgba(255, 255, 255, 0.8); /* 背景色（半透明の白） */
        display: flex; /* フレックスボックスレイアウト */
        justify-content: center; /* 中央寄せ（水平方向） */
        align-items: center; /* 中央寄せ（垂直方向） */
    }

    /* スピンアニメーションを持つローディングアイコンのスタイリング */
    .spinner {
        width: 50px; /* 幅 */
        height: 50px; /* 高さ */
        border: 5px solid rgba(255, 255, 255, 0.3); /* 境界線のスタイル（半透明の白） */
        border-radius: 50%; /* 丸い形状 */
        border-top: 5px solid black; /* 上側の境界線の色 */
        animation: spin 1s linear infinite; /* スピンアニメーションの適用 */
    }

    /* スピンアニメーションの定義 */
    @keyframes spin {
        0% { transform: rotate(0deg); } /* 開始時点の角度 */
        100% { transform: rotate(360deg); } /* 終了時点の角度（1回転） */
    }
</style>
