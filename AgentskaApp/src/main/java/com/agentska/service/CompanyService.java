package com.agentska.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.Company;
import com.agentska.repository.CompanyRepository;
import com.agentska.repository.TokenRepository;

import com.agentska.model.VerificationToken;

@Service
public class CompanyService {
	
	@Autowired
	private CompanyRepository companyRepository;

	public Company registerCompany(Company company)
	{
		return companyRepository.save(company);
	}
	
	public Company findById(Integer id)
	{
		Optional<Company> foundCompany = companyRepository.findById(id);
		if(foundCompany.isEmpty())
			return null;
		else
			return foundCompany.get();
	}
	
	public Company findByName(String companyName)
	{
		Optional<Company> foundCompany = companyRepository.findByName(companyName);
		if(foundCompany.isEmpty())
			return null;
		else
			return foundCompany.get();
	}
	
	public List<Company> findByOwnerId(Integer ownerId)
	{
		return companyRepository.findByOwnerId(ownerId);
	}
	
	@Transactional
	public void deleteById(Integer id)
	{
		companyRepository.deleteById(id);
	}
	
	public List<Company> findAll()
	{
		return companyRepository.findAll();
	}
	
	public Company save(Company company)
	{
		return companyRepository.save(company);
	}
}
