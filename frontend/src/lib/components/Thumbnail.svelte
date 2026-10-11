<script lang="ts">
	import fileStore from '$lib/stores/file-store.svelte.js';
	import isWebImageExtension from '$lib/tools/isWebImageExtension.js';
	import ImagePlaceholder from './ImagePlaceholder.svelte';

	type Props = Partial<
		Pick<HTMLImageElement, 'className' | 'src' | 'alt'> & { style: string }
	>;
	const props: Props = $props();

	let failed = $state(false);

	function imageError() {
		failed = true;
	}

	const ext = $derived.by(() => {
		if (!props.src) return '';

		const fileInfo = fileStore.files.get(props.src);
		return fileInfo?.extension ?? props.src.split('.').at(-1) ?? '';
	});
</script>

{#if false && !failed && isWebImageExtension(ext)}
    <img class="image" {...props} onerror={imageError} />
{:else}
    <ImagePlaceholder
        {ext}
        className="{props.className} {failed ? 'failed' : ''}"
        style={props.style}
    />
{/if}

<style module>
    .image {
        display: block;
        object-fit: cover;
        object-position: 50% 50%;
        color: white;
        border-radius: 1rem;
    }
</style>
