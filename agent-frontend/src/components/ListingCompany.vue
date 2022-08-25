<template class="tem">
  <div class="adventure-registration">
 <div class="container py-5">
    <!-- For demo purpose -->
    <header class="text-center text-white">
        <h1 class="display-4" style="color:black">All company</h1>  
    </header>
    <div v-for="c in companies" :key="c.id" class="row py-5">
        <div v-if="c.owner.username != loggedUser.username" class="col-lg-7 mx-auto">
            <div  class="card shadow mb-4">
                <div class="card-body p-5">
                    <h4 class="mb-4">{{c.name}}</h4>
                    <!-- Unstyled list -->
                    <ul class="list-unstyled">
                        <li class="mb-2">{{c.contactInfo}}</li>
                        <li class="mb-2">{{c.description}}</li>
                        <li class="mb-2">{{c.owner.username}}</li>  
  <button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#staticBackdrop">
              Info
            </button>                    </ul>
                </div>
               
            </div>
        </div>
        
    </div>

    <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" >
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" >Create report</button>
          </div>
        </div>
      </div>
    </div>

    
   
</div>


  </div>

</template>

<script>

import axios from 'axios'
import {devServer} from "../../vue.config";

export default{
    data(){
        return{
           companies: null
        }
    },
    mounted(){
    
   this.token = localStorage.getItem('token').substring(1, localStorage.getItem('token').length-1);
    axios.get(devServer.proxy+'api/auth/user_data', {
      headers: {
        'Authorization' : 'Bearer ' + this.token,
      }
    })
    .then(response => {
      this.loggedUser =response.data
      console.log("Ovaj user je ulogovan:", this.loggedUser.roles[0].name)
      axios.get(devServer.proxy+'api/company')
      .then(response1=>{
        this.companies=response1.data;
         console.log("Ovo su sve kompanije:", this.companies)
      })
    })
    }

}
</script>

<style scoped>
.adventure-registration{

  width: 60%;
  horiz-align: center;
  margin-left: 20%;
  margin-top: 2%;
  border-radius: 3%;
  padding: 3%;
  background-blend-mode: lighten;
}
.removeImageBtn{
    position: absolute;
}
.double-field{
  display: flex;
  flex-direction: row;
}

.double-field .form-check{
  margin-left: 0.5%;
}

.adventure-registration .form-label{
}

.container {
  position: relative;
}
.side {
 position: absolute;
 right: 0;
 top:0;
 bottom: 0;
 width: 250px;
 z-index: 1;
 background-color: #fff;
}

.content {
  margin-right: 250px;
}



</style>