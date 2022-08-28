<template class="tem">
  <div class="adventure-registration">
 <div class="container py-5">
    <!-- For demo purpose -->
    <header class="text-center text-white">
        <h1 class="display-4" style="color:black">All company</h1>  
    </header>
    <div v-for="c in companies" :key="c.id" class="row py-5">
        <div v-if="c.owner.username != loggedUser.username && c.validated==true" class="col-lg-7 mx-auto">
            <div  class="card shadow mb-4">
                <div class="card-body p-5">
                    <h4 class="mb-4">{{c.name}}</h4>
                    <!-- Unstyled list -->
                    <ul class="list-unstyled">
                        <li class="mb-2">{{c.contactInfo}}</li>
                        <li class="mb-2">{{c.description}}</li>
                        <li class="mb-2">{{c.owner.username}}</li>  
  <button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#staticBackdrop"  v-on:click="SetFields(c)">
              Info
            </button>                    </ul>
                </div>
               
            </div>
        </div>
        
    </div>
    
    <div class="modal fade " id="staticBackdrop"  data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" >
      <div class="modal-dialog  modal-lg">
        
        <div class="modal-content" >
           <div v-if="addComment==false && addInterview==false && addSalary==false">
             <button class="btn btn-primary" v-on:click="DisplayAddComment">
              Add comment
            </button>    
             <button class="btn btn-primary" v-on:click="DisplayAddinterview">
              Add interview comment
            </button>    
             <button class="btn btn-primary" v-on:click="DisplayAddSalary">
              Add salary comment
            </button>  
           
</div>
             <div class="modal-body" v-if="addComment==true">
        <form>
      <h3 class="mb-3">Add comment</h3>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Text</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "textComment" required>
      </div>
      <br>
       <div class="modal-footer">
          <button class="btn btn-primary " type="button" v-on:click="AddComment">Add </button>
        
             <button type="button" class="btn btn-secondary" v-on:click="Close">Close</button>
       </div>
          
      </form>
      
      </div>

      <div class="modal-body" v-if="addInterview==true">
        <form>
      <h3 class="mb-3">Add interview</h3>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Text</label>
        <input type="text" class="form-control" id="adventure-name"  v-model = "textInterview" required>
      </div>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Rating</label>
        <input type="number" min="1" class="form-control" id="adventure-name"  v-model = "rating" required>
      </div>
      <div class="col-11">
        <label for="difficulty_box" class="form-label">difficulty</label>
        <select class="form-control" name="difficulty_box" id="difficulty_box" v-model="difficulty" required>
          <option value="EASY">Easy</option>
          <option value="MEDIUM">Medium</option>
          <option value="HARD">Hard</option>
        </select>
      </div>
      <br>
       <div class="modal-footer">
          <button class="btn btn-primary " type="button" v-on:click="AddInterview">Add interview comment </button>
        
             <button type="button" class="btn btn-secondary" v-on:click="Close">Close</button>
       </div>
          
      </form>
      
      </div>

      <div class="modal-body" v-if="addSalary==true">
        <form>
      <h3 class="mb-3">Add salary</h3>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Position</label>
        <input type="text" class="form-control" id="adventure-name"  v-model = "position" required>
      </div>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Salary</label>
        <input type="text" class="form-control" id="adventure-name"  v-model = "salary" required>
      </div>
      <br>
       <div class="modal-footer">
          <button class="btn btn-primary " type="button" v-on:click="AddSalary">Add salary comment</button>
        
             <button type="button" class="btn btn-secondary" v-on:click="Close">Close</button>
       </div>
      </form>
      </div>
          <ul class="nav" v-if="addComment==false && addInterview==false && addSalary==false">
  <li class="nav-item">
    <a class="nav-link active" href="#" v-on:click="DisplayComment">Comments</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" href="#" v-on:click="DisplayJobOffer">Job offer</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" href="#" v-on:click="DisplayInterview">Interview</a>
  </li>
  <li class="nav-item">
    <a class="nav-link " href="#" v-on:click="DisplaySalaries">Salaries</a>
  </li>
</ul>
    <!--TABELA COMMENTS-->
    <div class="modal-footer" v-if="addComment==false && addInterview==false && addSalary==false && commentsCard==true && jobOfferCard==false && interviewCard==false && salariesCard==false">
            
<table  class="table table-striped table-hover">
        <thead>
        <tr>
          <th scope="col">Comment</th>
          <th scope="col">Username</th>
        </tr>
        </thead>
        <tbody>
        <tr style="align-content: center" data-toggle="modal" v-for="c in comments" :key="c.id">
          <td> {{c.text}}</td>
          <td  style="align-content: center"> {{c.user.username}}</td>
        </tr>
        </tbody>
      </table>
            
          </div>
          <!--TABELA Interview-->
    <div class="modal-footer" v-if="addComment==false && addInterview==false && addSalary==false && commentsCard==false && jobOfferCard==false && interviewCard==true && salariesCard==false">
            
<table  class="table table-striped table-hover">
        <thead>
        <tr>
          <th scope="col">Comment</th>
          <th scope="col">Username</th>
          <th scope="col">Rating</th>
          <th scope="col">Difficulty</th>
        </tr>
        </thead>
        <tbody>
        <tr style="align-content: center" data-toggle="modal" v-for="c in interviews" :key="c.id">
          <td> {{c.text}}</td>
          <td  style="align-content: center"> {{c.user.username}}</td>
          <td> {{c.rating}}</td>
          <td> {{c.difficulty}}</td>
        </tr>
        </tbody>
      </table>
            
          </div>

          <!--TABELA JOBOFFER-->
    <div class="modal-footer" v-if="addComment==false && addInterview==false && addSalary==false && commentsCard==false && jobOfferCard==true && interviewCard==false && salariesCard==false">
            
