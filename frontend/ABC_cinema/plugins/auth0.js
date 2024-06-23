import { createAuth0 } from '@auth0/auth0-vue'

export default defineNuxtPlugin((nuxtApp) => {
  // const config = useRuntimeConfig()
  const auth0 = createAuth0({
    domain: "<domain name>",
    clientId: "<your client Id>",
    authorizationParams: {
      redirect_uri: 'http://localhost:3000',
      audience: '<your api audience>'
    },
    cacheLocation: 'localstorage'
  })

  if (process.client) {
    nuxtApp.vueApp.use(auth0)
  }
})
