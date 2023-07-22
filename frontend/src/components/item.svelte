<script lang="ts">
    import type { main } from "./../../wailsjs/go/models";
    import { path } from "./../store";
    export let dirEntry: main.DirEntryJson;

    function changePath(newEntryName) {
        path.set($path + newEntryName + "/");
    }

    function handleDirEntryClick(dirEntry: main.DirEntryJson) {
        if (dirEntry.isDir)
            changePath(dirEntry.name);
    }

</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div on:click={() => handleDirEntryClick(dirEntry)}>
    {dirEntry.name}
</div>
