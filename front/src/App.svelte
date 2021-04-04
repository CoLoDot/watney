<script lang="ts">
  import Apod from "./components/Apod.svelte";
  import Neo from "./components/Neo.svelte";
</script>

<main>
  <Apod />
  <Neo />
</main>

<style>
  main {
    display: flex;
    justify-content: space-evenly;
  }

  @media (max-width: 680px) {
    main {
      flex-direction: column;
      align-items: center;
      margin-left: 56px;
      margin-right: 56px;
    }
  }
</style>
