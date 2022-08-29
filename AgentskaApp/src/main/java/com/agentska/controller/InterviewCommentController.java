package com.agentska.controller;

import java.security.Principal;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import com.agentska.model.Role;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.agentska.dto.InterviewCommentCreationDTO;
import com.agentska.model.InterviewComment;
import com.agentska.model.Company;
import com.agentska.model.User;
import com.agentska.service.InterviewCommentService;
import com.agentska.service.CompanyService;
import com.agentska.service.UserDetailsServiceImpl;
import com.agentska.service.UserService;

@RestController
@RequestMapping("/api")
public class InterviewCommentController {
	@Autowired
	InterviewCommentService interviewCommentService;
	@Autowired
	CompanyService companyService;
	@Autowired
	UserService userService;
	@Autowired
	UserDetailsServiceImpl userDetailsServiceImpl;
	
	@GetMapping("/interview_comment")
	public ResponseEntity<List<InterviewComment>> getAllInterviewComments(@RequestParam(required = false) String title) {
		try {
			List<InterviewComment> interviewComments = interviewCommentService.findAll();
			return new ResponseEntity<>(interviewComments, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/interview_comment/{id}")
	public ResponseEntity<InterviewComment> getInterviewCommentById(@PathVariable("id") Integer id) {
		InterviewComment interviewCommentData = interviewCommentService.findById(id);
		if (interviewCommentData != null) {
			return new ResponseEntity<>(interviewCommentData, HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	
	@GetMapping("/interview_comment/company/{id}")
	//@PreAuthorize("authentication.principal.id == #id")
	public ResponseEntity<List<InterviewComment>> getAllInterviewCommentsByCompany(@PathVariable int id)
	{
		List<InterviewComment> interviewCommentData = interviewCommentService.findByCompanyId(id);
		return new ResponseEntity<>(interviewCommentData, HttpStatus.OK);
	}
	
	@PostMapping("/interview_comment/{companyId}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<InterviewComment> createInterviewComment(@PathVariable int companyId, @RequestBody InterviewCommentCreationDTO interviewCommentDTO) {
		try {
			if (interviewCommentDTO.getRating() < 1 || interviewCommentDTO.getRating() > 5)
				return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
			User currentUser = getCurrentUser();
			Company company = companyService.findById(companyId);
			if (company.getOwner().getId() == getCurrentUser().getId()) {
				return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
			}
			InterviewComment interviewComment = new InterviewComment(company, currentUser, interviewCommentDTO);
			InterviewComment _interviewComment = interviewCommentService.createInterviewComment(interviewComment);
			return new ResponseEntity<>(_interviewComment, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	/*
	@PutMapping("/interview_comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<InterviewComment> updateInterviewComment(@PathVariable("id") Integer id, @RequestBody InterviewCommentDTO interviewCommentDTO) {
		try {
			User currentUser = getCurrentUser();
			InterviewComment interviewComment = interviewCommentService.findById(id);
			if (interviewComment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != interviewComment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN); // Da li je forbidden kad menja tudji resurs?
			if (!interviewComment.isValidated())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			interviewComment.setName(interviewCommentDTO.getName());
			interviewComment.setDescription(interviewCommentDTO.getDescription());
			interviewComment.setContactInfo(interviewCommentDTO.getContactInfo());
			return new ResponseEntity<>(interviewCommentService.save(interviewComment), HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	/*
	@DeleteMapping("/interview_comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<HttpStatus> deleteInterviewComment(@PathVariable("id") Integer id) {
		try {
			System.out.println("FFFF1");
			User currentUser = getCurrentUser();
			InterviewComment interviewComment = interviewCommentService.findById(id);
			if (interviewComment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != interviewComment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			System.out.println("FFFF2");
			System.out.println(id);
			interviewCommentService.deleteById(id);
			System.out.println("FFFF3");
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			e.printStackTrace();
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}