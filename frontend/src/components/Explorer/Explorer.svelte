<script lang="ts">
    import Navbar from "./Navbar.svelte";
    import Folder from "../../components/DirEntry/folder.svelte";
    import { ListFolder } from "../../../wailsjs/go/main/App";
    import { path } from "../store/store";
    let result = [];

    // List folder executes any time path is changed.
    $: ListFolder($path).then((res) => {
        result = res;
    });
</script>

<Navbar />

<div class="ExplorerBody">
    <Folder itemList={result} />
    <div></div>
</div>

<style lang="scss">
    .ExplorerBody {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
    }
</style>