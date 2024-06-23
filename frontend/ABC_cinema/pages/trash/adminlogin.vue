<template>
    <div class="h-screen bg-black text-white flex justify-center">
        <form class="flex flex-col space-y-10 items-center self-center">
            <p class="text-2xl underline">Enter admins password</p>
            <input type="password" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <button class="w-[380px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Login</button>
        </form>
    </div>
</template>

<script setup>
    definePageMeta({
            layout:""
        })
</script>

<style>

</style>