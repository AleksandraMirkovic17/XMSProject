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

import com.agentska.dto.JobCreationDTO;
import com.agentska.model.Company;
import com.agentska.model.Job;
import com.agentska.model.JobRequirement;
import com.agentska.model.User;
import com.agentska.model.enums.ERole;
import com.agentska.service.CompanyService;
import com.agentska.service.JobService;
import com.agentska.service.UserDetailsServiceImpl;
import com.agentska.service.UserService;

@RestController
@RequestMapping("/api")
public class JobController {
	@Autowired
	JobService jobService;
	@Autowired
	UserService userService;
	@Autowired
	CompanyService companyService;
	@Autowired
	UserDetailsServiceImpl userDetailsServiceImpl;
	
	@GetMapping("/job")
	public ResponseEntity<List<Job>> getAllJobs() {
		try {
			List<Job> jobs = jobService.findAll();
			return new ResponseEntity<>(jobs, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/job/{id}")
	public ResponseEntity<Job> getJobById(@PathVariable("id") Integer id) {
		Job jobData = jobService.findById(id);
		if (jobData != null) {
			return new ResponseEntity<>(jobData, HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	
	@GetMapping("/job/all/{id}")
	//@PreAuthorize("authentication.principal.id == #id")
	public ResponseEntity<List<Job>> getAllJobsByCompany(@PathVariable int id)
	{
		System.out.println(id);
		List<Job> jobData = jobService.findByCompanyId(id);
		return new ResponseEntity<>(jobData, HttpStatus.OK);
	}
	
	@PostMapping("/job/{companyId}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<Job> createJob(@PathVariable int companyId, @RequestBody JobCreationDTO jobDTO) {
		try {
			Company company = null;
			User currentUser = getCurrentUser();
			company = companyService.findById(companyId);
			/*
			List<Company> companies = companyService.findByOwnerId(currentUser.getId()); //Ovo je jako glupo napisano TODO: ispraviti
			for (Company c : companies) {
				if (c.getId() == companyId) {
					company = c;
					break;
				}
			}
			if (company == null) {
				return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
			}
			*/
			if (company.getOwner().getId() == currentUser.getId()) {
				return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
			}
			Job job = new Job(jobDTO, company);
			Job _job = jobService.createJob(job, jobDTO.getRequirements());
			return new ResponseEntity<>(_job, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}