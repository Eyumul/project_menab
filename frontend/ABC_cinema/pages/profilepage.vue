<template>
    <div class="bg-black h-screen text-white">
        <div class="pt-[100px] flex flex-col bg-black items-center">
            <h1 class="text-center text-5xl text-[#0089D0] font-bold">My Profile</h1>
            <div v-if="isAuthenticated">
                <div class="flex flex-col my-[50px] py-[100px] space-y-[80px] w-[1200px]">
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="text-3xl text-gray-500 font-bold">Name </h2>
                        <p class="text-3xl text-[#90D9FF] font-medium">{{ username.nickname }}</p>
                    </div>
                    <div class="w-full bg-gray-700 h-[1px]"></div>
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="flex space-x-2 text-3xl text-gray-500 font-bold">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="self-center size-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
                            </svg>
                            <p>Email</p>
                        </h2>
                        <p class="text-3xl text-[#90D9FF] font-medium">{{ username.email }}</p>
                    </div>
                    <div class="w-full bg-gray-700 h-[1px]"></div>
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="flex text-3xl space-x-2 text-gray-500 font-bold">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="self-center size-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z" />
                            </svg>
                            <p>Phone N<span class="underline">o</span></p>
                        </h2>
                        <p class="text-3xl text-[#90D9FF] font-medium">+251 912345678</p>
                    </div>
                    <div @click="logout" class="cursor-pointer text-[#0089D0] text-right text-xl underline">Logout</div>
                </div>
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