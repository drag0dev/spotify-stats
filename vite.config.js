import { sveltekit } from '@sveltejs/kit/vite';
import autoPreprocess from 'svelte-preprocess';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()]
};

export default config;
