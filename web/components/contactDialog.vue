<template>
  <div>
    <v-dialog v-model="dialog">
      <v-card>
        <v-card-title>お問い合わせ</v-card-title>
        <v-card-text>
          <p>お問い合わせは、<span>K1号館 106号室 ソフトウェア工房</span>までお越しください。</p>
        </v-card-text>

        <v-card-actions>
          <v-flex xs3 offset-xs9 align-end>
            <v-btn class="ml-auto" color="accent" right @click="dialog = false">閉じる</v-btn>
          </v-flex>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    interface contactModalInterface {
        dialog: boolean
    }

    export default Vue.extend({
        name: "contactDialog",
        props: ['value'],
        data() {
            return {
                dialog: false,
            }
        },
        watch: {
            value(value: boolean) {
                if (value) {
                    this.dialog = value;
                }
            },
            dialog(value: boolean) {
                this.$emit('input', value);
            }
        }
    });
</script>

<style scoped lang="scss">

</style>
