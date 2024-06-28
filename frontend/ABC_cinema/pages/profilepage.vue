<template>
    <div class="bg-black h-screen text-white">
        <div class="pt-[100px] flex flex-col bg-black items-center">
            <h1 class="text-center text-5xl text-[#0089D0] font-bold">My Profile</h1>
            <div v-if="isAuthenticated">
                <div class="flex flex-col py-[100px] space-y-[70px] w-[1200px]">
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="text-3xl text-gray-500 font-bold">Name </h2>
                        <p class="text-3xl text-[#90D9FF] font-medium">{{ name }}</p>
                    </div>
                    <div class="w-full bg-gray-700 h-[1px]"></div>
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="flex space-x-4 text-3xl text-gray-500 font-bold">
                            <p>Email</p>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="self-center size-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
                            </svg>
                        </h2>
                        <p class="text-3xl text-[#90D9FF] font-medium">{{ email }}</p>
                    </div>
                    <div class="w-full bg-gray-700 h-[1px]"></div>
                    <div class="flex justify-between mx-[200px]">
                        <h2 class="flex text-3xl space-x-4 text-gray-500 font-bold">
                            <p>Phone N<span class="underline">o</span></p>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="self-center size-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z" />
                            </svg>
                        </h2>
                        <div v-if="number == null" class="flex items-center space-x-2">
                            <input v-model="phoneNumber" :="phoneNumberProps" @keyup.enter="addnumber" class="text-white bg-black border-2 border-[#0089D0] rounded-[20px] px-4" type="text" placeholder="0912345678"/>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#0089D0" class="cursor-pointer size-7" @click="addnumber">
                                <path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25ZM12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25V9Z" clip-rule="evenodd" />
                            </svg>
                        </div>
                        <p v-else class="text-3xl text-[#90D9FF] font-medium">{{"+251 " + number}}</p>
                    </div>
                    <p class="text-center text-red-500">{{ errors.phoneNumber }}</p>
                    <p class="text-center text-green-500">{{ confirm }}</p>
                    <div @click="logout" class="cursor-pointer text-[#0089D0] text-right text-xl underline">Logout</div>
                </div>
            </div>

        </div>
    </div>
</template>
<script setup>
    import { useForm } from 'vee-validate';
    import { useAuth0 } from '@auth0/auth0-vue';
    import {jwtDecode} from 'jwt-decode';
    const auth0 = process.client ? useAuth0() : undefined
    const role = ref('')
    const hasuraId = ref('')
    function required(value) {
        const regex = /^(07|09)\d{8}$/
        if (regex.test(value)) return true
        return "Enter a valid number in the format 09******** or 07********"
    }

    //create form
    const { defineField, handleSubmit, errors } = useForm({
        validationSchema: {
            phoneNumber: required,
        },
    });
    const [phoneNumber, phoneNumberProps] = defineField('phoneNumber')
    const confirm = ref('')
    const logout = () => {
        auth0?.logout()
    }
    const isAuthenticated = computed(() => {
        if (process.client) return auth0?.isAuthenticated.value
    })
    const username = computed(() => {
        if (process.client) return auth0?.user.value
    })
    if (process.client) {
        getHasuraToken()
        const tokentext = localStorage.getItem('hasura-token')
        if (tokentext) {
        const decodedToken = jwtDecode(tokentext);
        role.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-default-role'];
        hasuraId.value = decodedToken['https://hasura.io/jwt/claims']['x-hasura-user-id']
        }
    }
    const profilequery = gql`
    query MyQuery {
        profile {
            name
            email
            phone_number
            username
        }
    }`
    const name = ref(null)
    const email = ref(null)
    const number = ref(null)
    setTimeout( async () => {
        if (process.client){
            
            const {data} = await useAsyncQuery(profilequery)
            name.value = data.value.profile[0].username
            email.value = data.value.profile[0].email
            number.value = data.value.profile[0].phone_number
        }}, 1000);
    const updatenumbermutation = gql`
    mutation MyMutation($phone_number: Int!, $_eq: String!) {
        update_profile(where: {user_id: {_eq: $_eq}}, _set: {phone_number: $phone_number}) {
            returning {
            phone_number
            }
        }
    }`
    const addnumber = handleSubmit(async () => {
        setTimeout( async () => {
            const { data } = await updatenumber({ _eq: hasuraId.value, phone_number: phoneNumber.value })
            confirm.value = "Phone number added: " + data.update_profile.returning[0].phoneNumber
            console.log(data,"numberadded")
            setTimeout(function(){window.location.reload();}, 5000);
        }, 1000);
    });
    const {mutate:updatenumber} = useMutation(updatenumbermutation)
</script>

<style scoped>

</style>