<script>
  // Only import what you need!
  import { AirplayIcon } from "svelte-feather-icons";
  import { path, removeEntryFromPath } from "./../store/store";
  import Search from "./Search.svelte";
  const handleBackClick = () => {
    path.update((value) => removeEntryFromPath(value));
  };
</script>

<div>
  <div id="navrow1" class="navrow">
    <div on:dblclick={() => handleBackClick()} class="navitem">
      <AirplayIcon size="24" on:click={() => handleBackClick()} />
    </div>
    <div class="navitem">
      <h2>{$path}</h2>
    </div>
  </div>
  <div class="navrow">
    <div class="navitem">
      <Search />
    </div>
  </div>
</div>

<style lang="scss">

  #navrow1 {
    justify-content: space-between;
    margin: 0 50px;
  }
  .navrow {
    display: flex;
    // justify-content: space-evenly;
  }

  .navitem {
    display: flex;
    align-items: center;
  }
</style>
