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

import com.agentska.dto.CommentCreationDTO;
import com.agentska.model.Comment;
import com.agentska.model.Company;
import com.agentska.model.User;
import com.agentska.model.enums.ERole;
import com.agentska.service.CommentService;
import com.agentska.service.CompanyService;
import com.agentska.service.UserDetailsServiceImpl;
import com.agentska.service.UserService;

@RestController
@RequestMapping("/api")
public class CommentController {
	@Autowired
	CommentService commentService;
	@Autowired
	CompanyService companyService;
	@Autowired
	UserService userService;
	@Autowired
	UserDetailsServiceImpl userDetailsServiceImpl;
	
	@GetMapping("/comment")
	public ResponseEntity<List<Comment>> getAllComments(@RequestParam(required = false) String title) {
		try {
			List<Comment> comments = commentService.findAll();
			return new ResponseEntity<>(comments, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/comment/{id}")
	public ResponseEntity<Comment> getCommentById(@PathVariable("id") Integer id) {
		Comment commentData = commentService.findById(id);
		if (commentData != null) {
			return new ResponseEntity<>(commentData, HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	
	@GetMapping("/comment/company/{id}")
	//@PreAuthorize("authentication.principal.id == #id")
	public ResponseEntity<List<Comment>> getAllCommentsByCompany(@PathVariable int id)
	{
		List<Comment> commentData = commentService.findByCompanyId(id);
		return new ResponseEntity<>(commentData, HttpStatus.OK);
	}
	
	@PostMapping("/comment/{companyId}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<Comment> createComment(@PathVariable int companyId, @RequestBody CommentCreationDTO commentDTO) {
		try {
			User currentUser = getCurrentUser();
			Company company = companyService.findById(companyId);
			if (company.getOwner().getId() == getCurrentUser().getId()) {
				return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
			}
			Comment comment = new Comment(company, currentUser, commentDTO);
			Comment _comment = commentService.createComment(comment);
			return new ResponseEntity<>(_comment, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	/*
	@PutMapping("/comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<Comment> updateComment(@PathVariable("id") Integer id, @RequestBody CommentDTO commentDTO) {
		try {
			User currentUser = getCurrentUser();
			Comment comment = commentService.findById(id);
			if (comment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != comment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN); // Da li je forbidden kad menja tudji resurs?
			if (!comment.isValidated())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			comment.setName(commentDTO.getName());
			comment.setDescription(commentDTO.getDescription());
			comment.setContactInfo(commentDTO.getContactInfo());
			return new ResponseEntity<>(commentService.save(comment), HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	/*
	@DeleteMapping("/comment/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<HttpStatus> deleteComment(@PathVariable("id") Integer id) {
		try {
			System.out.println("FFFF1");
			User currentUser = getCurrentUser();
			Comment comment = commentService.findById(id);
			if (comment == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != comment.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			System.out.println("FFFF2");
			System.out.println(id);
			commentService.deleteById(id);
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