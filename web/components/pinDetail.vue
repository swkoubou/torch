<template>
  <div>
    <v-dialog fullscreen v-model="dialog">
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="dialogClose">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>{{ spotName }}の情報</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="dialogClose">閉じる</v-btn>
        </v-toolbar-items>
      </v-toolbar>

      <v-img height="300" :src="spotImageUrl"></v-img>

      <v-flex class="d-flex px-2">
        <!-- 各種SNSシェア -->
        <v-btn fab class="mx-1">
          <v-icon color="#00acee">mdi-twitter</v-icon>
        </v-btn>
        <v-btn fab class="mx-1">
          <v-icon color="#3B5998">mdi-facebook</v-icon>
        </v-btn>
        <v-btn fab class="mx-1">
          <v-icon color="#262626">mdi-instagram</v-icon>
        </v-btn>
        <v-btn fab class="mx-1">
          <v-icon color="#1dcd00">fab fa-line</v-icon>
        </v-btn>

        <!-- いいね -->
        <v-btn fab class="ml-auto">
          <v-icon color="primary">mdi-heart</v-icon>
        </v-btn>
      </v-flex>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface pinDetailInterface {
        dialog: boolean
    }

    export default Vue.extend({
        name: 'pinDetail',
        data(): pinDetailInterface {
            return {
                // TODO: DEBUG用
                dialog: true
            };
        },
        computed: {
            spotName(): string {
                //TODO: スポットの名前を返す
                return 'test';
            },
            spotImageUrl(): string {
                //TODO: スポットの画像を返す
                return 'https://dummyimage.com/600x400/000/fff';
            }
        },
        methods: {
            dialogClose(): void {
                this.dialog = false;
            }
        }
    });
</script>

<style lang="scss" scoped>

</style>
