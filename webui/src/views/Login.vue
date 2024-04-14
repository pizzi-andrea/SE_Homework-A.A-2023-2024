<script>
import ErrorMsg from "@/components/ErrorMsg.vue";

export default {
	data() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			errorInUsername: false,
			msgErrorData:
				"Errore stringa 'username' mal formattata. Un username valido deve avere una lunghezza compresa tra 3 e 16 caratteri di qualunque tipo.",
			token: 0,
		};
	},
	methods: {
		 refresh() {
			
			
				this.loading = true;
				this.errormsg = null;
				this.errorInUsername = false;
				this.username = "";
				this.token = 0;
				localStorage.clear();
			
		},

		async login() {
			this.dataValidator();
			let response = null;

			try {
				if (this.errorInUsername) {
					throw "Username non valido.";
				}
				response = await this.$axios.post("/session", {
					name: `${this.username}`,
				});

				switch (response.status) {
					case 200:
					case 201:
						this.token = response.data.Value;
						this.$axios.defaults.headers.common[
							"Authorization"
						] = `Bearer ${this.token}`;
						localStorage.setItem("token", this.token);
						this.$router.push("/users/" + this.token + "/");
						break;
					
				}
				this.loading = false;
			} catch (e) {
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

		dataValidator() {
			this.errorInUsername = !(
				this.username.length >= 3 &&
				this.username.length <= 16 &&
				/^.*$/.test(this.username)
			);
		},
	},
	mounted() {
		this.refresh();
	},
	components: {
		ErrorMsg,
	},
};
</script>

<template>
	<title>WasaPhoto - login</title>

	<div class="container form-group  mr-4 mt-4">
		<div class="row">
			<div class="container">
				<div class="row  flex-row d-inline-flex">
					<h4>Username</h4>
					<div class="col-sm m-1 d-sm-inline-flex">
						<input
							v-model="username"
							id="usernameField"
							placeholder="username"
							minlength="3"
							maxlength="16"
							class="m-1 form-control"
							@input="dataValidator"
							@focus="refresh"
						/>
						<button class="btn btn-primary m-4" @click="login">Login/Sigin</button>
					</div>
				</div>
			</div>
		</div>
	</div>

	<hr>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<br>
	<ErrorMsg v-if="errorInUsername" :msg="msgErrorData"></ErrorMsg>
</template>

<style>
</style>
