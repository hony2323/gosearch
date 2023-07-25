<script>
  // Only import what you need!
  import { AirplayIcon } from "svelte-feather-icons";
  import { path, removeEntryFromPath } from "./../store/store";
  import Search from "./Search.svelte";
  const handleBackClick = () => {
    path.update((value) => removeEntryFromPath(value));
  };
</script>

<div class="navbar">
  <div on:dblclick={() => handleBackClick()}
    class="navitem">
    <AirplayIcon size="24" on:click={() => handleBackClick()} />
  </div>
  <div class="navitem">
    <h2>{$path}</h2>
  </div>
  <div class="navitem">
    <Search />
  </div>
</div>

<style lang="scss">
  .navbar {
    display: flex;
    justify-content: space-evenly;
  }

  .navitem {
    display: flex;
    align-items: center;
  }
</style>
