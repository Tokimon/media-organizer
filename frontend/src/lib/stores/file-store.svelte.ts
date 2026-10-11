import { onMount } from 'svelte';
import { SvelteMap } from 'svelte/reactivity';
import { EventNames } from '$lib/events.js';
import { GetExtensions, RequestAllFiles } from '$lib/wailsjs/go/api/Api.js';
import type { stores } from '$lib/wailsjs/go/models.js';
import { EventsOn } from '$lib/wailsjs/runtime/runtime.js';

class FileStore {
	files = new SvelteMap<string, stores.FileEntry>();
	extensions = $state<string[]>([]);
	loading = $state(false);

	resetFiles(files: stores.FileEntry[]) {
		this.files.clear();
		this.addFiles(files);
	}

	addFiles(files: stores.FileEntry[]) {
		for (const file of files) this.files.set(file.path, file);
	}

	init() {
		onMount(() => {
			GetExtensions().then((ext) => {
				this.extensions = ext;
			});

			const removeFileLoadBegin = EventsOn(EventNames.fileLoadingBegin, () => {
				this.loading = true;
			});

			const removeFileLoadEnd = EventsOn(EventNames.fileLoadingEnd, (files) =>
				this.resetFiles(files),
			);

			RequestAllFiles().then((files) => this.resetFiles(files));

			return () => {
				removeFileLoadBegin();
				removeFileLoadEnd();
			};
		});
	}
}

export default new FileStore();
