package com.agentska.controller;

import java.util.List;
import java.util.stream.Collectors;
import javax.validation.Valid;

import com.agentska.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.*;
import com.agentska.model.User;
import com.agentska.jwt.JwtResponse;
import com.agentska.repository.RoleRepository;
import com.agentska.repository.UserRepository;
import com.agentska.dto.LoginDTO;
import com.agentska.jwt.JwtUtils;
import com.agentska.service.UserDetailsImpl;

@RestController
@RequestMapping(
		value = {"/api"},
		produces = {"application/json"}

)
public class AuthController {
	@Autowired
	AuthenticationManager authenticationManager;
	@Autowired
	UserRepository userRepository;
	@Autowired
	RoleRepository roleRepository;
	@Autowired
	PasswordEncoder encoder;
	@Autowired
	JwtUtils jwtUtils;
	@Autowired
	UserService userService;
	@RequestMapping(
			method = {RequestMethod.POST},
			value = {"auth/login"}
	)

	public ResponseEntity<?> authenticateUser(@Valid @RequestBody LoginDTO loginRequest) {
		Authentication authentication = authenticationManager.authenticate(
				new UsernamePasswordAuthenticationToken(loginRequest.getEmail(), loginRequest.getPassword()));
		User user = userRepository.findByEmail(loginRequest.getEmail()).orElse(null);


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
	@RequestMapping(
			method = {RequestMethod.GET},
			value = {"/auth/user_data"},
			produces = {"application/json"}
	)
	@CrossOrigin(
			origins = {"*"}
	)
	public ResponseEntity<User> getUserData() {
		System.out.println("-----------------get user data-----------");
	User user=this.userService.getLoggedUser();
	System.out.println(user.toString());

		try {
			return ResponseEntity.ok(this.userService.getLoggedUser());
		} catch (Exception var2) {
			return new ResponseEntity((MultiValueMap)null, HttpStatus.BAD_REQUEST);
		}
	}
}
