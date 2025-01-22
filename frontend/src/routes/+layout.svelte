<script lang="ts">
  import { OnFileDrop } from '$lib/wailsjs/runtime/runtime';
  import { OnFileDrop as ApiFileDrop } from '$lib/wailsjs/go/api/Api';
  import '$lib/styles/index.css'
  const {children} = $props();

  for (const ev of ["dragenter", "dragover", "drop"])
    window.addEventListener(ev, (e) => e.preventDefault(), false)

  OnFileDrop(ApiFileDrop, true)
</script>

<div class="drop-target"></div>

{@render children()}

<style>
    :root {
      --wails-drop-target: none;
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
