<template>
  <div>
    <div class="page-header clear-filter" filter-color="orange">
      <div
          class="page-header-image"
          style="background-image: url('img/mountain.jpg')"
      ></div>
      <div style="width: 150%">
        <div class="container">
          <div class="col-md-12" style="width: 100%">
            <card type="login" plain>
              <div class="limiter">
                <div class="container-login100" style="background-image: /*savepage-url=images/bg-01.jpg*/ var(--savepage-url-15);">
                  <div class="wrap-login100">
                    <form class="login100-form validate-form">
                      <!--<span class="login100-form-logo">
                          <i class="zmdi zmdi-landscape"></i>
                      </span>-->
                      <span class="login100-form-title p-b-34 p-t-27">
                        <strong>Log in</strong>
                    </span>
                      <div class="wrap-input100 validate-input" data-validate="Enter username">
                        <input class="input100" type="text" name="username" placeholder="Username" value="" v-model="user.username">
                        <span class="focus-input100" data-placeholder=""></span>
                      </div>
                      <div class="wrap-input100 validate-input" data-validate="Enter password">
                        <input class="input100" type="password" name="pass" placeholder="Password" value="" v-model="user.password">
                        <span class="focus-input100" data-placeholder=""></span>
                      </div>
                      <!--<div class="contact100-form-checkbox">
                          <input class="input-checkbox100" id="ckb1" type="checkbox" name="remember-me">
                          <label class="label-checkbox100" for="ckb1">
                              Remember me
                          </label>
                      </div>-->
                      <div class="container-login100-form-btn">
                        <button class="login100-form-btn" type="button" v-on:click="login()">
                          <strong>Login</strong>
                        </button>
                      </div>
                      <div class="text-center p-t-90">
                        <a class="txt1" href="#">
                          Forgot Password?
                        </a>
                        <br>
                        <p class="txt1">
                          Don't have an account?
                          <a href="/register">
                            Register.
                          </a>
                        </p>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </card>
          </div>
        </div>
      </div>
      <main-footer></main-footer>
    </div>


  </div>
</template>

<script>

import UserService from '../services/UserService';
import { Card, Button, FormGroupInput } from '@/components';
import MainFooter from '@/layout/MainFooter';

export default {
    name: 'Login',
  bodyClass: 'login-page',
  components: {
    Card,
    MainFooter,
    [Button.name]: Button,
    [FormGroupInput.name]: FormGroupInput
  },
    data(){
        return {
            user: {
                username: "",
                password: ""
            }
        }
    },
    methods: {
        async login(){
            UserService.login(this.user).then(response => {
                console.log(response.data.token)
                console.log(response.data.username)
                if (response.data.token) {
                  console.log("Ovde")
                  localStorage.setItem('user', JSON.stringify(response.data));
                  this.$router.push("/profile/"+response.data.username)
                }
            })
            .catch(error => {
                  console.log(error)
                  alert("It is impossible to login!")
                }
            );
        }
    }
}

</script>

<style scoped>

/*//////////////////////////////////////////////////////////////////
[ RESTYLE TAG ]*/

* {
	margin: 0px;
	padding: 0px;
}

body, html {
	height: 100%;
}

/*---------------------------------------------*/
a {
	font-size: 14px;
	line-height: 1.7;
	color: #666666;
	margin: 0px;
	transition: all 0.4s;
	-webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
}



a:hover {
	text-decoration: none;
  color: #fff;
}

/*---------------------------------------------*/
h1,h2,h3,h4,h5,h6 {
	margin: 0px;
}

p {
	font-size: 14px;
	line-height: 1.7;
	color: #666666;
	margin: 0px;
}

ul, li {
	margin: 0px;
	list-style-type: none;
}


/*---------------------------------------------*/
input {
	outline: none;
	border: none;
}

textarea {
  outline: none;
  border: none;
}

textarea:focus, input:focus {
  border-color: transparent !important;
}

