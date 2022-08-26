package com.agentska.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.agentska.dto.SalaryCommentCreationDTO;

@Entity
@Table(name = "salary_comment")
public class SalaryComment {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	protected Integer id;
	
	@ManyToOne
    @JoinColumn(name = "salary_id")
    protected Company company;
	
	@ManyToOne
    @JoinColumn(name = "user_id")
    protected User user;
	
	private String salary;
	private String position;
	
	public SalaryComment() {}

	public SalaryComment(Integer id, Company company, User user, String salary, String position) {
		this.id = id;
		this.company = company;
		this.user = user;
		this.salary = salary;
		this.position = position;
	}
	
	public SalaryComment(Company company, User user, SalaryCommentCreationDTO salaryCommentDTO) {
		this.company = company;
		this.salary = salaryCommentDTO.getSalary();
		this.user = user;
		this.position = salaryCommentDTO.getPosition();
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

	public User getUser() {
		return user;
	}

	public void setUser(User user) {
		this.user = user;
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
