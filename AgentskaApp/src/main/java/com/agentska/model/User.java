package com.agentska.model;
import javax.persistence.*;

import com.agentska.dto.UserDTO;

@Entity
@Table(name = "user_table")
public class User {
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private long id;
	
	@Column(name = "username")
	private String username;
	@Column(name = "email")
	private String email;
	@Column(name = "password")
	private String password;
	
	public User() {
	}
	
	public User(String username, String email, String password) {
		this.username = username;
		this.email = email;
		this.password = password;
	}
	
	public User(UserDTO userDTO) {
		this.username = userDTO.getUsername();
		this.email = userDTO.getEmail();
		this.password = userDTO.getPassword();
	}
	
	public long getId() {
		return id;
	}
	public String getUsername() {
		return username;
	}
	public void setUsername(String username) {
		this.username = username;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public String getPassword() {
		return password;
	}
	public void setPassword(String password) {
		this.password = password;
	}
	@Override
	public String toString() {
		return "User [id=" + id + ", username=" + username + ", email=" + email + ", password=" + password + "]";
	}
}