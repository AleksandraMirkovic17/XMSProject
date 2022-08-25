package com.agentska.repository;

import java.util.List;
import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import com.agentska.model.Comment;
public interface CommentRepository extends JpaRepository<Comment, Long> {
	Optional<Comment> findById(Integer id);
	List<Comment> findByCompanyId(Integer companyId);
	List<Comment> findAll();
	void deleteById(Integer id);
}