<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {FetchURL, GetValue, ListTables} from '../wailsjs/go/backend/App.js'

  let resultText: string = "Please enter your name below 👇"
  let tables: string[] = []
  let accessKey: string = ""
  let url: string

  function fetchUrl(): void {
    FetchURL(url).then(response => accessKey = response)    
  }

  function listTables(): void {
    ListTables().then(response => tables = (response))
  }

  function getValue(): void {
    GetValue(1, "credentials").then(response => accessKey = response)
  }
  
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{tables}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={url} class="input" id="url" type="text"/>
    <button class="btn" on:click={listTables}>List Tables</button>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 80px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
