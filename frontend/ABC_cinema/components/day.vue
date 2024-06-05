<template>
    <div>
        <fieldset  class="flex flex-col space-y-4 bg-black border-2 border-[rgb(0,137,208,0.5)] rounded-[20px] p-5">
                <legend> {{ days }} </legend>
                <div class="flex justify-between items-center">
                    <label>2:00 PM</label>
                    <select class = "w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option>Brave</option>
                        <option>It</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>5:00 PM</label>
                    <select class = "w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option>Brave</option>
                        <option>It</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>9:00 PM</label>
                    <select class = "w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option>Brave</option>
                        <option>It</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>12:00 PM</label>
                    <select class = "w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option>Brave</option>
                        <option>It</option>
                    </select>
                </div>
            </fieldset>
    </div>
</template>

<script setup>
const props = defineProps(["days"]);
</script>

<style scoped>

</style>