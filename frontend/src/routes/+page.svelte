<script lang="ts">
    import Icon from "$lib/components/icon/Icon.svelte";
    import { SelectFiles, SelectDirectory } from "$lib/wailsjs/go/api/Api";
    import fileStore from "$lib/stores/file-store.svelte.js";

    async function selectFiles() {
        const files = await SelectFiles();
        fileStore.addFiles(files);
    }

    async function selectDirectory() {
        const files = await SelectDirectory();
        fileStore.addFiles(files);
    }

    type FileNameGroup = {
      name: string
      count: number
      paths: string[]
    }

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
              group.count++
            }
        }

        const names = Array.from(groups.keys()).toSorted((a, b) => {
            const A = a.toLowerCase();
            const B = b.toLowerCase();

            if (A === B) return 0;
            return A > B ? 1 : -1;
        });

        return { names, groups };
    });
</script>

<h1>Media organizer</h1>

<button onclick={selectFiles}>
    <Icon icon="add" />
    Click to open file selector
</button>

<button onclick={selectDirectory}>
    <Icon icon="add" />
    Click to open directory selector
</button>

<div class="file-list">
    {#each names as name}
        {@const group = groups.get(name)!}
        <figure class="file">
            <img class="file-thumb" src='wails://?thumb={group.paths[0]}' style="width: 100px; aspect-ratio: 1; background: gray;" alt="" />

            <figcaption class="file-label">
                {name}
                {#if group.count > 1}<span class="file-group-count">{group.count}</span>{/if}
            </figcaption>
        </figure>
    {:else}
        No files selected yet
    {/each}
</div>

<style>
    .file-list {
        text-align: left;
        color: white;
        display: grid;
        grid-template-columns: repeat(
            auto-fit,
            calc(10rem * var(--pct, 100) / 100)
        );
        justify-content: center;
        gap: 1rem;
        user-select: none;
        padding: 1rem;
    }

    .file {
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .file-thumb {
        max-width: 100%;
        aspect-ratio: 1;
        object-fit: cover;
        overflow: clip;
        border-radius: 1rem;
    }

    .file-label {
        padding: 0.02em;
        position: relative;
    }

    .file-group-count {
        font-size: 1rem;
        border-radius: 50%;
        inline-size: 2rem;
        aspect-ratio: 1;
        align-content: center;
        text-align: center;
        background: olivedrab;
        translate: -0.4rem -1.3rem;
        position: absolute;
        inset: auto 0 100% auto;
        color: white;
        box-shadow: inset 0 0 0.3rem rgb(0 0 0 / 50%);
    }
</style>
