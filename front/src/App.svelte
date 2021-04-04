<script lang="ts">
  import Apod from "./components/Apod.svelte";
  import Neo from "./components/Neo.svelte";
</script>

<main>
  <div class="menu">
    <h1>watney</h1>
    <div class="links">
      <a href="#apod">apod</a> | <a href="#neo">neo</a>
    </div>
  </div>
  <Apod />
  <Neo />
</main>

<style>
  main {
    display: flex;
    justify-content: space-evenly;
    flex-direction: column;
    align-items: center;
    margin-left: 56px;
    margin-right: 56px;
    margin-top: 120px;
  }

  .menu {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-decoration: none;
    border-bottom: 1px solid #c0c0c0;
    padding-bottom: 20px;
    width: 100%;
    position: fixed;
    top: 0;
    background-color: white;
  }

  h1 {
    text-align: center;
  }
</style>
