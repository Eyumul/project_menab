<template>
    <div class="bg-black h-screen text-white">
        <div class="pt-[72px]">
            <h1>Auth0 Integration with Nuxt 3</h1>
            <div v-if="!isAuthenticated">
                <button @click="login">Login</button>
            </div>
            <div v-else>
                <p>Welcome, {{ username }}</p>
                <NuxtImg :src="username.picture"></NuxtImg>
                <button @click="logout">Logout</button>
            </div>
        </div>
    </div>
</template>
<script setup>
    import { useAuth0 } from '@auth0/auth0-vue';
    const token = ref(null)
    const auth0 = process.client ? useAuth0() : undefined
    const logout = () => {
        auth0?.logout()
    }
    token.value = auth0?.getAccessTokenSilently()
    console.log(token.value)
    const isAuthenticated = computed(() => {
    return auth0?.isAuthenticated.value
    })
    const username = computed(() => {
        return auth0?.user.value
    })
</script>

<style scoped>

</style>