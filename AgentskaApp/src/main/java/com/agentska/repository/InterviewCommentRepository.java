package com.agentska.repository;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import com.agentska.model.InterviewComment;

public interface InterviewCommentRepository extends JpaRepository<InterviewComment, Long> {
	Optional<InterviewComment> findById(Integer id);
	List<InterviewComment> findByCompanyId(Integer companyId);
	List<InterviewComment> findAll();
	void deleteById(Integer id);
}