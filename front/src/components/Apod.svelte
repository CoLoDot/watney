<script lang="ts">
  import type { ApodResponse } from "../types/apod.type";

  const getApodResults = async (): Promise<ApodResponse> => {
    const response = await fetch("http://localhost:8080/api/apod");
    const result = await response.json();
    if (response.ok) {
      return result;
    } else {
      throw new Error(result);
    }
  };

  $: apodPromise = getApodResults();
</script>

<div class="apod-results" id="apod">
  {#await apodPromise}
    <span class="apod-results__await">...</span>
  {:then value}
    <span class="apod-results__title">{value?.title}</span>
    <span class="apod-results__description">{value?.description}</span>
    <img
      class="apod-results__image"
      src={value?.url}
      alt={`Image of ${value?.title}`}
    />
  {:catch error}
    <p class="apod-results__error">Something went wrong: {error.message}</p>
  {/await}
</div>

<style>
  .apod-results {
    display: flex;
    flex-direction: column;
    border: 1px solid #c0c0c0;
    border-radius: 5px;
    padding: 50px;
    text-align: center;
    align-items: center;
    max-width: 500px;
    width: 100%;
    overflow: scroll;
    margin: 10px;
  }

  .apod-results__error,
  .apod-results__title {
    font-size: 15px;
    margin-bottom: 20px;
  }

  .apod-results__description {
    font-size: 12px;
  }

  .apod-results__image {
    width: 100%;
    margin-top: 20px;
    max-width: 500px;
  }
</style>
