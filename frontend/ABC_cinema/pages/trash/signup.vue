<template>
    <div class="h-screen bg-black text-white flex flex-col justify-center space-y-24 items-center">
        <p class="logo font-black text-6xl">ABC_Cinema</p>
        <form class="flex flex-col space-y-4 items-center self-center">
            <p class="text-2xl underline">Create an account</p>
            <input v-model="profilename" :="profilenameProps" placeholder="Username" type="text" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">{{ errors.profilename  }}</p>
            <input v-model="profilemail" :="profilemailProps" placeholder="Email" type="text" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">{{ errors.profilemail  }}</p>
            <input v-model="profilepassword" :="profilepasswordProps" placeholder="Password" type="password" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">{{ errors.profilepassword  }}</p>
            <input v-model="repassword" :="repasswordProps" placeholder="Re-enter" type="password" class="w-[380px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
            <p class="text-red-800">{{ errors.repassword  }}</p>
            <div  @click="createProfile" class="cursor-pointer content-center w-[380px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Sign up</div>
        </form>
        <p class="text-sm text-gray-500">{{ proname }}</p>
        <p>Already have an account <NuxtLink to = "login" class="textColor">Sign in</NuxtLink></p>
    </div>
</template>

<script setup>
definePageMeta({
        layout:""
    })
import { useForm } from 'vee-validate';
import * as Yup from 'yup';

function passwordRequired(value) {
  if (!value) {
    return 'Password is a required field';
  }  
  if (value.length < 8) {
    return 'Password is too short';
  }
    return true;
}  

function repassRequired(value) {
  if (!value) {
    return 'Password confirmation is required';
  }
  if (value != profilepassword.value){
    return 'Confirmation failed'
  }  
  if (value.length < 8) {
    return 'Password is too short';
  }
    return true;
}
function namerequired(value) {
  return value ? true : 'Profile name is required';
}
function emailrequired(value) {
  return value ? true : 'Email is required';
} 

// Create the form
const { defineField, handleSubmit, errors } = useForm({
  validationSchema: {
    profilename: namerequired,
    profilemail: emailrequired,
    profilepassword: passwordRequired,
    repassword: repassRequired
    // password: passwordRequired
  },
});

const profileinsertion = gql`
mutation MyMutation($username: String , $email: String , $password: String ) {
    insert_profile_one(object: {username: $username, email: $email, password: $password}) {
        username
        id
        password
    }
}
`
const createProfile = handleSubmit(async () => {
    const{ data }= await profilemutate({username: profilename.value, email: profilemail.value, password: profilepassword.value})
    proname.value = data.insert_profile_one.username + " Congradulation! you have created an account sign in link is below"
    // console.log(proname.value)
});
const [profilename, profilenameProps] = defineField('profilename') 
const [profilemail, profilemailProps] = defineField('profilemail')
const [profilepassword, profilepasswordProps] = defineField('profilepassword')
const [repassword, repasswordProps] = defineField('repassword')
const proname = ref("")
const{mutate: profilemutate} = useMutation(profileinsertion)
</script>

<style>

</style>