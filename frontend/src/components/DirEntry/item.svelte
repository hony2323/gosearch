<script lang="ts">
    import { FolderIcon, FileIcon } from "svelte-feather-icons";

    import type { types } from "../../../wailsjs/go/models";
    import { path, addEntryToPath } from "../store/store";

    export let dirEntry: types.DirEntry;

    function handleDirEntryClick(dirEntry: types.DirEntry) {
        if (dirEntry.isDir)
            path.update((value) => addEntryToPath(dirEntry.name, value));
    }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="item">
    {#if dirEntry.isDir}
        <FolderIcon />
    {:else}
        <FileIcon />
    {/if}
    <div class="itemText" on:dblclick={() => handleDirEntryClick(dirEntry)}>
        {dirEntry.name}
    </div>
</div>

<style lang="scss">
    .item {
        display: flex;
    }
    .itemText {
        margin: 5px;
        text-align: start;
    }
</style>
