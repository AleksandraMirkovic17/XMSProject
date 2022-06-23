package com.xws.userservice.controller;

import java.util.List;
import java.util.stream.Collectors;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.xws.userservice.dto.LoginDTO;
import com.xws.userservice.dto.RegistrationDTO;
import com.xws.userservice.jwt.JwtResponse;
import com.xws.userservice.jwt.JwtUtils;
import com.xws.userservice.model.User;
import com.xws.userservice.model.misc.UserDetailsImpl;
import com.xws.userservice.repository.RoleRepository;
import com.xws.userservice.service.UserService;

@RestController
@RequestMapping("/api/users")
@CrossOrigin("http://localhost:8081/")
public class UserController {
	
	@Autowired
	AuthenticationManager authenticationManager;
	@Autowired
	UserService userService;
	@Autowired
	RoleRepository roleRepository;
	@Autowired
	PasswordEncoder encoder;
	@Autowired
	JwtUtils jwtUtils;
	
	@GetMapping("/all")
	public List<User> getUsers()
	{
		return userService.findAll();
	}

	@PostMapping("/login")
	public ResponseEntity<?> authenticateUser(@Valid @RequestBody LoginDTO loginRequest) {
		Authentication authentication = authenticationManager.authenticate(
				new UsernamePasswordAuthenticationToken(loginRequest.getEmail(), loginRequest.getPassword()));
		User user = userService.findByEmail(loginRequest.getEmail());
		if(user == null || !user.isActivated())
			return new ResponseEntity<>(
				      "User not authenticated or does not exist!", 
				      HttpStatus.NOT_FOUND);
		
		SecurityContextHolder.getContext().setAuthentication(authentication);
		String jwt = jwtUtils.generateJwtToken(authentication);
		
		UserDetailsImpl userDetails = (UserDetailsImpl) authentication.getPrincipal();		
		List<String> roles = userDetails.getAuthorities().stream()
				.map(item -> item.getAuthority())
				.collect(Collectors.toList());
		return ResponseEntity.ok(new JwtResponse(jwt, 
												 userDetails.getId(), 
												 userDetails.getUsername(), 
												 userDetails.getEmail(), 
												 roles));
	}
	
	@PostMapping
	@PreAuthorize("permitAll()")
	public ResponseEntity<String> register(@RequestBody RegistrationDTO signUpRequest, HttpServletRequest request)
	{
		if (userService.findByEmail(signUpRequest.getEmail()) != null) {
			return ResponseEntity
					.badRequest()
					.body("Error: Email is already taken!");
		}
		signUpRequest.setPassword(encoder.encode(signUpRequest.getPassword()));
		User registeredUser = new User(signUpRequest);
		userService.save(registeredUser);
		return new ResponseEntity<>(
			      "Registration successful!", 
			      HttpStatus.OK);
	}
}
