<script lang="ts">
    import Folder from "./components/folder.svelte";
    import { ListFolder } from "./../wailsjs/go/main/App";
    import { path } from "./store";
    let result = [];
    
    // List folder executes any time path is changed.
    $: ListFolder($path).then((res) => {result = res;})
</script>

<Folder itemList={result} />
