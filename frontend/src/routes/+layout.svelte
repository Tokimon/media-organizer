<script lang="ts">
  import AppHeader from '$lib/components/AppHeader.svelte';
  import fileStore from '$lib/stores/file-store.svelte.js';
  import { SelectFromDrop } from '$lib/wailsjs/go/api/Api';
  import { OnFileDrop } from '$lib/wailsjs/runtime/runtime';
  import { onMount } from 'svelte';

  import '$lib/styles/index.css';

  const { children } = $props();
  let showDropArea = $state(false)

  window.addEventListener("dragover", (e) => e.preventDefault())

  window.addEventListener("drop", (e) => e.preventDefault(), false)
  document.body.addEventListener("dragleave", (e) => e.preventDefault(), false)

  window.addEventListener("dragenter", (e) => {
    e.preventDefault()
    showDropArea = true
  }, false)

  OnFileDrop(async (...args) => {
    try {
      const files = await SelectFromDrop(...args)
      fileStore.update(files)
    } catch(err) {
      // TODO: Better error handling
      console.error(err)
    }

    showDropArea = false
  }, true)

  onMount(() => fileStore.init())
</script>

<AppHeader />

{#if showDropArea}
    <div class="drop-target"></div>
{/if}

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
