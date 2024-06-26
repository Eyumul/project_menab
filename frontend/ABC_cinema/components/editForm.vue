<template>
    <div>
        <div class="flex">
      <div class="border-2 p-5 rounded-[20px] space-x-8 border-[#90D9FF] bg-gray-800 flex">
        <p v-if="!isEditing" class="p-2 mr-8 ml-4 text-xl font-bold">{{ oldtext }}</p>
        <input
          v-else
          v-model="internalNewtext"
          class="p-2 pl-6 mr-12 text-xl text-white bg-black border-gray-500 rounded-[10px] font-bold"
          @keyup.enter="saveText"
          type="text"
        />
        <div class="self-center cursor-pointer" @click="editText">
          <svg v-if="!isEditing" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class=" text-[#0089D0] size-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L6.832 19.82a4.5 4.5 0 0 1-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 0 1 1.13-1.897L16.863 4.487Zm0 0L19.5 7.125" />
          </svg>
        </div>
        <p v-if="isEditing" @click="saveText" class="self-center cursor-pointer text-[#90D9FF] size-6">Save</p>
        <p v-if="isEditing" @click="cancel" class="cursor-pointer text-gray-300 self-center">Cancel</p>
        <div @click="deleteText" class="cursor-pointer self-center">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="text-red-500 size-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
            </svg>
        </div>
      </div>
    </div>
    </div>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue'

const props = defineProps({
  oldtext: {
    type: String,
    required: true
  },
  newtext: {
    type: String,
    required: true
  }
})

const emits = defineEmits(['update:newtext', 'update:oldtext', 'save', 'delete'])

const isEditing = ref(false)
const internalNewtext = ref(props.newtext)

// Watch for changes in the props and update the internal state accordingly
watch(() => props.newtext, (newValue) => {
  internalNewtext.value = newValue
})

function editText() {
  isEditing.value = !isEditing.value
  if (isEditing.value) {
    internalNewtext.value = props.oldtext // Initialize the input with current text
  }
}

function saveText() {
  if (internalNewtext.value.trim() !== '') {
    emits('update:oldtext', internalNewtext.value)
    emits('save')
  }
  isEditing.value = false
}
function deleteText() {
    emits('delete')
}
function cancel() {
  isEditing.value = false
}
</script>

<style scoped>

</style>