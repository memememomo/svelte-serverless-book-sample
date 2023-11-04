<script lang="ts">
    // svelteと関連ライブラリから必要な関数やコンポーネントをインポート
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    // FontAwesomeのアイコンをインポート
    import { faPencilAlt, faTrash } from '@fortawesome/free-solid-svg-icons';
    import Fa from 'svelte-fa';
    // ノートの型とノート関連の関数をインポート
    import type { NoteType } from "$lib/api";
    import { listNotes, deleteNote } from "$lib/api";  // deleteNote関数も存在すると仮定
    // ローディングコンポーネントをインポート
    import Loading from "$lib/components/Loading.svelte";

    // 初期のローディング状態とノートリスト、削除中のノートIDを定義
    let isLoading = true;
    let notes: NoteType[] = [];
    let deletingNoteId: string | null = null;

    // コンポーネントがマウントされたときにノートのロードを開始
    onMount(async () => {
        loadNotes();
    });

    // すべてのノートをロードする関数
    async function loadNotes() {
        try {
            // ノートリストをAPIから取得
            notes = await listNotes() || [];
        } catch (err) {
            // エラー発生時にコンソールにログを出力
            console.error("Error fetching notes:", err);
        } finally {
            // ローディング状態をfalseに更新
            isLoading = false;
        }
    }

    // 指定されたノートを削除する関数
    async function handleDelete(noteId: string) {
        // ユーザーにノートの削除を確認
        const confirmed = window.confirm("Are you sure you want to delete this note?");
        if (confirmed) {
            try {
                // 削除中のノートIDを更新
                deletingNoteId = noteId;
                // ノートを削除
                await deleteNote(noteId);
                // 削除後、ノートリストを再ロード
                await loadNotes();
            } catch (err) {
                // エラー発生時にコンソールにログを出力
                console.error("Error deleting note:", err);
            } finally {
                // 削除中のノートIDを初期化
                deletingNoteId = null;
            }
        }
    }

    // 日付文字列をフォーマットする関数
    function formatDate(str: undefined | string): string {
        // 日付が存在しない場合は空文字を返す、存在する場合は日時の形式に変換
        return !str ? "" : new Date(str).toLocaleString();
    }
</script>


<style>
    /* ノートのリストを表示するコンテナのスタイル定義 */
    .notes-container {
        max-width: 800px; /* 画面の幅に応じて最大幅を設定 */
        margin: 0 auto;   /* 水平方向に中央寄せするためのマージン設定 */
        padding: 1rem;   /* コンテンツの周りに余白を追加 */
    }
</style>

<div class="notes-container">
    {#if isLoading}
        <!-- データのロード中の場合の処理 -->
        <Loading/>
    {:else}
        <!-- データがロードされた後の処理 -->
        <!-- 新しいノートを作成するためのボタン -->
        <div
            role="button"
            tabindex="0"
            on:click={() => goto('/notes/new')}
            on:keydown={(event) => {
                if (event.key === 'Enter') goto('/notes/new');
            }}
            class="flex items-center py-3 px-4 cursor-pointer hover:bg-gray-100 border-b"
        >
        <Fa icon={faPencilAlt} class="text-gray-600 mr-2"/> <!-- ペンシルアイコンを表示 -->
        <span class="text-lg text-gray-700">Create a new note</span>
        </div>

        <!-- すべてのノートをリスト表示 -->
        {#each notes as { noteId, content, createdAt }}
            <div class="flex justify-between py-3 px-4 cursor-pointer hover:bg-gray-100 border-b">
                <!-- ノートの内容と作成日時を表示 -->
                <div
                    role="button"
                    tabindex="0"
                    on:click={() => goto(`/notes/${noteId}`)}
                    on:keydown={(event) => {
                        if (event.key === 'Enter') goto(`/notes/${noteId}`);
                    }}
                    class="flex flex-col"
                >
                    <span class="text-lg text-gray-800 mb-1">{content.trim().split("\n")[0]}</span> <!-- ノートの最初の行のみを表示 -->
                    <span class="text-sm text-gray-600">Created: {formatDate(createdAt)}</span> <!-- ノートの作成日時を表示 -->
                </div>
                <!-- 削除ボタンの表示と処理 -->
                {#if deletingNoteId === noteId} <!-- 削除処理中のノートの場合 -->
                    <span class="text-red-600">削除中...</span> <!-- 削除中のテキストを表示 -->
                {:else} <!-- それ以外の場合、削除ボタンを表示 -->
                    <button on:click={() => noteId && handleDelete(noteId)} class="text-red-600 hover:text-red-800 focus:outline-none">
                        <Fa icon={faTrash} />
                    </button>
                {/if}
            </div>
        {/each}
    {/if}
</div>