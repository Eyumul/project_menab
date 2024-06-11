<template>
    <div class="rounded-[18px] flex flex-col w-[400px] h-[300px] bg-[#0a0a0a] border-[rgb(0,137,208,0.1)] border-2 items-center">
        <div class=" w-full rounded-t-[18px] overflow-hidden h-[220px]"><NuxtImg class="" :src="moviethumbnail"></NuxtImg></div>
        <div class=" text-xl flex justify-between w-full h-[80px] items-center ">
            <p class=" font-black title ml-4 w-[240px]">{{movietitle}}</p>
            <NuxtLink :to=movielink ><p class="text-lg mr-2 w-[85px] textColor font-normal underline">See detail</p></NuxtLink>
        </div>
    </div>
</template>

<script setup>
const props = defineProps(["movietitle", "movielink", "moviethumbnail"])
</script>

<style>
.scard{
    width:400px;
    height:300px;
    background-color:#262626;
    border:solid 3px rgb(0,137,208,0.2);
    border-radius: 20px;
    
}
.title {
    background: -webkit-linear-gradient(right, #24B4FF, #90D9FF);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
</style>