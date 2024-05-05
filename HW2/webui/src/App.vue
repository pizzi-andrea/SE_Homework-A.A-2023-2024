<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
		return {
			navBarShow: true,
			user: "",
			photo: "",
			users: ""
		}
	},

	methods: {
		showNavBar(){
			this.navBarShow = true 
		},
		hideNavBar(){
			this.navBarShow = false
		},

		refresh(){
			if(localStorage.getItem('token') && localStorage.getItem('token') != '0'){
				this.user =  "/users/" + localStorage.getItem('token') + "/"
				this.photo = this.user + "myPhotos/"
				this.users = "/users/";
			}
			else
			{
				this.user  = "/session";
				this.photo = "/session";
				this.users = "/session";
			}
			
		}
	},

	mounted(){
		this.refresh();
	}
	
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" @mouseover="refresh"><RouterLink :to=user  class="nav-link">WasaPhoto, <i>share your life</i></RouterLink></a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div  class="container-fluid">
		
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky" v-show="navBarShow">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item" @mouseover="refresh">
							<RouterLink :to=user  class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Il tuo profilo
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to=users class="nav-link" @mouseover="refresh">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								Gestione Amici
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to=photo class="nav-link" @mouseover="refresh">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								I tuoi Post
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
