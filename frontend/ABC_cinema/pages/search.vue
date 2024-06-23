<template>
    <div>
      <div class="flex flex-col pt-[100px] h-screen bg-black text-white">
        <div class="flex justify-center space-x-4">
          <input
            v-model="searchtext"
            type="text"
            placeholder="Search"
            class="w-[300px] h-[30px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]"
          />
          <div @click="searchmovie" class="cursor-pointer hover:text-[#0089D0]">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="3"
              stroke="currentColor"
              class="h-[30px] w-5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"
              />
            </svg>
          </div>
        </div>
        <div class="flex flex-wrap w-[100%] pt-[77px] px-10" >
          <div v-for="movie in moviematches" :key="movie.title">
            <MovieCardOne class="m-5" :movielink="movie.title" :moviethumbnail="movie.thumbnail" :movietitle=" movie.title "/>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  
  const searchtext = ref("")
  const moviematches = ref([])
  
  const moviequery = gql`
    query myquery {
      movie {
        title
        genre
        thumbnail
      }
    }
  `
  
  const { data: movies } = await useAsyncQuery(moviequery)
  
  const searchmovie = () => {
    moviematches.value = []
    if (searchtext.value == "" || searchtext.value == " "){
      alert("No movie inserted")
      return
    }
    for (let movie of movies.value.movie) {
      if (movie.title.toUpperCase().includes(searchtext.value.toUpperCase())) {
        moviematches.value.push({
          title: movie.title,
          thumbnail: movie.thumbnail
        })
      }
    }
    if (moviematches.value.length == 0){
      alert("No match found")
    }
  }
  </script>
  
  <style>
  
  </style>
  