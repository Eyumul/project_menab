<template>
    <div class="flex flex-col justify-between items-center bcard">
        <div class="bcardthumbnail relative">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class=" hover:text-[#0089D0] absolute bottom-0 right-0 size-6 bg-[rgb(0,0,0,0.5)]">
                <path stroke-linecap="round" stroke-linejoin="round" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z" />
            </svg>
        </div>
        <div class="flex justify-between w-[500px]">
            <p class="text-2xl w-[350px]"><NuxtLink to="/movies/movieOne" class="underline">Kingdom of the planet of the apes</NuxtLink>: <span>(Action, Sci-fi)</span></p>
            <div>
                <p class="text-[24px] text-right">2h 25m</p>
                <p class="s16 textColor text-rightr">May 23 - 7:00PM</p>
            </div>
        </div>
        <div class="flex space-x-[32px] bg-[#0089D0] w-[550px] h-[60px] rounded-b-[20px] items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 text-black">
                <path fill-rule="evenodd" d="M1.5 6.375c0-1.036.84-1.875 1.875-1.875h17.25c1.035 0 1.875.84 1.875 1.875v3.026a.75.75 0 0 1-.375.65 2.249 2.249 0 0 0 0 3.898.75.75 0 0 1 .375.65v3.026c0 1.035-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 17.625v-3.026a.75.75 0 0 1 .374-.65 2.249 2.249 0 0 0 0-3.898.75.75 0 0 1-.374-.65V6.375Zm15-1.125a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0V6a.75.75 0 0 1 .75-.75Zm.75 4.5a.75.75 0 0 0-1.5 0v.75a.75.75 0 0 0 1.5 0v-.75Zm-.75 3a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0v-.75a.75.75 0 0 1 .75-.75Zm.75 4.5a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0v-.75ZM6 12a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H6.75A.75.75 0 0 1 6 12Zm.75 2.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5h-3Z" clip-rule="evenodd" />
            </svg>
            <p class="text-2xl font-black">Get Ticket</p>
        </div>
    </div>
</template>

<script setup>

</script>

<style>
.bcard{
    width:550px;
    height:610px;
    background-color:#262626;
    border-radius:20px;
    border:solid 3px rgb(0,137,208,0.2);
    border-bottom: 0px;
    padding-top: 25px;
}
.bcardthumbnail{
    height: 400px;
    width:500px;
    border-radius:20px;
    background-image: url('/figmaImage/thumbnail1.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
</style>