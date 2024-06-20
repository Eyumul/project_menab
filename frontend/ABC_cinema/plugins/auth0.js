import { createAuth0 } from '@auth0/auth0-vue'

export default defineNuxtPlugin((nuxtApp) => {
  // const config = useRuntimeConfig()
  const auth0 = createAuth0({
    domain: "dev-7dv1hf0d0zv281zb.us.auth0.com",
    clientId: "H07gJ2qsCtzOQ2WYVYqlEzmdwSkYwF33",
    authorizationParams: {
      redirect_uri: 'http://localhost:3000'
    },
    // cacheLocation: 'localstorage'
  })

  if (process.client) {
    nuxtApp.vueApp.use(auth0)
  }
})
