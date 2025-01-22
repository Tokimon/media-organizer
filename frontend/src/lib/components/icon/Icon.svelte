<script lang="ts">
    import type { IconNames } from "$lib/icons/icons";

    type Props = { icon: IconNames };
    const { icon, ...props }: Props = $props();
</script>

<i class="icon-{icon}"></i>
