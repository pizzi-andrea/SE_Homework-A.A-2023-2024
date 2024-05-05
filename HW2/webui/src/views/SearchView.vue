<script>
export default {
	
    data: function () {
        return {
            errormsg: null,
            usernameSrc: "",
            usersFound: [],  
        };
    },
    methods: {
        async refresh() {
            this.errormsg = null;
            
        },

        async getUsers(){
            this.refresh();
            try{
                let response = null
                response = await this.$axios.get("/users/", {
                    params: {
                        limit: 120,
                        username: this.usernameSrc
                    }
                });
                switch(response.status){
                    case 200:
                        this.usersFound = response.data;
                        this.usersFound = this.usersFound.filter(user => user.uid != localStorage.getItem('token'));
                        break;
                    default:
                        this.usersFound = []
                }
            }catch(e){
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.errormsg = "Errore nella richiesta";
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        this.errormsg = "Nessun utente trovato";
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }
                

            }
        },
    },
    mounted(){
        this.$axios.defaults.headers.common[
							"Authorization"
						] = `Bearer ${localStorage.getItem("token")}`;
        this.refresh();
        
    }

}



</script>

<template>
	<title>WasaPhoto</title>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		</div>
        <div class="container">
            <div class="row">
                <div class="col">
                    
                </div>
                <div class="col-9 m-4">
                    <div class="d-flex">
                        <input class="form-control me-2" type="search" placeholder="Cerca un amico" aria-label="Search" v-model="usernameSrc" id="usernameSrc">
                        <button class="btn btn-outline-success" @click="getUsers">Cerca</button>
                    </div>
                </div>
                <div class="col"></div>
            </div>

            <div class="row">
                <div class="col"></div>
                <div class="col-4">
                    <ol class="list-group">
                        <UserEntry v-for="user in usersFound" :usr="user" :key="user.uid"></UserEntry>
                    </ol>
                </div>
                <div class="col"></div>
            </div>

            <div class="row">
                <div class="col"></div>
                <div class="col"></div>
                <div class="col"></div>

            </div>
        </div>
		
		
				
	</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>