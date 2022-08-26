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

import com.agentska.dto.SalaryCommentCreationDTO;
import com.agentska.model.SalaryComment;
import com.agentska.model.Company;
import com.agentska.model.User;
import com.agentska.service.SalaryCommentService;
import com.agentska.service.CompanyService;
import com.agentska.service.UserDetailsServiceImpl;
import com.agentska.service.UserService;

@RestController
@RequestMapping("/api")
public class SalaryCommentController {
	@Autowired
	SalaryCommentService salaryCommentService;
	@Autowired
	CompanyService companyService;
	@Autowired
	UserService userService;
	@Autowired
	UserDetailsServiceImpl userDetailsServiceImpl;
	
	@GetMapping("/salary_comment")
	public ResponseEntity<List<SalaryComment>> getAllSalaryComments(@RequestParam(required = false) String title) {
		try {
			List<SalaryComment> salaryComments = salaryCommentService.findAll();
			return new ResponseEntity<>(salaryComments, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/salary_comment/{id}")
	public ResponseEntity<SalaryComment> getSalaryCommentById(@PathVariable("id") Integer id) {
		SalaryComment salaryCommentData = salaryCommentService.findById(id);
		if (salaryCommentData != null) {
			return new ResponseEntity<>(salaryCommentData, HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	
	@GetMapping("/salary_comment/company/{id}")
	//@PreAuthorize("authentication.principal.id == #id")
	public ResponseEntity<List<SalaryComment>> getAllSalaryCommentsByCompany(@PathVariable int id)
	{
		List<SalaryComment> salaryCommentData = salaryCommentService.findByCompanyId(id);
		return new ResponseEntity<>(salaryCommentData, HttpStatus.OK);
	}
	
	@PostMapping("/salary_comment/{companyId}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<SalaryComment> createSalaryComment(@PathVariable int companyId, @RequestBody SalaryCommentCreationDTO salaryCommentDTO) {
		try {
			User currentUser = getCurrentUser();
			Company company = companyService.findById(companyId);
			if (company.getOwner().getId() == getCurrentUser().getId()) {
				return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
			}
			SalaryComment salaryComment = new SalaryComment(company, currentUser, salaryCommentDTO);
			SalaryComment _salaryComment = salaryCommentService.createSalaryComment(salaryComment);
			return new ResponseEntity<>(_salaryComment, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	/*
	@PutMapping("/salary_comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<SalaryComment> updateSalaryComment(@PathVariable("id") Integer id, @RequestBody SalaryCommentDTO salaryCommentDTO) {
		try {
			User currentUser = getCurrentUser();
			SalaryComment salaryComment = salaryCommentService.findById(id);
			if (salaryComment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != salaryComment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN); // Da li je forbidden kad menja tudji resurs?
			if (!salaryComment.isValidated())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			salaryComment.setName(salaryCommentDTO.getName());
			salaryComment.setDescription(salaryCommentDTO.getDescription());
			salaryComment.setContactInfo(salaryCommentDTO.getContactInfo());
			return new ResponseEntity<>(salaryCommentService.save(salaryComment), HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	/*
	@DeleteMapping("/salary_comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<HttpStatus> deleteSalaryComment(@PathVariable("id") Integer id) {
		try {
			System.out.println("FFFF1");
			User currentUser = getCurrentUser();
			SalaryComment salaryComment = salaryCommentService.findById(id);
			if (salaryComment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != salaryComment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			System.out.println("FFFF2");
			System.out.println(id);
			salaryCommentService.deleteById(id);
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