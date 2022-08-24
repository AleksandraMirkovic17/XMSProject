package com.agentska.dto;

import java.time.LocalDateTime;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonFormat;

public class JobCreationDTO {
	private String position;
	private List<String> requirements;
	private String description;
	@JsonFormat(pattern="yyyy-MM-dd HH:mm:ss")
	private LocalDateTime creationDate;
	@JsonFormat(pattern="yyyy-MM-dd HH:mm:ss")
	private LocalDateTime dueDate;
	
	public JobCreationDTO() {}

	public JobCreationDTO(String position, List<String> requirements, String description, LocalDateTime creationDate,
			LocalDateTime dueDate) {
		this.position = position;
		this.requirements = requirements;
		this.description = description;
		this.creationDate = creationDate;
		this.dueDate = dueDate;
	}

	public String getPosition() {
		return position;
	}
	public void setPosition(String position) {
		this.position = position;
	}
	public List<String> getRequirements() {
		return requirements;
	}
	public void setRequirements(List<String> requirements) {
		this.requirements = requirements;
	}
	public String getDescription() {
		return description;
	}
	public void setDescription(String description) {
		this.description = description;
	}
	public LocalDateTime getCreationDate() {
		return creationDate;
	}
	public void setCreationDate(LocalDateTime creationDate) {
		this.creationDate = creationDate;
	}
	public LocalDateTime getDueDate() {
		return dueDate;
	}
	public void setDueDate(LocalDateTime dueDate) {
		this.dueDate = dueDate;
	}
	
	
}
