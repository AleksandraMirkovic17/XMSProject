package com.agentska.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.agentska.dto.InterviewCommentCreationDTO;
import com.agentska.model.enums.EDifficulty;

@Entity
@Table(name = "interview_comment")
public class InterviewComment {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	protected Integer id;
	
	@ManyToOne
    @JoinColumn(name = "company_id")
    protected Company company;
	
	@ManyToOne
    @JoinColumn(name = "user_id")
    protected User user;
	
	private String text;
	
	private EDifficulty difficulty;
	
	private int rating;

	public InterviewComment() {}
	
	public InterviewComment(Integer id, Company company, User user, String text, EDifficulty difficulty, int rating) {
		this.id = id;
		this.company = company;
		this.user = user;
		this.text = text;
		this.difficulty = difficulty;
		this.rating = rating;
	}

	public InterviewComment(Company company, User user, InterviewCommentCreationDTO interviewCommentDTO) {
		this.company = company;
		this.text = interviewCommentDTO.getText();
		this.user = user;
		this.difficulty = EDifficulty.valueOf(interviewCommentDTO.getDifficulty());
		this.rating = interviewCommentDTO.getRating();
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

	public String getText() {
		return text;
	}

	public void setText(String text) {
		this.text = text;
	}
	
}
