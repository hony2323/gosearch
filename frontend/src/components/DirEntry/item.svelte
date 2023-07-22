<script lang="ts">
    import type { main } from "../../../wailsjs/go/models";
    import { path, addEntryToPath } from "../store/store";


    export let dirEntry: main.DirEntryJson;

    function handleDirEntryClick(dirEntry: main.DirEntryJson) {
        if (dirEntry.isDir)
            path.update(value => addEntryToPath(dirEntry.name, value))
    }

</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div on:dblclick={() => handleDirEntryClick(dirEntry)}>
    {dirEntry.name}
</div>
