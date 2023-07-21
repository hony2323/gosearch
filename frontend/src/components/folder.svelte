<script lang="ts">
    import Item from "./item.svelte";

    export let itemList;
</script>

{#each itemList as itm}
    <Item date="0" icon="asd" name={itm.name} size=3 />
{/each}
