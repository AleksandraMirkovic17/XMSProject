package com.agentska.dto;

public class SalaryCommentCreationDTO {
	private String salary;
	private String position;
	
	public SalaryCommentCreationDTO() {}

	public SalaryCommentCreationDTO(String salary, String position) {
		this.salary = salary;
		this.position = position;
	}

	public String getSalary() {
		return salary;
	}

	public void setSalary(String salary) {
		this.salary = salary;
	}

	public String getPosition() {
		return position;
	}

	public void setPosition(String position) {
		this.position = position;
	}
	
}
