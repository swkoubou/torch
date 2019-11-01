<template>
  <v-app>
    <v-content>
      <v-container>
        <nuxt/>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend({
        data() {
            return {}
        }
    })
</script>

<style lang="scss">
  html, body {
    overflow: hidden;
    width: 100vw;
    height: 100vh;
  }
</style>

