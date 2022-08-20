package com.agentska.controller;

import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;
import java.util.Optional;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationEventPublisher;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.crypto.password.PasswordEncoder;
//import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import com.agentska.dto.UserDTO;
import com.agentska.model.Role;
import com.agentska.model.User;
import com.agentska.repository.UserRepository;
import com.agentska.service.UserService;
import com.agentska.repository.TokenRepository;
import com.agentska.model.VerificationToken;
import com.agentska.model.enums.ERole;
import com.agentska.event.OnRegistrationCompleteEvent;
import com.agentska.jwt.JwtUtils;


@RestController
@RequestMapping(
		value = {"/api"},
		produces = {"application/json"}
)
public class UserController {
	@Autowired
	UserService userService;
	@Autowired
	PasswordEncoder encoder;
	@Autowired
	JwtUtils jwtUtils;
	@Autowired
	UserRepository UserRepository; // TODO: Remove
	@Autowired
	TokenRepository tokenRepository;
	
	@Autowired
	ApplicationEventPublisher eventPublisher;
	
	@GetMapping("/users")
	public ResponseEntity<List<User>> getAllUsers(@RequestParam(required = false) String username) {
		try {
			List<User> Users = new ArrayList<User>();
			if (username == null)
				UserRepository.findAll().forEach(Users::add);
			else
				UserRepository.findByUsername(username);
			if (Users.isEmpty()) {
				return new ResponseEntity<>(HttpStatus.NO_CONTENT);
			}
			return new ResponseEntity<>(Users, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/users/{id}")
	public ResponseEntity<User> getUserById(@PathVariable("id") long id) {
		Optional<User> UserData = UserRepository.findById(id);
		if (UserData.isPresent()) {
			return new ResponseEntity<>(UserData.get(), HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	@PostMapping("/users")
	public ResponseEntity<User> createUser(@RequestBody User User) {
		try {
			User _User = UserRepository
					.save(new User(User.getUsername(), User.getEmail(), User.getPassword()));
			return new ResponseEntity<>(_User, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}


	@RequestMapping(
			method = {RequestMethod.POST},
			value = {"/users/register"},
			consumes = {"application/json"},
			produces = {"application/json"}
	)
	@CrossOrigin(
			origins = {"*"}
	)
	//@PostMapping("users/register")
	@PreAuthorize("not(isAuthenticated())")
	public ResponseEntity<String> register(@RequestBody UserDTO signUpRequest, HttpServletRequest request)
	{
		System.out.println("here1");
		if (userService.findByEmail(signUpRequest.getEmail()) != null) {
			return ResponseEntity
					.badRequest()
					.body("Error: Email is already taken!");
		}
		if (userService.findByUsername(signUpRequest.getUsername()) != null) {
			return ResponseEntity
					.badRequest()
					.body("Error: Username is already taken!");
		}
		signUpRequest.setPassword(encoder.encode(signUpRequest.getPassword())); // TODO: Add encoder
		signUpRequest.setPassword(signUpRequest.getPassword());
		System.out.println("before");
		User registeredUser = new User(signUpRequest);
		
		System.out.println("after0");
		userService.registerUser(registeredUser);
		System.out.println("after");
		String appUrl = request.getContextPath();

		System.out.print(appUrl);

        eventPublisher.publishEvent(new OnRegistrationCompleteEvent(registeredUser,
          request.getLocale(), appUrl));
		return new ResponseEntity<>(
			      "Registration successful!", 
			      HttpStatus.OK);
	}
	
	@PutMapping("/users/{id}")
	public ResponseEntity<User> updateUser(@PathVariable("id") long id, @RequestBody User User) {
		Optional<User> UserData = UserRepository.findById(id);
		if (UserData.isPresent()) {
			User _User = UserData.get();
			_User.setUsername(User.getUsername());
			_User.setEmail(User.getEmail());
			_User.setPassword(User.getPassword());
			return new ResponseEntity<>(UserRepository.save(_User), HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	@DeleteMapping("/users/{id}")
	public ResponseEntity<HttpStatus> deleteUser(@PathVariable("id") long id) {
		try {
			UserRepository.deleteById(id);
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@DeleteMapping("/users")
	public ResponseEntity<HttpStatus> deleteAllUsers() {
		try {
			UserRepository.deleteAll();
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("user/registration_confirm")
	public ResponseEntity<String> confirmRegistration(@RequestParam("token") String token) {
		
	    VerificationToken verificationToken = tokenRepository.findByToken(token).orElse(null);
	    if (verificationToken == null) {
	    	return new ResponseEntity<>(
				      "Verification token not found!", 
				      HttpStatus.NOT_FOUND);
	    }
	    
	    User user = verificationToken.getUser();
	    Calendar cal = Calendar.getInstance();
	    if ((verificationToken.getExpiryDate().getTime() - cal.getTime().getTime()) <= 0) {
	    	return new ResponseEntity<>(
				      "Verification token expired!", 
				      HttpStatus.BAD_REQUEST);
	    } 
	    
	    user.setActivated(true); 
	    userService.save(user); 
	    return new ResponseEntity<>(
			      "Successful verification!", 
			      HttpStatus.OK);
	}
}