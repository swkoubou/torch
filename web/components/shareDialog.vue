<template>
  <div>
    <v-dialog v-model="dialog" persistent>
      <v-card>
        <v-card-title>
          <span>シェア</span>
        </v-card-title>
        <v-card-text>
          <v-flex class="d-flex px-2 mt-4 mb-2">
            <!-- 各種SNSシェア -->
            <v-btn fab x-large class="mx-1 ml-auto" @click="shareTwitter">
              <v-icon color="#00acee">mdi-twitter</v-icon>
            </v-btn>
            <v-btn fab x-large class="mx-1 ml-auto" @click="shareFaceBook">
              <v-icon color="#3B5998">mdi-facebook</v-icon>
            </v-btn>
            <v-btn fab x-large class="mx-1 ml-auto" @click="shareLine">
              <v-icon color="#1dcd00">fab fa-line</v-icon>
            </v-btn>
          </v-flex>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="parentDialog">
            閉じる
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
  import Vue from 'vue';

  interface shareDialogInterface {
      dialog: boolean
  }

  export default Vue.extend({
      name: 'shareDialog',
      props: {
          dialog: {
              type: Boolean,
              default: true
          }
      },
      methods: {
          shareTwitter(): void {
            let shareURL = 'https://twitter.com/intent/tweet?text=torch.swkoubou.com&hashtags=torch';
            open(shareURL, "_blank");
          },
          shareFaceBook(): void {
              let shareURL = 'https://www.facebook.com/sharer/sharer.php?u=torch.swkoubou.com/'
              open(shareURL, "_blank");
          },
          shareLine(): void {
              let baseUrl = encodeURIComponent('torch.swkoubou.com');
              let shareURL = `https://social-plugins.line.me/lineit/share?url=${baseUrl}`;
              open(shareURL, "_blank");
          },
          parentDialog(): void {
              this.$emit('change');
          }
      }
  })
</script>

<style lang="scss" scoped>

</style>

<!--TODO: シェアリングはtorch.swkoubou.com-->
