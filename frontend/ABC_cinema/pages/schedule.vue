<template>
    <div class="bg-black text-white">
        <div class="pt-[130px]">
            <h1 class="text-4xl font-black text-[#90D9FF] text-center underline">{{ formatDate(today) + " (Today's)" }}</h1>
            <div v-for="schedule in schedules.schedule" :key="schedule">
                <div v-if="schedule.date == today" class="flex justify-around my-12 mx-36 py-4 rounded-[20px] border-[rgb(0,137,208,0.2)] border-2 bg-[#000e14] ">
                    <h2 class="self-center text-3xl w-[150px] text-[#0089D0]">{{ convertTo12HourFormat(schedule.time) }}</h2>
                    <NuxtImg :src="schedule.movie.thumbnail" class="w-[400px] h-[300px] rounded-[12px]"/>
                    <h3 class="self-center w-[400px] text-xl text-[#24B4FF] font-black"><NuxtLink :to="schedule.movie.title">{{ schedule.movie.title }}</NuxtLink></h3>
                </div>
            </div>
            <div v-if="istodayempty()" class="text-gray-500 font-black text-2xl text-center m-12">No Movie on this day </div>
        </div>     
        <div class="w-full bg-gray-700 h-[1px]"></div>
        <div class="pt-[130px]">
            <h1 class="text-4xl font-black text-[#90D9FF] text-center underline">{{ formatDate(tomorrow) + " (Tomorrow's)" }}</h1>
            <div v-for="schedule in schedules.schedule" :key="schedule">
                <div v-if="schedule.date == tomorrow" class="flex justify-around my-12 mx-36 py-4 rounded-[20px] border-[rgb(0,137,208,0.2)] border-2 bg-[#000e14] ">
                    <h2 class="self-center text-3xl w-[150px] text-[#0089D0]">{{ convertTo12HourFormat(schedule.time) }}</h2>
                    <NuxtImg :src="schedule.movie.thumbnail" class="w-[400px] h-[300px] rounded-[12px]"/>
                    <h3 class="self-center w-[400px] text-xl text-[#24B4FF] font-black"><NuxtLink :to="schedule.movie.title">{{ schedule.movie.title }}</NuxtLink></h3>
                </div>
            </div>
            <div v-if="istomorrowempty()" class="text-gray-500 font-black text-2xl text-center m-12">No Movie on this day </div>
        </div>
        <div class="w-full bg-gray-700 h-[1px]"></div>
        <div class="pt-[130px]">
            <h1 class="text-4xl font-black text-[#90D9FF] text-center underline">{{ formatDate(dayAfterTomorrow) + " (Day after tomorrow's)" }}</h1>
            <div v-for="schedule in schedules.schedule" :key="schedule">
                <div v-if="schedule.date == dayAfterTomorrow" class="flex justify-around my-12 mx-36 py-4 rounded-[20px] border-[rgb(0,137,208,0.2)] border-2 bg-[#000e14] ">
                    <h2 class="self-center text-3xl w-[150px] text-[#0089D0]">{{ convertTo12HourFormat(schedule.time) }}</h2>
                    <NuxtImg :src="schedule.movie.thumbnail" class="w-[400px] h-[300px] rounded-[12px]"/>
                    <h3 class="self-center w-[400px] text-xl text-[#24B4FF] font-black"><NuxtLink :to="schedule.movie.title">{{ schedule.movie.title }}</NuxtLink></h3>
                </div>
            </div>
            <div v-if="isdayafterempty()" class="text-gray-500 font-black text-2xl text-center m-12">No Movie on this day </div>
        </div>
        <div class="w-full bg-gray-700 h-[1px]"></div>
        <div class="pt-[130px] pb-12" v-if="!commingSoon.length == 0">
            <h1 class="text-4xl font-black text-[#90D9FF] text-center underline">{{ "Comming Soon" }}</h1>
            <div class="flex flex-wrap ml-32">
                <div v-for="movie in commingSoon" :key="movie">
                    <div class="my-12 mx-2 p-4">
                        <NuxtImg class="w-[380px] h-[290px] rounded-t-[20px]" :src="movie.thumbnail"/>
                        <h1 class="text-xl text-[#24B4FF] font-black my-2"><NuxtLink :to="movie.title">{{ movie.title }}</NuxtLink></h1>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
const schedulequery = gql`
query MyQuery {
  schedule {
    date
    movie_id
    id
    time
    movie {
      thumbnail
      title
    }
  }
}
`
const {data: schedules} = await useAsyncQuery(schedulequery)
function isPastDate(inputDate) {
  // Parse the input date
  const date = new Date(inputDate);

  // Get the current date and time
  const now = new Date();

  // Compare the input date with the current date
  return date < now;
}
function convertTo12HourFormat(time) {
  let [hours, minutes, seconds] = time.split(':').map(Number);
  const period = hours >= 12 ? 'PM' : 'AM';
  hours = hours % 12 || 12;
  hours = String(hours).padStart(2, '0');
  minutes = String(minutes).padStart(2, '0');
  seconds = String(seconds).padStart(2, '0');
  return `${hours}:${minutes} ${period}`;
}
function formatDate(inputDate) {
  const date = new Date(inputDate);
  const days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  const months = [
    "January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
  ];
  const dayName = days[date.getUTCDay()]; // Get the day of the week
  const monthName = months[date.getUTCMonth()]; // Get the month
  const day = date.getUTCDate(); // Get the day of the month
  const year = date.getUTCFullYear(); // Get the year
  return `${dayName} ${monthName} ${day} - ${year}`;
}
// Today's date
let today = new Date();
let dd = String(today.getDate()).padStart(2, '0');
let mm = String(today.getMonth() + 1).padStart(2, '0'); // January is 0!
let yyyy = today.getFullYear();

today = yyyy + '-' + mm + '-' + dd;

// Tomorrow's date
let tomorrow = new Date();
tomorrow.setDate(tomorrow.getDate() + 1);
dd = String(tomorrow.getDate()).padStart(2, '0');
mm = String(tomorrow.getMonth() + 1).padStart(2, '0'); // January is 0!
yyyy = tomorrow.getFullYear();

tomorrow = yyyy + '-' + mm + '-' + dd;

// Day after tomorrow's date
let dayAfterTomorrow = new Date();
dayAfterTomorrow.setDate(dayAfterTomorrow.getDate() + 2);
dd = String(dayAfterTomorrow.getDate()).padStart(2, '0');
mm = String(dayAfterTomorrow.getMonth() + 1).padStart(2, '0'); // January is 0!
yyyy = dayAfterTomorrow.getFullYear();

dayAfterTomorrow = yyyy + '-' + mm + '-' + dd;
const commingSoon = ref([])
for(const schedule of schedules.value.schedule){
    if (schedule.date != today && schedule.date != tomorrow && schedule.date != dayAfterTomorrow && !isPastDate(schedule.date)){
        commingSoon.value.push({"title": schedule.movie.title, "thumbnail": schedule.movie.thumbnail})
    }
}

function istodayempty() {
    for(const schedule of schedules.value.schedule){
        if (schedule.date == today) {return false}
    }
    return true
}
function istomorrowempty() {
    for(const schedule of schedules.value.schedule){
        if (schedule.date == tomorrow) {return false}
    }
    return true
}
function isdayafterempty() {
    for(const schedule of schedules.value.schedule){
        if (schedule.date == dayAfterTomorrow) {return false}
    }
    return true
}
</script>

<style scoped>
body {
    background-color: black;
}
</style>