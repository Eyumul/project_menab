<template>
    <div class="bg-black h-screen text-white"> 
        <div>
            <nav class=" pl-[138px] pt-[50px] space-x-4 flex justify-between">
                <NuxtLink to = "/admins" class="flex hover:text-[#0089D0]" >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10 self-center">
                        <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06l-7.5-7.5Z" clip-rule="evenodd" />
                    </svg>
                    <p class="self-center text-2xl font-black">Back</p>
                </NuxtLink>
                <p class="text-4xl content-center pr-[800px]">Edit Directors</p>
            </nav>
        </div>
        <div class="flex flex-col text-xl space-y-12 py-24 w-full items-center">
            <AddDirectorForm class="w-[600px]" />
        </div>
        <p class="text-red-500 text-center">{{ conflictError }}</p> 
        <div class="flex flex-wrap pl-36 bg-black">
            <div class="m-5" v-for="(director, index) in oldtext" :key="index">
                <EditForm :oldtext="director" :newtext="newtext" @update:oldtext="updateOldText" @save="handleSave(index)" @delete="handleDelete(index)"/>
            </div>
        </div>
    </div>
</template>

<script setup>
definePageMeta({
    layout:""
})

const dirquery = gql`
query MyQuery {
    director {
        name
        id
        movies {
            title
        }
    }
}`
const { data:directors } = await useAsyncQuery(dirquery)
const deleteDirectorMutation = gql`
  mutation MyMutation($_eq: String!) {
    delete_director(where: { name: { _eq: $_eq } }) {
      returning {
        name
      }
    }
  }
`
const { mutate: deleteDirector } = useMutation(deleteDirectorMutation)
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

let currentvalue = ref(null)
const newtext = ref('')
const oldtext = ref([])
const conflictError = ref('')
for (const director of directors.value.director){
    oldtext.value.push(director.name)
}
function updateOldText(value) {
    currentvalue = value
}

async function handleSave(index) {
    console.log(currentvalue)
    console.log('Saved text:', oldtext.value)
    try {
    const { data } = await updateDirector({ _eq: oldtext.value[index], name: currentvalue })
    console.log('Updated director:', data.update_director.returning[0].name)
} catch (error) {
    console.error('Error updating director:', error)
}
oldtext.value[index] = currentvalue
}
async function handleDelete(index) {
    console.log(index)
    if (!directors.value.director[index].movies.length == 0) {
        conflictError.value = "Cannot delete '" + oldtext.value[index] + "', it is linked to a movie"
        return
    }
    try {
    const { data } = await deleteDirector({ _eq: oldtext.value[index] })
    console.log('Deleted director:', data.delete_director.returning[0].name)
    location.reload()
  } catch (error) {
    console.error('Error deleting director:', error)
  } 
}
</script>

<style>

</style>