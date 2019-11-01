<template>
  <div>
    <p class="area-text body-1 text-center">{{ areaName }}</p>

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
                return 'area1';
            }
        }
    });
</script>

<style lang="scss" scoped>
  .area-text {
    width: 10rem;
    background-color: rgba(229,57,53,0.5);
    border-radius: 1em;
  }
</style>
