<script lang="ts">
  import Router from "svelte-spa-router";
  import TailwindCss from './TailwindCSS.svelte'
  import Tables from "./routes/Tables.svelte"; 
  import Lambdas from "./routes/Lambdas.svelte";
</script>

<TailwindCss/>

<main>
  <div class="flex flex-col h-screen">
    <div role="tablist" class="tabs tabs-lifted mb-1" data-theme="nord">
      <div role="tab" class="tab tab-active">Tables</div>
      <div role="tab" class="tab">Lambdas</div>
      <div role="tab" class="tab">Settings</div>
    </div>

    <Router
      routes={{
          '/': Tables,
          '/lambdas': Lambdas,
      }}
    />
  </div>
</main>

