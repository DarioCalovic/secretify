import path from 'path'
import fs from 'fs'

/*
 ** This is for GitHub pages
 */
const routerBase = {}

export default {
  /*
   ** Concat router base setting
   */
  ...routerBase,
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.png' },
      { rel: 'dns-prefetch', href: 'https://fonts.gstatic.com' },
      {
        rel: 'stylesheet',
        type: 'text/css',
        href: 'https://fonts.googleapis.com/css2?family=Poppins:wght@400;500&display=swap',
      },
      {
        rel: 'stylesheet',
        type: 'text/css',
        href: 'https://fonts.googleapis.com/css2?family=Roboto+Mono:wght@300;400&display=swap',
      },
      {
        rel: 'stylesheet',
        type: 'text/css',
        href:
          'https://cdn.materialdesignicons.com/4.9.95/css/materialdesignicons.min.css',
      },
    ],
  },
  /*
  * Server
  */
  server: {
    https: {
      key: fs.readFileSync(path.resolve(__dirname, 'cert.key')),
      cert: fs.readFileSync(path.resolve(__dirname, 'cert.pem'))
    }
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: ['./assets/scss/main.scss'],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [
    { src: '~/plugins/repositories.js' },
    { src: '~/plugins/axios.js' },
    { src: '~/plugins/filters.js' },
    { src: '~/plugins/vee-validate.js', mode: 'client' },
    { src: '~/plugins/crypto.js', mode: 'client' },
    { src: '~/plugins/track.js', mode: 'client' },
  ],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: false,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module',
    //'@nuxtjs/moment',
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://buefy.github.io/#/documentation
    ['nuxt-buefy', { materialDesignIcons: false }],
    '@nuxtjs/axios',
    'nuxt-clipboard',
    '@nuxt/content',
    // Nuxt 3 only '@nuxt/content',
    // Doc: https://axios.nuxtjs.org/usage
  ],
  /*
  * Nuxt content
  */
  content: {
    navigation: false,
  },
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    proxy: false,
  },
  /*
  ** Proxy module configuration
  */
  proxy: {
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {
    loaders: {
      sass: {}
    },
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {
      if (!config.externals) {
        config.externals = {}
      }
      if (ctx.isDev && ctx.isClient && config.module) {
        config.module.exprContextCritical = false;
      }
    },
    plugins: [
    ]
  },
  publicRuntimeConfig: {
    apiURL: process.env.API_URL || 'http://localhost:8800/api/v1',
    uiURL: process.env.UI_URL || 'https://localhost:3000',
    apiKey: process.env.API_KEY || '',
    marketing: {
      show: process.env.MARKETING_SHOW || 'true',
    },
    branding: {
      primary_color: process.env.BRANDING_PRIMARY_COLOR || '#669eff',
      logo: process.env.BRANDING_LOGO || ''
    },
    track: {
      enabled: process.env.TRACK_ENABLED || 'false',
      domain: process.env.TRACK_PLAUSIBLE_DOMAIN || 'localhost:3000',
      localhost: process.env.TRACK_PLAUSIBLE_LOCALHOST || 'true',
      apiURL: process.env.TRACK_PLAUSIBLE_API_URL || 'http://localhost:8000',
    }
  },
}
