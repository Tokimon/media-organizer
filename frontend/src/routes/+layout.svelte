<script lang="ts">
	import AppHeader from '$lib/components/AppHeader.svelte';
	import LoadingDialog from '$lib/components/dialogs/LoadingDialog.svelte';
	import DnD from '$lib/scripts/DnD.svelte.js';
	import FileStore from '$lib/stores/file-store.svelte.js';

	import '$lib/styles/index.css';

	const { children } = $props();

	DnD.init();
	FileStore.init();
</script>

<AppHeader />

{#if DnD.dragOver}
    <div class="drop-target"></div>
{/if}

<LoadingDialog open={FileStore.loading} />

<main>{@render children()}</main>

<style>
    :root {
      --wails-drop-target: none;
      font-size: 62.5%;
      position: relative;
    }

    :global(body) {
        font-size: 1.4rem;
        display: flex;
        flex-direction: column;
        height: 100dvh;
    }

    .drop-target {
        --wails-drop-target: drop;
        border: 1px dashed #ccc;
        position: fixed;
        inset: 2em;
        background: rgb(255 255 255 / 50%);
        border-radius: 2em;
    }

    main {
        display: flex;
        flex-direction: column;
        flex: 1;
        overflow: clip;
    }
</style>
