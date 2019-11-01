<template>
  <div>
    <!-- Admin Plus Button -->
    <v-btn fixed bottom right fab large color="accent" v-show="!dialogShow" @click="dialogShow=true">
      <v-icon>mdi-plus</v-icon>
    </v-btn>

    <!-- Admin Dialog -->
    <v-dialog v-model="dialogShow" fullscreen>
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="dialogShow = false">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>新しいスポットの追加</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="dialogShow = false">Save</v-btn>
        </v-toolbar-items>
      </v-toolbar>

      <v-form class="px-3">
        <v-file-input label="写真"></v-file-input>
        <v-text-field label="タイトル "></v-text-field>
        <p>期間</p>
        <v-checkbox label="すべての期間" v-model="allTime"></v-checkbox>
        <v-radio-group :disabled="allTime">
          <v-radio label="土曜日" value="1"></v-radio>
          <v-radio label="日曜日" value="2"></v-radio>
          <v-checkbox label="終日" v-model="allDate"></v-checkbox>
          <v-btn :disabled="allTime || allDate" color="accent" @click="timeDialogShow=true">開始時間を設定</v-btn>
          <p>{{ startAndEndTime.startTime }} 〜 {{ startAndEndTime.endTime }}</p>
        </v-radio-group>

        <v-textarea label="説明"></v-textarea>

        <v-col class="text-right" cols="12" sm="4">
          <v-btn color="accent">SAVE</v-btn>
        </v-col>
      </v-form>

    </v-dialog>

    <v-dialog v-model="timeDialogShow">
      <v-card>
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
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface startAndEndTimeInterface {
        startTime: string
        endTime: string
    }

    interface AdminInterface {
        dialogShow: boolean,
        timeDialogShow: boolean,
        allTime: boolean,
        allDate: boolean,
        startAndEndTime: startAndEndTimeInterface,
        startAndEndTimeTemp: startAndEndTimeInterface,
    }

    export default Vue.extend({
        name: 'admin',
        data(): AdminInterface {
            return {
                dialogShow: false,
                timeDialogShow: false,
                allTime: true,
                allDate: true,
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
            }
        }
    });
</script>

<style scoped lang="scss">

</style>
