<template>
    <div class="mt-72">
      <h1>Image Upload</h1>
      <input type="file" @change="onFileChange" />
      <button @click="uploadFile">Upload</button>
      <div v-if="imageUrl">
        <h2>Uploaded Image:</h2>
        <img :src="imageUrl" alt="Uploaded Image" />
      </div>
      
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

    </div>
  </template>
  
  <script setup>
import { useAuth0 } from '@auth0/auth0-vue';
const auth0 = process.client ? useAuth0() : undefined
const login = () => {
    auth0?.loginWithRedirect()
}
const logout = () => {
    auth0?.logout()
}
const isAuthenticated = computed(() => {
  return auth0?.isAuthenticated.value
})
const username = computed(() => {
    return auth0?.user.value
})
// console.log(username)



  const file = ref(null)
  const imageUrl = ref('')
  
  function onFileChange(event) {
    file.value = event.target.files[0]
  }
  
  async function uploadFile() {
    if (!file.value) {
      alert('Please select a file to upload')
      return
    }
  
    const formData = new FormData()
    formData.append('file', file.value)
  
    const { data, error } = await useFetch('http://localhost:3041/upload', {
      method: 'POST',
      body: formData,
    }).json()
  
    //console.log(data)
    if (error.value) {
      console.error('Error uploading file:', error.value)
      return
    }
  
    imageUrl.value = data.value.url
  }
  </script>
  