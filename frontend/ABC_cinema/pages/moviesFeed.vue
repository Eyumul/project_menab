<template>
    <div class="h-screen bg-black text-white pb-40">
        <div class="bg-black pb-[50px]">
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
                <div  class="flex flex-wrap w-[100%] pt-[77px] px-10">
                    <div v-for="(movie, index) in visibleItems" :key="index" class="m-5"><MovieCardOne :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/></div>            
                    <button class="self-center bg-[#0089D0] flex w-40 ml-6 mt-6 pl-4 justify-center items-center text-xl rounded-2xl h-12 font-black cursor-pointer" v-if="visibleItems.length < items.length" @click="loadMore">
                        SEE MORE
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10">
                            <path fill-rule="evenodd" d="M16.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L14.69 12 7.72 5.03a.75.75 0 0 1 1.06-1.06l7.5 7.5Z" clip-rule="evenodd" />
                        </svg>
                    </button>
                </div>
            </div>
            <div v-else-if = "browsetype == 'Genre'">
                <div v-for="genre in genres"  class="flex flex-col" :key="genre">
                    <h1 class="text-[#0089D0] text-3xl px-16 pt-14 underline">{{genre}}</h1>
                    <div class="flex flex-wrap w-[100%] pb-10 px-10">
                        <div v-for="movie in movies.movie" :key="movie.title">
                            
                                <MovieCardOne class="m-5" v-if="movie.genre == genre" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
                            
                        </div>            
                        <div  class="w-full bg-gray-700 h-[1px]"></div>
                    </div>
                </div>
            </div>
            <div v-else-if = "browsetype == 'Directors'">
                <div class="flex pl-[138px] pt-10 space-x-4">
                    <div @click="searchdirector" class="cursor-pointer hover:text-[#0089D0]">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-5">
                            <path fill-rule="evenodd" d="M3.792 2.938A49.069 49.069 0 0 1 12 2.25c2.797 0 5.54.236 8.209.688a1.857 1.857 0 0 1 1.541 1.836v1.044a3 3 0 0 1-.879 2.121l-6.182 6.182a1.5 1.5 0 0 0-.439 1.061v2.927a3 3 0 0 1-1.658 2.684l-1.757.878A.75.75 0 0 1 9.75 21v-5.818a1.5 1.5 0 0 0-.44-1.06L3.13 7.938a3 3 0 0 1-.879-2.121V4.774c0-.897.64-1.683 1.542-1.836Z" clip-rule="evenodd" />
                        </svg>
                    </div>
                    <input v-model="searchtext" @keyup.enter="searchdirector" type="text" placeholder="Filter by director" class="w-[150px] h-[20px] border-none bg-black ring-[2px] ring-[rgb(0,137,208,0.5)] text-[18px] text-center rounded-[5px]"/>
                    <div @click="clearinput" class="cursor-pointer hover:text-[#0089D0] underline" v-if="searchtext != ''">Clear</div>
                </div>
    
    
                <div v-for="director in directormatches" :key="director.name" class="flex flex-col ">
                    <div v-if="director.movies.length != 0">
                        <h1 class="text-[#0089D0] text-3xl px-16 pt-14 underline">{{director.name}}</h1>
                        <div class="flex flex-wrap w-[100%] pb-10 px-10">
                            <div v-for="movie in movies.movie" :key="movie.title">
                                    <MovieCardOne class="m-5" v-if="movie.director.name == director.name" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
                            </div>            
                            <div class="w-full bg-gray-700 h-[1px]"></div>
                        </div>
                    </div>
                </div>
    
                <div v-if="searchtext == ''">
                    <div v-for="director in directors.director" :key="director.name" class="flex flex-col ">
                        <div v-if="director.movies.length != 0">
                            <h1 class="text-[#0089D0] text-3xl px-16 pt-14 underline">{{director.name}}</h1>
                            <div class="flex flex-wrap w-[100%] pb-10 px-10">
                                <div v-for="movie in movies.movie" :key="movie.title">
                                        <MovieCardOne class="m-5" v-if="movie.director.name == director.name" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
                                </div>            
                                <div  class="w-full bg-gray-700 h-[1px]"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
    import { ref, onMounted } from 'vue';
    const browsetype = ref("No categories")
    const items = ref([]);         // Full list of items
    const visibleItems = ref([]);  // Items currently displayed
    const itemsToShow = ref(4);    // Number of items to show initially and each time load more is clicked  
    const searchtext = ref("")
    const directormatches = ref([])
    
    const clearinput = () => {
        searchtext.value = ''
        directormatches.value = []
    }
    const searchdirector = () => {
    directormatches.value = []
    if (searchtext.value == "" || searchtext.value == " "){
      alert("No director inserted")
      return
    }
    for (let director of directors.value.director) {
      if (director.name.toUpperCase().includes(searchtext.value.toUpperCase())) {
        directormatches.value.push({
            __typename: "director",
            name: director.name,
            movies: director.movies
        })
      }
    }
    if (directormatches.value.length == 0){
      alert('Director "' + searchtext.value + '" Not found')
    }
  }
    
    
    
    const loadMore = () => {
        const newLength = visibleItems.value.length + itemsToShow.value;
        visibleItems.value = items.value.slice(0, newLength);
    };

    onMounted(async () => {
        items.value = movies.value.movie;
        visibleItems.value = items.value.slice(0, itemsToShow.value);
    });


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
            movies {
                title
            }
        }
    }`
    const {data:directors} = await useAsyncQuery(directorquery)
    const { data:movies } = await useAsyncQuery(moviequery)

    const avaliablegenres = async () => {
        const genrelist = []
        for(const movie of movies.value.movie){
            if (genrelist.includes(movie.genre)){
                continue
            }
            genrelist.push(movie.genre)
        }
        return genrelist;
    };
    const genres = await avaliablegenres()
</script>

<style>

</style>