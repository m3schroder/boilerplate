const defaultTheme = require('tailwindcss/defaultTheme')

const config = {
	content: [
		"./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",
		"./src/routes/**/*.{html,js,svelte,ts}",
		"./src/components/**/*.{html,js,svelte,ts}"
	],
	safelist: [
		"!bg-amber-100",
		"!bg-neutral-100",
		"!bg-orange-100"
	],
	theme: {
		extend: {
			keyframes: {
				rotate180: {
					'0%': { transform: 'rotate(180deg)' },
					'100%': { transform: 'rotate(0deg)' }
				}
			},
			fontFamily: {
				'sans': ['Merriweather', 'serif'],
			},
			animation: {
				'rotate-180': 'rotate180 300ms linear infinite'
			},
			borderWidth: {
				"1": "1px"
			},
			boxShadow: {
				neu: '2px 3px 0px #96969625',
			},
			colors: {
				'blue': {
					DEFAULT: "#c3a765",
					'50': '#fbfaf1',
					'100': '#f5f2df',
					'200': '#e9e2bf',
					'300': '#dbcd96',
					'400': '#cbb26c',
					'500': '#c5a55c',
					'600': '#b38943',
					'700': '#956e39',
					'800': '#785834',
					'900': '#62492c',
				},
				"secondary": {
					DEFAULT: "#A1A1A1",
					dark: "#A1A1A1"
				},
				"paper": {
					DEFAULT: "#EAEAEA"
				},
			}
		}
	},
	plugins: [
		require('flowbite/plugin')
	],
	darkMode: 'class',
};

module.exports = config;