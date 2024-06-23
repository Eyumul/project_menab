<template>
    <div class="flex flex-col justify-center items-center bg-black text-white text-4xl h-screen space-y-12">
        <h1 class="text-white text-6xl font-bold">Admin Pannel</h1>
        <div class="flex space-x-12">
            <NuxtLink to="/admins/directors">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center">Directors</p>
                </div>
            </NuxtLink>
            <NuxtLink to="/admins/stars">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center">Stars</p>
                </div>
            </NuxtLink>
        </div>
        <div class="flex space-x-12">
            <NuxtLink to="/admins/movies">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]  hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center">Movies</p>
                </div>
            </NuxtLink>
            <NuxtLink to="/admins/schedules">
                <div class="flex border-2 border-white w-[500px] justify-center rounded-[18px] h-[300px]   hover:text-[#0089D0] hover:border-[#0089D0]">
                    <p class="content-center">Schedules</p>
                </div>
            </NuxtLink>
        </div>
    </div>
</template>

<script setup>
definePageMeta({
        layout:""
    })
</script>

<style>

</style>