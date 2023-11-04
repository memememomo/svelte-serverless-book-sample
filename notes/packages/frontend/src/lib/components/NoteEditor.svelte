<script lang="ts">
    // 必要なライブラリやモジュールをインポート
    import {s3Upload} from "$lib/upload";
    import type {NoteType} from "$lib/api";
    import {createNote, saveNote} from "$lib/api";
    import { Storage } from 'aws-amplify';
    import {goto} from "$app/navigation";
    import Fa from "svelte-fa";
    import { faPencilAlt, faPlus } from "@fortawesome/free-solid-svg-icons";

    // 外部からの入力を受けるnote変数を定義
    export let note: NoteType | null = null;

    // 各種変数を初期化
    let content: string = note ? note.content : '';   // ノートの内容
    let file: File | null = null;                     // アップロードするファイル
    let message: string = '';                         // 表示するメッセージ
    let showMessage = false;                          // メッセージ表示フラグ
    let isLoading = false;                            // ローディングフラグ
    let errorMessage: string = '';                    // エラーメッセージ

    // ファイルのアップロードイベントハンドラ
    const handleFileUpload = async (event: Event) => {
        const target = event.target as HTMLInputElement;  // イベント発火元の要素を取得
        file = target.files ? target.files[0] : null;    // 選択されたファイルを取得
    }

    // ノートの内容をバリデーションする関数
    function validateContent(): boolean {
        if (!content.trim()) {
            errorMessage = 'コメントを入力してください';  // エラーメッセージを設定
            return false;
        }
        errorMessage = '';   // エラーメッセージをリセット
        return true;
    }

    // 提出イベントのハンドラ
    const handleSubmit = async () => {
        if (!validateContent()) {
            isLoading = false;  // ローディングフラグをリセット
            return;
        }

        isLoading = true;  // ローディングフラグをオン
        try {
            const attachment = file ? await s3Upload(file) : undefined;  // ファイルが存在すればアップロード

            if (note) {
                // ノートが存在すれば、ノートを更新
                note = await saveNote({noteId: note.noteId, content, attachment});
                if (note?.attachment) {
                    note.attachmentURL = await Storage.vault.get(note.attachment);  // ノートの添付ファイルURLを更新
                }
                message = '更新しました';  // 成功メッセージ
                showMessage = true;  // メッセージ表示フラグをオン
            } else {
                // ノートが存在しなければ新しいノートを作成
                await createNote({ content, attachment });
                navigateToList();  // 一覧ページへリダイレクト
            }
        } catch (err) {
            console.log(err);  // エラーをログに出力
        } finally {
            isLoading = false;  // ローディングフラグをリセット
        }
    }

    // 一覧ページへのリダイレクト関数
    const navigateToList = () => {
        goto('/notes');
    };
</script>


