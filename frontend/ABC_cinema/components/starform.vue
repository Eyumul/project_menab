<template>
    <div>
        <div class="space-y-3">
                <form class="flex items-center w-full justify-between">
                    <label>Insert movie star: </label>
                    <input v-model="starname" :="starnameProps" type="text" class="w-[280px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
                    <div @click="createStar" class="cursor-pointer content-center w-[80px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Add</div>
                </form>
                <p class="text-red-800 text-sm">{{ errors.starname  }}</p>
                <p class="text-sm text-gray-500">{{ strname }}</p>
        </div>
    </div>
</template>

<script setup>
//import necessaries
import { useForm } from 'vee-validate';
import * as Yup from 'yup';

//validation rule function
function required(value) {
  return value ? true : 'field is required';
}

// Create the form
const { defineField, handleSubmit, errors } = useForm({
  validationSchema: {
    starname: required,
  },
});

//create mutation(insertion)
const strinsertion = gql`
mutation MyMutation ($name: String) {
  insert_star_one(object: {name: $name}) {
    name
  }
}
`
//create a director if it passes validation 
const createStar = handleSubmit(async () => {
    const{ data }= await strmutate({name: starname.value})
    strname.value = data.insert_star_one.name + " is added as a star"
});

//declare consts
const [starname, starnameProps] = defineField('starname')
const{mutate: strmutate} = useMutation(strinsertion)
const strname = ref("")
</script>

<style scoped>

</style>