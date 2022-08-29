package com.agentska.service;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.agentska.model.Job;
import com.agentska.model.JobRequirement;
import com.agentska.repository.JobRepository;
import com.agentska.repository.JobRequirementRepository;

@Service
public class JobService {
	
	@Autowired
	private JobRepository jobRepository;
	@Autowired
	private JobRequirementRepository jobRequirementRepository;

	public Job createJob(Job job, List<String> requirementTexts)
	{
		Job savedJob = jobRepository.save(job); //Ako ne uspe treba mozda obrisati?
		System.out.println("11");
		for (String r : requirementTexts) {
			System.out.println("11.");
			JobRequirement jobRequirement = new JobRequirement(job, r);
			System.out.println("12.");
			jobRequirementRepository.save(jobRequirement);
			System.out.println("13.");
		}
		System.out.println("22");
		return savedJob;
	}
	
	public Job findById(Integer id)
	{
		Optional<Job> foundJob = jobRepository.findById(id);
		if(foundJob.isEmpty())
			return null;
		else
			return foundJob.get();
	}
	
	public List<JobRequirement> getRequirementsByJobId(Integer id) {
		return jobRequirementRepository.findByJobId(id);
	}
	
	public List<Job> findByCompanyId(Integer companyId)
	{
		return jobRepository.findByCompanyId(companyId);
	}
	
	@Transactional
	public void deleteById(Integer id)
	{
		jobRepository.deleteById(id);
	}
	
	public List<Job> findAll()
	{
		return jobRepository.findAll();
	}
	
	public Job save(Job job)
	{
		return jobRepository.save(job);
	}
	/*
	private User getCurrentUser() {
		String currentUserEmail = SecurityContextHolder.getContext().getAuthentication().getName();
		return userService.findByEmail(currentUserEmail);
	}
	*/
}
