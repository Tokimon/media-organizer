<script lang="ts">
    import stringEndSplit from "~/tools/stringEndSplit";
    import type { Stylable } from "~/types.d";

    type Props = Stylable & {
        path: string;
        separator?: string;
        placeholder?: string;
    };

    const {
        path,
        separator = "/",
        placeholder = "No path ...",
        style,
        ...props
    }: Props = $props();

    const [beginStr, endStr] = $derived(stringEndSplit(path, separator));
</script>

<span
    class="path-string {props.class}"
    class:empty={!endStr && !beginStr}
    {style}
>
    {#if beginStr}<span class="beginning">{beginStr}</span>{/if}
    {#if endStr}<span class="end">{separator + endStr}</span>{/if}
    {#if !endStr && !beginStr}<span class="end">{placeholder}</span>{/if}
</span>

<style>
    .path-string {
        white-space: nowrap;
        display: flex;
        box-sizing: border-box;
        overflow: hidden;

        &.empty {
            color: var(--light-text);
        }
    }

    .beginning {
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .end {
        flex: 0;
    }
</style>
