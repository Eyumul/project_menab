<template>
    <div class="h-screen bg-black"> 
        <div class="text-white">
            <nav class=" pl-[138px] pt-[50px] space-x-4 flex justify-between">
                <NuxtLink to = "/admins" class="flex hover:text-[#0089D0]" >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 self-center">
                        <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06l-7.5-7.5Z" clip-rule="evenodd" />
                    </svg>
                    <p class="self-center text-2xl font-black">Back</p>
                </NuxtLink>
                <p class="text-4xl content-center pr-[800px]">Edit Movies</p>
            </nav>
        </div>
        <div class="flex flex-col text-xl space-y-12 py-24 w-full bg-black items-center">
            <AddMovieForm class="w-[600px]"/>
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