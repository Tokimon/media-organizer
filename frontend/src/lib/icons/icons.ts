export const iconNames = [
	'add',
	'close',
	'delete',
	'search',
	'settings'
] as const;

export type IconNames = typeof iconNames[number];