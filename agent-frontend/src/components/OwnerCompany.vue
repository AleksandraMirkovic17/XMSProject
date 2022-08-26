<template class="tem">
  <div class="adventure-registration">
 <div class="container py-5">
    <!-- For demo purpose -->
    <header class="text-center text-white">
        <h1 class="display-4" style="color:black">All company</h1>  
    </header>
    <div v-for="c in companies" :key="c.id" class="row py-5">
        <div v-if="c.owner.username == loggedUser.username && c.validated==true" class="col-lg-7 mx-auto">
            <div  class="card shadow mb-4">
                <div class="card-body p-5">
                    <h4 class="mb-4">{{c.name}}</h4>
                    <!-- Unstyled list -->
                    <ul class="list-unstyled">
                        <li class="mb-2">{{c.contactInfo}}</li>
                        <li class="mb-2">{{c.description}}</li>
                        <li class="mb-2">{{c.owner.username}}</li>  
  <button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#staticBackdrop" v-on:click="SetFields(c)">
              Info
            </button>                    </ul>
                </div>
               
            </div>
        </div>
        
    </div>
    
    <div class="modal fade " id="staticBackdrop"  data-bs-backdrop="static"  tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="false">
      <div class="modal-dialog  modal-lg">
        
        <div class="modal-content">
             <button class="btn btn-primary" v-on:click="DisplayEdit">
              Edit
            </button>    
             <button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#staticBackdrop">
              Add job offer
            </button>
            
            <div class="modal-body" v-if="edit==true">
        <form>
      <h3 class="mb-3">Company registration</h3>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Company id</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "id" disabled>
      </div>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Company name</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "name" required>
      </div>

      <div class="col-11">
        <label for="adventure-name" class="form-label">Company description</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "description" required>
      </div>

      <div class="col-11">
        <label for="adventure-name" class="form-label">Company contact info</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "contactInfo" required>
      </div>
      <br>
       <div class="modal-footer">
          <button class="btn btn-primary " type="button" v-on:click="SaveEdit">Save </button>
        
             <button type="button" class="btn btn-secondary" v-on:click="CloseEdit">Close</button>
       </div>
          

      </form>
      </div>
          <ul class="nav" v-if="edit==false">
  <li class="nav-item">
    <a class="nav-link active" href="#">Comments</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" href="#">Job offer</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" href="#">Interview</a>
  </li>
  <li class="nav-item">
    <a class="nav-link " href="#">Salaries</a>
  </li>
</ul>
    <div class="modal-footer" v-if="edit==false">
            
<table class="table table-striped table-hover">
        <thead>
        <tr>
          <th scope="col">Comment</th>
          <th scope="col">Username</th>
     
        </tr>
        </thead>
        <tbody>
        <tr style="align-content: center" data-toggle="modal" >
          <td> Nesto ovde pise</td>
          <td  style="align-content: center"> i ovde isto </td>
          
        </tr>
         <tr style="align-content: center" data-toggle="modal" >
          <td> Nesto ovde pise</td>
          <td  style="align-content: center"> i ovde isto </td>
          
        </tr>
         <tr style="align-content: center" data-toggle="modal" >
          <td> Nesto ovde pise</td>
          <td  style="align-content: center"> i ovde isto </td>
        
        </tr>
         <tr style="align-content: center" data-toggle="modal" >
          <td> Nesto ovde pise</td>
          <td  style="align-content: center"> i ovde isto </td>
        </tr>
        
        </tbody>
      </table>
            
          </div>
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" v-if="edit==false">Close</button>
        </div>
      </div>
    </div>

    
<!-- Modal edita -->

    
   
</div>


  </div>

</template>
<script>
import axios from 'axios'
import {devServer} from "../../vue.config";

export default{
    data(){
        return{
           companies: null,
           loggedUser: null,
           edit: false,
           name: '',
           description: '',
           contactInfo: '',
           id: 0,
           token: ''
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
    },
    methods:
    {
        DisplayEdit()
        {
            this.edit=true;
        },
       
        SetFields(company)
        {
            this.name=company.name;
            this.description=company.description;
            this.contactInfo=company.contactInfo;
            this.id=company.id;
        },
        CloseEdit()
        {
            this.edit=false;
        },
        SaveEdit()
        {
            console.log(this.id)
            axios.put(devServer.proxy+'api/company/'+this.id,
            {
              "name": this.name,
              "description": this.description,
              "contactInfo": this.contactInfo
            },  {
            headers: {
              'Authorization' : 'Bearer ' + this.token,
            }
            })
            .then(response=>
            {
                 
                 console.log(response.data)
                 axios.get(devServer.proxy+'api/company')
      .then(response1=>{
        this.companies=response1.data;
         console.log("Ovo su sve kompanije:", this.companies)
      })
            })
            this.edit=false;
        }

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