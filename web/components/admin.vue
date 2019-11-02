<template>
  <div>
    <!-- Admin Plus Button -->
    <v-btn fixed bottom right fab large color="accent" v-show="!dialogShow" @click="dialogShow=true">
      <v-icon>mdi-plus</v-icon>
    </v-btn>

    <!-- Admin Dialog -->
    <v-dialog v-model="dialogShow" fullscreen hide-overlay>
      <v-card>
        <v-toolbar dark color="primary">
          <v-btn icon dark @click="dialogShow = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>新しいスポットの追加</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark text @click="save">Save</v-btn>
          </v-toolbar-items>
        </v-toolbar>

        <v-card-text class="mt-4">
          <v-form>
            <v-text-field label="緯度" type="number" v-model="latAndLon.lon"></v-text-field>
            <v-text-field label="経度" type="number" v-model="latAndLon.lat"></v-text-field>
            <div style="overflow-x: hidden">
              <v-file-input label="写真" accept="image/*" ref="photo" show-size style="width:100%"></v-file-input>
            </div>
            <v-text-field label="タイトル" v-model="title"></v-text-field>
            <p>期間</p>
            <v-checkbox label="すべての期間" v-model="allTime"></v-checkbox>
            <v-radio-group :disabled="allTime" v-model="targetDate">
              <v-radio label="土曜日" value="1"></v-radio>
              <v-radio label="日曜日" value="2"></v-radio>
              <v-checkbox label="終日" v-model="allDate"></v-checkbox>
              <v-btn :disabled="allTime || allDate" color="accent" @click="timeDialogShow=true">期間を設定</v-btn>
              <p>{{ startAndEndTime.startTime }} 〜 {{ startAndEndTime.endTime }}</p>
            </v-radio-group>

            <v-textarea label="説明" v-model="description"></v-textarea>

            <v-col class="text-right" cols="12" sm="4">
              <v-btn color="accent" @click="save">SAVE</v-btn>
            </v-col>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-dialog v-model="timeDialogShow">
      <v-card-title>時間を設定</v-card-title>
      <v-card-text>
        <p>開始時刻</p>
        <v-time-picker format="24hr" v-model="startAndEndTimeTemp.startTime"></v-time-picker>
        <p class="mt-4">終了時刻</p>
        <v-time-picker format="24hr" v-model="startAndEndTimeTemp.endTime"></v-time-picker>
      </v-card-text>
      <v-card-actions>
        <v-btn block color="accent" large @click="applyTimeSettings">設定して閉じる</v-btn>
      </v-card-actions>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';
    import Api from "~/module/api";

    interface startAndEndTimeInterface {
        startTime: string
        endTime: string
    }

    interface AdminInterface {
        latAndLon: {
            lat: number
            lon: number
        }
        dialogShow: boolean
        timeDialogShow: boolean
        title: string
        description: string
        allTime: boolean
        allDate: boolean
        targetDate: string,
        startAndEndTime: startAndEndTimeInterface
        startAndEndTimeTemp: startAndEndTimeInterface
    }

    export default Vue.extend({
        name: 'admin',
        props: ['value'],
        watch: {
            value(value) {
                this.latAndLon = Object.assign({}, value);
            }
        },
        data(): AdminInterface {
            return {
                latAndLon: {
                    lat: 0,
                    lon: 0,
                },
                dialogShow: false,
                timeDialogShow: false,
                title: '',
                description: '',
                allTime: true,
                allDate: true,
                targetDate: '1',
                startAndEndTime: {
                    startTime: '',
                    endTime: '',
                },
                startAndEndTimeTemp: {
                    startTime: '',
                    endTime: '',
                }
            };
        },
        methods: {
            applyTimeSettings(): void {
                this.startAndEndTime = Object.assign({}, this.startAndEndTimeTemp);
                this.timeDialogShow = false;
            },
            save(): void {
                const satDate = '2019-11/02';
                const sunDate = '2019-11/03';

                let start = new Date();
                let end = new Date();

                if (this.allTime) {
                    start = new Date(satDate + ' 00:00:00');
                    end = new Date(sunDate + ' 23:59:59');
                } else {
                    let startTime = '00:00:00';
                    let endTime = '23:59:59';

                    if (!this.allDate) {
                        startTime = this.startAndEndTime.startTime + ':00';
                        endTime = this.startAndEndTime.endTime + ':00';
                    }

                    if (this.targetDate == '1') {
                        start = new Date(satDate + ' ' + startTime);
                        end = new Date(satDate + ' ' + endTime);
                    } else if (this.targetDate == '2') {
                        start = new Date(sunDate + ' ' + startTime);
                        end = new Date(sunDate + ' ' + endTime);
                    }
                }

                let postData = new FormData();

                // @ts-ignore
                const imageInput = this.$refs['photo'].$el.querySelector('input');
                const file = imageInput.files[0];
                postData.append("image", file);

                let option = {
                    name: this.title,
                    description: this.description,
                    loc_latitude: this.latAndLon.lat,
                    lon_latitude: this.latAndLon.lon,
                    area_id: "!!!!!!!!!!!!!!!!!!",
                    start_year: start.getFullYear(),
                    start_month: start.getMonth(),
                    start_day: start.getDate(),
                    start_hour: start.getHours(),
                    start_minute: start.getMinutes(),
                    start_second: start.getSeconds(),
                    end_year: end.getFullYear(),
                    end_month: end.getMonth(),
                    end_day: end.getDate(),
                    end_hour: end.getHours(),
                    end_minute: end.getMinutes(),
                    end_second: end.getSeconds(),
                };

                postData.append("option", JSON.stringify(option));

                Api.addSpot(postData).then(() => {

                })
                this.dialogShow = false
            }
        },
    });
</script>

<style scoped lang="scss">

</style>
