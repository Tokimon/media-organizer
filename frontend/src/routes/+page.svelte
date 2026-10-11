<script lang="ts">
	import FileList from '$lib/components/FileList.svelte';
	import fileStore from '$lib/stores/file-store.svelte.js';

	type FileNameGroup = {
		name: string;
		count: number;
		paths: string[];
	};

	const { names, groups } = $derived.by(() => {
		const groups = new Map<string, FileNameGroup>();

		for (const file of fileStore.files.values()) {
			const group = groups.get(file.name);

			if (!group) {
				groups.set(file.name, {
					name: file.name,
					count: 1,
					paths: [file.path],
				});
			} else {
				group.count++;
			}
		}

		const names = Array.from(groups.keys()).toSorted((a, b) =>
			a.toLocaleLowerCase().localeCompare(b.toLocaleLowerCase()),
		);

		return { names, groups };
	});

	let itemWidth = $state(50);
	let narrow = $state(false);

	function setItemWidth() {
		itemWidth = itemWidth === 50 ? 157 : 50;
	}

	function setNarrow() {
		narrow = !narrow;
	}
</script>

<div class="buttons">
    <button onclick={setItemWidth}>Change item size</button>
    <button onclick={setNarrow}>Change list size</button>
</div>

<FileList files={[]} {itemWidth} {narrow} />
