package com.agentska.repository;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import com.agentska.model.Job;
public interface JobRepository extends JpaRepository<Job, Long> {
	Optional<Job> findById(Integer id);
	//Optional<Job> findByName(String name);
	List<Job> findByCompanyId(Integer owner);
	List<Job> findAll();
	void deleteById(Integer id);
}