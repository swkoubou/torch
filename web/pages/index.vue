<template>
  <div>
    <v-btn fixed top left disabled outlined rounded large color="primary" width="10rem">{{areaName}}</v-btn>

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
    }

    export default Vue.extend({
        name: 'index',
        components: {admin, pinDetail},
        data(): indexData {
            return {
                isAdmin: false,
            };
        },
        created(): void {
            const query = this.$route.query;
            if ('admin' in query) {
                this.isAdmin = query['admin'] === 'hoo0Jaek8jooTeeti0eiciedeithougee4aexooGhaiNgieDa9gio6jaipeevach';
            }
        },
        computed: {
            areaName(): string {
                return 'さばんなちほー';
            }
        }
    });
</script>

<style lang="scss" scoped>
</style>
