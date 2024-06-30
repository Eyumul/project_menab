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
        <div v-for = "trendings in trendingmovie.movie" :key = "trendings.title" >
            <div v-if="trendings.trending">
                <trending :movieid="trendings.id" :movietitle="trendings.title" :moviethumbnail="trendings.thumbnail" :bgone="prefix + trendings.featured_images[0].image_one + sufix" :bgtwo="prefix + trendings.featured_images[0].image_two + sufix" :bgthree="prefix + trendings.featured_images[0].image_three + sufix" :directorname="trendings.director.name" :description="trendings.description" :genre="trendings.genre" :duration="trendings.duration" :star-one="trendings.movie_stars[0].star.name" :star-two="trendings.movie_stars[1].star.name" :star-three="trendings.movie_stars[2].star.name" :star-four="trendings.movie_stars[3].star.name" :star-five="trendings.movie_stars[4].star.name" :rate="trendings.ratings_aggregate.aggregate.avg.rating"  />
            </div>
        </div>
        <div v-if="isAuthenticated && role == 'user'">
            <div class="flex space-x-8 justify-center pt-16 content end">
                <img src="/public/figmaImage/dashboard.png"/>
                <p id="here" class="s32 font-normal">My Dashboard</p>
            </div>
            <div class="flex flex-col mx-[138px] py-10 space-y-[125px]">
                <div class="capitalize text-center text-sm text-yellow-400">
                    {{welcomemessage}}
                </div>
            </div>
            <div>
                <h1 class="ml-5 px-10 underline text-[#0089D0] text-3xl">Saved movies</h1>
                <div  class="flex flex-wrap w-[100%] pt-[77px] px-10">
                    <div class="m-5" v-for = "movies in savedmovieresult" :key="movies.movie.title">
                        <MovieCardOne :movielink="movies.movie.title" :moviethumbnail="movies.movie.thumbnail" :movietitle=" movies.movie.title "/>
                    </div>
                </div>
            </div>
            <div>
                <h1 class="ml-5 px-10 mt-36 underline text-[#0089D0] text-3xl">Booked Schedules</h1>
                <div  class="flex flex-wrap w-[100%] pt-[77px]">
                    <div v-for="ticket in tickets" :key="ticket.tx_ref">
                        <div class="flex w-[600px] justify-around mx-12 py-4 rounded-[20px] border-[rgb(0,137,208,0.2)] border-2 bg-[#000e14] ">
                            <div class="self-center">
                                <h2 class="self-center text-lg w-[75px] text-[#0089D0]">{{ formatDateshort(ticket.schedule.date) }}</h2>
                                <h2 class="self-center text-lg w-[75px] text-[#0089D0]">{{ convertTo12HourFormat(ticket.schedule.time) }}</h2>
                            </div>
                            <NuxtImg :src="ticket.schedule.movie.thumbnail" class="w-[200px] h-[150px] rounded-[12px]"/>
                            <h3 class="self-center w-[200px] text-xl text-[#24B4FF] font-black"><NuxtLink :to="'./tickets/'+ticket.tx_ref">{{ ticket.schedule.movie.title }}</NuxtLink></h3>
                        </div>
                    </div>
                </div>
                
                    
                
            </div>
        </div>
    </div>
</template>

<script setup>
import { useAuth0 } from '@auth0/auth0-vue';
import {jwtDecode} from 'jwt-decode';
const auth0 = process.client ? useAuth0() : undefined
const welcomemessage = ref("")
const savedmovieresult = ref([])
const tickets = ref([])
const login = () => {
    auth0?.loginWithRedirect()
}
const isAuthenticated = computed(() => {
  return auth0?.isAuthenticated.value
})
const username = computed(() => {
    return auth0?.user.value
})
getHasuraToken()
const hasuraId = ref('')
const role = ref('')
setTimeout( async () => {
    if(process.client){
        const tokentext = localStorage.getItem('hasura-token')
            if (tokentext) {
            const decodedToken = jwtDecode(tokentext);
            role.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-default-role'];
            hasuraId.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-user-id']
            }
}}, 1000);
async function syncUsers() {
    setTimeout( async () => {
      // mutation for adding user or updating a user if it exist
      const addUsers = gql `
        mutation MyMutation($email: String, $name: String, $phone_number: Int) {
            insert_profile_one(object: {email: $email, name: $name, phone_number: $phone_number}, on_conflict: {constraint: profile_user_id_key, update_columns: [name, email]}) {
                email
                name
                phone_number
                username
            }
        }`
      const { mutate:addUser } = useMutation(addUsers)
      const userData = {
        name: auth0?.user.value.name,
        phone_number: auth0?.user.value.phone ? auth0?.user.value.phone: null,
        email: auth0?.user.value.email,
      }
      const { data } = await addUser(userData)
      welcomemessage.value = "Welcome to your dashboard " + data.insert_profile_one.username + " !!!"
      console.log(welcomemessage.value)
      savedmovieresult.value = (await savedmoviefunction()).value.saved_movie
      tickets.value = (await ticketfunction()).value.ticket 
    }, 1000);
}
if(process.client){
    if(auth0?.isAuthenticated.value) {
        syncUsers()
    }

}


const trendingquery = gql`
query myquery {
  movie {
    id
    title
    thumbnail
    director {
      name
      id
    }
    description
    genre
    duration
    movie_stars {
      star {
        name
        id
      }
    }
    featured_images {
      image_one
      image_three
      image_two
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
const { data:trendingmovie } = await useAsyncQuery(trendingquery)

async function savedmoviefunction (){
    const savedquery = gql`
    query MyQuery {
      saved_movie {
        movie {
          title
          thumbnail
        }
      }
    }
    `
    const {data:savedMovies} = await useAsyncQuery(savedquery)
    return savedMovies
}
async function ticketfunction (){
    const ticketquery = gql`
    query MyQuery {
  ticket {
    tx_ref
    schedule {
      date
      time
      movie {
        title
        thumbnail
      }
    }
  }
}
    `
    const {data:movietickets} = await useAsyncQuery(ticketquery)
    return movietickets
}

const prefix = "url('"
const sufix = "')"
function convertTo12HourFormat(time) {
  let [hours, minutes, seconds] = time.split(':').map(Number);
  const period = hours >= 12 ? 'PM' : 'AM';
  hours = hours % 12 || 12;
  hours = String(hours).padStart(2, '0');
  minutes = String(minutes).padStart(2, '0');
  seconds = String(seconds).padStart(2, '0');
  return `${hours}:${minutes} ${period}`;
}
function formatDateshort(inputDate) {
  const date = new Date(inputDate);
  const days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  const months = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
  ];
  const dayName = days[date.getUTCDay()]; // Get the day of the week
  const monthName = months[date.getUTCMonth()]; // Get the month
  const day = date.getUTCDate(); // Get the day of the month
  const year = date.getUTCFullYear(); // Get the year
  return `${dayName} ${day} ${monthName}`;
}
// const featuredone = "url('"+ trendingmovie.value.movie[0].featured_images[0].image_one + "')"
// const featuredtwo = "url('" + trendingmovie.value.movie[0].featured_images[0].image_two + "')"
// const featuredthree = "url('" + trendingmovie.value.movie[0].featured_images[0].image_three + "')"
definePageMeta({
        layout:"home"
    })
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