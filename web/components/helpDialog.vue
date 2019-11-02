<template>
  <div>
    <v-dialog v-model="dialog" persistent>
      <v-carousel light hide-delimiters hide-delimiter-background :cycle="false" v-model="circleStatus">
        <v-carousel-item>
          <v-card height="100%">
            <v-img src="/logo.png"></v-img>

            <v-card-text>
              <p class="text-center">TORCHへようこそ。</p>
              <p>TORCHはいいねをして幾徳祭を盛り上げるサービスです。</p>
            </v-card-text>
          </v-card>
        </v-carousel-item>

        <v-carousel-item>
          <v-card height="100%">
            <div class="max-large-icon">
              <v-icon color="primary">mdi-heart</v-icon>
            </div>

            <v-card-text>
              <p>「いいね」を押すことでエリアを盛り上げることができます。</p>
              <p>「いいね」は何度でも押すことができます。感動したり楽しかったら何度でも「いいね」を押して、エリアを盛り上げましょう！</p>
            </v-card-text>
          </v-card>
        </v-carousel-item>

        <v-carousel-item>
          <v-card height="100%">
            <div class="max-large-icon">
              <v-icon color="primary">mdi-map-marker</v-icon>
            </div>

            <v-card-text>
              <p>マーカーをタップすると、スポットの詳細を見ることができます。</p>
              <p>マーカーでも「いいね」することができます。このスポットを盛り上げましょう！</p>
              <p>マーカーをタップして「いいね」をするとコメントや写真をアップロードできます。</p>
            </v-card-text>
          </v-card>
        </v-carousel-item>

        <v-carousel-item>
          <v-card height="100%">
            <div class="max-large-icon">
              <v-icon color="primary">mdi-map-check</v-icon>
            </div>

            <v-card-text>
              <p>赤くなっているエリアは、人気のスポットです。</p>
              <p>ぜひホットなスポットへ行ってみてください。ホットなスポットへ行った際には、マーカーをタップしてその場の写真やコメントでシェアしよう！</p>
            </v-card-text>
          </v-card>
        </v-carousel-item>

        <v-carousel-item>
          <v-card height="100%">
            <div class="max-large-icon">
              <v-icon color="primary">mdi-account-question</v-icon>
            </div>

            <v-card-text>
              <p>「スポットを追加したい！」と思った方はK1号館106号室のソフトウェア工房までおこしください。</p>
              <p>また、バグや質問などがあった場合も、K1号館106号室のソフトウェア工房までおこしください。</p>
            </v-card-text>
          </v-card>
        </v-carousel-item>

        <v-carousel-item>
          <v-card height="100%">
            <div class="max-large-icon">
              <v-icon color="info">mdi-check</v-icon>
            </div>

            <v-card-text class="py-5">
              <p class="text-center">さぁ、始めましょう！</p>
            </v-card-text>

            <v-card-actions>
              <v-btn x-large block color="accent" @click="closeModal">はじめる</v-btn>
            </v-card-actions>
          </v-card>
        </v-carousel-item>
      </v-carousel>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface helpDialogInterface {
        dialog: boolean
        circleStatus: number
    }

    export default Vue.extend({
        name: "help-dialog",
        props: ['value'],
        data(): helpDialogInterface {
            return {
                dialog: true,
                circleStatus: 0
            }
        },
        watch: {
            dialog(value: boolean) {
                this.$emit('input', value);
            },
            value(value: boolean) {
                this.dialog = value;
                this.circleStatus = 0;
            },
            circleStatus(value: number, old: number) {
                const diff = value - old;
                if (diff < 0 && value === 0 && old === 5) {
                    this.closeModal();
                }
            },
        },
        methods: {
            closeModal() {
                this.dialog = false;
            }
        }
    });
</script>

<style scoped lang="scss">
  .max-large-icon {
    display: flex;
    justify-content: center;
    align-items: center;

    i {
      font-size: 250px;
    }
  }
</style>
