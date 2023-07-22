<script lang="ts">
    import type { main } from "./../../wailsjs/go/models";
    import { path } from "./../store";
    export let dirEntry: main.DirEntryJson;
    function changePath(newEntryName) {
        path.set($path + newEntryName + "/");
    }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
    on:click={() => {
        debugger;
        changePath(dirEntry.name);
    }}>
    {dirEntry.name}
</div>
