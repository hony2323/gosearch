<script lang="ts">
    import Folder from "./components/folder.svelte";
    import { ListFolder } from "./../wailsjs/go/main/App";
    let result = [];
    ListFolder("Z:/code/Go").then((res) => {
        debugger
        result = res;
    });
</script>

<Folder itemList={result} />
