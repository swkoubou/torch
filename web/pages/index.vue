<template>
  <div>
    <v-btn fixed top left disabled outlined rounded large color="success" width="10rem">{{areaName}}</v-btn>

    <div class="map-parent" ref="map-parent" :style="mapParentStyle">
      <img src="/map.png" class="map" ref="map" :style="mapStyle" alt="map" @load="imageLoaded">
    </div>

    <!-- いいねボタン -->
    <v-btn fixed bottom right fab large outlined color="primary" v-if="!isAdmin">
      <v-icon color="primary">mdi-heart</v-icon>
    </v-btn>

    <pin-detail></pin-detail>

    <admin v-if="isAdmin"></admin>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';
    import admin from '../components/admin.vue';
    import pinDetail from '../components/pinDetail.vue';

    interface indexData {
        isAdmin: boolean
        pins: Array<any>
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
        scaleFlag: boolean
        touches: number
    }

    export default Vue.extend({
        name: 'index',
        components: {admin, pinDetail},
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
                scaleFlag: false,
                touches: 0,
            };
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
            loadPins() {
                this.pins = [
                    {
                        left: 0,
                        right: 0
                    }
                ]
            },
        },
        computed: {
            areaName(): string {
                return 'さばんなちほー';
            }
        }
    })
</script>

<style lang="scss" scoped>
  .map-parent {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 2;
  }

  .theme--light.v-btn.v-btn--disabled {
    color: #d32f2f !important;
  }
</style>
