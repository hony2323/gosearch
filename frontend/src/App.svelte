<script lang="ts">
    import Folder from "./components/folder.svelte";
    import { ListFolder } from "./../wailsjs/go/main/App";
    import { path } from "./store";
    let result = [];
    
// todo: find out how to do the same thing but now it the 
// html, because ypu see line 11 in the page 
</script>

{ListFolder($path).then((res) => {
    debugger;
    result = res;
})}
<Folder itemList={result} />
