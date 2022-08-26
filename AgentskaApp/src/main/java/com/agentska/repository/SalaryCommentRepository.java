package com.agentska.repository;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import com.agentska.model.SalaryComment;

public interface SalaryCommentRepository extends JpaRepository<SalaryComment, Long> {
	Optional<SalaryComment> findById(Integer id);
	List<SalaryComment> findByCompanyId(Integer companyId);
	List<SalaryComment> findAll();
	void deleteById(Integer id);
}