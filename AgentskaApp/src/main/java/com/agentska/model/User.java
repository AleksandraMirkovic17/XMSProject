package com.agentska.model;
import java.util.HashSet;
import java.util.Set;

import javax.persistence.*;

import com.agentska.dto.UserDTO;
import com.agentska.model.Role;
import com.agentska.model.enums.ERole;

@Entity
@Table(name = "user_table")
public class User {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	protected Integer id;
	
	@Column(name = "username")
	private String username;
	@Column(name = "email")
	private String email;
	@Column(name = "password")
	private String password;
	
	@ManyToMany(fetch = FetchType.LAZY)
	@JoinTable(	name = "user_roles", 
				joinColumns = @JoinColumn(name = "user_id"), 
				inverseJoinColumns = @JoinColumn(name = "role_id"))
	private Set<Role> roles = new HashSet<>();
	
	boolean activated = false;
	
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
	
	public Integer getId() {
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
	public Set<Role> getRoles() {
		return roles;
	}
	public void setRoles(Set<Role> roles) {
		this.roles = roles;
	}
	public boolean isActivated() {
		return activated;
	}
	public void setActivated(boolean activated) {
		this.activated = activated;
	}
	public boolean isOwner() {
		for (Role r : roles) {
		if (r.getName() == ERole.OWNER)
			return true;
		}
		return false;
	}

	@Override
	public String toString() {
		return "User [id=" + id + ", username=" + username + ", email=" + email + ", password=" + password + "]";
	}
}