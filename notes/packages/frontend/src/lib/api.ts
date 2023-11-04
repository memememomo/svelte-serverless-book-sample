// AWS AmplifyのAPIモジュールをインポート
import { API } from "aws-amplify";

// Noteの型定義
export type NoteType = {
    noteId?: string;        // ノートのID（オプショナル）
    content: string;        // ノートの内容
    createdAt?: string;     // ノートの作成日時（オプショナル）
    attachment?: string;    // 添付ファイル名
    attachmentURL?: string; // 添付ファイルのURL
};

// APIの名前とベースURLを定数として定義
const apiName = "notes";
const baseUrl = "/notes";

// ノートIDを元にエンドポイントURLを生成する関数
const endpointWithId = (id: string) => `${baseUrl}/${id}`;

// ノートを新規作成する関数
export const createNote = async (note: NoteType) => {
    return API.post(apiName, baseUrl, {
        body: note,
    });
};

// すべてのノートをリストとして取得する関数
export const listNotes = async (): Promise<NoteType[]> => {
    return await API.get(apiName, baseUrl, {});
};

// 既存のノートを更新する関数
export const saveNote = async (note: NoteType) => {
    await API.put(apiName, endpointWithId(note.noteId!), {
        body: note,
    });
    // 更新後のノート情報を取得して返す
    return await getNoteById(note.noteId!);
};

// ノートを削除する関数
export const deleteNote = async (id: string) => {
    await API.del(apiName, endpointWithId(id), {});
};

// IDを指定してノート情報を取得する関数
export const getNoteById = async (id: string): Promise<NoteType> => {
    return await API.get(apiName, endpointWithId(id), {});
}
