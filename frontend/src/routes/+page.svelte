<script lang="ts">
    import Icon from "$lib/components/icon/Icon.svelte";
    import { AddFiles, AddDirectory } from "$lib/wailsjs/go/api/Api";

    async function loadFiles() {
        const result = await AddFiles();
        console.log({ result });
    }

    async function loadDirectory() {
        const result = await AddDirectory();
        console.log({ result });
    }
</script>

<h1>Media organizer</h1>

<button onclick={loadFiles}>
    <Icon icon="add" />
    Click to open file selector
</button>

<button onclick={loadDirectory}>
    <Icon icon="add" />
    Click to open directory selector
</button>
