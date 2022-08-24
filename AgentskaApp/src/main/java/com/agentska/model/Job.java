package com.agentska.model;

import java.time.LocalDateTime;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.agentska.dto.JobCreationDTO;

@Entity
@Table(name = "job")
public class Job {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Integer id;
	
	@ManyToOne
    @JoinColumn(name = "company_id")
    protected Company company;
	
	private String position;
	private String description;
	private LocalDateTime creationDate;
	private LocalDateTime dueDate;
	
	public Job() {}
	
	public Job(Integer id, Company company, String position, String description, LocalDateTime dueDate,
			LocalDateTime creationDate) {
		this.id = id;
		this.company = company;
		this.position = position;
		this.description = description;
		this.dueDate = dueDate;
		this.creationDate = creationDate;
	}
	
	public Job(JobCreationDTO jobDTO, Company company) {
		this.company = company;
		this.position = jobDTO.getPosition();
		this.description = jobDTO.getDescription();
		this.dueDate = jobDTO.getDueDate();
		this.creationDate = jobDTO.getCreationDate();
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

	public LocalDateTime getDueDate() {
		return dueDate;
	}

	public void setDueDate(LocalDateTime dueDate) {
		this.dueDate = dueDate;
	}

	public LocalDateTime getCreationDate() {
		return creationDate;
	}

	public void setCreationDate(LocalDateTime creationDate) {
		this.creationDate = creationDate;
	}
	
}
