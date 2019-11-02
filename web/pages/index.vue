<template>
  <div>
    <v-btn fixed top left disabled outlined rounded large color="success" width="10rem">{{areaName}}</v-btn>

    <div class="map-parent" ref="map-parent" :style="mapParentStyle">
      <img src="/map.png" class="map" ref="map" :style="mapStyle" alt="map" @load="imageLoaded">
      <div class="pin-parent" :style="mapStyle">
        <div v-for="p in pins" class="pin" :style="{ 'top': p.y + 'px', 'left': p.x + 'px' }"></div>
      </div>
      <div class="user-location-parent" :style="mapStyle">
        <div class="user-location"
             :style="{ 'transform': 'translate('+userLocation.x + 'px, '+ userLocation.y + 'px)' }"></div>
      </div>
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

    interface pinInfo {
        x: number
        y: number
    }

    interface userLocation {
        x: number
        y: number
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
        userLocation: userLocation
        locationWatchId: any
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
                realScale: 0,
                scaleFlag: false,
                touches: 0,
                testPins: [],
                userLocation: {
                    x: 0,
                    y: 0,
                },
                locationWatchId: 0,
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
            });

            this.watchMyLocation();
        },
        beforeDestroy(): void {
            navigator.geolocation.clearWatch(this.locationWatchId);
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
                this.pins = [];
                this.testPins.forEach((testPin) => {
                    const xy = this.getGeo2Px(testPin);
                    const pxX = xy.x;
                    const pxY = xy.y;
                    this.pins.push({x: pxX, y: pxY});
                })
            },
            getGeo2Px(testPin: object): object {
                const map: any = this.$refs.map;

                const iw = map.offsetWidth;
                const ih = map.offsetHeight;

                const start = {
                    lat: 35.48832,
                    lon: 139.34024,
                };
                const end = {
                    lat: 35.48491,
                    lon: 139.34596,
                };

                const startPos = this.convertPos(start.lat, start.lon);
                const endPos = this.convertPos(end.lat, end.lon);
                const currentXY = this.convertPos(testPin.lat, testPin.lon);

                const bx = iw / (endPos.x - startPos.x);
                const by = ih / (endPos.y - startPos.y);

                currentXY.x -= startPos.x;
                currentXY.y -= startPos.y;

                const pxX = currentXY.x * bx;
                const pxY = currentXY.y * by;
                return {
                    x: pxX,
                    y: pxY,
                }
            },
            convertPos(lat: number, lon: number): any {
                const z = 40;
                const L = 85.05112878;
                const pointX = Math.pow(2, (z + 7)) * ((lon / 180) + 1);
                const pointY = Math.pow(2, (z + 7) / Math.PI) * (-1 * Math.atanh(Math.sin(lat * Math.PI / 180)) + Math.atanh(Math.sin(L * Math.PI / 180)));
                return {
                    x: pointX,
                    y: pointY,
                }
            },
            loadPins() {
                //TODO: あとで消す
                this.testPins.push({
                    lat: 35.48755,
                    lon: 139.34112,
                });
                this.testPins.push({
                    lat: 35.48655,
                    lon: 139.34442,
                });

                // TODO: あとでAPIに変える
                this.updatePins();
                navigator.geolocation.getCurrentPosition((position) => {
                    this.setUserLocation(position);
                }, () => {
                }, {
                    enableHighAccuracy: true,
                    maximumAge: 5,
                });
            },
            watchMyLocation() {
                this.locationWatchId = navigator.geolocation.watchPosition((position) => {
                    this.setUserLocation(position);
                });
            },
            setUserLocation(position: any) {
                const userLocation = {
                    lat: position.coords.latitude,
                    lon: position.coords.longitude
                };
                const xy = this.getGeo2Px(userLocation);

                this.userLocation = {
                    x: xy.x,
                    y: xy.y
                };
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

    .map {
      margin: 0;
      display: block;
    }

    .pin-parent {
      position: relative;
      top: -100%;

      .pin {
        position: absolute;
        width: 8px;
        height: 8px;
        top: 0;
        left: 0;
        border-radius: 50%;
        background-color: red;
      }
    }

    .user-location-parent {
      position: relative;
      top: -200%;

      .user-location {
        $size: 15px;
        position: absolute;
        top: 0;
        left: 0;
        width: $size;
        height: $size;
        background-color: rgb(40, 53, 147);
        border-radius: 50%;
        transition: ease .1s transform;

        &::before {
          display: block;
          content: '';
          width: $size * 4;
          height: $size* 4;
          border-radius: 50%;
          background-color: rgba(57, 73, 171, .3);
          margin: -($size*3/2);
          border: solid thin rgba(57, 73, 171, .6);
          border-left-color: transparent;
          border-right-color: transparent;
          animation: 1.8s linear opacity-blink-animate infinite;
        }
      }
    }
  }

  .theme--light.v-btn.v-btn--disabled {
    color: #d32f2f !important;
    z-index: 3;
  }

  @keyframes opacity-blink-animate {
    0% {
      opacity: 1;
      transform: rotate(0deg);
    }
    50% {
      opacity: 0.4;
    }
    100% {
      opacity: 1;
      transform: rotate(180deg);
    }
  }
</style>
