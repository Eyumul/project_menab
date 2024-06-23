<template>
    <div>
        <div class="space-y-3">
                <form class="flex items-center w-full justify-between">
                    <label>Update a director: </label>
                    <select v-model="olddir" class="w-[280px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]">
                        <option :value=" director.id" v-for="director in directors.director" :key="director">{{director.name}}</option>
                    </select>
                    <div @click="updateDirector" class="cursor-pointer content-center w-[80px] h-[50px] bg-[#0089D0] text-[24px] text-center text-white rounded-[20px]">Edit</div>
                </form>
                <input placeholder="new director name" v-model="directorname" type="text" class="w-[280px] ml-[200px] h-[50px] border-none bg-black ring-[3px] ring-[rgb(0,137,208,0.5)] text-[24px] text-center rounded-[20px]" />
                <p class="text-red-800 text-sm">{{ olddir }}</p>
                <p class="text-sm text-gray-500">{{ directorname }}</p>
        </div>
    </div>
</template>

<script setup>
const dirquery = gql`
query MyQuery {
    director {
        name
        id
    }
}`
const { data:directors } = await useAsyncQuery(dirquery)

const updateMutation = gql`
mutation MyMutation($id: Int, $name: String) {
  update_director_by_pk(pk_columns: {id: $id}, _set: {name: $name}) {
    name
  }
}
`

function updateDirector() {
    const { data } = updateDirectorMutation({name: directorname.value}, { id: olddir.value});
}
const directorname = ref("")
const dirname = ref("")
const olddir = ref("")
const { mutate: updateDirectorMutation } = useMutation(updateMutation);
</script>

<style scoped>

</style>