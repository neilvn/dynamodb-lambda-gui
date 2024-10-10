<script lang="ts">
  import { onMount } from 'svelte';
  import {ListTables, ScanTable} from '../../wailsjs/go/backend/App.js'

  let data = []
  let tables: string[] = []
  let headers: string[] = []
  let prod: boolean = false

  function listTables(): void {
    ListTables().then(response => {
      tables = (response)
    })
  }

  function scanTable(name: string): void {
    ScanTable(name).then(response => {
      const sortedData = response.map(item => alphabetize(item));
      data = sortedData;
    })
  }

  function getHeaders(num: number) {
    const setHeaders = new Set();
    const rows = data.slice(0, num);
    
    for (let row of rows) {
      const keys = Object.keys(row);
      for (const key of keys) {
        setHeaders.add(key);
      }
    }
    headers = Array.from(headers).sort().reverse();
  }

  function alphabetize(input: object) {
    return Object.keys(input).sort().reduce(function (result, key) {
        result[key] = input[key];
        return result;
    }, {});
  }

  $: filteredTables = tables.filter(table => table.endsWith(prod ? '-staging' : '-dev'))

  onMount(() => listTables())
</script>

<div class="grow bg-white">
  <div class="flex flex-row">
    <div>
      <div class="form-control w-64" data-theme="nord">
        <label class="label cursor-pointer">
          <span class="label-text text-gray-600">Dev / Prod</span>
          <input bind:checked={prod} type="checkbox" class="toggle" />
        </label>
      </div>
      <ul class="menu bg-base-200 rounded-box w-64" data-theme="nord">
        {#each filteredTables as table}
          <li><button class="menu-md" on:click={() => scanTable(table)}>{table}</button></li>
          {:else}
          <li>Empty list</li>
        {/each}
      </ul>
    </div>

    <div class="grow">
      {#if data?.length}

      	<div class="overflow-x-auto">
          <table class="table-sm">
            <!-- head -->
            <thead>
              <tr class="h-5">
                {#each headers as header}
                  <th>{header}</th>
                {/each}
              </tr>
            </thead>
            <tbody>
              {#each data as row}
                <tr >
                  {#each Object.keys(row) as key}
                    <td class="h-5 min-w-[100px] max-w-[300px] overflow-hidden truncate whitespace-nowrap">{row[key]}</td>
                  {/each}
                </tr>
              {/each}
            </tbody>
          </table>
        </div>

      {:else}

      	<h2>{"No data found"}</h2>

      {/if}
    </div>

  </div>
</div>
