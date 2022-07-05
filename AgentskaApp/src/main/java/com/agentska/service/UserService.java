package com.agentska.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.agentska.model.User;
import com.agentska.repository.UserRepository;

@Service
public class UserService {
	
	@Autowired
	private UserRepository userRepository;

	public User registerUser(User user)
	{
		return userRepository.save(user);
	}
	
	public User findByEmail(String email)
	{
		Optional<User> foundUser = userRepository.findByEmail(email);
		if(foundUser.isEmpty())
			return null;
		else
			return foundUser.get();
	}
	
	public User findByUsername(String username)
	{
		Optional<User> foundUser = userRepository.findByUsername(username);
		if(foundUser.isEmpty())
			return null;
		else
			return foundUser.get();
	}
}
