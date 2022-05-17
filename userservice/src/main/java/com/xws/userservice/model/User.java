package com.xws.userservice.model;

import java.time.LocalDateTime;
import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.OneToMany;
import javax.persistence.Table;
import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.xws.userservice.enums.EGender;
import com.xws.userservice.enums.EPrivacyLevel;

import javax.persistence.JoinColumn;


@Entity
@Table(name="\"users\"")
public class User {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Integer id;
	
	@NotBlank
	@Column(unique=true)
	@Size(max = 50)
	@Email
	String email;
	
	@NotBlank
	@Column(unique=true)
	@Size(max = 50)
	String username;
	
	@NotBlank
	@Size(max = 120)
	String password;
	
	@NotBlank
	@Size(max = 120)
	String firstName;
	String lastName;
	String contactPhone;
	LocalDateTime birthDate;
	EGender gender;
	@Column(length=4095)
	private String biography;
	
	@OneToMany(cascade = CascadeType.DETACH, fetch = FetchType.EAGER)
    @JoinColumn(name = "user_id")
    private Set<Experience> experiences = new HashSet<>();
	
	@ManyToMany(fetch = FetchType.LAZY)
	@JoinTable(	name = "blocked_users", 
				joinColumns = @JoinColumn(name = "blocker_id"), 
				inverseJoinColumns = @JoinColumn(name = "blocked_id"))
	private Set<User> blockedUsers = new HashSet<>();
	EPrivacyLevel privacy;
	
	@ManyToMany(fetch = FetchType.LAZY)
	@JoinTable(	name = "user_roles", 
				joinColumns = @JoinColumn(name = "user_id"), 
				inverseJoinColumns = @JoinColumn(name = "role_id"))
	private Set<Role> roles = new HashSet<>();
	boolean activated = false;
	
	public User() { }
	
	public Integer getId() {
		return id;
	}

	public void setId(Integer id) {
		this.id = id;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	@JsonIgnore
	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getLastName() {
		return lastName;
	}

	public void setLastName(String lastName) {
		this.lastName = lastName;
	}

	public String getContactPhone() {
		return contactPhone;
	}

	public void setContactPhone(String contactPhone) {
		this.contactPhone = contactPhone;
	}

	public boolean isActivated() {
		return activated;
	}

	public void setActivated(boolean activated) {
		this.activated = activated;
	}
	
	public Set<Role> getRoles() {
		return roles;
	}

	public void setRoles(Set<Role> userRoles) {
		this.roles = userRoles;
	}

	public String getBiography() {
		return biography;
	}

	public void setBiography(String biography) {
		this.biography = biography;
	}

	public String getUsername() {
		return username;
	}

	public LocalDateTime getBirthDate() {
		return birthDate;
	}

	public EGender getGender() {
		return gender;
	}

	public Set<Experience> getExperiences() {
		return experiences;
	}

	public Set<User> getBlockedUsers() {
		return blockedUsers;
	}

	public EPrivacyLevel getPrivacy() {
		return privacy;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public void setBirthDate(LocalDateTime birthDate) {
		this.birthDate = birthDate;
	}

	public void setGender(EGender gender) {
		this.gender = gender;
	}

	public void setExperiences(Set<Experience> experiences) {
		this.experiences = experiences;
	}

	public void setBlockedUsers(Set<User> blockedUsers) {
		this.blockedUsers = blockedUsers;
	}

	public void setPrivacy(EPrivacyLevel privacy) {
		this.privacy = privacy;
	}
}
