<script lang='ts' module>
console.log('MODULE INNIT')
console.time('files')
const files = Array.from(Array(25000), (_, i) => i + ".png");
console.timeEnd('files')
</script>

<script lang="ts">
    import Thumbnail from "$lib/components/Thumbnail.svelte";

    type Props = {
        items?: string[];
        itemWidth?: number;
        narrow?: boolean;
    };

    const { items, itemWidth = 100, narrow = false }: Props = $props();

    const SPACING = 10;

    // biome-ignore lint/style/useConst: this will be set with the bind further down
    let itemHeight = $state(0);
    // biome-ignore lint/style/useConst: this will be set with the bind further down
    let listHeight = $state(0);
    // biome-ignore lint/style/useConst: this will be set with the bind further down
    let listWidth = $state(0);

    // biome-ignore lint/style/useConst: this will be set with the onscroll event further down
    let scrollY = $state(0)

    const fullItemWidth = $derived(itemWidth + SPACING)
    const fullItemHeight = $derived(itemHeight + SPACING)

    const view = $derived.by(() => {
      const cols = itemWidth ? Math.floor((listWidth - (SPACING * 2)) / fullItemWidth) : 1;
      const rows = itemHeight ? Math.ceil((listHeight - (SPACING * 2)) / fullItemHeight) : 1;
      const width = (fullItemWidth * cols) - SPACING;
      const height = fullItemHeight * Math.ceil(files.length / cols)
      const extra = Math.min(Math.ceil(rows / 2), 3)

      return { cols, rows, extra, width, height }
    })

    const bounds = $derived.by(() => {
      const rowsInSpace = Math.floor((scrollY + SPACING) / fullItemHeight)
      const startRow = Math.max(rowsInSpace - view.extra, 0)
      const viewFullRowCount = view.rows + (view.extra * 2)
      const start = startRow * view.cols;
      const end = Math.min(start + (viewFullRowCount * view.cols), files.length);

      return { start, end, startRow }
    });
</script>

<!--
    TODO:
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


<!-- <div class="size" style="width: {view.width}px; height: {itemHeight}px;">
    <p>List: {listWidth} / {listHeight} / {SPACING}</p>
    <p>Item: {itemWidth} / {itemHeight} / {fullItemHeight}</p>
    <p>Info: {view.cols} / {view.rows}/ {view.width} / {view.height}</p>
    <p>Bounds: {bounds.start} / {bounds.end}</p>
</div> -->

<div
    class="file-list-container"
    bind:offsetWidth={listWidth}
    bind:offsetHeight={listHeight}
    onscroll={(e) => { scrollY = (e.target as HTMLElement).scrollTop; console.log('setting scroll!') }}
    onscrollend={() => {console.log('END!'); requestAnimationFrame(() => console.log('REAL END!'))}}
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
        <!-- {@const start = (space.rowsBefore * view.cols) || 0} -->
        <div class="file-list" style="--size: {itemWidth}px; --spacing: {SPACING}px; --position-offset: {bounds.startRow * fullItemHeight}px;">
            {#each files.slice(bounds.start, bounds.end) as file, i (file)}
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

    /* .size {
        position: absolute;
        padding: 10px;
        background: rgb(200 100 0 / 70%);
        inset: 10px 0 auto 0;
        margin: 0 auto;
        color: black;
        z-index: 9;
        box-sizing: border-box;
        gap: 30px;
        display: flex;
        justify-content: center;
        align-items: center;
        text-shadow: 0 0 8px white;
        font-weight: bolder;
        pointer-events: none;

        p {
            margin: 0;
        }
    } */
</style>
