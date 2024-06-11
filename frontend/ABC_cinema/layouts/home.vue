<template>
    <div class="flex flex-col content-center text-white" id="slideshow5">
        <header class="h-full flex flex-col justify-between">
            <nav class="flex h-20 justify-between trans fixed w-full">
                <div class="px-3 self-center "><NuxtLink to = "/"><p class="logo">ABC_Cinema</p></NuxtLink></div>
                <ul class = "flex gap-x-12 px-9 items-center ">
                    <li><NuxtLink to = "">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="h-4 w-5">
                            <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                        </svg>
                    </NuxtLink></li>
                    <li><NuxtLink to = "#here">Dashboard</NuxtLink></li>
                    <li><NuxtLink to = "/moviesFeed">Movies</NuxtLink></li>
                    <li><NuxtLink to = "/schedule">Schedule</NuxtLink></li>
                    <li><NuxtLink to = "">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-7">
                            <path fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 1 1 9 0 4.5 4.5 0 0 1-9 0ZM3.751 20.105a8.25 8.25 0 0 1 16.498 0 .75.75 0 0 1-.437.695A18.683 18.683 0 0 1 12 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 0 1-.437-.695Z" clip-rule="evenodd" />
                        </svg>
                    </NuxtLink></li>
                </ul>
            </nav>
            <div class="mt-[250px]">
                <ul class="flex justify-between px-8">
                    <li @click="previous">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10">
                            <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06l-7.5-7.5Z" clip-rule="evenodd" />
                        </svg>
                    </li>
                    <li @click="next">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-10">
                            <path fill-rule="evenodd" d="M16.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L14.69 12 7.72 5.03a.75.75 0 0 1 1.06-1.06l7.5 7.5Z" clip-rule="evenodd" />
                        </svg>
                    </li>
                </ul>

                <div class="flex gap-x-8 mt-[300px] pb-2 justify-center items-center">
                    <span class="dot" id="dot1"></span>
                    <span class="dot" id="dot2"></span>
                    <span class="dot" id="dot3"></span>
                    <span class="dot" id="dot4"></span>
                    <span class="dot" id="dotActive"></span>
                </div>
            </div>
        </header>
    </div>
    <div>
        <slot />
    </div>
</template>

<script setup>
let imagenumber = 5

function previous(){
    let currentimage = document.getElementById( `slideshow${imagenumber}`)
    let previousdot = document.getElementById("dotActive")
    imagenumber--
    if (imagenumber < 1) {
        imagenumber = 5
        }
    let currentdot = document.getElementById(`dot${imagenumber}`)
    currentimage.id = `slideshow${imagenumber}`
    previousdot.id = `dot${imagenumber+1}`
    currentdot.id = "dotActive"
    console.log("previous ",previousdot.id, "current ", currentdot.id)
}
function next(){
    let currentimage = document.getElementById(`slideshow${imagenumber}`)
    imagenumber++
    if (imagenumber > 5){
        imagenumber = 1
    }
    currentimage.id = `slideshow${imagenumber}`
}
// setInterval(previous, 4000);

</script>

<style>
/*

dark blue: #0089D0
light blue: #24B4FF 

*/
#slideshow1 {
    height: 637px;
    width:100%;
    background-image: url('/figmaImage/slideshow1.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshow2 {
    height: 637px;
    width:100%;
    background-image: url('/figmaImage/slideshow2.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshow3 {
    height: 637px;
    width:100%;
    background-image: url('/figmaImage/slideshow3.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshow4 {
    height: 637px;
    width:100%;
    background-image: url('/figmaImage/slideshow4.jpg');
    background-repeat: no-repeat;
    background-size:cover;
}
#slideshow5 {
    height: 637px;
    width:100%;
    background-image: url('/figmaImage/slideshow5.webp');
    background-repeat: no-repeat;
    background-size:cover;
}
.trans {
    background-color: rgb(0, 0, 0,0);
    background-image: linear-gradient(rgb(0, 0, 0),rgba(110, 110, 110, 0));
}
.logo {
    font-family: sans-serif;
    font-size: 32px;
    background: -webkit-linear-gradient(right, white, #0089D0);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
li:hover {
    color:#0089D0;
    cursor: pointer;
}
.dot {
    height: 10px;
    width: 10px;
    background-color: #959494;
    border-radius: 50%;
    display: inline-block;  
    cursor: pointer;  
}
.dot:hover{
    background-color:#0089D0;
    height: 12px;
    width: 12px
}
#dotActive {
    height: 15px;
    width: 15px;
    background-color: #DAD8D8;
    border-radius: 50%;
    display: inline-block;
}
#dotActive:hover {
    background-color: #0089D0;
    height: 17px;
    width: 17px;
    cursor: pointer;
}
</style>