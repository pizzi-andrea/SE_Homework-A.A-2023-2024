<script>
import Profile from "../components/Profile.js"
export default {
	
    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: Profile,
			myStream: [],
            pathUser: '/users/' + this.$route.params.id + '/',
            
        };
    },
    methods: {
        async refresh() {
            let response;
            try {
                response = await this.$axios.get( this.pathUser);
                switch (response.status) {
                case 200:
                    this.profile = new Profile();
                    this.profile = response.data;
                    break;
                }
            
            }catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.$router.push("/error/400");
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        $router.push("error/404");
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }
            }
            
               

			try {
                response = await this.$axios.get(this.pathUser + "myPhotos/");
                switch (response.status) {
                case 200:
				case 204:
                    this.myStream = response.data;
                    break;
                }
              
            }catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.$router.push("/error/400");
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        $router.push("error/404");
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }
            }
            
            
        },

        logout(){
            this.$axios.defaults.headers.common['Authorization'] = ''
            localStorage.removeItem('token')
            this.$router.push("/");

        },

        async uploadUsername(username){
            
            try{
                await this.$axios.put(this.pathUser, {
                    username: username
                });
            }catch(e){
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.$router.push("/error/400");
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        $router.push("error/404");
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }

            }
        }
    },
    mounted() {
        this.$axios.defaults.headers.common[
							"Authorization"
						] = `Bearer ${localStorage.getItem("token")}`;
        this.refresh();
    },

    

}



</script>

<template>
	<title>WasaPhoto</title>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profile di {{profile.user.username}} </h1>
            <div class="">
                    <p> Followers  {{profile.follower}}</p>
					<p> Following  {{profile.following}}</p>
                </div>
		</div>

		<div class = 'container float-start'>
			<div class="row" id="top">
			</div>

			<div class="row d-inline p-2 flex-row justify-content-start justify-content-between col-3 position-absolute top-25 start-50 translate-middle-x" id="body">
				
                <div class="">
                    <div v-for="post in myStream" :key="post.refer">
                                <Post :post="post"></Post>
                    </div>
                </div>
			</div>
			
			

			<div class="row" id="footer"></div>
		</div>

		
				
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>