package com.agentska.controller;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
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
import com.agentska.model.Company;
import com.agentska.repository.CompanyRepository;
@CrossOrigin(origins = "http://localhost:8081")
@RestController
@RequestMapping("/api")
public class CompanyController {
	@Autowired
	CompanyRepository companyRepository;
	@GetMapping("/companies")
	public ResponseEntity<List<Company>> getAllCompanies(@RequestParam(required = false) String title) {
		try {
			List<Company> companies = companyRepository.findAll();
			return new ResponseEntity<>(companies, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/companies/{id}")
	public ResponseEntity<Company> getCompanyById(@PathVariable("id") Integer id) {
		Optional<Company> companyData = companyRepository.findById(id);
		if (companyData.isPresent()) {
			return new ResponseEntity<>(companyData.get(), HttpStatus.OK);
		} else {
			return new ResponseEntity<>(HttpStatus.NOT_FOUND);
		}
	}
	@GetMapping("/companies/user/{id}")
	public ResponseEntity<List<Company>> getCompanyByOwnerId(@PathVariable("id") Integer ownerId) {
		List<Company> companyData = companyRepository.findByOwnerId(ownerId);
		return new ResponseEntity<>(companyData, HttpStatus.OK);
	}
	/*
	@PostMapping("/companies")
	public ResponseEntity<Company> createCompany(@RequestBody Company company) {
		try {
			Company _company = companyRepository
					.save(new Company(company.getTitle(), company.getDescription(), false));
			return new ResponseEntity<>(_company, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	*/
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
}