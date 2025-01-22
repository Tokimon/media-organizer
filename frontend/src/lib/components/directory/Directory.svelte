<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";
    import type { Stylable } from "~/types";
    import FocusDot from "~/components/focus-dot/FocusDot.svelte";
    import Path from "~/components/path/Path.svelte";
    import SvgIcon from "~/components/svg-icon/SVGIcon.svelte";

    type Props = Stylable & {
        onclick?: MouseEventHandler<HTMLButtonElement>;
        path: string;
        count?: number;
    };

    const { path = "", count = 0, onclick, style, ...props }: Props = $props();

    const unixPath = $derived(path.replaceAll("\\", "/").replace(/\/+$/, ""));
</script>

<button type="button" class="display {props.class}" {onclick} {style}>
    <SvgIcon svg="folder" class="icon" />
    {#if unixPath}<FocusDot class="counter">{count}</FocusDot>{/if}
    <Path class="folder" path={unixPath} placeholder="No folder selected ..." />
</button>

<style>
    .icon {
        width: auto;
        flex: 0 0 auto;
        display: block;
        fill: var(--primary);
        grid-area: icon;
        padding: 0.5rem;
    }

    .counter {
        grid-area: icon;
        color: var(--secondary);
        background: rgb(255 255 255 / 85%);
        place-self: center;
        translate: 0 2px;
    }

    .folder {
        grid-area: folder;
        padding: 0 0.3rem;
    }

    .display {
        display: grid;
        grid-template-columns: 4.3rem 1fr;
        grid-template-areas: "icon folder";
        align-items: center;
        margin: 0;
        overflow: hidden;
        background: none;
        border: none;
        cursor: pointer;
        color: var(--light-text);
        font-size: 2rem;
    }
</style>
