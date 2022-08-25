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

import com.agentska.dto.CompanyDTO;
import com.agentska.dto.CompanyRegisterDTO;
import com.agentska.model.Company;
import com.agentska.model.User;
import com.agentska.model.enums.ERole;
import com.agentska.service.CompanyService;
import com.agentska.service.UserDetailsServiceImpl;
import com.agentska.service.UserService;

@RestController
@RequestMapping("/api")
public class CompanyController {
	@Autowired
	CompanyService companyService;
	@Autowired
	UserService userService;
	@Autowired
	UserDetailsServiceImpl userDetailsServiceImpl;
	
	@GetMapping("/company")
	public ResponseEntity<List<Company>> getAllCompanies() {
		try {
			List<Company> companies = companyService.findAll();
			return new ResponseEntity<>(companies, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/company/{id}")
	public ResponseEntity<Company> getCompanyById(@PathVariable("id") Integer id) {
		Company companyData = companyService.findById(id);
		if (companyData != null) {
			return new ResponseEntity<>(companyData, HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	
	@GetMapping("/company/all/{id}")
	@PreAuthorize("authentication.principal.id == #id")
	public ResponseEntity<List<Company>> getAllCompaniesByOwner(@PathVariable int id)
	{
		System.out.println(SecurityContextHolder.getContext().getAuthentication().getPrincipal());
		System.out.println(id);
		List<Company> companyData = companyService.findByOwnerId(id);
		return new ResponseEntity<>(companyData, HttpStatus.OK);
	}
	
	@PostMapping("/company")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<Company> createCompany(@RequestBody CompanyRegisterDTO companyDTO) {
		try {
			User currentUser = getCurrentUser();
			Company company = new Company(companyDTO, currentUser);
			Company _company = companyService.registerCompany(company);
			/*Company _company = companyService
					.save(company);*/
			return new ResponseEntity<>(_company, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	@PutMapping("/company/validate/{id}")
	@PreAuthorize("hasRole('ADMINISTRATOR') and isAuthenticated()")
	public ResponseEntity<Company> validateCompany(@PathVariable("id") Integer id) {
		System.out.println("ffffff");
		try {
			Company company = companyService.findById(id);
			if (company.isValidated())
				return new ResponseEntity<>(null, HttpStatus.NOT_MODIFIED); // Da li je ovaj statusni kod?
			company.setValidated(true);
			return new ResponseEntity<>(companyService.validateCompany(company), HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	@PutMapping("/company/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<Company> updateCompany(@PathVariable("id") Integer id, @RequestBody CompanyDTO companyDTO) {
		try {
			User currentUser = getCurrentUser();
			Company company = companyService.findById(id);
			if (company == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != company.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN); // Da li je forbidden kad menja tudji resurs?
			if (!company.isValidated())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			company.setName(companyDTO.getName());
			company.setDescription(companyDTO.getDescription());
			company.setContactInfo(companyDTO.getContactInfo());
			return new ResponseEntity<>(companyService.save(company), HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	@DeleteMapping("/company/{id}")
	@PreAuthorize("isAuthenticated()")
	public ResponseEntity<HttpStatus> deleteCompany(@PathVariable("id") Integer id) {
		try {
			System.out.println("FFFF1");
			User currentUser = getCurrentUser();
			Company company = companyService.findById(id);
			if (company == null)
				return new ResponseEntity<>(HttpStatus.NOT_FOUND);
			if (currentUser.getId() != company.getOwner().getId())
				return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
			System.out.println("FFFF2");
			System.out.println(id);
			companyService.deleteById(id);
			System.out.println("FFFF3");
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			e.printStackTrace();
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	/*
	@PutMapping("/companies/{id}")
	public ResponseEntity<Company> updateCompany(@PathVariable("id") long id, @RequestBody Company company) {
		Optional<Company> companyData = companyRepository.findById(id);
		if (companyData.isPresent()) {
			Company _company = companyData.get();
			_company.setTitle(company.getTitle());
			_company.setDescription(company.getDescription());
			_company.setPublished(company.isPublished());
			return new ResponseEntity<>(companyRepository.save(_company), HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	*/
	/*
	@DeleteMapping("/companies/{id}")
	public ResponseEntity<HttpStatus> deleteCompany(@PathVariable("id") long id) {
		try {
			companyRepository.deleteById(id);
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	/*
	@DeleteMapping("/companies")
	public ResponseEntity<HttpStatus> deleteAllCompanies() {
		try {
			companyRepository.deleteAll();
			return new ResponseEntity<>(HttpStatus.NO_CONTENT);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	/*
	@GetMapping("/companies/published")
	public ResponseEntity<List<Company>> findByPublished() {
		try {
			List<Company> companies = companyRepository.findByPublished(true);
			if (companies.isEmpty()) {
				return new ResponseEntity<>(HttpStatus.NO_CONTENT);
			}
			return new ResponseEntity<>(companies, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}