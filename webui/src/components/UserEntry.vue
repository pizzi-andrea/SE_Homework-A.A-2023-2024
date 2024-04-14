<script>
import { RouterLink, RouterView } from 'vue-router'

export default {
    props: ['usr'],

    data: function() {
        return {
            errormsg: "",
            id: this.usr.uid,
            username: this.usr.username,
            l_followers: [],
            followers: 0,
            isBanned: false,
            isFollower: false,
        }
    },

    methods: {
        async putBan() {
            if (this.isBanned) {
                return
            }

            let response = null
            try {
                response = await this.$axios.put("/users/" + localStorage.getItem('token') + "/banned/" + this.id);
                this.isBanned = true;
                this.isFollower = false;
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.errormsg = "Errore generico nella richiesta";
                        break;
                    case 401:
                        this.$router.push("/error/401");
                        break;
                    case 403:
                        this.$router.push("/error/403");
                        break;
                    case 404:
                        errormsg = "Utente non trovato";
                        break;
                    case 500:
                        this.$router.push("/error/500");
                        break;
                }
            }
        },

        async deleteBan() {
            if (!this.isBanned) {
                return
            }
            let response = null
            try {
                response = await this.$axios.delete("/users/" + localStorage.getItem('token') + "/banned/" + this.id);
                switch (response.status) {
                    case 204:
                        this.isBanned = false;
                        this.isFollower = false;
                        break;
                }
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        $router.push("/error/400");
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

        async refresh() {
            this.errormsg = null;
            this.isBannedUser();
            this.getFollowers();
            this.isFollowerUser();
        },

        async putFollow() {
            if (this.isFollower) {
                return
            }

            let response = null
            try {
                response = await this.$axios.put("/users/" + this.id + "/followers/" + localStorage.getItem('token'));
                switch (response.status) {
                    case 204:
                        this.refresh()
                        break;
                    case 201:
                        this.isFollower = true;
                        break;
                }
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        $router.push("/error/400");
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
            this.getFollowers()
        },

        async deleteFollow() {
            if (!this.isFollower) {
                return
            }
            let response = null
            try {
                response = await this.$axios.delete("/users/" + this.id + "/followers/" + localStorage.getItem('token'));
                switch (response.status) {
                    case 204:
                        this.isFollower = false;
                        break;
                }
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        $router.push("/error/400");
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

        async isBannedUser() {
            let response = null
            try {
                response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/banned/" + this.id);
                switch (response.status) {
                    case 200:
                        this.isBanned = true;
                        break;
                    case 204:
                        this.isBanned = false;
                        break;
                }
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        this.errormsg = "Errore generico nella richiesta";
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

        async getFollowers() {
            let response = null;
            try {
                response = await this.$axios.get("/users/" + this.id + "/followers/");
                switch (response.status) {
                    case 200:
                        this.l_followers = response.data;
                        this.followers = this.l_followers.length
                        break;
                    case 204:
                        this.isFollower = false;
                        break;
                }
            } catch (e) {
                console.log(e);
                switch (e.response.status) {
                    case 400:
                        $router.push("/error/400");
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
            this.isFollowerUser();
        },

        isFollowerUser() {
            let t = localStorage.getItem('token');
            this.isFollower = false;
            for (let i = 0; i < this.l_followers.length; i++) {
                if (this.l_followers[i].uid == t) {
                    this.isFollower = true;
                    break;
                }
            }
        }
    },

    mounted() {
        this.refresh();
    },
}
</script>

<template>
    <li class="list-group-item d-flex justify-content-between align-items-start rounded m-2 ">
        <div class="ms-2 me-auto">
            <RouterLink :to="'/users/' + this.id + '/profile'">
                <div class="fw-bold">{{username}}</div>
            </RouterLink>
        </div>
        <div class="btn-group ms-1 me-auto" role="group">
            <button type="button" class="btn btn-danger" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-if="isFollower" @click="deleteFollow">
                Non seguire
                <span class="badge text-bg-primary rounded-pill">{{followers}}</span>
            </button>
            <button type="button" class="btn btn-primary" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-else @click="putFollow" :disabled="isBanned">
                Segui
                <span class="badge text-bg-primary rounded-pill">{{followers}}</span>
            </button>
            <button type="button" class="btn btn-warning" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-if="isBanned" @click="deleteBan">
                Sblocca
            </button>
            <button type="button" class="btn btn-danger" style="--bs-btn-padding-y: .25rem; --bs-btn-padding-x: .5rem; --bs-btn-font-size: .75rem;" v-else @click="putBan">
                Blocca
            </button>
        </div>
    </li>

    <ErrorMeg :msg="errormsg" v-if="errormsg"></ErrorMeg>
</template>

<style>
</style>
