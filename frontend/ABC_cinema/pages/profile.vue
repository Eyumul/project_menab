<template>
    <div class="mt-[72px]">
        <p> Hello world</p>
        <p>{{ data.profile }}</p>
    </div>
</template>
<script setup>
    const query = gql`
    query myquery {
    profile {
        username,
        password,
        email
     }
    }
    `
    const variables = { limit: 10 }
    const { data } = await useAsyncQuery(query)
</script>

<style scoped>

</style>