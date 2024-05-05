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
    <img class="" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Vienna_Convention_road_sign_B2a.svg/800px-Vienna_Convention_road_sign_B2a.svg.png" alt="403 non autorizzato">
    <button class="btn btn-primary" @click="$router.back()">Torna indietro</button>


</div>
	
</template>

<style>
</style>

