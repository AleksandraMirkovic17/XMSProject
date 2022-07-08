package com.agentska.repository;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import com.agentska.model.Company;
public interface CompanyRepository extends JpaRepository<Company, Long> {
	Optional<Company> findById(Integer id);
	Optional<Company> findByName(String name);
	List<Company> findByOwnerId(Integer owner);
	List<Company> findAll();
}