<style>
    /* コンテナ: ページ全体のスタイリングを設定 */
    .container {
        max-width: 800px;  /* コンテナの最大幅を800pxに制限 */
        margin: 0 auto;  /* コンテナをページの中央に配置 */
        padding: 0 20px;  /* コンテナの左右に20pxのパディングを設定 */
    }

    /* セクション: 各コンテンツ部分をグルーピングするスタイル */
    .section {
        margin-bottom: 30px;  /* セクション間の下部の余白を30pxに設定 */
    }

    /* スピンアニメーション: ローディングアニメーションの回転動作を定義 */
    @keyframes spin {
        0% { transform: rotate(0deg); }   /* 開始時は0度の位置 */
        100% { transform: rotate(360deg); }  /* 終了時は360度の位置（1回転） */
    }

    /* ローディングインジケータ: ページ読み込み中やデータ取得中のインジケータスタイル */
    .loading-indicator {
        border: 4px solid rgba(0, 0, 255, 0.3);  /* ボーダーの色と透明度を設定 */
        border-radius: 50%;  /* 円形にする */
        border-top: 4px solid #0000ff;  /* 上部のボーダーの色を設定 */
        animation: spin 1s linear infinite;  /* 定義したspinアニメーションを1秒で無限に繰り返し適用 */
        width: 20px;  /* インジケータの幅 */
        height: 20px;  /* インジケータの高さ */
        margin-right: 15px; /* 右側の余白を15pxに設定 */
    }

    /* オーバーレイメッセージ: 画面全体に表示されるメッセージの背景スタイル */
    .overlay-message {
        position: fixed;  /* 位置を固定 */
        top: 0;  /* 上辺の位置 */
        left: 0;  /* 左辺の位置 */
        width: 100vw;  /* 画面全体の幅 */
        height: 100vh;  /* 画面全体の高さ */
        background-color: rgba(0, 0, 0, 0.3);  /* 背景色と透明度 */
        display: flex;  /* フレックスボックスを使用 */
        align-items: center;  /* 垂直方向に中央に配置 */
        justify-content: center;  /* 水平方向に中央に配置 */
        z-index: 1000;  /* zインデックス値（他の要素の上に表示） */
    }

    /* メッセージコンテナ: オーバーレイメッセージ内のテキストやボタンをグルーピング */
    .message-container {
        display: flex;  /* フレックスボックスを使用 */
        flex-direction: column;  /* 子要素を縦方向に配置 */
        align-items: center;  /* 垂直方向に中央に配置 */
        gap: 10px;  /* 子要素間のスペース */
    }

    /* メッセージコンテンツ: オーバーレイメッセージ内の具体的なテキストやアイコンのスタイル */
    .message-content {
        padding: 20px 25px;  /* 上下左右のパディング */
        background-color: #68D391; /* 背景色を青緑に */
        color: white;  /* テキスト色を白に */
        border-radius: 10px;  /* 角を丸く */
        box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);  /* ボックスの影を追加 */
        position: relative;  /* 位置を相対的に */
        font-size: 1.2em;  /* フォントサイズを指定 */
        display: flex;  /* フレックスボックスを使用 */
        align-items: center;  /* 垂直方向に中央に配置 */
        justify-content: space-between;  /* 子要素を均等に配置 */
    }


    /* メッセージ内のスパン: メッセージのテキスト部分のスタイル */
    .message-content span {
        flex: 1;  /* フレックスアイテムとして均等に空間を取る */
        text-align: center;  /* テキストを中央に配置 */
    }

    /* クローズアイコン: メッセージを閉じるアイコンのスタイル */
    .close-icon {
        padding-left: 15px;  /* 左側のパディングを15pxに設定 */
        cursor: pointer;  /* マウスカーソルをポインタに変更 */
    }


    /* 登録・更新ボタンのスタイル */
    .button-submit {
        background-color: #68D391; /* Soft Blue-Green: ボタンの背景色 */
        color: white; /* ボタン内の文字の色 */
        min-height: 50px;  /* ボタンの最小の高さを50pxに指定 */
        min-width: 200px;  /* ボタンの最小の横幅を200pxに指定 */
        display: flex;  /* フレックスボックスのスタイルを適用: 要素の中身を中央配置するために使用 */
        align-items: center;  /* 子要素を中央に垂直配置 */
        justify-content: center;  /* 子要素を中央に水平配置 */
        padding: 10px 20px;  /* ボタンの内部の余白を上下10px、左右20pxに設定 */
    }

    /* ボタンがホバーされたときの背景色変更 */
    .button-submit:hover {
        background-color: #48BB78; /* Slightly Darker Blue-Green: ホバー時の背景色を少し暗くする */
    }

    /* ボタンがdisabled属性を持っているとき（例：ローディング時）のスタイル */
    .button-submit[disabled] {
        background-color: #A3E9C0; /* Lighter Blue-Green: 背景色をより薄くする */
        color: #FFFFFF; /* 文字の色 */
        cursor: not-allowed; /* カーソルをnot-allowedに変更して、クリックできないことを示す */
        display: flex;  /* フレックスボックスのスタイルを適用: 要素の中身を中央配置するために使用 */
        align-items: center;  /* 子要素を中央に垂直配置 */
        justify-content: center;  /* 子要素を中央に水平配置 */
    }

    /* ボタンを中央配置するためのコンテナのスタイル */
    .button-container {
        display: flex;        /* フレックスボックスを使用: 要素の中身を中央配置するために使用 */
        justify-content: center;  /* 子要素（ボタン）を中央に水平配置 */
        align-items: center;  /* 子要素（ボタン）を中央に垂直配置 */
        width: 100%;         /* コンテナの横幅を100%に設定 */
    }

    /* ヘッダーコンテナの基本的なスタイル */
    .header-container {
        padding: 5px; /* 内部の余白を設定 */
        display: flex; /* フレックスボックスを使用: 要素の中身を配置するために使用 */
        align-items: center; /* 子要素を中央に垂直配置 */
        gap: 10px; /* 子要素間の距離を10pxに設定 */
        margin-bottom: 20px; /* 下の要素との間隔を20pxに設定 */
    }

    /* ヘッダーコンテナ内のh2タグのスタイル */
    .header-container h2 {
        margin: 0; /* 余白を0に設定 */
        font-weight: bold;    /* フォントの太さをboldに設定 */
        color: #333;          /* タイトルの色を深いグレーに変更 */
    }

    /* 編集モードのヘッダーコンテナのスタイル */
    .header-container.editing {
        background-color: #f0f0f0; /* ソフトなグレーの背景色 */
        border: 1px solid #ccc; /* ボーダーを薄いグレーで追加 */
    }

    /* 作成モードのヘッダーコンテナのスタイル */
    .header-container.creating {
        background-color: #e6f7ff; /* ソフトなブルーの背景色 */
        border: 1px solid #b3d1ff; /* ボーダーを薄いブルーで追加 */
    }
