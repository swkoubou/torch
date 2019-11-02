<template>
  <div >
    <v-btn fixed top left outlined rounded disabled color="accent" width="10rem" style="z-index: 3">{{areaName}}</v-btn>

    <v-menu offset-y fixed top right v-model="menuValue">
      <template v-slot:activator="{ on }">
        <v-btn fixed top right fab outlined large color="primary" v-on="on">
          <v-icon color="primary">mdi-menu</v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item>
          <v-list-item-title>ヘルプ</v-list-item-title>
        </v-list-item>
        <v-list-item @click="changeShare">
          <v-list-item-title>シェア</v-list-item-title>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item>
          <v-list-item-title>問い合わせ</v-list-item-title>
        </v-list-item>
        <v-list-item @click="$router.push('term')">
          <v-list-item-title>利用規約</v-list-item-title>
        </v-list-item>
        <v-list-item @click="$router.push('privacy')">
          <v-list-item-title>プライバシーポリシー</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <div class="map-parent" ref="map-parent" :style="mapParentStyle">
      <img src="/map.png" class="map" ref="map" :style="mapStyle" alt="map" @load="imageLoaded">
      <div class="pin-parent" :style="mapStyle">
        <div v-for="p in pins" class="pin" :style="{ 'top': p.y + 'px', 'left': p.x + 'px' }">
          <v-icon :class="p.class">mdi-map-marker</v-icon>
        </div>
      </div>
    </div>

    <!-- いいねボタン -->
    <v-btn fixed bottom right fab large outlined color="primary" v-if="!isAdmin">
      <v-icon color="primary">mdi-heart</v-icon>
    </v-btn>

    <pin-detail></pin-detail>
    <share-dialog v-if="shareFlag" @change="changeShare"></share-dialog>

    <admin v-if="isAdmin"></admin>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';
    import admin from '../components/admin.vue';
    import pinDetail from '../components/pinDetail.vue';
    import shareDialog from "../components/shareDialog.vue";

    interface pinInfo {
        x: number
        y: number
        class: string
    }

    interface indexData {
        isAdmin: boolean
        pins: Array<pinInfo>
        touchStartPos: {
            x: number
            y: number
            backPosX: number
            backPosY: number
        }
        scaleMeta: {
            timeoutId: any
            beseDistance: number
            baseScale: number
            baseImageWidth: number
            baseImageHeight: number
            defaultSize: {
                width: number
                height: number
            }
        }
        mapParentStyle: object
        mapStyle: object
        realScale: number
        scaleFlag: boolean
        touches: number
        testPins: Array<any>
        menuValue: boolean
        shareFlag: boolean
    }

    export default Vue.extend({
        name: 'index',
        components: {admin, pinDetail, shareDialog},
        data(): indexData {
            return {
                isAdmin: false,
                pins: [],
                touchStartPos: {
                    x: 0,
                    y: 0,
                    backPosX: 0,
                    backPosY: 0,
                },
                scaleMeta: {
                    timeoutId: 0,
                    beseDistance: 0,
                    baseScale: 0,
                    baseImageWidth: 0,
                    baseImageHeight: 0,
                    defaultSize: {
                        width: 0,
                        height: 0,
                    }
                },
                mapParentStyle: {
                    transform: '',
                    width: '',
                    height: ''
                },
                mapStyle: {
                    width: '',
                    height: ''
                },
                realScale: 0,
                scaleFlag: false,
                touches: 0,
                testPins: [],
                menuValue: false,
                shareFlag: false,
            };
        },
        computed: {
            areaName(): string {
                return 'さばんなちほー';
            }
        },
        created(): void {
            const query = this.$route.query;
            if ('admin' in query) {
                this.isAdmin = query['admin'] === 'hoo0Jaek8jooTeeti0eiciedeithougee4aexooGhaiNgieDa9gio6jaipeevach';
            }
        },
        mounted() {
            const mapParent: any = this.$refs['map-parent'];

            mapParent.addEventListener("touchstart", (e: any) => {
                if (e.changedTouches.length == 1) {
                    this.menuValue = false;
                    const touch = e.changedTouches[0];
                    this.touchStartPos = {
                        x: touch.pageX,
                        y: touch.pageY,
                        backPosX: this.touchStartPos.backPosX,
                        backPosY: this.touchStartPos.backPosY,
                    };
                }
                this.touches += e.changedTouches.length;

                this.scaleFlag = this.touches > 1;

                e.preventDefault();
            });

            mapParent.addEventListener("touchmove", (e: any) => {
                const touches = e.changedTouches;
                if (touches.length == 1 && !this.scaleFlag) {
                    this.mapMoveEventHandler(touches);
                } else if (touches.length > 1) {
                    this.mapScaleEventHandler(touches);
                }
                e.preventDefault();
            });

            mapParent.addEventListener('touchend', (e: any) => {
                this.touches -= e.changedTouches.length;
                if (this.touches == 0) {
                    this.scaleMeta.beseDistance = 0;
                    this.scaleMeta.baseImageWidth = 0;
                    this.scaleMeta.baseImageWidth = 0;
                }
            })
        },
        methods: {
            makeTransformStr(x: number, y: number): string {
                return 'translate(' + x + 'px, ' + y + 'px)';
            },
            imageLoaded(): void {
                const map: any = this.$refs.map;
                this.scaleMeta.defaultSize.width = map.offsetWidth;
                this.scaleMeta.defaultSize.height = map.offsetHeight;

                this.scaleMeta.baseImageWidth = map.offsetWidth;
                this.scaleMeta.baseImageHeight = map.offsetHeight;

                this.mapScale(0.4, true);
                setTimeout(() => {
                    this.loadPins();
                }, 250);
            },
            mapScale(scale: number, force?: boolean): void {
                const truthScale = (this.scaleMeta.baseImageWidth * scale) / this.scaleMeta.defaultSize.width;

                if ((truthScale > 2 || truthScale < 0.3) && (force === undefined || !force)) {
                    return;
                }

                if (scale && scale != Infinity) {
                    this.$set(this.mapStyle, 'width', (this.scaleMeta.baseImageWidth * scale) + 'px');
                    this.$set(this.mapStyle, 'height', (this.scaleMeta.baseImageHeight * scale) + 'px');
                    this.$set(this.mapParentStyle, 'width', (this.scaleMeta.baseImageWidth * scale) + 'px');
                    this.$set(this.mapParentStyle, 'height', (this.scaleMeta.baseImageHeight * scale) + 'px');

                    this.realScale = truthScale;
                    this.updatePins();
                }
            },
            mapMoveEventHandler(touches: Touch[]): void {
                const touch = touches[0];

                const moveX = this.touchStartPos.x - touch.pageX;
                const moveY = this.touchStartPos.y - touch.pageY;

                this.touchStartPos.backPosX -= moveX;
                this.touchStartPos.backPosY -= moveY;

                const backgroundPos = this.makeTransformStr(this.touchStartPos.backPosX, this.touchStartPos.backPosY);
                this.$set(this.mapParentStyle, 'transform', backgroundPos);

                this.touchStartPos = {
                    x: touch.pageX,
                    y: touch.pageY,
                    backPosX: this.touchStartPos.backPosX,
                    backPosY: this.touchStartPos.backPosY,
                };
            },
            mapScaleEventHandler(touches: Touch[]): void {
                const map: any = this.$refs.map;

                this.scaleFlag = true;
                // 座標1 (1本目の指)
                var x1 = touches[0].pageX;
                var y1 = touches[0].pageY;

                // 座標2 (2本目の指)
                var x2 = touches[1].pageX;
                var y2 = touches[1].pageY;

                // 2点間の距離
                const distance = Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2));
                clearTimeout(this.scaleMeta.timeoutId);

                if (this.scaleMeta.beseDistance && this.scaleMeta.baseImageWidth && this.scaleMeta.baseImageHeight) {
                    const scale = distance / this.scaleMeta.beseDistance;
                    this.mapScale(scale);
                } else {
                    this.scaleMeta.beseDistance = distance;
                    this.scaleMeta.baseImageWidth = map.offsetWidth;
                    this.scaleMeta.baseImageHeight = map.offsetHeight;
                }

            },
            updatePins() {
                const map: any = this.$refs.map;

                const start = {
                    lat: 35.48832,
                    lon: 139.34024,
                };
                const end = {
                    lat: 35.48491,
                    lon: 139.34596,
                };

                const iw = map.offsetWidth;
                const ih = map.offsetHeight;

                this.pins = [];
                this.testPins.forEach((testPin) => {
                    const pxX = iw - (testPin.lat - start.lat) * iw / (end.lat - start.lat);
                    const pxY = ih - (testPin.lon - start.lon) * ih / (end.lon - start.lon);
                    this.pins.push({
                        x: pxX,
                        y: pxY,
                        class: testPin.class,
                    });
                })
            },
            loadPins() {
                //TODO: あとで消す
                this.testPins.push({
                    lat: 35.48560,
                    lon: 139.34135,
                    class: 'active',
                });
                this.testPins.push({
                    lat: 35.48655,
                    lon: 139.34287,
                    class: 'disabled',
                });
                this.testPins.push({
                    lat: 35.48763,
                    lon: 139.34382,
                    class: 'hot1',
                });
                this.testPins.push({
                    lat: 35.48559,
                    lon: 139.34436,
                    class: 'hot2',
                });
                this.testPins.push({
                    lat: 35.48625,
                    lon: 139.34375,
                    class: 'hot3',
                });

                // TODO: あとでAPIに変える
                navigator.geolocation.getCurrentPosition((position) => {
                    this.testPins.push({
                        lat: position.coords.latitude,
                        lon: position.coords.longitude
                    });
                    this.updatePins();
                }, () => {
                    this.updatePins();
                }, {
                    enableHighAccuracy: true,
                    maximumAge: 5,
                });
            },
            changeShare() {
                this.shareFlag = !this.shareFlag;
            },
        },
    })
