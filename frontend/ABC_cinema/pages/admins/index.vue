<template>
    <div class="flex flex-col justify-center items-center bg-black text-white text-4xl h-screen space-y-12">
        <h1 class="text-white text-6xl font-bold font-[limelight]">Admin Pannel</h1>
        <div class="flex space-x-12">
            <NuxtLink to="/admins/directors">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center font-[oswald]">Directors</p>
                </div>
            </NuxtLink>
            <NuxtLink to="/admins/stars">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center font-[oswald]">Stars</p>
                </div>
            </NuxtLink>
        </div>
        <div class="flex space-x-12">
            <NuxtLink to="/admins/movies">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center font-[oswald]">Movies</p>
                </div>
            </NuxtLink>
            <NuxtLink to="/admins/schedules">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]   hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center font-[oswald]">Schedules</p>
                </div>
            </NuxtLink>
        </div>
        <NuxtLink class="text-xl items-center gap-3 hover:font-black hover:scale-110 text-[#0089D0] hover:underline flex" to="./">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="size-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
            </svg>
            <p>View</p>
        </NuxtLink>
    </div>
</template>

<script setup>
definePageMeta({
        layout:""
    })
</script>

<style>

</style>