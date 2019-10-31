<template>
  <div>
    <!-- Admin Plus Button -->
    <v-btn fixed bottom right fab large color="accent" v-show="!dialogShow" @click="dialogShow=true">
      <v-icon>mdi-plus</v-icon>
    </v-btn>

    <!-- Admin Dialog -->
    <v-dialog v-model="dialogShow" fullscreen>
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="dialogShow = false">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>新しいスポットの追加</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="dialogShow = false">Save</v-btn>
        </v-toolbar-items>
      </v-toolbar>


    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface AdminInterface {
        dialogShow: boolean,
    }

    export default Vue.extend({
        name: "admin",
        data(): AdminInterface {
            return {
                dialogShow: false
            }
        }
    });
</script>

<style scoped lang="scss">

</style>
