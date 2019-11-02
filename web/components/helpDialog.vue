<template>
  <div>
    <v-dialog v-model="dialog">
      <v-card>
        <v-img src="/logo.png"></v-img>

        <v-card-text>
          <p class="text-center">TORCHへようこそ。</p>
          <p>TORCHはいいねをして幾徳祭を盛り上げるサービスです。</p>
        </v-card-text>

        <v-card-actions>
          <v-btn block large color="accent">次へ</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface helpDialogInterface {
        dialog: boolean
    }

    export default Vue.extend({
        name: "helpDialog",
        data(): helpDialogInterface {
            return {
                dialog: true,
            }
        }
    });
</script>

<style scoped lang="scss">

</style>
