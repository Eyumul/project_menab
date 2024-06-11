<template>
    <div class="mt-72">
      <h1>Image Upload</h1>
      <input type="file" @change="onFileChange" />
      <button @click="uploadFile">Upload</button>
      <div v-if="imageUrl">
        <h2>Uploaded Image:</h2>
        <img :src="imageUrl" alt="Uploaded Image" />
      </div>
    </div>
  </template>
  
  <script setup>
  import { useFetch } from '@vueuse/core'
  
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
  