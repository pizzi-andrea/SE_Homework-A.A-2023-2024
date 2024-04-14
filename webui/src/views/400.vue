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
    
    <h1 style="color: coral;">400</h1>
    <b>{{ errormsg }}</b>
    <img class="" src="https://www.shutterstock.com/image-vector/broken-robot-fix-technology-hand-260nw-527043607.jpg" alt="400 errore del client">
    <button class="btn btn-primary" @click="$router.back()">Indietro</button>


</div>
	
</template>

<style>
</style>

