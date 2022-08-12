<template>
<section class="bg-dark text-light p-5 text-center text-sm-start">
<div id="registrationForm">
<br>
<br>
   <h1>Registration</h1>
   <input placeholder="Username" class="inputKredencijali" type="text" />            
   <br>
   <input placeholder="Email" class="inputKredencijali" type="text" />
   <br> 
   <input placeholder="Password" class="inputKredencijali" type="password" />
   <input placeholder="Repeat your password" class="inputKredencijali" type="password" />
   <br>
   <br>
   <br>
   <br>
 <button class="btn btn-outline-success reg mr-sm-2 reg" type="submit" ><i class="bi bi-person-plus-fill"></i>

Registration</button>
   <br>
   <br>
   <br>
 
</div>
</section>
</template>

<style>
 #registrationForm{
      width: 30%;
      margin-left: 30% ;
      text-align:center;
  }

  #registrationForm input{
  width: 60%;
  margin:5px 0;
	padding:10px;
	border-radius:20px;
	border: 2px solid #eee;
	box-shadow:0 0 15px 4px rgba(0,0,0,0.06);
  }

  #registrationForm .radio{
    display: flex;
    flex-direction: row;
    padding: 10px;
    border-radius:20px;
    background: #2ECC71;
  }

  #registrationForm input[type='radio']{
    width: 30%;
    margin: 0;
    border-color: purple;
  }



  #registrationForm input:focus{
	background: #FCFCFC;
	outline: none;
	
  }
  
 .buttonRegister {
	padding:10px;
	border:none;
	background-color:#2ECC71;
	color:#fff;
	font-weight:600;
	border-radius:20px;
	width:300px;
  }

  .buttonLogin:hover{
	background: #333;
  }

  #registrationForm textarea{
    font-family: inherit;
    color: #bbbbbb;
    width: 70%;
    height: 200px;
    padding: 10px;
    border-radius: 20px;
  }
</style>