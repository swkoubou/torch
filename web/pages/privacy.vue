<template>
  <div>
    <v-app-bar app color="primary" dark>
      <nuxt-link to="/">
        <v-icon>mdi-arrow-left</v-icon>
      </nuxt-link>
      <v-toolbar-title class="mx-2 mt-1">プライバシーポリシー</v-toolbar-title>
    </v-app-bar>
    <v-container id="scroll-target" style="max-height: 600px" class="overflow-y-auto">
      <v-row v-scroll:#scroll-target="onScroll">
        <v-card>
          <v-card-title>利用規約</v-card-title>
          <v-card-text class="caption">
            ソフトウェア工房（以下，「swk」といいます。）は，本ウェブサイト上で提供するサービス（以下,「本サービス」といいます。）における，ユーザーの個人情報の取扱いについて，以下のとおりプライバシーポリシー（以下，「本ポリシー」といいます。）を定めます。
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第1条（個人情報）</v-card-title>
          <v-card-text class="caption">
            「個人情報」とは，個人情報保護法にいう「個人情報」を指すものとし，生存する個人に関する情報であって，当該情報に含まれる氏名，生年月日，住所，電話番号，連絡先その他の記述等により特定の個人を識別できる情報及び容貌，指紋，声紋にかかるデータ，及び健康保険証の保険者番号などの当該情報単体から特定の個人を識別できる情報（個人識別情報）を指します。
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第2条（利用者情報の取得と利用目的）</v-card-title>
          <v-card-text class="caption">
            1. swkが本アプリケーションにおいて取得する利用者情報（以下、「利用者情報」といいます）は以下のとおりです。swkは、利用者情報を、アプリケーション経由で自動的に取得し、またはお客様の手入力により登録していただくことにより取得します。<br>
            1-1. 取得方法：アプリケーションによる自動取得 - 本サービスの利用状況の分析および品質改善のため/サーバ連携機能の提供のため<br>
            2. 本アプリケーションおよび本サービスにおける利用者情報の利用目的は上記に定めるとおりです。swkは、swkが別に定める「お客様の個人情報に関するプライバシーポリシー」その他お客様とswkの間の契約において別段の定めがある場合には、利用者情報を、これらに記載の利用目的の範囲で利用することがあります。<br>
            3. swkは、本アプリケーションおよび本サービスの提供のために、本アプリケーション以外でswkが適正に取得したお客様の個人情報を利用することがあります。また、swkは利用者情報とこれらの個人情報とを組み合わせた上で、上記の利用目的の範囲で利用することがあります。<br>
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第3条（情報収集モジュールによる取得）</v-card-title>
          <v-card-text class="caption">
            swkは、前条に記載の利用者情報を、本アプリケーション内に組み込まれた下記の情報収集モジュールにて取得します。情報収集モジュールとは、第三者が提供するプログラムであって、利用者情報を取得・解析するための機能をもつものをいいます。なお、swkは情報収集モジュールで取得した利用者情報を、お客様個人を識別するIDと組み合わせた上で、お客様の属性情報等を付加し、利用する場合があります。<br>
            Google Analyticsを使用し、本アプリケーションのご使用状況を取得します。また、利用者情報はGoogleの管理するサーバシステムに格納されます。なお、Googleは当該利用者情報を、Googleが定める利用目的の範囲で利用します。Google Analyticsの詳細およびGoogleが定める利用目的は別ウインドウが開きますhttps://policies.google.com/privacy?hl=jaをご参照ください。<br>
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第4条（利用者情報の第三者提供）</v-card-title>
          <v-card-text class="caption">
            利用者情報は以下各号に定める場合を除き、第三者に送信され、または提供されることはありません。<br>
            1. 本サービスの利用規約に定める目的で提携事業者へ提供する場合（第1条②に定める利用者情報について）<br>
            2. 第3条定める情報収集モジュール提供者へ送信する場合<br>
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第5条（利用者情報の取得・蓄積・利用に関する同意）</v-card-title>
          <v-card-text class="caption">
            お客様は本アプリケーションを使用する際に、本アプリケーション・プライバシーポリシーの内容を確認し、同意を付与する機会が与えられます。本アプリケーションおよび本サービスは本アプリケーション・プライバシーポリシーをご確認いただき、内容を理解したうえでご利用ください。また、本アプリケーション・プライバシーポリシーは本アプリの3点リーダー→利用規約に公開いたします。
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第6条（利用者情報の取得停止等）</v-card-title>
          <v-card-text class="caption">
            本アプリケーションでは、利用者情報の取得、利用を停止する手段を提供しておりません。これらを停止したい場合には、本アプリケーションをアンインストールしてください。
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-title>第7条（本アプリケーション・プライバシーポリシーの変更）</v-card-title>
          <v-card-text class="caption">
            swkは、本アプリケーションのバージョンアップに伴って、本アプリケーション・プライバシーポリシーを変更することがあります。変更内容については、新バージョンのアプリケーションに付随するアプリケーション・プライバシーポリシーをご参照ください。
          </v-card-text>
        </v-card>
      </v-row>
    </v-container>
  </div>
</template>

<script>
    export default {
        name: "privacy"
    }
</script>

<style scoped>

</style>
