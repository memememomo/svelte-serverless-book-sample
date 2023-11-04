import { expect, test } from '@playwright/test';

test('index page redirects to /notes', async ({ page }) => {
	// トップページにアクセス
	await page.goto('/');

	// ページが正しく'/notes'にリダイレクトされたことを確認
	const currentURL = page.url();
	expect(currentURL).toContain('/login');
});
