import { writable } from 'svelte/store';

// ログイン状態を管理するストア
export const loggedIn = writable(false);