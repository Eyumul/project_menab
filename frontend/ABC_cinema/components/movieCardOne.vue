<template>
    <div class=" flex flex-col scard">
        <div class="scardthumbnail w-full h-5/6"></div>
        <div class="flex justify-between w-full h-1/6 items-center ">
            <p class="s16 font-normal">{{movietitle}}</p>
            <NuxtLink :to=movielink ><p class="textColor s16 font-normal underline">See detail</p></NuxtLink>
        </div>
    </div>
</template>

<script setup>
const props = defineProps(["movietitle", "movielink"])
</script>

<style>
.scard{
    width:400px;
    height:300px;
    background-color:#262626;
    border:solid 3px rgb(0,137,208,0.2);
    
}
.scardthumbnail{
    background-image: url('/figmaImage/thumbnail1.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
</style>