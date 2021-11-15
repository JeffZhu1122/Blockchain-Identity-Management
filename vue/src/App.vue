<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
<style>
    @media (min-width: 1400px){
      .container {
          max-width: 100% !important;
          margin-top: 50px;
      }
    }
</style>
