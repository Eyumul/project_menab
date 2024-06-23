<template>
    <div class="h-full bg-black text-white pb-40">
        <div class="flex justify-between items-end px-[138px] pt-[100px]">
            <div class="flex space-x-8 justify-center pt-16 content end">
                <img src="/public/figmaImage/movie.png"/>
                <p class="s32 font-normal">Movies</p>
            </div>
            <div class="flex space-x-[34px] items-center">
                <p class="text-[24px] font-light">Browse movies by</p>
                <form>
                    <select v-model="browsetype" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] focus:ring-[rgb(0,137,208)] text-[24px] text-center  rounded-[20px]">
                        <option>No categories</option>
                        <option>Genre</option>
                        <option>Directors</option>
                    </select>
                </form>
            </div>
        </div>
        <div v-if = "browsetype == 'No categories'">
            <div class="flex flex-wrap w-[100%] pt-[77px] px-10">
                <div v-for="movie in movies.movie" :key="movie.title" class="m-5"><MovieCardOne :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/></div>            
            </div>
        </div>
        <div v-else-if = "browsetype == 'Genre'">
            <div v-for="genre in genres"  class="flex flex-col" :key="genre">
                <h1 class="text-[#0089D0] text-3xl px-5 pt-10">{{genre}}</h1>
                <div class="flex flex-wrap w-[100%] pb-10 px-10">
                    <div v-for="movie in movies.movie" :key="movie.title" class="">
                        <div>
                            <MovieCardOne class="m-5" v-if="movie.genre == genre" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
                        </div>
                    </div>            
                </div>
            </div>
        </div>
        <div v-else-if = "browsetype == 'Directors'">
            <div v-for="director in directors.director" :key="director.name" class="flex flex-col ">
                <h1 class="text-[#0089D0] text-3xl px-5 pt-10">{{director.name}}</h1>
                <div class="flex flex-wrap w-[100%] pb-10 px-10">
                    <div v-for="movie in movies.movie" :key="movie.title" class="">
                        <div>
                            <MovieCardOne class="m-5" v-if="movie.director.name == director.name" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
                        </div>
                    </div>            
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
    const browsetype = ref("No categories")
    const genres = ["Scifi","Action","Drama","Horror","Comedy","Adventure","Thriller"]
    const moviequery = gql`
    query myquery {
        movie {
            title
        genre
        thumbnail
        director {
        name
        }
      }
    }`
    const directorquery = gql`
    query MyQuery {
        director {
            name
        }
    }`
    const {data:directors} = await useAsyncQuery(directorquery)
    const { data:movies } = await useAsyncQuery(moviequery)
</script>

<style>

</style>