</script>

<style lang="scss" scoped>
  .map-parent {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 2;

    .map {
      margin: 0;
      display: block;
    }

    .pin-parent {
      position: relative;
      top: -100%;

      .pin {
        $size: 30px;
        position: absolute;
        width: $size;
        height: $size;

        i {
          display: flex;
          width: $size;
          height: $size;
          font-size: $size;
          justify-items: center;
          align-items: center;
          color: rgb(0, 121, 107);
          border-radius: 50%;

          &.active {
            color: rgb(25, 117, 210);
          }

          &.disabled {
            color: rgb(120, 144, 156);
          }

          /** 一番低いホット **/
          &.hot1 {
            color: rgb(229, 57, 53);
          }

          /** 真ん中ホット **/
          &.hot2 {
            color: rgb(230, 74, 25);
          }

          /** 一番高いホット **/
          &.hot3 {
            color: rgb(198, 40, 40);
            border: solid 1px rgb(198, 40, 40);
            animation: 1s linear blink-animate infinite;
          }
        }

      }
    }
  }

  .theme--light.v-btn.v-btn--disabled {
    color: #d32f2f !important;
    z-index: 3;
  }

  @keyframes blink-animate {
    0% {
      border-color: rgba(198, 40, 40, 1);
    }
    50% {
      border-color: rgba(198, 40, 40, 0);
    }
    100% {
      border-color: rgba(198, 40, 40, 1);
    }
  }

  @-webkit-keyframes blink-animate {
    0% {
      border-color: rgba(198, 40, 40, 1);
    }
    50% {
      border-color: rgba(198, 40, 40, 0);
    }
    100% {
      border-color: rgba(198, 40, 40, 1);
    }
  }
</style>
