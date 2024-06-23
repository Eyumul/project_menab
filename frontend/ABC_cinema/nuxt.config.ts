// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/apollo', '@nuxt/image'],
  apollo: {
    clients: {
      default: {
        httpEndpoint: 'http://localhost:8080/v1/graphql',
        tokenStorage: 'localStorage',
        tokenName: 'hasura-token'
      }
    },
  },
  image: {
    domains: ['nuxtjs.org']
  },
  plugins: [
    '~/plugins/auth0.js'
  ]
})