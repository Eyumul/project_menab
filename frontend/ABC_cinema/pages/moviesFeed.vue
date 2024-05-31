<template>
    <div class="h-full bg-black text-white pb-40">
        <div class="flex justify-between items-end px-[138px] pt-[100px]">
            <div class="flex space-x-8 justify-center pt-16 content end">
                <img src="/public/figmaImage/movie.png"/>
                <p class="s32 font-normal">Movies</p>
            </div>
            <div class="flex space-x-[34px] items-center">
                <p class="text-[24px] font-light">Browse movies by</p>
                <Dropdownmenu/>
            </div>
        </div>
        <div class="my-[144px] flex space-x-[27px] justify-center">
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
        </div>
        <div class="my-[144px] flex space-x-[27px] justify-center">
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
        </div>
        <div class="my-[144px] flex space-x-[27px] justify-center">
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
            <MovieCardOne/>
        </div>
    </div>
</template>

<script setup>

</script>

<style>

</style>