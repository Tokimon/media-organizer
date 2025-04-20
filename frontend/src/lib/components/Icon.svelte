<script lang="ts">
    import type { IconNames } from "$lib/icons/icons";

    type Props = { icon: IconNames; className?: string };
    const { icon, className }: Props = $props();
</script>

<i class={["icon", "icon-" + icon, className]}></i>
