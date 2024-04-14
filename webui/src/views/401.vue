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
    
    <h1 style="color: rgb(238, 255, 80);">401</h1>
    <b>{{ errormsg }}</b>
    <img class="" src="https://cdn-01.media-brady.com/store/emus/media/catalog/product/cache/4/image/85e4522595efc69f496374d01ef2bf13/1618578674/p/a/padlock-labels-pklb2-lg.jpg" alt="401 non identificato">
    <button class="btn btn-primary" @click="$router.push('/')">Torna al login</button>


</div>
	
</template>

<style>
</style>
