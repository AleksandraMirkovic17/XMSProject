package com.xws.userservice.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.xws.userservice.model.User;
import com.xws.userservice.repository.RoleRepository;
import com.xws.userservice.service.UserService;

@RestController
@RequestMapping("/api/users")
@CrossOrigin("http://localhost:8081/")
public class UserController {
	
	@Autowired
	RoleRepository roleRepository;
	@Autowired
	UserService userService;
	
	@GetMapping("/all")
	public List<User> getUsers()
	{
		return userService.findAll();
	}

}
