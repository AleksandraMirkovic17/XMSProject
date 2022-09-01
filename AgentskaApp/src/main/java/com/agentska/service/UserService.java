package com.agentska.service;

import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import com.agentska.model.User;
import com.agentska.repository.UserRepository;
import com.agentska.model.Role;
import com.agentska.model.enums.ERole;
import com.agentska.repository.RoleRepository;
import com.agentska.repository.TokenRepository;
import com.agentska.model.VerificationToken;

@Service
public class UserService {
	
	@Autowired
	private UserRepository userRepository;
	@Autowired
	private TokenRepository tokenRepository;
	@Autowired
	private RoleRepository roleRepository;

	public User registerUser(User user)
	{
		Role userRole = roleRepository.findByName(ERole.USER)
				.orElseThrow(() -> new RuntimeException("Error: Role is not found."));
		user.getRoles().add(userRole);
		user.setActivated(true);
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
	
	public User findById(Integer id)
	{
		Optional<User> foundUser = userRepository.findById(id);
		if(foundUser.isEmpty())
			return null;
		else
			return foundUser.get();
	}
	
    public void createVerificationTokenForUser(final User user, final String token) {
        final VerificationToken myToken = new VerificationToken(token, user);
        tokenRepository.save(myToken);
    }

    public VerificationToken generateNewVerificationToken(final String existingVerificationToken) {
        VerificationToken vToken = tokenRepository.findByToken(existingVerificationToken).orElse(null);
        if(vToken == null)
        	return null;
        vToken.updateToken(UUID.randomUUID()
            .toString());
        vToken = tokenRepository.save(vToken);
        return vToken;
    }
    
	public User save(User user)
	{
		Role userRole = roleRepository.findByName(ERole.USER)
				.orElseThrow(() -> new RuntimeException("Error: Role is not found."));
		user.getRoles().add(userRole);
		return userRepository.save(user);
	}
	public User getLoggedUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		System.out.print(currentUserEmail);
		return findByEmail(currentUserEmail);
	}
}
