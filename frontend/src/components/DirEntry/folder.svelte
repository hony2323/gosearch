<script lang="ts">
    import type { types } from "../../../wailsjs/go/models";

    import Item from "./item.svelte";

    export let itemList: types.DirEntry[];
</script>

<div class="Folder">
    {#each itemList as itm}
        <Item dirEntry={itm} />
    {/each}
</div>

<!--        Styling        -->
<style lang="scss">
    .Folder {
        display: flex;
        flex-direction: column;
    }
</style>
