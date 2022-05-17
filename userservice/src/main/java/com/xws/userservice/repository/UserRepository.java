package com.xws.userservice.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import com.xws.userservice.model.User;

public interface UserRepository extends JpaRepository<User, Integer> {

	Optional<User> findByEmailAndPassword(String email, String password);
	Optional<User> findByEmail(String email);
}
