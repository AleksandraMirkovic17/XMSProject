package com.agentska.model;

import javax.persistence.*;

import com.agentska.dto.CommentCreationDTO;

@Entity
@Table(name = "comment")
public class Comment {
	@Id
	@SequenceGenerator(name = "comment_sequence_generator", sequenceName = "comment_sequence", initialValue = 100)

	@GeneratedValue(strategy = GenerationType.IDENTITY,generator="comment_sequence_generator")
	protected Integer id;
	
	@ManyToOne
    @JoinColumn(name = "company_id")
    protected Company company;
	
	@ManyToOne
    @JoinColumn(name = "user_id")
    protected User user;
	
	private String text;
	
	public Comment() {}

	public Comment(Integer id, Company company, String text, User user) {
		this.id = id;
		this.company = company;
		this.text = text;
		this.user = user;
	}
	
	public Comment(Company company, User user, CommentCreationDTO commentDTO) {
		this.company = company;
		this.text = commentDTO.getText();
		this.user = user;
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