input:focus::-webkit-input-placeholder { color:transparent; }
input:focus:-moz-placeholder { color:transparent; }
input:focus::-moz-placeholder { color:transparent; }
input:focus:-ms-input-placeholder { color:transparent; }

textarea:focus::-webkit-input-placeholder { color:transparent; }
textarea:focus:-moz-placeholder { color:transparent; }
textarea:focus::-moz-placeholder { color:transparent; }
textarea:focus:-ms-input-placeholder { color:transparent; }

input::-webkit-input-placeholder { color: #fff;}
input:-moz-placeholder { color: #fff;}
input::-moz-placeholder { color: #fff;}
input:-ms-input-placeholder { color: #fff;}

textarea::-webkit-input-placeholder { color: #fff;}
textarea:-moz-placeholder { color: #fff;}
textarea::-moz-placeholder { color: #fff;}
textarea:-ms-input-placeholder { color: #fff;}

label {
  margin: 0;
  display: block;
}

/*---------------------------------------------*/
button {
	outline: none !important;
	border: none;
	background: transparent;
}

button:hover {
	cursor: pointer;
}

iframe {
	border: none !important;
}


/*//////////////////////////////////////////////////////////////////
[ Utility ]*/
.txt1 {
  font-size: 13px;
  color: #e5e5e5;
  line-height: 1.5;
}


/*//////////////////////////////////////////////////////////////////
[ login ]*/

.limiter {
  width: 100%;
  margin: 0 auto;
}

.container-login100 {
  width: 120%;
  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  margin-left: 35%;
  margin-top: 35%;

  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  position: relative;
  z-index: 2;
}

.container-login100::before {
  content: "";
  display: block;
  position: absolute;
  z-index: -1;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background-color: rgba(255,255,255,0.9);
}

.wrap-login100 {
  width: 500px;
  overflow: hidden;
  border: 1pt whitesmoke solid;
  padding: 55px 55px 37px 55px;

  background: #e85a4f;
  background: -webkit-linear-gradient(to top, #d56161, #756261);
  background: -o-linear-gradient(to top, #d56161, #756261);
  background: -moz-linear-gradient(to top, #d56161, #756261);
  background: linear-gradient(to top, #d56161, #756261);
}


/*------------------------------------------------------------------
[ Form ]*/

.login100-form {
  width: 100%;
}

.login100-form-logo {
  font-size: 60px;
  color: #333333;

  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background-color: #fff;
  margin: 0 auto;
}

.login100-form-title {
  font-size: 30px;
  color: #fff;
  line-height: 1.2;
  text-align: center;
  text-transform: uppercase;

  display: block;
}


/*------------------------------------------------------------------
[ Input ]*/

.wrap-input100 {
  width: 100%;
  position: relative;
  border-bottom: 2px solid rgba(255,255,255,0.24);
  margin-bottom: 30px;
}

.input100 {
  font-size: 16px;
  color: #fff;
  line-height: 1.2;

  display: block;
  width: 100%;
  height: 45px;
  background: transparent;
  padding: 0 5px 0 38px;
}

/*---------------------------------------------*/
.focus-input100 {
  position: absolute;
  display: block;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  pointer-events: none;
}

.focus-input100::before {
  content: "";
  display: block;
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 0;
  height: 2px;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;

  background: #fff;
}

.focus-input100::after {
  font-size: 22px;
  color: #fff;

  content: attr(data-placeholder);
  display: block;
  width: 100%;
  position: absolute;
  top: 6px;
  left: 0px;
  padding-left: 5px;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}

.input100:focus {
  padding-left: 5px;
}

.input100:focus + .focus-input100::after {
  top: -22px;
  font-size: 18px;
}

.input100:focus + .focus-input100::before {
  width: 100%;
}

.has-val.input100 + .focus-input100::after {
  top: -22px;
  font-size: 18px;
}

.has-val.input100 + .focus-input100::before {
  width: 100%;
}

.has-val.input100 {
  padding-left: 5px;
}


/*==================================================================
[ Restyle Checkbox ]*/

.contact100-form-checkbox {
  padding-left: 5px;
  padding-top: 5px;
  padding-bottom: 35px;
}

.input-checkbox100 {
  display: none;
}

.label-checkbox100 {
  font-size: 13px;
  color: #fff;
  line-height: 1.2;

  display: block;
  position: relative;
  padding-left: 26px;
  cursor: pointer;
}

.label-checkbox100::before {
  content: "\f26b";
  font-size: 13px;
  color: transparent;

  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  width: 16px;
  height: 16px;
  border-radius: 2px;
  background: #fff;
  left: 0;
  top: 50%;
  -webkit-transform: translateY(-50%);
  -moz-transform: translateY(-50%);
  -ms-transform: translateY(-50%);
  -o-transform: translateY(-50%);
  transform: translateY(-50%);
}

.input-checkbox100:checked + .label-checkbox100::before {
  color: #555555;
}


/*------------------------------------------------------------------
[ Button ]*/
.container-login100-form-btn {
  width: 100%;
  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.login100-form-btn {
  font-size: 16px;
  color: #555555;
  line-height: 1.2;

  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 20px;
  min-width: 120px;
  height: 50px;
  border-radius: 25px;

  background: #e85a4f;
  background: -webkit-linear-gradient(bottom, #e85a4f, #e98074);
  background: -o-linear-gradient(bottom, #e85a4f, #e98074);
  background: -moz-linear-gradient(bottom, #e85a4f, #e98074);
  background: linear-gradient(bottom, #e85a4f, #e98074);
  position: relative;
  z-index: 1;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}

.login100-form-btn::before {
  content: "";
  display: block;
  position: absolute;
  z-index: -1;
  width: 100%;
  height: 100%;
  border-radius: 25px;
  background-color: #fff;
  top: 0;
  left: 0;
  opacity: 1;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}

.login100-form-btn:hover {
  color: #fff;
}

.login100-form-btn:hover:before {
  opacity: 0;
}


/*------------------------------------------------------------------
[ Responsive ]*/

@media (max-width: 576px) {
  .wrap-login100 {
    padding: 55px 15px 37px 15px;
  }
}



/*------------------------------------------------------------------
[ Alert validate ]*/

.validate-input {
  position: relative;
}

.alert-validate::before {
  content: attr(data-validate);
  position: absolute;
  max-width: 70%;
  background-color: #fff;
  border: 1px solid #c80000;
  border-radius: 2px;
  padding: 4px 25px 4px 10px;
  top: 50%;
  -webkit-transform: translateY(-50%);
  -moz-transform: translateY(-50%);
  -ms-transform: translateY(-50%);
  -o-transform: translateY(-50%);
  transform: translateY(-50%);
  right: 0px;
  pointer-events: none;

  color: #c80000;
  font-size: 13px;
  line-height: 1.4;
  text-align: left;

  visibility: hidden;
  opacity: 0;

  -webkit-transition: opacity 0.4s;
  -o-transition: opacity 0.4s;
  -moz-transition: opacity 0.4s;
  transition: opacity 0.4s;
}

.alert-validate::after {
  content: "\f12a";
  font-size: 16px;
  color: #c80000;

  display: block;
  position: absolute;
  top: 50%;
  -webkit-transform: translateY(-50%);
  -moz-transform: translateY(-50%);
  -ms-transform: translateY(-50%);
  -o-transform: translateY(-50%);
  transform: translateY(-50%);
  right: 5px;
}

.alert-validate:hover:before {
  visibility: visible;
  opacity: 1;
}

@media (max-width: 992px) {
  .alert-validate::before {
    visibility: visible;
    opacity: 1;
  }
}


.p-t-27 {padding-top: 27px;}
.p-t-90 {padding-top: 90px;}
.p-b-34 {padding-bottom: 34px;}

</style>