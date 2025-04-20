<script lang="ts">
    import Icon from "$lib/components/Icon.svelte";
    import fileStore from "$lib/stores/file-store.svelte.js";
    import extensionToIcon from "$lib/tools/extensionToIcon.js";

    type Props = {
        ext: string;
        className?: string;
        style?: string;
    };

    const { ext, className, style }: Props = $props();
</script>

<div class={["image-placeholder", className]} {style}>
    <Icon icon={extensionToIcon(ext)} />
</div>

<style>
    .image-placeholder {
        color: var(--light-text);
        align-content: center;
        justify-content: center;
        container-type: inline-size;
        background: rgb(163 162 162 / 70%);
        aspect-ratio: 1;
        border-radius: 1rem;

        > :global(.icon) {
            display: block;
            width: clamp(1rem, 50%, 35rem);
            margin: 0 auto;
        }
    }
</style>
