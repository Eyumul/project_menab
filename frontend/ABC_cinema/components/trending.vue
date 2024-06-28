<template>
    <div class="flex items-center px-[138px] py-20">
                <div class="h-[734px] w-[828px] rounded-[50px] overflow-hidden" id = "slideshoow3">
                    <div class="mt-[330px]">
                        <ul class="flex justify-between px-8">
                            <li @click="previous">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 cursor-pointer hover:text-[#0089D0]">
                                    <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06l-7.5-7.5Z" clip-rule="evenodd" />
                                </svg>
                            </li>
                            <li @click="next">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 cursor-pointer hover:text-[#0089D0]">
                                    <path fill-rule="evenodd" d="M16.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L14.69 12 7.72 5.03a.75.75 0 0 1 1.06-1.06l7.5 7.5Z" clip-rule="evenodd" />
                                </svg>
                            </li>
                        </ul>

                        <!-- <div class="flex gap-x-8 mt-[300px] pb-2 justify-center items-center">
                            <span class="dot" id="dot1"></span>
                            <span class="dot" id="dot2"></span>
                            <span class="dot" id="dot3"></span>
                            <span class="dot" id="dot4"></span>
                            <span class="dot" id="dotActive"></span>
                        </div> -->
                    </div>
                </div>
            <div class="line ml-12"></div>
            <div class="movieDetails flex flex-col ml-28 space-y-6 items-center">
                <div class="flex flex-col space-y-6">
                    <p class="movieTitle self-center ">{{ movietitle }}</p>
                    <p class="movieDirector">Directed by <span class = "directorName">{{directorname}}</span></p>
                    <p class="movieDescription">{{ description }}</p>
                    <p class="movieGenre">Genre: <span class = "genreName">{{genre}}</span></p>
                </div>
                <div class="flex justify-between font-medium w-full">
                    <div class="flex  self-start space-x-1.5 items-center">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-16">
                            <path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25ZM12.75 6a.75.75 0 0 0-1.5 0v6c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-3.75V6Z" clip-rule="evenodd" />
                        </svg>
                        <p class="textColor ">{{(duration)}}</p>
                    </div>
                    <div class="flex items-start space-x-1.5">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-16">
                            <path d="M12.75 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM7.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM8.25 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM9.75 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM10.5 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM12.75 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM14.25 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 13.5a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" />
                            <path fill-rule="evenodd" d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75Zm13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5Z" clip-rule="evenodd" />
                        </svg>
                        <div>
                            <p class="textColor">Date</p>
                            <ul class="list-disc list-inside indent-3">
                                <li>3:00 PM</li>
                                <li>6:00 PM</li>
                                <li>9:00 PM</li>
                            </ul>
                        </div>
                    </div>
                    <div class="">
                        <p>Starring</p>
                        <ul class="list-disc list-inside indent-3 textColor">
                            <li>{{starOne}}</li>
                            <li>{{starTwo}}</li>
                            <li>{{starThree}}</li>
                            <li>{{starFour}}</li>
                            <li>{{starFive}}</li>
                        </ul>
                    </div>
                </div>
                <p v-if="rate" class="w-full font-black text-[14px]">ABC_cinema Rating <span class="textColor ml-4 text-[24px]">{{rate.toFixed(1)}}/5</span></p>
                <div v-if="auth0?.isAuthenticated.value" class="w-full">
                    <div v-if="role == 'user'" class="flex justify-between mt-4">
                        <div  class="flex">
                            <label>Your rating: </label><input v-model="rating" @change="rateMovie" class=" border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[18px] text-center rounded-[10px] ml-4" type="number" min="1" max="5"/>
                        </div>
                        <svg v-if="!isSaved" @click="save" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 cursor-pointer hover:text-[#0089D0]">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z" />
                        </svg>
                        <svg v-else-if="isSaved" @click="remove" xmlns="http://www.w3.org/2000/svg" fill="#0089D0" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 cursor-pointer text-[#0089D0]">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z" />
                        </svg>
                    </div>
                </div>
                <div v-if="auth0?.isAuthenticated.value && role == 'user'"  class="flex space-x-7 hover:bg-[#016699] cursor-pointer button justify-center items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 text-black">
                        <path fill-rule="evenodd" d="M1.5 6.375c0-1.036.84-1.875 1.875-1.875h17.25c1.035 0 1.875.84 1.875 1.875v3.026a.75.75 0 0 1-.375.65 2.249 2.249 0 0 0 0 3.898.75.75 0 0 1 .375.65v3.026c0 1.035-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 17.625v-3.026a.75.75 0 0 1 .374-.65 2.249 2.249 0 0 0 0-3.898.75.75 0 0 1-.374-.65V6.375Zm15-1.125a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0V6a.75.75 0 0 1 .75-.75Zm.75 4.5a.75.75 0 0 0-1.5 0v.75a.75.75 0 0 0 1.5 0v-.75Zm-.75 3a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0v-.75a.75.75 0 0 1 .75-.75Zm.75 4.5a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0v-.75ZM6 12a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H6.75A.75.75 0 0 1 6 12Zm.75 2.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5h-3Z" clip-rule="evenodd" />
                    </svg>
                    <div class="text-2xl font-black">Buy Ticket</div>
                </div>
                <p v-if="auth0?.isAuthenticated.value"   class="hidden">{{ movieId = movieid }}</p>
            </div>
        </div>
</template>

<script setup>
import { useAuth0 } from '@auth0/auth0-vue';
import {jwtDecode} from 'jwt-decode';
const auth0 = process.client ? useAuth0() : undefined
const hasuraId = ref('')
const role = ref('')
if(process.client){
    const tokentext = localStorage.getItem('hasura-token')
        if (tokentext) {
        const decodedToken = jwtDecode(tokentext);
        role.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-default-role'];
        hasuraId.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-user-id']
        }
}
    
        const movieId = ref('')
        const saveinsertion = gql`
        mutation MyMutation($movie_id: Int!, $user_id: Int!) {
          insert_saved_movie_one(object: {movie_id: $movie_id, user_id: $user_id}) {
            movie {
              title
            }
            profile {
              name
            }
          }
        }`
        const{mutate: savemutate} = useMutation(saveinsertion)
        const deleteinsertion = gql`
        mutation MyMutation($_eq: Int) {
          delete_saved_movie(where: {movie_id: {_eq: $_eq}}) {
            returning {
              movie_id
              movie {
                    title
                }
            }
          }
        }
        `
        const { mutate: removemutate } = useMutation(deleteinsertion)
        
            const id = ref(null)
            const isSaved = ref(null)
            let movies = []
            setTimeout( async () => {
                if (process.client && auth0?.isAuthenticated.value){
                    const profilequery = gql`
                        query MyQuery {
                            profile {
                                id
                                saved_movies {
                                    movie_id
                                }
                            }
                        }`
                    const {data, error} = await useAsyncQuery(profilequery)
                    if (error.value && role.value == 'user') location.reload()
                    id.value = data.value.profile[0].id
                    
                    for (const movie of data.value.profile[0].saved_movies){
                        movies.push(movie.movie_id)
                    }
                    isSaved.value = movies.includes(movieId.value)
                }}, 1000);
                
        async function save(){
            // Saved_movie insertion code goes here
            const{ data }= await savemutate({movie_id: movieId.value,user_id: id.value})
            console.log("You Bookmarked a movie: ",data.insert_saved_movie_one.movie.title)
            isSaved.value = true
        }
        async function remove(){
            // Saved_movie deletion code goes here
            const { data } = await removemutate({ _eq: movieId.value })
            console.log("You unsaved movie: ",data.delete_saved_movie.returning[0].movie.title)
            isSaved.value = false
        }
                
        const rateinsertion = gql`
        mutation MyMutation($movie_id: Int!, $rating: Int!, $user_id: Int!) {
          insert_rating_one(object: {movie_id: $movie_id, rating: $rating, user_id: $user_id}) {
            movie_id
            rating
            user_id
            movie {
                title
            }
          }
        }
        `
        const{mutate: ratemutate} = useMutation(rateinsertion)
        const updaterating = gql`
        mutation MyMutation($rating: Int!, $_eq: Int!) {
            update_rating(where: {movie_id: {_eq: $_eq}}, _set: {rating: $rating}) {
                returning {
                movie {
                    title
                }
                rating
                }
            }
        }`
        const { mutate: updateratemutate } = useMutation(updaterating)
        const rating = ref(null)
        const currentRate = ref(null)
        setTimeout( async () => {
            if (process.client && auth0?.isAuthenticated.value){
            const RATING_QUERY = gql`
                query MyQuery{
                    rating {
                        rating
                        movie_id
                    }
                }`
            const { data, error} = await useAsyncQuery(RATING_QUERY);
            if (error.value) location.reload()
            for (const rate of data.value.rating){
                if(rate.movie_id == movieId.value) {
                    rating.value = rate.rating
                    currentRate.value = rate.rating
                }
            }
            }}, 1000);
        async function rateMovie (){
            if(currentRate.value == null){
                // rating insertion goes here
                const{ data }= await ratemutate({movie_id: movieId.value,user_id: id.value, rating: rating.value})
                console.log("(New) You rated '",data.insert_rating_one.movie.title,"': ",data.insert_rating_one.rating,"/5")
                currentRate.value = rating.value
            }else{
                // update insertion goes here
                const { data } = await updateratemutate({ _eq: movieId.value, rating: rating.value })
                console.log("You rated '",data.update_rating.returning[0].movie.title,"': ",data.update_rating.returning[0].rating,"/5")
            }
        }

     

   






const props = defineProps(["movieid","movietitle", "moviethumbnail", "bgone", "bgtwo", "bgthree", "directorname", "description", "genre", "duration" , "date", "starOne", "starTwo", "starThree", "starFour", "starFive", "rate"])
let imagenumber = 3
function previous(){
    let currentimage = document.getElementById( `slideshoow${imagenumber}`)
    imagenumber--
    if (imagenumber < 1) {
        imagenumber = 3
        }
    currentimage.id = `slideshoow${imagenumber}`
}
function next(){
    let currentimage = document.getElementById(`slideshoow${imagenumber}`)
    imagenumber++
    if (imagenumber > 3){
        imagenumber = 1
    }
    currentimage.id = `slideshoow${imagenumber}`
}
</script>

<style scoped>
#slideshoow1 {
    background-image: v-bind(bgone);
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshoow2 {
    background-image: v-bind(bgtwo);
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshoow3 {
    background-image: v-bind(bgthree);
    background-repeat: no-repeat;
    background-size:cover;
}
</style>