<table  class="table table-striped table-hover">
        <thead>
        <tr>
          <th scope="col">Position</th>
          <th scope="col">Description</th>
          <th scope="col">Due Date</th>
          <th scope="col">Company name</th>
        </tr>
        </thead>
        <tbody>
        <tr style="align-content: center" data-toggle="modal" v-for="c in jobs" :key="c.id">
          <td> {{c.position}}</td>
          <td  style="align-content: center"> {{c.description}}</td>
          <td> {{c.dueDate}}</td>
          <td  style="align-content: center"> {{c.company.name}}</td>
        </tr>
        </tbody>
      </table>
            
          </div>
           <div class="modal-footer" v-if="addComment==false && addInterview==false && addSalary==false && commentsCard==false && jobOfferCard==false && interviewCard==false && salariesCard==true">
            
<table  class="table table-striped table-hover">
        <thead>
        <tr>
          <th scope="col">Position</th>
          <th scope="col">Salary</th>
            <th scope="col">Username</th>
        </tr>
        </thead>
        <tbody>
        <tr style="align-content: center" data-toggle="modal" v-for="c in salaries" :key="c.id">
          <td> {{c.position}}</td>
           <td> {{c.salary}}</td>
          <td  style="align-content: center"> {{c.user.username}}</td>
        </tr>
        </tbody>
      </table>
            
          </div>
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
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
          token: '',
           companies: null,
           loggedUser: null,
           addComment: false,
           addInterview: false,
           addSalary: false, 
           textComment: '',
           textInterview: '',
           name: '',
           description: '',
           contactInfo: '',
           id: 0,
           rating: 1,
           difficulty: '',
           salary: '',
           position: '',
            commentsCard: false,
           jobOfferCard: false,
           salariesCard: false,
           interviewCard: false,
           comments: null,
           interviews: null,
           jobs: null,
           salaries: null


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
      console.log("Ovaj user je ulogovan:", this.loggedUser.roles[0].name) //Treba videti da li postoji odgovarajuca rola
      axios.get(devServer.proxy+'api/company')
      .then(response1=>{
        this.companies=response1.data;
         console.log("Ovo su sve kompanije:", this.companies)
      })
    })
    },
    methods:
    {
      DisplayAddComment()
      {
        this.addComment=true;
      },
      SetFields(company)
        {
            this.name=company.name;
            this.description=company.description;
            this.contactInfo=company.contactInfo;
            this.id=company.id;
        },
        AddComment()
        {
          let text=this.textComment
          axios.post(devServer.proxy+'api/comment/'+this.id,
          {
            "text": text
          },
             {
            headers: {
              'Authorization' : 'Bearer ' + this.token,
            }
          })
          .then(response=>
          {
            console.log(response.data)
            this.addComment=false;
          })
        },
        Close()
        {
           this.addComment=false;
            this.addInterview=false;
            this.addSallary=false;

        },
        DisplayAddinterview()
        {
          this.addInterview=true;
        },
        AddInterview()
        {
          let text=this.textInterview
          axios.post(devServer.proxy+'api/interview_comment/'+this.id,
          {
            "text": text,
            "rating": this.rating,
            "difficulty": this.difficulty
          },
          {
             headers: {
              'Authorization' : 'Bearer ' + this.token,
            }
          }).
          then(response=> {
            console.log(response.data)
            this.addInterview=false;
          })
        },
        DisplayAddSalary()
        {
          this.addSalary=true;
        },
        AddSalary()
        {
          axios.post(devServer.proxy+'api/salary_comment/'+this.id,
          {
            "position": this.position,
            "salary": this.salary
          },
            {
             headers: {
              'Authorization' : 'Bearer ' + this.token,
            }
          }).then(response=>
          {
            console.log(response.data)
            this.addSalary=false;
          })
        },
        DisplayComment()
        {
          this.commentsCard=true;
           axios.get(devServer.proxy+'api/comment/company/'+this.id)
      .then(response=>
      {
      this.comments=response.data;
      console.log(this.comments)
      this.commentsCard=true;
      this.salariesCard=false;
      this.interviewCard=false;
      this.jobOfferCard=false;
      }
      )
        },
       DisplayJobOffer(){
     
        axios.get(devServer.proxy+'api/job/all/'+this.id)
      .then(response=>
      {
        this.jobs=response.data;
  
         this.commentsCard=false;
      this.salariesCard=false;
      this.interviewCard=false;
      this.jobOfferCard=true;
      })
    },
    DisplaySalaries()
    {
      axios.get(devServer.proxy+'api/salary_comment/company/'+this.id)
      .then(response=>
      {
        this.salaries=response.data
          this.commentsCard=false;
      this.salariesCard=true;
      this.interviewCard=false;
      this.jobOfferCard=false;
      })
    },
    DisplayInterview()
    {
      axios.get(devServer.proxy+'api/interview_comment/company/'+this.id)
      .then(response=>
      {
        this.interviews=response.data
        this.commentsCard=false;
      this.salariesCard=false;
      this.interviewCard=true;
      this.jobOfferCard=false;
      })
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