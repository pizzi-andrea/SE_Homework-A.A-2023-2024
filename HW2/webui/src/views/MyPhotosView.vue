<script>
export default {
    data: function () {
        return {
            errormsg: null,
            myPhotos: [],
            description: "",
        };
    },
    methods: {
        async refresh() {
            let response = null;
            try {
                response = await this.$axios.get(
                    "/users/" + localStorage.getItem("token") + "/myPhotos/"
                );
                switch (response.status) {
                    case 200:
                        this.myPhotos = response.data;
                        break;
                    case 204:
                        this.myPhotos = [];
                        break;
                }
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
        async uploadPost() {
            const form = new FormData();
            this.img = document.getElementById("photo").files;
            if (this.img == null || this.img.length == 0) {
                return;
            }
           
            form.append("img", this.img[0]);
            form.append("desc", this.description);
            let response = null;
            try {
                response = await this.$axios.post(
                    "/users/" + localStorage.getItem("token") + "/" + "myPhotos/",
                    form,
                    {
                        headers: {
                            "Content-Type": "multipart/form-data",
                        },
                    }
                );
                switch (response.status) {
                    case 201:
                        this.refresh();
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
    },
    mounted() {
        this.$axios.defaults.headers.common[
							"Authorization"
						] = `Bearer ${localStorage.getItem("token")}`;
        this.refresh();
    },
};
</script>

<template>
    <title>WasaPhoto</title>
    <div>
        <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <div class="btn-group me-2"></div>
        </div>
    </div>

    <div class="container float-start">
        <div class="row" id="top"></div>

        <div class="container float-start">
            <div class="row" id="top"></div>

            <div
                class="row d-inline p-2 flex-row justify-content-start justify-content-between col-3 position-absolute top-25 start-50 translate-middle-x"
                id="body"
            >
                <div class="">
                    <div v-for="post in myPhotos" :key="post.refer">
                        <DeletablePost :post="post"></DeletablePost>
                    </div>
                </div>

                <div class="card m-4" style="width: 40rem;">
                    <div class="card-header">
                        <small>Postato da</small>
                    </div>

                    <div class="card-img-top img-fluid object-fit-xxl-contain border rounded"></div>

                    <div class="card-body">
                        <div class="input-group mb-3">
                            <span class="input-group-text mb-3">descrizione</span>
                            <textarea class="form-control" aria-label="descrizione" name="desc" v-model="description"></textarea>
                        </div>

                        <div class="input-group mb-3">
                            <input type="file" class="form-control" aria-describedby="inputGroupFileAddon04" aria-label="Upload" accept="image/png" id="photo" name="img" />
                            <button class="btn btn-outline-secondary" type="button" id="inputGroupFileAddon04" @click="uploadPost">Carica</button>
                        </div>

                        <p class="card-text"><small class="text-muted">Aggiunta il </small></p>

                        <div class="list-group">
                            <small><i>commenti</i></small>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="row" id="footer"></div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style></style>


