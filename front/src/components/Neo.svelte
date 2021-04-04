<script lang="ts">
  import type { NeoResponse } from "../types/neo.type";

  const getNeoResults = async (): Promise<NeoResponse> => {
    const response = await fetch("http://localhost:8080/api/neo");
    const result = await response.json();
    if (response.ok) {
      return result;
    } else {
      throw new Error(result);
    }
  };

  $: neoPromise = getNeoResults();
</script>

<div class="neo-results" id="neo">
  {#await neoPromise}
    <span class="neo-results__await">...</span>
  {:then value}
    <span>Total objects: {value?.totalObjects}</span>
    <span>Date: {value?.date}</span>
    {#each value?.nearEarthObjects as obj}
      <div class="neo-results__object">
        <p>Name: {obj?.name}</p>
        <p>Estimated diameters: {obj?.estimatedDiameterInMeters}</p>
        <p>Relative velocity per second: {obj?.relativeVelocityKMPerSecond}</p>
      </div>
    {/each}
  {:catch error}
    <p class="neo-results__error">Something went wrong: {error.message}</p>
  {/await}
</div>

<style>
  .neo-results {
    display: flex;
    flex-direction: column;
    border: 1px solid #c0c0c0;
    border-radius: 5px;
    padding: 50px;
    max-width: 500px;
    width: 100%;
    overflow: scroll;
    margin: 10px;
  }

  span,
  .neo-results__error {
    font-size: 15px;
    margin-bottom: 20px;
  }

  .neo-results__object {
    font-size: 15px;
    line-height: 0;
    border: 1px solid #c0c0c0;
    border-radius: 5px;
    display: flex;
    flex-direction: column;
    padding: 12px;
    margin-top: 10px;
  }
</style>
