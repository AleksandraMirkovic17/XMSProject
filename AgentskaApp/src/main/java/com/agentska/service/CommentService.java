package com.agentska.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.Comment;
import com.agentska.model.User;
import com.agentska.repository.CommentRepository;
import com.agentska.repository.RoleRepository;

@Service
public class CommentService {
	
	@Autowired
	private CommentRepository commentRepository;
	@Autowired
	private UserService userService; // Ovde koristi servis

	public Comment createComment(Comment comment)
	{
		return commentRepository.save(comment);
	}
	
	public Comment findById(Integer id)
	{
		Optional<Comment> foundComment = commentRepository.findById(id);
		if(foundComment.isEmpty())
			return null;
		else
			return foundComment.get();
	}
	
	public List<Comment> findByCompanyId(Integer companyId)
	{
		return commentRepository.findByCompanyId(companyId);
	}
	
	@Transactional
	public void deleteById(Integer id)
	{
		commentRepository.deleteById(id);
	}
	
	public List<Comment> findAll()
	{
		return commentRepository.findAll();
	}
	
	public Comment save(Comment comment)
	{
		return commentRepository.save(comment);
	}
	
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}
