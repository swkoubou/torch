<template>
  <div>
    <v-btn fixed top left outlined rounded disabled color="accent" width="10rem" style="z-index: 3">{{areaName}}</v-btn>

    <v-menu offset-y fixed top right v-model="menuValue">
      <template v-slot:activator="{ on }">
        <v-btn fixed top right fab outlined large color="primary" v-on="on">
          <v-icon color="primary">mdi-menu</v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item @click="helpFlag=true">
          <v-list-item-title>ヘルプ</v-list-item-title>
        </v-list-item>
        <v-list-item @click="changeShare">
          <v-list-item-title>シェア</v-list-item-title>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item @click="contactFlag = true">
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
      <div class="user-location-parent" :style="mapStyle">
        <div class="user-location"
             :style="{ 'transform': 'translate('+userLocation.x + 'px, '+ userLocation.y + 'px)' }"></div>
      </div>
      <div class="area-range" :style="mapStyle">
        <div class="area-rect" v-for="a in areas" :style="{
        transform: 'translate('+a.leftUp.x + 'px, '+a.leftUp.y + 'px)',
        width: a.width + 'px',
        height: a.height + 'px',
        opacity: a.opacity ,
        }"></div>

        <div class="like-effect" v-for="i in likeEffects" :key="i.createdTime"
             :style="{ 'top': i.y + 'px', 'left': i.x + 'px' }"
             :class="'like-effect'+i.effectType">
          <v-icon color="primary">mdi-heart</v-icon>
        </div>
      </div>

      <div class="pin-parent" ref="map-action" :style="mapStyle">
        <div v-for="p in pins" class="pin" :style="{ 'top': p.y + 'px', 'left': p.x + 'px' }"
             v-on:touchstart="showDetail(p)">
          <v-icon :class="p.class">mdi-map-marker</v-icon>
        </div>
        <div class="pin admin-pin" v-if="isAdmin"
             :style="{ 'top': adminLocation.y + 'px', 'left': adminLocation.x + 'px' }"></div>
      </div>
    </div>

    <!-- いいねボタン -->
    <v-btn fixed bottom right fab large outlined color="primary" v-if="!isAdmin" v-on:touchstart="sendLike">
      <v-icon color="primary">mdi-heart</v-icon>
    </v-btn>

    <share-dialog v-if="shareFlag" @change="changeShare"></share-dialog>
    <help-dialog v-model="helpFlag"></help-dialog>
    <contact-dialog v-model="contactFlag"></contact-dialog>

    <admin v-if="isAdmin" v-model="adminLatAndLon"></admin>

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
    import admin from '../components/admin.vue';
    import shareDialog from "../components/shareDialog.vue";
    import helpDialog from "../components/helpDialog.vue";
    import Api from "~/module/api";
    import contactDialog from "../components/contactDialog.vue";
    import GeoUtils from "~/utils/geoUtils";
    import {structs} from "~/proto/web";
    import PinUtils from "~/utils/pinUtils";

    interface pinInfo {
        id: any
        x: number
        y: number
        class: string
    }

    interface userLocation {
        x: number
        y: number
    }

    interface indexData {
        isAdmin: boolean
        pinInfoArr: Array<structs.ISpotInfo>
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
        mapParentStyle: any
        mapStyle: any
        realScale: number
        scaleFlag: boolean
        touches: number
        userLocation: userLocation
        locationWatchId: any
        userTruthLocation: any
        menuValue: boolean
        shareFlag: boolean
        helpFlag: boolean
        areas: Array<any>
        contactFlag: boolean
        adminLatAndLon: {
            lat: number
            lonL: number
            area: structs.AreaInfo | undefined
        },
        adminLocation: userLocation
        likeEffects: Array<any>
        errorMessage: string
        errorDialog: boolean
        poolingTimer: any
    }

    export default Vue.extend({
        name: 'index',
        components: {admin, shareDialog, helpDialog, contactDialog},
        data(): indexData {
            return {
                isAdmin: false,
                pinInfoArr: [],
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
                userLocation: {
                    x: 0,
                    y: 0,
                },
                userTruthLocation: {
                    lat: 0,
                    lon: 0,
                },
                locationWatchId: 0,
                menuValue: false,
                shareFlag: false,
                helpFlag: false,
                contactFlag: false,
                areas: [],
                adminLatAndLon: {
                    lat: 0,
                    lonL: 0,
                    area: undefined
                },
                adminLocation: {
                    x: 0,
                    y: 0,
                },
                likeEffects: [],
                errorMessage: '',
                errorDialog: false,
                poolingTimer: 0,
            };
        },
        computed: {
            areaName(): string {
                const defaultName = '未開の地';
                const userLat = this.userTruthLocation.lat;
                const userLon = this.userTruthLocation.lon;

                const info = GeoUtils.containArea(this.areas, userLat, userLon);
                if (typeof info === "undefined") {
                    return defaultName;
                } else {
                    return info.name;
                }
            },
            currentAreaId(): number | undefined {
                const userLat = this.userTruthLocation.lat;
                const userLon = this.userTruthLocation.lon;

                const info = GeoUtils.containArea(this.areas, userLat, userLon);
                if (typeof info === "undefined") {
                    return undefined;
                } else {
                    return info.areaId;
                }
            }
        },
        created(): void {
            const query = this.$route.query;
            if ('admin' in query) {
                this.isAdmin = query['admin'] === 'hoo0Jaek8jooTeeti0eiciedeithougee4aexooGhaiNgieDa9gio6jaipeevach';
            }

            this.helpFlag = localStorage.getItem('help-dialog') != 'true';
            this.loadAreas();

            this.poolingTimer = setInterval(() => {
                this.loadAreas();
                this.loadPins();
            }, 5000);
        },
        mounted() {
            const mapAction: any = this.$refs['map-action'];

            mapAction.addEventListener("touchstart", (e: any) => {
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
            }, false);

            mapAction.addEventListener("touchmove", (e: any) => {
                const touches = e.changedTouches;
                if (touches.length == 1 && !this.scaleFlag) {
                    this.mapMoveEventHandler(touches);
                } else if (touches.length > 1) {
                    this.mapScaleEventHandler(touches);
                }
                e.preventDefault();
            }, false);

            mapAction.addEventListener('touchend', (e: any) => {
                this.touches -= e.changedTouches.length;
                if (this.touches == 0) {
                    this.scaleMeta.beseDistance = 0;
                    this.scaleMeta.baseImageWidth = 0;
                    this.scaleMeta.baseImageWidth = 0;
                }
            }, false);

            if (this.isAdmin) {
                mapAction.addEventListener("touchstart", (e: any) => {
                    if (e.changedTouches.length == 1) {
                        const touch = e.changedTouches[0];
                        this.getMapPosHandler(touch.clientX, touch.clientY);
                    }
                });
                mapAction.addEventListener("mousedown", (e: any) => {
                    this.getMapPosHandler(e.clientX, e.clientY);
                });
            }

            this.watchMyLocation();
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
        beforeDestroy(): void {
            navigator.geolocation.clearWatch(this.locationWatchId);
            clearInterval(this.poolingTimer);
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
                    this.updateUserLocation(this.userTruthLocation);
                    this.updateAreas();
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
                this.pinInfoArr.forEach((pin: structs.ISpotInfo) => {
                    const domOps = PinUtils.convertUiPin(pin);

                    const xy = this.getGeo2Px(domOps['geoPos']);
                    const pxX = xy.x;
                    const pxY = xy.y;
                    this.pins.push({
                        id: domOps['id'],
                        x: pxX,
                        y: pxY,
                        class: domOps['class'],
                    });
                })
            },
            getGeo2Px(testPin: any): any {
                const map: any = this.$refs.map;

                const iw = map.offsetWidth;
                const ih = map.offsetHeight;

                const startPos = GeoUtils.convertPos(GeoUtils.start.lat, GeoUtils.start.lon);
                const endPos = GeoUtils.convertPos(GeoUtils.end.lat, GeoUtils.end.lon);
                const currentXY = GeoUtils.convertPos(testPin.lat, testPin.lon);

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

            loadPins() {
                Api.getPins().then((res) => {
                    this.pinInfoArr = res.spotInfos;
                    this.updatePins();
                }).catch(() => {
                    this.errorMessage = 'ネットワークエラー';
                });

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
                this.updateUserLocation(userLocation);
            },
            updateUserLocation(userLocation: any) {
                const xy = this.getGeo2Px(userLocation);

                this.userLocation = {
                    x: xy.x,
                    y: xy.y
                };
                this.userTruthLocation = {
                    lat: userLocation.lat,
                    lon: userLocation.lon
                }
            },
            changeShare() {
                this.shareFlag = !this.shareFlag;
            },
            loadAreas() {
                Api.getAreas().then(res => {
                    this.areas = res.areaInfos;
                    this.updateAreas();
                }).catch(() => {
                    this.errorMessage = 'ネットワークエラー';
                });
            },
            updateAreas() {
                this.areas.forEach((v: any, k: number) => {
                    const leftUp = v.region.leftUp;
                    const leftUpPx = this.getGeo2Px({
                        lat: leftUp.latitude,
                        lon: leftUp.longitude
                    });
                    const rightBottom = v.region.rightBottom;
                    const rightBottomPx = this.getGeo2Px({
                        lat: rightBottom.latitude,
                        lon: rightBottom.longitude
                    });

                    let w = rightBottomPx.x - leftUpPx.x;
                    let h = rightBottomPx.y - leftUpPx.y;
                    if (w > h) {
                        h = w;
                    } else {
                        w = h;
                    }

                    this.$set(this.areas[k], 'leftUp', leftUpPx);
                    this.$set(this.areas[k], 'width', w);
                    this.$set(this.areas[k], 'height', h);
                    this.$set(this.areas[k], 'opacity', v.hotScore / 100.0);
                });
            },
            getMapPosHandler(x: number, y: number) {
                const map: any = this.$refs.map;

                const startPos = GeoUtils.convertPos(GeoUtils.start.lat, GeoUtils.start.lon);
                const endPos = GeoUtils.convertPos(GeoUtils.end.lat, GeoUtils.end.lon);

                let cx = x - this.touchStartPos.backPosX;
                let cy = y - this.touchStartPos.backPosY;

                let iw = map.offsetWidth;
                let ih = map.offsetHeight;

                const bx = iw / (endPos.x - startPos.x);
                const by = ih / (endPos.y - startPos.y);

                cx /= bx;
                cy /= by;
                cx += startPos.x;
                cy += startPos.y;

                const pos = GeoUtils.convertPosFromPx(cx, cy);

                this.$set(this.adminLatAndLon, 'lat', pos.lat);
                this.$set(this.adminLatAndLon, 'lon', pos.lon);
                this.$set(this.adminLatAndLon, 'area', GeoUtils.containArea(this.areas, pos.lat, pos.lon));
                this.adminLocation = {
                    x: x - this.touchStartPos.backPosX,
                    y: y - this.touchStartPos.backPosY,
                };
            },
            showDetail(pin: any) {
                const id = pin['id'];
                this.$router.push('/spot/' + id);
            },
            sendLike() {
                const areaId = this.currentAreaId;
                if (typeof areaId === "undefined") {
                    return;
                }

                const xRand = 30 - Math.round(Math.random() * 30);
                const yRand = 30 - Math.round(Math.random() * 30);
                const effectRand = Math.ceil(Math.random() * 2);

                const now = new Date().getTime();
                let e = {
                    createdTime: now,
                    effectType: effectRand,
                    x: this.userLocation.x + xRand,
                    y: this.userLocation.y + yRand,
                };
                this.likeEffects.push(e);

                if (this.likeEffects.length > 300) {
                    let deleteKeys = Array<number>();
                    this.likeEffects.forEach((v, index) => {
                        const isDelete = now - v.createdTime > 800;
                        if (isDelete) {
                            deleteKeys.push(index);
                        }
                    });
                    deleteKeys.forEach((i) => {
                        this.$delete(this.likeEffects, i);
                    });
                }
                Api.areaLike(areaId).then(res => {
                    if (res.message !== 'success') {
                        this.errorMessage = '不明なエラーです'
                    }
                }).catch((e) => {
                    this.errorMessage = 'ネットワークエラー'
                })
            },
            reload() {
                location.reload();
            }
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

    .user-location-parent {
      position: relative;
      top: -100%;

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

    .area-range {
      position: relative;
      top: -200%;

      .area-rect {
        position: absolute;
        overflow: hidden;
        font-size: 8px;
        border-radius: 50%;
        border: solid 1px rgba(red, .3);
        background-color: rgba(red, .1);
      }

      .like-effect {
        display: flex;
        position: absolute;
        width: 200px;
        height: 200px;
        border-radius: 50%;
        top: 200px;
        left: 200px;
        margin: -100px;
        border: solid 1px rgba(red, .5);
        background-color: rgba(red, .15);
        justify-content: center;
        align-items: center;

        &.like-effect1 {
          animation: 0.5s ease like-effect forwards;
        }

        &.like-effect2 {
          animation: 0.5s ease like-effect2 forwards;
        }

        i {
          font-size: 120px;
        }
      }
    }


    .pin-parent {
      position: relative;
      top: -300%;

      .pin {
        $size: 30px;
        position: absolute;
        width: $size;
        height: $size;
        top: 0;
        left: 0;
        border-radius: 50%;

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

      .admin-pin {
        width: 15px;
        height: 15px;
        transform: translate(-50%, -50%);
        background-color: rgba(red, .8);
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

  @keyframes like-effect {
    0% {
      transform: scale(0.2) rotate(80deg);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    65% {
      opacity: 1;
    }
    80% {
      transform: scale(0.9);
    }
    100% {
      transform: scale(1) rotate(0deg);
      opacity: 0;
    }
  }

  @keyframes like-effect2 {
    0% {
      transform: scale(0.2) rotate(-80deg);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    65% {
      opacity: 1;
    }
    80% {
      transform: scale(0.9);
    }
    100% {
      transform: scale(1) rotate(0deg);
      opacity: 0;
    }
  }
</style>
