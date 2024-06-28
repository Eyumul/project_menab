<template>
    <div class="bg-black text-white h-screen w-full">
        <nav class=" pl-[138px] pt-[50px] space-x-4 flex justify-between">
                <NuxtLink to = "/admins" class="flex hover:text-[#0089D0]" >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 self-center">
                        <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06l-7.5-7.5Z" clip-rule="evenodd" />
                    </svg>
                    <p class="self-center text-2xl font-black">Back</p>
                </NuxtLink>
                <p class="text-4xl content-center pr-[800px]">Edit Schedules</p>
            </nav>
            <div class="flex flex-col">
                <h1 class="text-2xl self-center py-16">Insert a schedule </h1>
                <form class="flex self-center">
                    <label class="mr-4 text-xl">Movie:</label>
                            <select v-model="movieidinput" :="movieidProps" class="w-[380px] h-[35px] mr-12 px-8 border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                                <option :value=" movie.id" v-for="movie in movies.movie" :key="movie">{{movie.title}}</option>
                            </select >
                    <label class="mr-4 text-xl">Date:</label>
                    <input v-model="dateinput" :="dateProps" class="w-[220px] h-[35px] mr-12 px-8 border-none bg-gray-700 ring-[3px] ring-[rgb(0,137,208,0.5)] text-[20px] text-center rounded-[15px]" type="date"/>
                    <label class="mr-4 text-xl">Time:</label>
                    <input v-model="timeinput" :="timeProps" class="w-[220px] h-[35px] mr-12 px-8 border-none bg-gray-700 ring-[3px] ring-[rgb(0,137,208,0.5)] text-[20px] text-center rounded-[15px]" type="time"/>
                    <div @click="createSchedule" class="cursor-pointer content-center w-[80px] h-[35px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Add</div>
                </form>
                <div class="flex w-[1250px] self-center justify-between py-4">
                    <p class="text-green-500 self-start">
                        {{ schedule }}
                    </p>
                    <p class="text-red-800 mx-24">
                        {{ errors.movieidinput }}
                    </p>
                    <p class="text-red-800 mx-24">
                        {{ errors.dateinput }}
                    </p>
                    <p class="text-red-800 mx-24">
                        {{ errors.timeinput }}
                    </p>
                </div>
            </div>
    </div>
</template>

<script setup>
import { useForm } from 'vee-validate';
    definePageMeta({
            layout:""
        })
        function movierequired(value) {
            return value ? true : 'Movie is required;  ';
        }
        function daterequired(value) {
            return value ? true : 'Date is required;  ';
        }
        function timerequired(value) {
            return value ? true : 'Time is required;  ';
        }
    const { defineField, handleSubmit, errors } = useForm({
        validationSchema: {
            movieidinput: movierequired,
            dateinput: daterequired,
            timeinput: timerequired
        },
    });
    const [timeinput, timeProps] = defineField('timeinput')
    const [dateinput, dateProps] = defineField('dateinput')
    const [movieidinput, movieidProps] = defineField('movieidinput')

    const moviequery = gql`
    query MyQuery {
        movie {
            title
            id
        }
     }`
const { data:movies } = await useAsyncQuery(moviequery)
const scheduleinsertion = gql`
mutation MyMutation($date: date, $movie_id: Int, $time: time) {
  insert_schedule_one(object: {date: $date, movie_id: $movie_id, time: $time}) {
    movie_id
    date
    time
    movie {
      title
    }
  }
}
`
const createSchedule = handleSubmit(async () => {
    const {data} = await schedulemutate({date: dateinput.value, movie_id: movieidinput.value, time: timeinput.value})
    schedule.value = "SCHEDULE ADDED: " + '"' + data.insert_schedule_one.movie.title + '" is Scheduled for ' + data.insert_schedule_one.date + " " + data.insert_schedule_one.time
});
const schedule = ref("")
const{mutate: schedulemutate} = useMutation(scheduleinsertion)
</script>

<style scoped>

</style>