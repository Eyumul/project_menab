<template>
    <div>
        <div class="space-y-3">
                <div class="flex items-center w-full justify-between">
                    <label>Insert a director: </label>
                    <input v-model="directorname" :="directornameProps" @keyup.enter="createDirector" type="text" class="w-[280px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
                    <div @click="createDirector" class="cursor-pointer content-center w-[80px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Add</div>
                </div>
                <p class="text-red-800 text-sm">{{ errors.directorname  }}</p>
                <p class="text-sm text-green-500">{{ dirname }}</p>
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

//create form
const { defineField, handleSubmit, errors } = useForm({
  validationSchema: {
    directorname: required,
  },
});

//create mutation(insertion)
const dirinsertion = gql`
mutation MyMutation ($name: String) {
  insert_director_one(object: {name: $name}) {
    name
  }
}
`
//create a director if it passes validation 
const createDirector = handleSubmit(async () => {
    const{ data }= await dirmutate({name: directorname.value})
    dirname.value = data.insert_director_one.name + " is added as a movie director"
    setTimeout(function(){
      window.location.reload();
    }, 5000);
    // location.reload()
});

//declare consts
const [directorname, directornameProps] = defineField('directorname')
const dirname = ref("")
const{mutate: dirmutate} = useMutation(dirinsertion)

</script>

<style scoped>

</style>