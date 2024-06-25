<template>
    <div class="text-white bg-black">
        <header class="h-full flex flex-col justify-between ">
            <nav class="flex h-20 justify-between trans fixed w-full">
                <div class="px-3 self-center"><NuxtLink to = "/"><p class="logo">ABC_Cinema</p></NuxtLink></div>
                <ul class = "flex gap-x-12 px-9 items-center ">
                    <li class="cursor-pointer hover:text-[#0089D0]"><NuxtLink to = "/search">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="h-4 w-5">
                            <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                        </svg>
                    </NuxtLink></li>
                    <li class="cursor-pointer hover:text-[#0089D0]" v-if="!isAuthenticated"><span @click="login" class="cursor-pointer hover:text-[#0089D0]">Dashboard</span></li>
                    <li class="cursor-pointer hover:text-[#0089D0]" v-else><NuxtLink to = "/#here">Dashboard</NuxtLink></li>
                    <li class="cursor-pointer hover:text-[#0089D0]"><NuxtLink to = "/moviesFeed">Movies</NuxtLink></li>
                    <li class="cursor-pointer hover:text-[#0089D0]"><NuxtLink to = "/schedule">Schedule</NuxtLink></li>
                    <div v-if="!isAuthenticated">
                        <li class="cursor-pointer hover:text-[#0089D0]" @click = "login">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-7">
                                <path fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 1 1 9 0 4.5 4.5 0 0 1-9 0ZM3.751 20.105a8.25 8.25 0 0 1 16.498 0 .75.75 0 0 1-.437.695A18.683 18.683 0 0 1 12 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 0 1-.437-.695Z" clip-rule="evenodd" />
                            </svg>
                        </li>
                    </div>
                    <div v-else>
                        <NuxtLink to="/profilepage">
                            <NuxtImg :src="username.picture" class="size-8 rounded-[100%]"></NuxtImg>
                        </NuxtLink>
                    </div>
                </ul>
            </nav>
        </header>
        <slot />
    </div>
</template>

<script setup>
import { useAuth0 } from '@auth0/auth0-vue';
const auth0 = process.client ? useAuth0() : undefined
const login = () => {
    auth0?.loginWithRedirect()
}
const isAuthenticated = computed(() => {
  return auth0?.isAuthenticated.value
})
const username = computed(() => {
    return auth0?.user.value
})
</script>

<style>
/*

dark blue: #0089D0
light blue: #24B4FF 

*/
.trans {
    background-color: rgb(0, 0, 0,0);
    background-image: linear-gradient(rgb(0, 0, 0),rgba(110, 110, 110, 0));
}
.logo {
    font-family: sans-serif;
    font-size: 32px;
    background: -webkit-linear-gradient(right, white, #0089D0);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
</style>