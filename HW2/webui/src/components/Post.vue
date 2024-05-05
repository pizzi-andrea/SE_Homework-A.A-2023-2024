<script>
import { RouterLink, RouterView } from 'vue-router'
export default {
	props: ['post'],

  data: function() {
		return {
      errormsg: "",
      username: "",
      myUsername: "",
      txtComment: "",
      users_likes: [],
      likes: 0,
      img: this.$axios.defaults.baseURL + "/images/" + this.post.refer,
      like: false,
      commentPerform: false,
      loading: false,
      comments: [],
      i: localStorage.getItem('token'),
      is_owner: false



    }
  },

  methods: {
        async refresh() {

          this.refreshComments();
          this.getUsername();
          this.isLiked();
          this.getMyUsername();
        },

        async getMyUsername(){
          let response = null;
          try {
            response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/")
            switch(response.status){
              case 200:
                this.myUsername = response.data.user.username
                break;
            }
          }catch(e){
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
            this.$router.push("error/404");
            break;
          case 500:
            this.$router.push("/error/500");
            break;
        }
          }

          
            
          },

        
        async delLike(){
          let response = null
          try{
            
            response = await this.$axios.delete("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'))
            
            switch(response.status){
              case 204:
                this.like = false;
                this.countLikes();
                break;
              case 404:
                //TODO
                break;
              case 400:
                //todo
                break;
              case 500:
                //todo
                break;
              default:
                //todo
                
            }
          }catch(e){
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
        async putLike(){
          let response = null
          try{
            
            response = await this.$axios.put("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'))
            switch(response.status){
              case 201:
                this.like = true;
                this.countLikes();
                break;
              case 204:
                break;
                
            }
          }catch(e){
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

        async putComment(){
          let response = null
          try{
            
            response = await this.$axios.post("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/comments/", {
              author: {
                username: this.myUsername
              },
              text: this.txtComment,

            });
            switch(response.status){
              case 201:
                this.disableComment();
                this.refreshComments();
                break;
                
            }
          }catch(e){
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

        enableComment(){
          this.commentPerform = true;
        },

        disableComment(){
          this.commentPerform = false;
          this.txtComment = ""
        },

        clearTextComment(){
          this.txtComment = ""
        },

        async getLikes(){
          let response = null;
          try{
            
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/")
           
            
            switch(response.status){
              case 200:
                this.users_likes = response.data;
                break;
              case 204:
                this.users_likes = [];
                break;
              
                
            }

            

          }catch(e){
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

        
        async getUsername(){
          let response = null
          try {
            response = await this.$axios.get("/users/" + this.post.owner + "/")
            switch(response.status){
            case 200:
              this.username =  response.data.user.username
              break;
            case 400:
              this.$router.push("/error/client");
              break;
            case 500:
              this.$router.push("/error/server");
              break;
            
          }
          }catch(e){
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

        async countLikes(){
          await this.getLikes()
          this.likes = this.users_likes ? this.users_likes.length : 0

        },

        async isLiked(){
          let response = null
          try{
            
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/likes/" + localStorage.getItem('token'));
            
            switch(response.status){
              case 200:
                this.like = true
                break;
              case 204:
                this.like = false;
            }
          }catch(e){
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

        async refreshComments(){
          let response = null
          try{
            
            response = await this.$axios.get("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/comments/");
            
            switch(response.status){
              case 200:
                this.comments = response.data
                break;
              case 204:
                this.comments = []
    
            }
          }catch(e){
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

        async delComment(commentId){
          try {
            let response = null; 
            response = await this.$axios.delete("/users/" + this.post.owner + "/myPhotos/" + this.post.refer + "/comments/" + commentId)

            switch(response.status) {
              case 204:
                document.getElementById(commentId).remove();
                this.refreshComments();
                break;

            }
          }catch(e){
            switch (e.response.status) {
          
                case 400: //bad request
                  this.errormsg = "Errore nella richiesta"
                  break; 
                case 401: //unauthorized    
                  this.errormsg = "Non autorizzato"
                  break;  
                  case 403: //forbidden   
                  this.errormsg = "Non hai i permessi per eseguire questa operazione"
                  break;
                case 404: //not found
                  this.errormsg = "Risorsa non trovata"
                  break;
                
                case 500: //server error  
                  this.errormsg = "Errore del server"
                  break;
          
          }

        }
      },

        youComment(commentId){
          return commentId == localStorage.getItem('token');
        }
        


        
      
        
   
  },
    mounted(){
        this.refresh();
      this.is_owner = this.post.owner == localStorage.getItem('token');
        this.users_likes = this.post.likes;
        this.likes = this.users_likes ? this.users_likes.length : 0;
        this.comments = this.post.comments;
    },
}
</script>

<template>
	<div class="card m-4" style="" id={{post.refer}}>
    <div class="card-header">

      <RouterLink :to=" '/users/' + post.owner + '/profile' ">
        <a> <i>Postata da {{ username }}</i> </a>
      </RouterLink>
      
      
    </div>
    <img class="card-img-top img-fluid object-fit-xxl-contain border rounded" :src="img" alt="img">
  <div class="card-body">
    <p class="card-text"> {{post.descriptionImg}} </p>
    
    <!-- -->
    <button type="button" @click="delLike" class="btn btn-danger" :disabled=is_owner v-if="like">
      <span class="badge badge-light"> <img src="./icons/lost-like.svg"> </span>  <span class="badge badge-light">{{likes}}</span>
    </button>
    <button type="button" @click="putLike" class="btn btn-primary" :disabled=is_owner v-else >
      <span class="badge badge-light"> <img src="./icons/need-like.svg"> </span> <span class="badge badge-light">{{likes}}</span>
    </button>

    <div class="btn btn-light m-3">
      Commenti: {{ comments ? comments.length : 0  }}
    </div>


    <p class="card-text"><small class="text-muted">Aggiunta il: {{ post.timeUpdate }}</small></p>
    
    
    <div class="list-group">
      <RouterLink :to=  "comment.author.uid == i ? '#' : '/users/' + comment.author.uid + '/profile' " v-for="comment in comments" :id="comment.commentId" :key="comment.commentId">
        <a  class="list-group-item list-group-item-action flex-column align-items-start border rounded" >
          <div class="d-flex w-100 justify-content-between">
            <h5 class="mb-1"></h5>
            <small><b><i>{{ comment.author.username}}</i></b></small>
            <small>{{comment.timeStamp}}</small>
          </div>
          <p class="mb-1">{{comment.text}}</p>
            <small></small>
            <div v-if="youComment(comment.author.uid)" class="align-items-end">
              <button type="button" class="btn btn-danger" @click="delComment(comment.commentId)"> <img src="./icons/trash.svg"> </button>
            </div>
        </a>
    </RouterLink>
    
      <div class="form-floating m-2" v-if="commentPerform">
        <textarea class="form-control" placeholder="Leave a comment here" id="floatingTextarea2" style="height: 100px" v-model="txtComment"></textarea>
        <label for="floatingTextarea2">Comments</label>

        <div class="btn-group" role="group m-4 p-2 d-flex">
          <button type="button" class="btn btn-primary" @click="putComment">Invia</button>
          <button type="button" class="btn btn-secondary" @click="clearTextComment">Cancella</button>
          <button type="button" class="btn btn-secondary"> <span class="badge badge-light" @click="disableComment"> <img src="./icons/cancel.svg"> </span> </button>
        </div>

      </div>
      <button type="button" @click="enableComment" :disabled=is_owner  class="btn btn-light m-2" v-else>
        <small> commenta </small> <span class="badge badge-light"><img src="./icons/write.svg"> </span>
      </button>
     
    </div>


    <div class="spinner-border" role="status" v-if="loading">
      <span class="visually-hidden">Loading...</span>
    </div>
    
    
  </div>
</div>
<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>