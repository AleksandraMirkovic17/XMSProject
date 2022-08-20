package com.agentska.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.Company;
import com.agentska.model.Role;
import com.agentska.model.User;
import com.agentska.repository.CompanyRepository;
import com.agentska.repository.RoleRepository;
import com.agentska.repository.TokenRepository;
import com.agentska.repository.UserRepository;
import com.agentska.model.VerificationToken;
import com.agentska.model.enums.ERole;

@Service
public class CompanyService {
	
	@Autowired
	private CompanyRepository companyRepository;
	@Autowired
	private RoleRepository roleRepository;
	@Autowired
	private UserService userService; // Ovde koristi servis

	public Company registerCompany(Company company)
	{
		return companyRepository.save(company);
	}
	
	public Company validateCompany(Company company)
	{
		User owner = company.getOwner();
		if (!owner.isOwner()) { // Ovo verovatno ne treba jer je hash tabela
			Role ownerRole = roleRepository.findByName(ERole.OWNER)
			.orElseThrow(() -> new RuntimeException("Error: Role is not found."));
			owner.getRoles().add(ownerRole);
		}
		userService.save(owner); // Nema provere da li je uspelo
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
	
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
}
