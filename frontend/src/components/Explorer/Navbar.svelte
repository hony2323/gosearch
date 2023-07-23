<script>
  // Only import what you need!
  import { AirplayIcon } from "svelte-feather-icons";
  import { path, removeEntryFromPath } from "./../store/store";
  const handleBackClick = () => {
    path.update((value) => removeEntryFromPath(value));
  };
</script>

<div class="navbar">
  <div on:dblclick={() => handleBackClick()}>
    <AirplayIcon size="24" on:click={() => handleBackClick()} />
  </div>
  <h2>{$path}</h2>
</div>

<style lang="scss">
  .navbar {
    display: flex;
    justify-content: start;
  }
</style>
