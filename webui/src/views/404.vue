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
    
    <h1 style="color: rgb(80, 121, 255);">403</h1>
    <b>{{ errormsg }}</b>
    <img class="" src="https://yoast.com/cdn-cgi/image/width=400%2Cheight=210%2Cfit=scale-down%2Cf=auto%2Cquality=75%2Conerror=redirect/app/uploads/2015/08/404_error_checking_FI.png" alt="404 non trovato">
    <button class="btn btn-primary" @click="$router.back()">Torna indietro</button>


</div>
	
</template>

<style>
</style>


