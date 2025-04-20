export const iconNames = [
	'add',
	'add_folder',
	'add_image',
	'book',
	'camera_roll',
	'close',
	'code',
	'data',
	'delete',
	'hd',
	'image_square',
	'play_circle',
	'save',
	'search',
	'settings',
	'settings_2',
	'settings_square',
	'storage',
	'unknown',
	'unknown_square',
	'vector',
	'vector_2',
	'volume_up'
] as const;

export type IconNames = typeof iconNames[number];