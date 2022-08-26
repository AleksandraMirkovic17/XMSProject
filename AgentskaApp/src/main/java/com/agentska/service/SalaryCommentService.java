package com.agentska.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.SalaryComment;
import com.agentska.model.User;
import com.agentska.repository.SalaryCommentRepository;

@Service
public class SalaryCommentService {
	
	@Autowired
	private SalaryCommentRepository salaryCommentRepository;
	@Autowired
	private UserService userService; // Ovde koristi servis

	public SalaryComment createSalaryComment(SalaryComment salaryComment)
	{
		return salaryCommentRepository.save(salaryComment);
	}
	
	public SalaryComment findById(Integer id)
	{
		Optional<SalaryComment> foundSalaryComment = salaryCommentRepository.findById(id);
		if(foundSalaryComment.isEmpty())
			return null;
		else
			return foundSalaryComment.get();
	}
	
	public List<SalaryComment> findByCompanyId(Integer companyId)
	{
		return salaryCommentRepository.findByCompanyId(companyId);
	}
	
	@Transactional
	public void deleteById(Integer id)
	{
		salaryCommentRepository.deleteById(id);
	}
	
	public List<SalaryComment> findAll()
	{
		return salaryCommentRepository.findAll();
	}
	
	public SalaryComment save(SalaryComment salaryComment)
	{
		return salaryCommentRepository.save(salaryComment);
	}
	
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}
