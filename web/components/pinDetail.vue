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

      <div class="my-3">
        <pre class="text-wrap">{{ spotDetailText }}</pre>
      </div>

      <!-- ユーザー投稿欄 -->
      <div>

      </div>
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
            },
            spotDetailText(): string {
                //TODO: スポットのテキストを返す
                return '　吾輩は猫である。名前はまだ無い。\n' +
                    '　どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙を吹く。どうも咽せぽくて実に弱った。これが人間の飲む煙草というものである事はようやくこの頃知った。';
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
