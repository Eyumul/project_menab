<template>
    <div class="h-screen bg-black text-white flex flex-col justify-center space-y-24 items-center">
        <p class="logo font-black text-6xl">ABC_Cinema</p>
        <form class="flex flex-col space-y-6 items-center self-center">
            <p class="text-2xl underline">Sign in to existing account</p>
            <input v-model="username" :="usernameProps" placeholder="Username" type="text" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">
                {{ errors.username  }}
            </p>
            <input v-model="password" :="passwordProps" placeholder="Password" type="password" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">
                {{ errors.password  }}
            </p>
            <div @click = "onSubmit" class="w-[380px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Login</div>
        </form>
        <p>Dont have an account? <NuxtLink to = "signup" class="textColor">Create account</NuxtLink></p>
    </div>
</template>

<script setup>
import { useForm } from 'vee-validate';
import * as Yup from 'yup';

definePageMeta({
        layout:""
    })

function passwordRequired(value) {
  if (!value) {
    return 'Password is a required field';
  }  
  if (value.length < 8) {
    return 'Password is too short';
  }
    return true;
}  

function required(value) {
  return value ? true : 'Username field is required';
}
    
// Create the form
const { defineField, handleSubmit, errors } = useForm({
  validationSchema: {
    username: required,
    password: passwordRequired
  },
});
const onSubmit = handleSubmit(async () => {
  // Submit to API
  console.log("values")
  console.log(username)
});
const [username, usernameProps] = defineField('username');
const [password, passwordProps] = defineField('password');
</script>

<style>

</style>