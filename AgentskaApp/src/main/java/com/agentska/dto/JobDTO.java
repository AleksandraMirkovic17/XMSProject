package com.agentska.dto;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

import com.agentska.model.Company;
import com.agentska.model.Job;
import com.agentska.model.JobRequirement;

public class JobDTO {
	private Integer id;
	
    protected Company company;
	
	private String position;
	private String description;
	private LocalDate creationDate;
	private LocalDate dueDate;
	private List<String> requirements;
	
	public JobDTO() {
		this.requirements = new ArrayList<String>();
	}
	
	public JobDTO(Integer id, Company company, String position, String description, LocalDate dueDate,
			LocalDate creationDate, List<String> requirements) {
		this.id = id;
		this.company = company;
		this.position = position;
		this.description = description;
		this.dueDate = dueDate;
		this.creationDate = creationDate;
		this.requirements = requirements;
	}
	
	public JobDTO(Job job, List<JobRequirement> requirements) {
		this.id = job.getId();
		this.company = job.getCompany();
		this.position = job.getPosition();
		this.description = job.getDescription();
		this.dueDate = job.getDueDate().toLocalDate();
		this.creationDate = job.getCreationDate().toLocalDate();
		this.requirements = new ArrayList<String>();
		for (JobRequirement r : requirements) {
			this.requirements.add(r.getRequirementText());
		}
	}
	
	public Integer getId() {
		return id;
	}

	public void setId(Integer id) {
		this.id = id;
	}

	public Company getCompany() {
		return company;
	}

	public void setCompany(Company company) {
		this.company = company;
	}

	public String getPosition() {
		return position;
	}

	public void setPosition(String position) {
		this.position = position;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public LocalDate getDueDate() {
		return dueDate;
	}

	public void setDueDate(LocalDate dueDate) {
		this.dueDate = dueDate;
	}

	public LocalDate getCreationDate() {
		return creationDate;
	}

	public void setCreationDate(LocalDate creationDate) {
		this.creationDate = creationDate;
	}

	public List<String> getRequirements() {
		return requirements;
	}

	public void setRequirements(List<String> requirements) {
		this.requirements = requirements;
	}
	
	
	
}
