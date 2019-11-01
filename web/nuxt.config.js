import colors from 'vuetify/es5/util/colors'

export default {
  mode: 'spa',
  server: {
    port: 2810,
    host: '0.0.0.0'
  },
  /*
  ** Headers of the page
  */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {hid: 'description', name: 'description', content: process.env.npm_package_description || ''}
    ],
    link: [
      {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: {color: '#fff'},
  /*
  ** Global CSS
  */
  css: [],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    'plugins/vuetifyFontawasome'
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxt/typescript-build',
    '@nuxtjs/vuetify',
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    '@nuxtjs/pwa',
  ],
  manifest: {
    name: 'torch',
    lang: 'ja',
    icons: [
      {
        src: "/pwa/logo_64.png",
        sizes: "64x64",
        type: "image/png",
      },
      {
        src: "/pwa/logo_120.png",
        sizes: "120x120",
        type: "image/png",
      },
      {
        src: "/pwa/logo_144.png",
        sizes: "144x144",
        type: "image/png",
      },
      {
        src: "/pwa/logo_152.png",
        sizes: "152x152",
        type: "image/png",
      },
      {
        src: "/pwa/logo_192.png",
        sizes: "192x192",
        type: "image/png",
      },
      {
        src: "/pwa/logo_384.png",
        sizes: "384x384",
        type: "image/png",
      },
      {
        src: "/pwa/logo_512.png",
        sizes: "512x512",
        type: "image/png",
      }
    ]
  },
  /*
  ** vuetify module configuration
  ** https://github.com/nuxt-community/vuetify-module
  */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: false,
      themes: {
        light: {
          primary: colors.red.darken2,
          accent: colors.red.darken3,
          secondary: colors.red.lighten3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
    },
    filenames: {
      app: ({isDev}) => isDev ? '[name].[hash].js' : '[chunkhash].js',
      chunk: ({isDev}) => isDev ? '[name].[hash].js' : '[chunkhash].js'
    }
  }
}
