<script lang="ts">
    // 必要なモジュールやコンポーネントをインポート
    import { onMount } from 'svelte'; // svelteのライフサイクルメソッド
    import { page } from '$app/stores'; // SvelteKitのページストア
    import NoteEditor from '$lib/components/NoteEditor.svelte'; // ノートエディターコンポーネント
    import type { NoteType } from "$lib/api"; // NoteType型のインポート
    import { getNoteById } from "$lib/api"; // 特定のIDによるノートの取得メソッド
    import { Storage } from 'aws-amplify'; // AWS Amplifyのストレージモジュール
    import Loading from "$lib/components/Loading.svelte"; // ローディングコンポーネント

    let note: NoteType | null = null; // ノートデータを格納する変数

    // pageストアからnoteIdを取得
    let noteId = $page.params.noteId;

    onMount(async () => { // コンポーネントがマウントされたときに実行される
        try {
            note = await getNoteById(noteId); // APIを使ってnoteIdに対応するノートを取得

            // ノートに添付ファイルがある場合、そのURLを取得
            if (note?.attachment) {
                note.attachmentURL = await Storage.vault.get(note.attachment);
            }
        } catch (error) { // エラー処理
            console.error("Error fetching note:", error);
            // エラーのハンドリングや、ユーザーを適切な場所にリダイレクトするロジックをここに記述
        }
    });
</script>

{#if note} <!-- ノートが存在する場合 -->
    <NoteEditor {note} /> <!-- NoteEditorコンポーネントにノートデータを渡して表示 -->
{:else} <!-- ノートがまだ読み込まれていない、またはエラーが発生した場合 -->
    <Loading /> <!-- ローディングコンポーネントを表示 -->
{/if}
