package com.agentska.controller;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
//import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.agentska.dto.UserDTO;
import com.agentska.model.User;
import com.agentska.repository.UserRepository;
import com.agentska.service.UserService;

@CrossOrigin(origins = "http://localhost:8081")
@RestController
@RequestMapping("/api")
public class UserController {
	@Autowired
	UserService userService;
	@Autowired
	UserRepository UserRepository; // TODO: Remove
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
	
	@PostMapping("/register")
	//@PreAuthorize("not(isAuthenticated())")
	public ResponseEntity<String> register(@RequestBody UserDTO signUpRequest, HttpServletRequest request)
	{
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
		//signUpRequest.setPassword(encoder.encode(signUpRequest.getPassword())); // TODO: Add encoder
		signUpRequest.setPassword(signUpRequest.getPassword());
		User registeredUser = new User(signUpRequest);
		userService.registerUser(registeredUser);
		
		String appUrl = request.getContextPath();
        //eventPublisher.publishEvent(new OnRegistrationCompleteEvent(registeredUser,  // WTF je ovo
        //  request.getLocale(), appUrl));
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
	/*
	@GetMapping("/users/published")
	public ResponseEntity<List<User>> findByPublished() {
		try {
			List<User> Users = UserRepository.findByPublished(true);
			if (Users.isEmpty()) {
				return new ResponseEntity<>(HttpStatus.NO_CONTENT);
			}
			return new ResponseEntity<>(Users, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
}