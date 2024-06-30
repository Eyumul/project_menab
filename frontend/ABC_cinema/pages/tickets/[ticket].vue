<template>
    <div class="flex flex-col justify-center bg-black text-white h-screen">
        <div class="bg-black">
            <div class="max-w-sm mx-auto bg-gray-900 text-gray-200 shadow-xl rounded-lg overflow-hidden p-6">
                <div class="text-center mb-6">
                    <h2 class="text-3xl font-bold text-yellow-500">Ticket</h2>
                </div>
                <div class="mb-4 border-b border-gray-700 pb-2">
                <p><span class="font-semibold text-yellow-500">Customer name:</span> {{ customerName }}</p>
            </div>
                <div class="mb-4 border-b border-gray-700 pb-2">
                <p><span class="font-semibold text-yellow-500">Ticket for movie:</span> {{movieName}}</p>
            </div>
            <div class="mb-4 border-b border-gray-700 pb-2">
                <p><span class="font-semibold text-yellow-500">Movie date:</span> {{date}}</p>
            </div>
                <div class="mb-4 border-b border-gray-700 pb-2">
                    <p><span class="font-semibold text-yellow-500">Movie time:</span> {{time}}</p>
                </div>
                <div class="mb-4 border-b border-gray-700 pb-2">
                    <p><span class="font-semibold text-yellow-500">Tx_ref:</span> {{Tx_ref}}</p>
                </div>
                <div class="mb-4 border-b border-gray-700 pb-2">
                    <p><span class="font-semibold text-yellow-500">Price:</span>{{ price }}</p>
                </div>
                <div class="mb-4">
                    <p class="text-green-500"><span class="font-semibold text-yellow-500">Status:</span> {{status}}</p>
                </div>
            </div>
            <NuxtLink class="flex justify-center" to="../"><span class="text-xl pt-4 hover:text-[#0089D0] underline">Go Back</span></NuxtLink>
        </div>
    </div>
</template>

<script setup>
definePageMeta({
            layout:""
        })
const {ticket} = useRoute().params
const scheduleId = ticket.slice(4,6)
const userId = ticket.slice(10,12)
const customerName = ref("")
const movieName = ref("")
const date = ref("")
const time = ref("")
const Tx_ref = ref("")
const price = ref(" ")
const status = ref(" ")
console.log(scheduleId,userId)
function convertTo12HourFormat(time) {
  let [hours, minutes, seconds] = time.split(':').map(Number);
  const period = hours >= 12 ? 'PM' : 'AM';
  hours = hours % 12 || 12;
  hours = String(hours).padStart(2, '0');
  minutes = String(minutes).padStart(2, '0');
  seconds = String(seconds).padStart(2, '0');
  return `${hours}:${minutes} ${period}`;
}
function formatDate(inputDate) {
  const date = new Date(inputDate);
  const days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  const months = [
    "January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
  ];
  const dayName = days[date.getUTCDay()]; // Get the day of the week
  const monthName = months[date.getUTCMonth()]; // Get the month
  const day = date.getUTCDate(); // Get the day of the month
  const year = date.getUTCFullYear(); // Get the year
  return `${dayName} ${monthName} ${day} - ${year}`;
}
setTimeout( async () => {
const addtickets = gql `
mutation MyMutation($schedule_id: Int, $user_id: Int, $tx_ref: String) {
  insert_ticket_one(object: {schedule_id: $schedule_id, user_id: $user_id, tx_ref: $tx_ref}, on_conflict: {constraint: ticket_tx_ref_key, update_columns: schedule_id}) {
    profile {
      username
    }
    schedule {
      movie {
        title
      }
      date
      time
    }
    tx_ref
    id
  }
}
`
const { mutate:addticket } = useMutation(addtickets)
const ticketData = {
    schedule_id: scheduleId,
    user_id: userId,
    tx_ref: ticket,
}
const { data } = await addticket(ticketData)
console.log(data.insert_ticket_one)
customerName.value = data.insert_ticket_one.profile.username
movieName.value = data.insert_ticket_one.schedule.movie.title
date.value = formatDate(data.insert_ticket_one.schedule.date)
time.value = convertTo12HourFormat(data.insert_ticket_one.schedule.time)
Tx_ref.value = ticket
price.value = "100 ETB"
status.value = "Payed"
}, 1000);
</script>

<style>

</style>