import preprocess from "svelte-preprocess";
import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/kit/vite';
import path from 'path'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: [vitePreprocess(), preprocess({
		postcss: true
	})],

	kit: {
		adapter: adapter(),
		alias: {
			"@lib": path.resolve("./src/lib"),
			"@ui": path.resolve("./src/components"),
			"@src": path.resolve("./src"),
		}
	}
};

export default config;
