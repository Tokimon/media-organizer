<script lang="ts">
	import FileStore from '$lib/stores/file-store.svelte.js';

	const { open } = $props();
</script>

<dialog open={open}>Loading files...</dialog>
