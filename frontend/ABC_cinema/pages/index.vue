<template>
    <div class="h-full bg-black text-white pb-40">
        <div class="flex space-x-8 justify-center pt-16">
            <div>
                <img src="/public/figmaImage/fire.png" />
            </div>
            <div class="flex flex-col items-center">
                <p class="s32 font-normal">Trending movies</p>
                <p class="s24">On cinema</p>
            </div>
        </div>
        <div v-for = "trendings in trendingmovie.movie" >
            <div v-if="trendings.trending">
                <trending :movietitle="trendings.title" :moviethumbnail="trendings.thumbnail" :directorname="trendings.director.name" :description="trendings.description" :genre="trendings.genre" :duration="trendings.duration" :star-one="trendings.movie_stars[0].star.name" :star-two="trendings.movie_stars[1].star.name" :star-three="trendings.movie_stars[2].star.name" :star-four="trendings.movie_stars[3].star.name" :star-five="trendings.movie_stars[4].star.name" :rate="trendings.ratings_aggregate.aggregate.avg.rating"  />
            </div>
        </div>

        <div class="flex space-x-8 justify-center pt-16 content end">
            <img src="/public/figmaImage/dashboard.png"/>
            <p id="here" class="s32 font-normal">My Dashboard</p>
        </div>
        <div class="flex flex-col mx-[138px] py-20 space-y-[125px]">
            <div class="text-center text-xs text-gray-400">
                Sign up to see your movies tickets and movies
            </div>
        </div>
    </div>
</template>

<script setup>
const query = gql`
query myquery {
  movie {
    title
    thumbnail
    director {
      name
      id
    }
    description
    genre
    duration
    schedules {
      showtime
      id
    }
    movie_stars {
      star {
        name
        id
      }
    }
    trending
    ratings_aggregate {
      aggregate {
        avg {
          rating
        }
      }
    }
  }
}
`
const { data:trendingmovie } = await useAsyncQuery(query)
definePageMeta({
        layout:"home"
    })
    // const trendings = []
    // const test = []
    // setTimeout(() => {
        
    //     for (let i=0; i < 6; i++){
    //         // if (trendingmovie.movie[i].trending) {
    //         //     trendings.push(trendingmovie.movie[])
    //         // }
    //         // if (true){
    //         //     trendings.push(i)
    //         // }
    //         test.push(trendingmovie.movie[0].trending)
    //     }
    // }, 500);
    
</script>

<style>
.line {
    height: 500px;
    width: 3px;
    border-radius:3px;
    background-color: rgb(0,137,208,50%);
}
.s32 {
    font-size:32px;
}
.s24 {
    font-size:24px;
    color:#0089D0;
}
.s16 {
    font-size:16px;
}
.textColor{
    color:#0089D0;
}
.bgColor{
    background-color:#0089D0;
}
.button{
    background-color:#0089D0;
    height: 60px;
    width: 276px;
    border-radius: 10px;
}
.movieDetails{
    width:582px;
    height:fit-content;
}
.movieTitle {
    width:450px;
    font-size:32px;
    text-align:center;
    text-transform: uppercase;
    font-weight:900;
    background: -webkit-linear-gradient(right, #24B4FF, #90D9FF);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
.movieDirector {
    width:582px;
    font-size:16px;
    font-weight: 900;
}
.directorName {
    color:#0089D0;
}
.movieDescription{
    width: 582px;
    text-indent: 30px;
    font-size:16px;
    font-weight:300;
}
.movieGenre {
    width: 582px;
    text-align: right;
    font-size:16px;
    font-weight:700;
}
.genreName {
    color:#0089D0;
}

</style>