package com.agentska.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.InterviewComment;
import com.agentska.model.User;
import com.agentska.repository.InterviewCommentRepository;

@Service
public class InterviewCommentService {
	
	@Autowired
	private InterviewCommentRepository interviewCommentRepository;
	@Autowired
	private UserService userService; // Ovde koristi servis

	public InterviewComment createInterviewComment(InterviewComment interviewComment)
	{
		return interviewCommentRepository.save(interviewComment);
	}
	
	public InterviewComment findById(Integer id)
	{
		Optional<InterviewComment> foundInterviewComment = interviewCommentRepository.findById(id);
		if(foundInterviewComment.isEmpty())
			return null;
		else
			return foundInterviewComment.get();
	}
	
	public List<InterviewComment> findByCompanyId(Integer companyId)
	{
		return interviewCommentRepository.findByCompanyId(companyId);
	}
	
	@Transactional
	public void deleteById(Integer id)
	{
		interviewCommentRepository.deleteById(id);
	}
	
	public List<InterviewComment> findAll()
	{
		return interviewCommentRepository.findAll();
	}
	
	public InterviewComment save(InterviewComment interviewComment)
	{
		return interviewCommentRepository.save(interviewComment);
	}
	
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}
