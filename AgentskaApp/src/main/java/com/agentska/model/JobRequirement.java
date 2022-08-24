package com.agentska.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "job_requirement")
public class JobRequirement {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Integer id;
	
	@ManyToOne
    @JoinColumn(name = "job_id")
    protected Job job;
	
	private String requirementText;
	
	public JobRequirement() {}

	public JobRequirement(Integer id, Job job, String requirementText) {
		this.id = id;
		this.job = job;
		this.requirementText = requirementText;
	}
	
	public JobRequirement(Job job, String requirementText) {
		this.job = job;
		this.requirementText = requirementText;
	}

	public Integer getId() {
		return id;
	}

	public void setId(Integer id) {
		this.id = id;
	}

	public Job getJob() {
		return job;
	}

	public void setJob(Job job) {
		this.job = job;
	}

	public String getRequirementText() {
		return requirementText;
	}

	public void setRequirementText(String requirementText) {
		this.requirementText = requirementText;
	}
	
}
