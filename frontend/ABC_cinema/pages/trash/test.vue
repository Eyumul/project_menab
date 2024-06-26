<template>
    <div class="mt-72">
      <h1>Image Upload</h1>
      <input type="file" @change="onThumbnailChange" />
      <button @click="uploadFile">Upload</button>
      <div v-if="imageUrl">
        <h2>Uploaded Image:</h2>
        <img :src="imageUrl" alt="Uploaded Image" />
      </div>
      
      <h1>Multiple Image Upload</h1>
      <input multiple type="file" @change="onFeaturedImagesChange" />
      <div v-if="imageUrl">
        <h2>Uploaded Image:</h2>
        <!-- <img :src="imageUrl" alt="Uploaded Image" /> -->
      </div>
      
      <input v-model="time" type="time"/>
      <input type="number" min="1" max="3" />
      {{ time }}

      <div>
        <h1>Auth0 Integration with Nuxt 3</h1>
        <div v-if="!isAuthenticated">
            <button @click="login">Login</button>
        </div>
        <div v-else>
            <p>Welcome, {{ username }}</p>
            <NuxtImg :src="username.picture"></NuxtImg>
            <button @click="logout">Logout</button>
        </div>
        </div>





        
        
        
        <div>
          <ul>
            <li v-for="(item, index) in visibleItems" :key="index">{{ item }}</li>
          </ul>
          <button v-if="visibleItems.length < items.length" @click="loadMore">Load More</button>
        </div>















        <div>
    <input v-model="oldDirectorName" placeholder="Enter current director name" />
    <input v-model="newDirectorName" placeholder="Enter new director name" />
    <button @click="handleUpdateDirector">Update Director</button>
  </div>

      </div>
    </template>
  
  <script setup>

// Define the mutation
const updateDirectorMutation = gql`
  mutation MyMutation($_eq: String!, $name: String!) {
    update_director(where: { name: { _eq: $_eq } }, _set: { name: $name }) {
      returning {
        name
      }
    }
  }
`
const { mutate: updateDirector } = useMutation(updateDirectorMutation)

// Reactive variables for the director names
const oldDirectorName = ref("")
const newDirectorName = ref("")

// Setup the mutation hook

// Function to update a director's name
async function handleUpdateDirector() {
  try {
    const { data } = await updateDirector({ _eq: oldDirectorName.value, name: newDirectorName.value })
    console.log('Updated director:', data.update_director.returning[0].name)
  } catch (error) {
    console.error('Error updating director:', error)
  }
}











































  import { useAuth0 } from '@auth0/auth0-vue';
import { ref, onMounted } from 'vue';
const time = ref("")

// Define the state variables
const items = ref([]);         // Full list of items
const visibleItems = ref([]);  // Items currently displayed
const itemsToShow = ref(3);    // Number of items to show initially and each time load more is clicked

// Fetch items from an API or any data source
const fetchItems = async () => {
  // Replace with your actual data fetching logic
  return ['Item 1', 'Item 2', 'Item 3', 'Item 4', 'Item 5', 'Item 6', 'Item 7', 'Item 8', 'Item 9'];
};

// Load more items
const loadMore = () => {
  const newLength = visibleItems.value.length + itemsToShow.value;
  visibleItems.value = items.value.slice(0, newLength);
};

// Fetch the items when the component is mounted
onMounted(async () => {
  items.value = await fetchItems();
  visibleItems.value = items.value.slice(0, itemsToShow.value);
});























  
import { useFetch } from '@vueuse/core'
  const thumnail = ref(null)
  const featuredImage = ref(null)
  const featuredImageOne = ref(null)
  const featuredImageTwo = ref(null)
  const featuredImageThree = ref(null)
  const imageUrl = ref('')
  const featuredImageOneUrl = ref('')
  const featuredImageTwoUrl = ref('')
  const featuredImageThreeUrl = ref('')
  
  function onThumbnailChange(event) {
    thumnail.value = event.target.files[0]
  }


  
  function onFeaturedImagesChange(event) {
    featuredImage.value = event.target.files
    console.log(featuredImage.value.length)

    featuredImageOne.value = event.target.files[0]
    featuredImageTwo.value = event.target.files[1]
    featuredImageThree.value = event.target.files[2]
  }
  




  async function uploadFile() {
    if (!thumnail.value) {
      alert('Please select a thumbnail')
      return
    }
    else if (!featuredImage.value) {
      alert('Please select a featured image')
      return
    }
    
    const formData = new FormData()
    formData.append('ImageOne', featuredImageOne.value)
    formData.append('file', thumnail.value)
    formData.append('ImageTwo', featuredImageTwo.value)
    formData.append('ImageThree', featuredImageThree.value)
    
    
    
    const { data, error } = await useFetch('http://localhost:3041/upload', {
      method: 'POST',
      body: formData,
    }).json()
    
    console.log(data)
    if (error.value) {
      console.error('Error uploading file:', error.value)
      return
    }
    
    imageUrl.value = data.value.url
    featuredImageOneUrl.value = data.value.urlone
    featuredImageTwoUrl.value = data.value.urltwo
    featuredImageThreeUrl.value = data.value.urlthree
    console.log(imageUrl.value)
    console.log(featuredImageOneUrl.value)
    console.log(featuredImageTwoUrl.value)
    console.log(featuredImageThreeUrl.value)
  }



  const auth0 = process.client ? useAuth0() : undefined
  const login = () => {
      auth0?.loginWithRedirect()
  }
  const logout = () => {
      auth0?.logout()
  }
  const token = ref(null)
  token.value = auth0?.getAccessTokenSilently()
 console.log(auth0?.user.value)
 console.log(auth0?.getAccessTokenSilently())

  const isAuthenticated = computed(() => {
    return auth0?.isAuthenticated.value
  })
  const username = computed(() => {
      return auth0?.user.value
  })
  </script>
  