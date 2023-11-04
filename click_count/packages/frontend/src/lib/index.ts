import {PUBLIC_API_ENDPOINT} from "$env/static/public";
const endpoint = () => `${PUBLIC_API_ENDPOINT}/count`;

// カウントAPIのレスポンスの型
type ResponseCount = {
    count: number;
};

// カウント取得APIの実行関数
export const getCount = async (): Promise<ResponseCount> => {
    // fetch関数を使ってAPIを実行
    const res = await fetch(endpoint(), {
        method: 'GET',
        mode: 'cors',
    });
    // レスポンスJSONを返す
    return res.json();
};

// カウント更新APIの実行関数
export const updateCount = async (): Promise<ResponseCount> => {
    // fetch関数を使ってカウント更新APIを実行
    const res = await fetch(endpoint(), {
        method: 'POST',
        mode: 'cors',
    });
    // レスポンスJSONを返す
    return res.json();
}