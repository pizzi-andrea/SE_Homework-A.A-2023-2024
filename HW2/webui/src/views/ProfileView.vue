<script>
export default {
	
    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: null,
			myStream: [],
            change: false,
            oldUsername: "",
            badUsername: false,
            
            
        };
    },
    methods: {
        async refresh() {
            let response;
            try {
                response = await this.$axios.get("/users/" + localStorage.getItem("token") + "/");
                switch (response.status) {
                case 200:
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
                response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/myStream/");
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
            localStorage.clear()
            this.$router.replace("/");

        },

        async uploadUsername(username){
            
            try{
                
                let response = await this.$axios.put("/users/" + localStorage.getItem('token') + "/", "\"" + username + "\"");
                switch(response.status){
                    case 200:
                    case 204:
                        this.refresh();
                        this.changeUsername();
                        break;
                }
            }catch(e){
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.badUsername = true;
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
        
        changeUsername() {
            
            if (this.change){
                this.change = false;
                this.profile.user.user = this.oldUsername;
                
                
            } else {
                this.change = true;
                this.oldUsername = this.profile.user.username

            }

            this.badUsername = false;



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
			<h1 class="h2">Benvenuto 
                <div class="input-group">
                    <input type="text" class="form-control" :placeholder="profile.user.username" aria-label="Recipient's username with two button addons" :disabled="!change" v-model="profile.user.username">
                    <button class="btn btn-outline-secondary" type="button" v-show="change" @click="uploadUsername(profile.user.username)">invia</button>
                    <button class="btn btn-outline-secondary" type="button" v-show="change" @click="changeUsername">back</button>
                    
                </div>
                <ErrorMsg v-if="badUsername" class="m-2" :msg="'Username non valido,\
                    un username valido contiene una qualsiasi sequenza di caratteri di lunghezza minima 3 e massima 16'">
                    </ErrorMsg>

                </h1>
            <div class="vr m-4"></div>
            <div class="vstack gap-2 col-md-5 mx-auto">

                    <h3> Followers</h3>  
                    <h5>{{profile.follower}}</h5>
					<h3> Following</h3> 
                    <h5> {{profile.following}}</h5>
                </div>
			<div class="btn-toolbar mb-2 mb-md-0">

                
				
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="refresh">
						Aggiorna stream
					</button>
                    <button type="button" class="btn btn-sm btn-warning" @click="changeUsername" :disabled="change">
						Aggiorna username
					</button>
					<button type="button" class="btn btn-sm btn btn-danger" @click="logout">
						Logout
					</button>
				</div>
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