</style>

<div class="container">
    <!-- メッセージのオーバーレイを表示する -->
    {#if showMessage}
        <div class="overlay-message">
            <div class="message-container">
                <!-- メッセージをクリックしたり、Enterキーを押した時にメッセージを閉じる -->
                <div
                        on:click={() => showMessage = false}
                        on:keydown={(event) => {
                    if (event.key === 'Enter') {
                        showMessage = false;
                    }
                }}
                        role="button"
                        tabindex="0"
                        class="message-content"
                >
                    <!-- メッセージ内容を表示 -->
                    <span>{message}</span>
                    <!-- メッセージを閉じるための×アイコン -->
                    <div class="close-icon">✕</div>
                </div>
            </div>
        </div>
    {/if}

    <!-- 一覧画面への遷移ボタンセクション -->
    <div class="section">
        <button on:click={navigateToList} class="p-2 bg-gray-600 text-white rounded-md mb-4">一覧に戻る</button>
    </div>

    <!-- 入力フォームと更新/登録ボタンのセクション -->
    <div class="section">
        <!-- ノートが存在する場合は編集モード、存在しない場合は作成モードのヘッダを表示 -->
        {#if note}
            <div class="header-container editing">
                <!-- 編集アイコン -->
                <Fa icon={faPencilAlt} class="text-gray-600" />
                <!-- 編集モードのタイトル -->
                <h2 class="text-xl font-bold mb-4">ノートの編集</h2>
            </div>
        {:else}
            <div class="header-container creating">
                <!-- 追加アイコン -->
                <Fa icon={faPlus} class="text-blue-500" />
                <!-- 作成モードのタイトル -->
                <h2 class="text-xl font-bold mb-4">新しいノートの作成</h2>
            </div>
        {/if}

        <!-- コメント入力フィールド -->
        <div class="mb-4">
            <label for="comment" class="block mb-2 font-bold text-lg">コメント</label>
            <!-- コメントテキスト入力エリア -->
            <input id="comment" type="text" bind:value={content} required placeholder="コメントを入力してください" class="p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600 w-full" />
            <!-- エラーメッセージ表示エリア -->
            {#if errorMessage}
                <p class="text-red-500 mt-2">{errorMessage}</p>
            {/if}
        </div>

        <!-- 画像アップロードエリア -->
        <div class="mb-4">
            <label for="file-upload" class="block mb-2 font-bold text-lg">画像アップロード</label>
            <!-- 画像ファイル入力エリア -->
            <input id="file-upload" type="file" on:change={handleFileUpload} class="p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600 w-full" />
        </div>

        <!-- 添付画像のリンク表示エリア -->
        {#if note?.attachmentURL}
            <div class="mb-4">
                <label for="attachment" class="block mb-2 font-bold text-lg">添付画像</label>
                <!-- 添付画像のリンク -->
                <a href="{note.attachmentURL}" target="_blank" rel="noopener noreferrer" class="text-blue-600 hover:text-blue-800">{note.attachment}</a>
            </div>
        {/if}

        <!-- 更新/登録ボタンエリア -->
        <div class="mt-4 button-container">
            <!-- ボタンの表示はローディング状態やノートの有無に基づいて変わる -->
            <button on:click={handleSubmit} class="button-submit" disabled={isLoading}>
                {#if isLoading}
                    <div class="loading-indicator"></div>
                    処理中...
                {:else if note}
                    更新
                {:else}
                    作成
                {/if}
            </button>
        </div>
    </div>
</div>
