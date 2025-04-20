<script lang='ts'>
  const props = $props()
</script>

<button class="button" type="button" {...props}>
    {@render props.children()}
</button>

<style>
    .button {
        all: unset;
        display: flex;
        gap: 0.5em;
        padding: 0.3em 0.5em;
        border: 2px solid currentColor;
        border-radius: 0.3em;
        align-items: center;
        cursor: pointer;

        &:focus-visible {
            border-color: limegreen;
            outline: none;
        }

        &:hover {
            background: rgb(255 255 255 / 30%);
        }
    }
</style>
