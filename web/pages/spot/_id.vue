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

        <v-card width="100%" flat>
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
            <div class="ml-auto mr-1">
              <v-btn fab @click="sendLike">
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

    <!-- エラー -->
    <v-dialog v-model="errorDialog">
      <v-card>
        <v-card-title>エラー</v-card-title>
        <v-card-text>
          <p>エラーが発生しました。</p>
          <p>{{ errorMessage }}</p>
        </v-card-text>

        <v-card-actions class="d-flex justify-end">
          <v-btn class="mr-3 px-3" left @click="reload">再読込</v-btn>
          <v-btn class="ml-3 px-3" right @click="errorDialog = false">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';
    import Api from "~/module/api";
    import {structs} from "~/proto/web";

    interface _idInterface {
        info: structs.ISpotInfo | undefined
        errorMessage: string
        errorDialog: boolean
        poolingTimer: any
    }

    export default Vue.extend({
        name: "index",
        data(): _idInterface {
            return {
                info: undefined,
                errorMessage: '',
                errorDialog: false,
                poolingTimer: 0,
            }
        },
        computed: {
            id(): number {
                return Number(this.$route.params['id']);
            },
            pinFavCount(): number {
                if (this.info === undefined) {
                    return 0;
                }
                const count = this.info.sumLikeCount;
                if (typeof count === "number") {
                    return count
                }
                return 0;
            },
            pinName(): string {
                if (this.info === undefined) {
                    return 'Unknown';
                }
                const name = this.info.name;
                if (typeof name === "string") {
                    return name
                }
                return 'Unknown';
            },
            pinImageUrl(): string {
                if (this.info === undefined) {
                    return '';
                }
                const image = this.info.photoFileName;
                if (typeof image === "string") {
                    return '/static-api/images/spots/' + image
                }
                return '';
            },
            pinDetailText(): string {
                if (this.info === undefined) {
                    return 'Unknown';
                }
                const detail = this.info.description;
                if (typeof detail === "string") {
                    return detail
                }
                return 'Unknown';
            }
        },
        watch: {
            errorDialog(val) {
                if (!val && this.errorMessage !== "") {
                    this.errorMessage = "";
                }
            },
            errorMessage(val) {
                this.errorDialog = val !== "";
            }
        },
        methods: {
            close(): void {
                this.$router.push('/');
            },
            loadParams() {
                Api.getPinInfo(this.id).then(res => {
                    const info = res.spotInfo;
                    if (info === undefined || info == null) {
                        this.errorMessage = '正常にデータを転送できませんでした';
                        return;
                    }
                    this.info = info;
                }).catch(() => {
                    this.errorMessage = 'ネットワークエラー';
                });
            },
            shareTwitter(): void {
                let text = this.pinName + '- TORCH 神奈川工科大学学園祭・幾徳祭';
                text = encodeURI(text);
                const url = encodeURI(location.href);
                let shareURL = `https://twitter.com/intent/tweet?url=${url}&hashtags=torch&text=${text}`;
                open(shareURL, "_blank");
            },
            shareFaceBook(): void {
                let shareURL = `https://www.facebook.com/sharer/sharer.php?u=torch.swkoubou.com`;
                open(shareURL, "_blank");
            },
            shareLine(): void {
                let text = this.pinName + '- TORCH 神奈川工科大学学園祭・幾徳祭';
                let message = encodeURIComponent(text + '\r\n' + location.href);
                let shareURL = `http://line.me/R/msg/text/?${message}`;
                open(shareURL, "_blank");
            },
            sendLike() {
                Api.pinLike(this.id).then(res => {
                    if (res.message !== 'success') {
                        this.errorMessage = '不明なエラーです'
                    } else {
                        this.loadParams();
                    }
                }).catch(() => {
                    this.errorMessage = 'ネットワークエラー'
                })
            },
            reload(){
                location.reload();
            }
        },
        beforeDestroy(): void {
            clearInterval(this.poolingTimer);
        },
        created(): void {
            this.loadParams();

            this.poolingTimer = setInterval(() => {
                this.loadParams();
            }, 3000);
        }
    });
</script>

<style scoped lang="scss">

</style>
