<template>
  <div>
    <v-app-bar app color="primary" dark>
      <v-btn icon dark @click="close">
        <v-icon>mdi-close</v-icon>
      </v-btn>
      <v-toolbar-title>{{ pinName }}の情報</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <v-btn dark text @click="close">閉じる</v-btn>
      </v-toolbar-items>
    </v-app-bar>

    <v-container id="scroll-target" style="max-height: 90vh" class="overflow-y-auto">
      <v-row>
        <v-card>
          <v-img height="300" :src="pinImageUrl"></v-img>
        </v-card>

        <v-card flat>
          <v-flex class="d-flex mt-4 mb-2">
            <!-- 各種SNSシェア -->
            <v-btn fab class="mx-1" @click="shareTwitter">
              <v-icon color="#00acee">mdi-twitter</v-icon>
            </v-btn>
            <v-btn fab class="mx-1" @click="shareFaceBook">
              <v-icon color="#3B5998">mdi-facebook</v-icon>
            </v-btn>
            <v-btn fab class="mx-1" @click="shareLine">
              <v-icon color="#1dcd00">fab fa-line</v-icon>
            </v-btn>

            <!-- いいね -->
            <div class="ml-auto">
              <v-btn fab>
                <v-icon color="primary">mdi-heart</v-icon>
              </v-btn>
              <p class="caption text-center" style="color: #888888">{{ pinFavCount }}</p>
            </div>
          </v-flex>

          <div class="my-3">
            <pre class="text-wrap">{{ pinDetailText }}</pre>
          </div>
        </v-card>

        <v-card flat>
          <!-- ユーザー投稿欄 -->
          <div>

          </div>
        </v-card>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend({
        name: "_id",
        computed: {
            id() {
                return this.$route.params['id'];
            },
            pinFavCount(): number {
                return 1204;
            },
            pinName(): string {
                //TODO: ピンの名前を返す
                return 'test';
            },
            pinImageUrl(): string {
                //TODO: ピンの画像を返す
                return 'https://dummyimage.com/600x400/000/fff';
            },
            pinDetailText(): string {
                //TODO: ピンのテキストを返す
                return '　吾輩は猫である。名前はまだ無い。\n' +
                    '　どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙を吹く。どうも咽せぽくて実に弱った。これが人間の飲む煙草というものである事はようやくこの頃知った。';
            }
        },
        methods: {
            close(): void {
                this.$router.push('/');
            },
            shareTwitter(): void {
                // TODO: なんかいい感じのテキストを....
                let text = 'sample'
                let shareURL = `https://twitter.com/intent/tweet?url=torch.swkoubou.com&hashtags=torch&text=${text}`;
                open(shareURL, "_blank");
            },
            shareFaceBook(): void {
                let shareURL = `https://www.facebook.com/sharer/sharer.php?u=torch.swkoubou.com`;
                open(shareURL, "_blank");
            },
            shareLine(): void {
                // TODO: なんかいい感じのテキストを....
                let text = 'Torchはじめました\nよろしく\nにゃーん\n'
                let message = encodeURIComponent(text + 'torch.swkoubou.com');
                let shareURL = `http://line.me/R/msg/text/?${message}`;
                open(shareURL, "_blank");
            }
        }
    });
</script>

<style scoped lang="scss">

</style>
