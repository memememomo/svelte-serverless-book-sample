// AWS AmplifyのStorageモジュールをインポート
import { Storage } from 'aws-amplify';

// S3へのアップロードを行う非同期関数
export async function s3Upload(file: File) {
    // ファイル名を一意にするために、現在のタイムスタンプを接頭語として付ける
    const filename = `${Date.now()}-${file.name}`;

    // AmplifyのStorageモジュールを使用して、ファイルをS3にアップロード
    // Storage.vault.putメソッドを使用することで、プライベートなS3バケットにファイルをアップロード
    // contentTypeには、ファイルのMIMEタイプを指定
    const stored = await Storage.vault.put(filename, file, {
        contentType: file.type,
    });

    // アップロードしたファイルのS3キーを返す
    return stored.key;
}
