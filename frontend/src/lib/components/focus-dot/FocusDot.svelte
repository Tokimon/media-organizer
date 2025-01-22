<script lang="ts">
    import type { PropsWithChildren } from "~/types.d";
    const props: PropsWithChildren = $props();
</script>

<div class="focus-dot {props.class}" style={props.style}>
    {@render props.children()}
</div>

<style>
    .focus-dot {
        font-size: 1rem;
        border-radius: 0.8rem;
        height: 1.6rem;
        min-width: 1.6rem;
        line-height: 1;
        padding: 0 5px;
        color: white;
        background: var(--secondary);
        display: inline-flex;
        box-sizing: border-box;
        align-items: center;
        justify-content: center;
        font-weight: bold;
    }
</style>
