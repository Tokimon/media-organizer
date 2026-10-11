<script lang="ts">
	import Thumbnail from '$lib/components/Thumbnail.svelte';
	import debounce from '$lib/tools/debounce.js';

	type Props = {
		files: string[];
		itemWidth?: number;
		narrow?: boolean;
	};

	const { files, itemWidth = 100, narrow = false }: Props = $props();

	const SPACING = 10;

	// biome-ignore lint/style/useConst: this will be set with the bind further down
	let itemHeight = $state(0);
	// biome-ignore lint/style/useConst: this will be set with the bind further down
	let listHeight = $state(0);
	// biome-ignore lint/style/useConst: this will be set with the bind further down
	let listWidth = $state(0);

	let scrollY = $state(0);

	const fullItemWidth = $derived(itemWidth + SPACING);
	const fullItemHeight = $derived(itemHeight + SPACING);

	const view = $derived.by(() => {
		const cols = itemWidth
			? Math.floor((listWidth - SPACING * 2) / fullItemWidth)
			: 1;
		const rows = itemHeight
			? Math.ceil((listHeight - SPACING * 2) / fullItemHeight)
			: 1;
		const width = fullItemWidth * cols - SPACING;
		const height = fullItemHeight * Math.ceil(files.length / cols);
		const extra = Math.min(Math.ceil(rows / 2), 3);

		return { cols, rows, extra, width, height };
	});

	const bounds = $derived.by(() => {
		const rowsInSpace = Math.floor((scrollY + SPACING) / fullItemHeight);
		const renderedRows = view.rows + view.extra * 2;

		const startRow = Math.max(rowsInSpace - view.extra, 0);
		const start = startRow * view.cols;
		const end = Math.min(start + renderedRows * view.cols, files.length);
		const count = end - start;

		return { start, end, startRow, count };
	});

	function initList(fileList: HTMLElement) {
		const scrollEnd = debounce(() => {
			console.log('SCROLL END');
		});

		function updateScroll() {
			scrollY = fileList.scrollTop;
			scrollEnd();
		}

		$effect(() => {
			fileList.addEventListener('scroll', updateScroll);
			updateScroll();

			return () => {
				fileList.removeEventListener('scroll', updateScroll);
			};
		});
	}
</script>

<!--
    TODO:
    - Add POC for the huge list problem in sveltekit (100K+ files)
    - Image Loading
    - Selection options (toggle selection mode ?):
        - click: No selection change
        - ctrl + click: Select one
        - shift + click: Select from last to here
        - shift + ctrl + click: Add from last to here to current selection
        - alt + click: 'block' selection (like drwing a square) from last to here
        - alt + ctrl + click: add 'block' selection (like drwing a square) from last to here to current selection
        - Same goes for keyboard selection (so replace click with moving 'cursor' around)
-->

<div
    class="file-list-container"
    bind:offsetWidth={listWidth}
    bind:offsetHeight={listHeight}
    use:initList
    style="--content-height: {view.height}px; width: {narrow ? '150rem' : 'auto'};"
>
    {#snippet fileContent(file: typeof files[number])}
        <Thumbnail src={file} />
        <!-- {#if group.count > 1}<span class="file-group-count">{group.count}</span>{/if} -->

        <!-- svelte-ignore a11y_figcaption_parent -->
        <figcaption class="file-label">
            {file}
        </figcaption>
    {/snippet}

    {#if files.length}
        <div class="file-list" style="--size: {itemWidth}px; --spacing: {SPACING}px; --position-offset: {bounds.startRow * fullItemHeight}px;">
            {#each { length: bounds.count }, i (i)}
                {@const file = files[i + bounds.start]}

                {#if i === 0}
                    <figure class="file" bind:offsetHeight={itemHeight} title={file}>
                        {@render fileContent(file)}
                    </figure>
                {:else}
                    <figure class="file" title={file}>
                        {@render fileContent(file)}
                    </figure>
                {/if}
            {/each}
        </div>
    {/if}
</div>

<style>
    .file-list-container {
        --content-height: 0;

        user-select: none;
        overflow: auto;
        flex: 1 1 100%;

        &:before {
            content: "";
            height: var(--content-height);
            width: 0;
            overflow: hidden;
            float: left;
        }
    }

    .file-list {
        --size: 100px;
        --spacing: 10px;
        --position-offset: 0;

        text-align: left;
        display: grid;
        grid-template-columns: repeat(auto-fit,var(--size));
        justify-content: center;
        gap: var(--spacing);
        padding: var(--spacing);
        box-sizing: border-box;
        overflow: auto;
        position: relative;
        translate: 0 var(--position-offset);
    }

    .file {
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        position: relative;
        color: white;
    }

    .file-label {
        padding: 0.02em;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        text-align: center;
    }

    .file-group-count {
        font-size: 1rem;
        border-radius: 50%;
        inline-size: 2rem;
        aspect-ratio: 1;
        align-content: center;
        text-align: center;
        background: olivedrab;
        position: absolute;
        inset: auto 0.5rem 2.5rem auto;
        color: white;
        box-shadow: inset 0 0 0.3rem rgb(0 0 0 / 50%);
    }
</style>
