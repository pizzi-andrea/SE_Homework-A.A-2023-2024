<script>

export default {
	data: function() {
		return {
			errormsg: null,	
		}
	},
	methods: {

        mounted(){

        },
    }
		
}



</script>

<template>
	
<div class="align-items-center d-flex justify-content-center">
    
    <h1 style="color: rgb(80, 121, 255);">500</h1>
    <b>{{ errormsg }}</b>
    <img class="" src="https://www.elearningworld.org/wp-content/uploads/2018/02/Broken_Server.jpg" alt="400 errore del client">
    <button class="btn btn-primary" @click="$router.back()">Indietro</button>


</div>
	
</template>

<style>
</style>

