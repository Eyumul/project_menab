<template>
    <div class="flex flex-col space-y-24 items-center bg-black text-white pt-64 pb-36">
        <div class="text-whie w-[600px] text-xl space-y-12 p-5">
            <Directorform />
            <Starform/>
        </div>
        <form class="text-white w-[600px] text-xl">
            <fieldset  class="flex flex-col space-y-4 bg-black border-2 border-[rgb(0,137,208,0.5)] rounded-[20px] p-5">
                <legend> Movie </legend>
                <div class="flex justify-between items-center">
                    <label>Title: </label>
                    <input v-model="movietitle" :="movietitleProps" type="text" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
                </div>
                <div class="flex justify-between items-center">
                    <label>Genre:</label>
                    <select v-model="moviegenre" :="moviegenreProps" class = "w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option>Scifi</option>
                        <option>Action</option>
                        <option>Drama</option>
                        <option>Horror</option>
                        <option>Comedy</option>
                        <option>Adventure</option>
                        <option>Thriller</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Duration:</label>
                    <input v-model="movieduration" :="moviedurationProps" type="text" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
                </div>
                <div class="flex justify-between items-center">
                    <label>Director:</label>
                    <select v-model="moviedirector" :="moviedirectorProps" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" director.id" v-for="director in directors.director">{{director.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star one:</label>
                    <select v-model="moviestarone" :="moviestaroneProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star two:</label>
                    <select v-model="moviestartwo" :="moviestartwoProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star three:</label>
                    <select v-model="moviestarthree" :="moviestarthreeProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star four:</label>
                    <select v-model="moviestarfour" :="moviestarfourProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star five:</label>
                    <select v-model="moviestarfive" :="moviestarfiveProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Description:</label>
                    <textarea v-model="moviedescription" :="moviedescriptionProps" rows="10" cols="40" class="w-[380px] p-2 border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[10px] text-left rounded-[20px]"></textarea>
                </div>
                <div class="flex justify-between items-center">
                    <label>Image</label>
                    <input type="file" class="w-[380px] h-[50px] text-[rgb(0,137,208)] bg-black text-[24px] p-2" />
                </div>
                <div class="flex space-x-12 items-center justify-end pr-32">
                    <input v-model="istrending" type="checkbox" class="border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)]" />
                    <label>This movie is trending</label>
                </div>
                <div @click="createMovie" class="cursor-pointer content-center w-full h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Add</div>
            </fieldset>
        </form>
        <div>
            <p class="text-sm text-gray-500">{{ movtitle }}</p>
            <p class="text-red-800 text-sm">{{errors.movietitle  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviegenre  }}</p>
            <p class="text-red-800 text-sm">{{errors.movieduration  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviedirector  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviestarone  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviestartwo  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviestarthree  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviestarfour  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviestarfive  }}</p>
            <p class="text-red-800 text-sm">{{errors.moviedescription  }}</p>
        </div>
    </div>
</template>
<script setup>
definePageMeta({
    layout:"admin"
})



import { useForm } from 'vee-validate';
import { handleError } from 'vue';
import * as Yup from 'yup';


function titlerequired(value) {
  return value ? true : 'Movie title is required;  ';
}
function genrerequired(value) {
  return value ? true : 'Genre is required;  ';
}
function durationrequired(value) {
    return value ? true : 'Duration is required;  ';
}
function directorrequired(value) {
  return value ? true : 'Director is required;  ';
}
function staronerequired(value) {
  return value ? true : 'Star one is required;  ';
}
function startworequired(value) {
  return value ? true : 'Star two is required;  ';
}
function starthreerequired(value) {
  return value ? true : 'Star three is required;  ';
}
function starfourrequired(value) {
  return value ? true : 'Star four is required;  ';
}
function starfiverequired(value) {
    return value ? true : 'Star five is required;  ';
}
function descriptionRequired(value) {
  if (!value) {
    return 'description is required';
  }  
  if (value.length < 20) {
    return 'description is too short';
  }
    return true;
}  

//create form
const { defineField, handleSubmit, errors } = useForm({
  validationSchema: {
    movietitle: titlerequired,
    moviegenre: genrerequired,
    movieduration: durationrequired,
    moviedirector: directorrequired,
    moviestarone: staronerequired,
    moviestartwo: startworequired,
    moviestarthree: starthreerequired,
    moviestarfour: starfourrequired,
    moviestarfive: starfiverequired,
    moviedescription: descriptionRequired
  },
});

const dirquery = gql`
query MyQuery {
    director {
        name
        id
    }
}`
const { data:directors } = await useAsyncQuery(dirquery)

const strquery = gql`
query MyQuery {
    star {
        name
        id
    }
}`
const { data:stars } = await useAsyncQuery(strquery)

// const [starname, starnameProps] = defineField('starname')
const [movietitle, movietitleProps] = defineField('movietitle')
const [moviegenre, moviegenreProps] = defineField('moviegenre')
const [movieduration, moviedurationProps] = defineField('movieduration')
const [moviedirector, moviedirectorProps] = defineField('moviedirector')
const [moviestarone, moviestaroneProps] = defineField('moviestarone')
const [moviestartwo, moviestartwoProps] = defineField('moviestartwo')
const [moviestarthree, moviestarthreeProps] = defineField('moviestarthree')
const [moviestarfour, moviestarfourProps] = defineField('moviestarfour')
const [moviestarfive, moviestarfiveProps] = defineField('moviestarfive')
const [moviedescription,moviedescriptionProps] = defineField('moviedescription')
const istrending = ref(false)

// const strname = ref("")
const movtitle = ref("")


const movinsertion = gql`
mutation MyMutation($title: String , $genre: String, $duration: String, $director_id: Int, $description: String, $thumbnail: String, $trending: Boolean) {
  insert_movie_one(object: {title: $title, genre: $genre, duration: $duration, director_id: $director_id, description: $description, thumbnail: $thumbnail, trending: $trending}) {
    title
    id
  }
}
`
const movstrinsertion = gql`
mutation MyMutation($movie_id: Int, $star_id: Int) {
  insert_movie_star_one(object: {movie_id: $movie_id, star_id: $star_id}) {
    movie_id
    star_id
  }
}
`


const createMovie = handleSubmit(async() => {
    const{ data: movdata }= await movmutate({title: movietitle.value, genre: moviegenre.value, duration: movieduration.value, director_id: moviedirector.value, description: moviedescription.value, trending: istrending.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarone.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestartwo.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarthree.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarfour.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarfive.value})
    movtitle.value = movdata.insert_movie_one.title + " is added as a movie"
    // console.log(movtitle.value)
});


const{mutate: movmutate} = useMutation(movinsertion)
const{mutate: movstrmutate} = useMutation(movstrinsertion)
</script>

<style>

</style>
