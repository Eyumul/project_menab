<template>
    <div class="bg-black text-white pt-64 pb-36">
        <div class="flex flex-wrap w-[100%] pt-[77px] px-10">
                <div v-for="movie in movies.movie" class="m-5"><MovieCardOne :movielink="movie.title" :movietitle=" movie.title "/></div>                
        </div>
    </div>
</template>

<script setup>
    definePageMeta({
            layout:""
        })
    const moviequery = gql`
    query myquery {
        movie {
            title
        genre
        director {
        name
        }
      }
    }`
    const { data:movies } = await useAsyncQuery(moviequery)
</script>

<style scoped>

</style>