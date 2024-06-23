<template>
    <div class="flex flex-col space-y-24 items-center bg-black text-white">
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
                        <option :value=" director.id" v-for="director in directors.director" :key="director">{{director.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star one:</label>
                    <select v-model="moviestarone" :="moviestaroneProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star" :key="star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star two:</label>
                    <select v-model="moviestartwo" :="moviestartwoProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star" :key="star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star three:</label>
                    <select v-model="moviestarthree" :="moviestarthreeProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star" :key="star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star four:</label>
                    <select v-model="moviestarfour" :="moviestarfourProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star" :key="star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Movie star five:</label>
                    <select v-model="moviestarfive" :="moviestarfiveProps" class = "w-[380px] h-[35px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" star.id" v-for="star in stars.star" :key="star">{{star.name}}</option>
                    </select>
                </div>
                <div class="flex justify-between items-center">
                    <label>Description:</label>
                    <textarea v-model="moviedescription" :="moviedescriptionProps" rows="10" cols="40" class="w-[380px] p-2 border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[10px] text-left rounded-[20px]"></textarea>
                </div>
                <div class="flex justify-between items-center">
                    <label>Thumbnail</label>
                    <input @change="onThumbnailChange" type="file" class="w-[380px] h-[50px] text-[rgb(0,137,208)] bg-black text-[24px] p-2" />
                </div>
                <div class="flex justify-between items-center">
                    <label>Featured images</label>
                    <input multiple @change="onFeaturedImagesChange" type="file" class="w-[380px] h-[50px] text-[rgb(0,137,208)] bg-black text-[24px] p-2" />
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
            <p class="text-red-800 text-sm">{{errors.thumbnail }}</p>
            <p class="text-red-800 text-sm">{{errors.featuredImageOne }}</p>
            <p class="text-red-800 text-sm">{{errors.featuredImageTwo }}</p>
            <p class="text-red-800 text-sm">{{errors.featuredImageThree }}</p>
        </div>
    </div>
</template>
<script setup>



import { useForm } from 'vee-validate';
import { handleError } from 'vue';
import * as Yup from 'yup';
import { useFetch } from '@vueuse/core'
  
const imageUrl = ref('')
const featuredImageOneUrl = ref('')
const featuredImageTwoUrl = ref('')
const featuredImageThreeUrl = ref('')
const featuredImage = ref(null)

function onThumbnailChange(event) {
    thumbnail.value = event.target.files[0]
}
function onFeaturedImagesChange(event) {
    featuredImage.value = event.target.files
    featuredImageOne.value = event.target.files[0]
    featuredImageTwo.value = event.target.files[1]
    featuredImageThree.value = event.target.files[2]
}
    
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
function validimageRequired(value) {
    if (!value){
        return 'image not selected'
    }
    if (value.type.slice(0,5) != "image"){
        return 'selected file is not an image'
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
    moviedescription: descriptionRequired,
    thumbnail: validimageRequired,
    featuredImageOne: validimageRequired,
    featuredImageTwo: validimageRequired,
    featuredImageThree: validimageRequired
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
const [thumbnail, thumbnailProps] = defineField('thumbnail')
const [featuredImageOne, featuredImageOneProps] = defineField('featuredImageOne')
const [featuredImageTwo, featuredImageTwoProps] = defineField('featuredImageTwo')
const [featuredImageThree, featuredImageThreeProps] = defineField('featuredImageThree') 
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
const featuredImageinsertion = gql`
mutation MyMutation($image_one: String, $image_three: String, $image_two: String, $movie_id: Int) {
  insert_featured_image_one(object: {image_one: $image_one, image_three: $image_three, image_two: $image_two, movie_id: $movie_id}) {
    image_one
    image_three
    image_two
    movie_id
  }
}
`

const createMovie = handleSubmit(async() => {
      
    const formData = new FormData()
    formData.append('file', thumbnail.value)
    formData.append('ImageOne', featuredImageOne.value)
    formData.append('ImageTwo', featuredImageTwo.value)
    formData.append('ImageThree', featuredImageThree.value)
    
    
    
    const { data, error } = await useFetch('http://localhost:3041/upload', {
      method: 'POST',
      body: formData,
    }).json()
    if (error.value) {
      console.error('Error uploading file:', error.value)
      return
    }
    
    imageUrl.value = data.value.url
    featuredImageOneUrl.value = data.value.urlone
    featuredImageTwoUrl.value = data.value.urltwo
    featuredImageThreeUrl.value = data.value.urlthree
    const{ data: movdata }= await movmutate({title: movietitle.value, genre: moviegenre.value, duration: movieduration.value, director_id: moviedirector.value, description: moviedescription.value, thumbnail: imageUrl.value, trending: istrending.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarone.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestartwo.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarthree.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarfour.value})
    await movstrmutate({movie_id: movdata.insert_movie_one.id, star_id: moviestarfive.value})
    await featuredImagemutate({image_one: featuredImageOneUrl.value, image_three: featuredImageThreeUrl.value, image_two: featuredImageTwoUrl.value, movie_id: movdata.insert_movie_one.id})
    movtitle.value = movdata.insert_movie_one.title + " is added as a movie"
    // console.log(movtitle.value)
});


const{mutate: movmutate} = useMutation(movinsertion)
const{mutate: movstrmutate} = useMutation(movstrinsertion)
const{mutate: featuredImagemutate} = useMutation(featuredImageinsertion)
</script>

<style>

</style>
