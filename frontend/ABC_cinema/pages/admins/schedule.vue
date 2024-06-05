<template>
    <div class="bg-black text-white pt-64 pb-36 w-full">
        <h1 class="text-4xl underline text-center">This weeks Schedule</h1>
        <form class="flex flex-wrap text-white text-xl w-[100%] p-10">
            <div v-for="day in days">
                <day :days="day" class="p-8"/>
            </div>
        </form>
    </div>
</template>

<script setup>
    definePageMeta({
            layout:"admin"
        })
    const days = ["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]
</script>

<style scoped>

</style>