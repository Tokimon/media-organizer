<script lang="ts">
    import Button from "$lib/components/Button.svelte";
    import Icon from "$lib/components/Icon.svelte";
    import fileStore from "$lib/stores/file-store.svelte.js";
    import { SelectDirectory, SelectFiles } from "$lib/wailsjs/go/api/Api";

    async function selectFiles() {
        const files = await SelectFiles();
        fileStore.update(files);
    }

    async function selectDirectory() {
        const files = await SelectDirectory();
        fileStore.update(files);
    }
</script>

<header>
    {fileStore.files.size} Files

    <Button onclick={selectFiles}>
        <Icon icon="add_image" />
    </Button>

    <Button onclick={selectDirectory}>
        <Icon icon="add_folder" />
    </Button>
</header>

<style>
    header {
        font-size: 2rem;
        display: flex;
        gap: 2rem;
        align-items: center;
        background: #333;
        padding: 0.5em 1em;
        position: sticky;
        inset: 0 0 auto 0;
        z-index: 9;
        box-shadow: 0 2px 1rem rgb(0 0 0 / 50%);
    }
</style>
