package com.xws.userservice.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.xws.userservice.model.Role;

public interface RoleRepository extends JpaRepository<Role, Integer> {

}
