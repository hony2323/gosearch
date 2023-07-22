<script lang="ts">
    import type { main } from "./../../wailsjs/go/models";

    import Item from "./item.svelte";

    export let itemList: main.DirEntryJson[];
</script>

<div class="Folder">
    {#each itemList as itm}
        <Item dirEntry={itm} />
    {/each}
</div>

<style lang="scss">
    .Folder {
        display: flex;
        flex-direction: column;
    }
</style>
