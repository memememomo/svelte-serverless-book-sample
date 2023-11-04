import type { PageLoad } from './$types';
import {getCount} from "$lib";

// load関数は、SvelteKitの機能により、コンポーネントがレンダリングされる前に実行される
export const load: PageLoad = async () => {
    // カウント取得APIでカウントを取得
    return await getCount();
};