<script lang="ts">
  import { onMount } from 'svelte';
  import {ListTables} from '../../wailsjs/go/backend/App.js'

  let tables: string[] = []
  let prod: boolean = false

  function listTables(): void {
    ListTables().then(response => {
      tables = (response)
    })
  }

  $: filteredTables = tables.filter(table => table.endsWith(prod ? '-staging' : '-dev'))

  onMount(() => listTables())
</script>

<div class="grow bg-white">
  <div class="flex-initial">
    <div class="form-control w-64" data-theme="nord">
      <label class="label cursor-pointer">
        <span class="label-text text-gray-600">Dev / Prod</span>
        <input bind:checked={prod} type="checkbox" class="toggle" />
      </label>
    </div>
    <ul class="menu bg-base-200 rounded-box w-64" data-theme="nord">
      {#each filteredTables as table}
        <li><div class="menu-md">{table}</div></li>
        {:else}
        <li>Empty list</li>
      {/each}
    </ul>
  </div>
</div>
