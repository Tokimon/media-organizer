 <script lang="ts">
  import { OnFileDrop } from '$lib/wailsjs/runtime/runtime';
  import { SelectFromDrop } from '$lib/wailsjs/go/api/Api';
  import fileStore from '$lib/stores/file-store.svelte.js'

  import '$lib/styles/index.css'

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
      fileStore.addFiles(files)
    } catch(err) {
      // TODO: Better error handling
      console.error(err)
    }

    showDropArea = false
  }, true)
</script>

{#if showDropArea}
    <div class="drop-target"></div>
{/if}

{@render children()}

<style>
    :root {
      --wails-drop-target: none;
      font-size: 62.5%;
    }

    .drop-target {
        --wails-drop-target: drop;
        border: 1px dashed #ccc;
        position: fixed;
        inset: 2em;
        background: rgb(255 255 255 / 50%);
        border-radius: 2em;
    }
</style>
