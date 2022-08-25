package com.agentska.model;

import javax.persistence.*;

import com.agentska.dto.CompanyDTO;
import com.agentska.dto.CompanyRegisterDTO;

@Entity
@Table(name = "company")
public class Company {
	
	@Id
	@SequenceGenerator(name = "company_sequence_generator", sequenceName = "company_sequence", initialValue = 100)

	@GeneratedValue(strategy = GenerationType.IDENTITY,generator = "company_sequence_generator")
	protected Integer id;
	
	@ManyToOne
    @JoinColumn(name = "owner_id")
    protected User owner;
	
	private String name;
	@Column(length=4095)
	private String description;
	@Column(length=100)
	private String contactInfo;
	
	private boolean validated = false;
	
	public Company() {}
	
	public Company(Integer id, User owner, String name, String description, String contactInfo, boolean validated) {
		this.id = id;
		this.owner = owner;
		this.name = name;
		this.description = description;
		this.contactInfo = contactInfo;
		this.validated = validated;
	}
	
	public Company(CompanyRegisterDTO companyDTO, User owner) {
		this.owner = owner;
		this.name = companyDTO.getName();
		this.description = companyDTO.getDescription();
		this.contactInfo = companyDTO.getContactInfo();
	}
	
	public Integer getId() {
		return id;
	}
	public void setId(Integer id) {
		this.id = id;
	}
	public User getOwner() {
		return owner;
	}
	public void setOwner(User owner) {
		this.owner = owner;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getDescription() {
		return description;
	}
	public void setDescription(String description) {
		this.description = description;
	}
	public String getContactInfo() {
		return contactInfo;
	}
	public void setContactInfo(String contactInfo) {
		this.contactInfo = contactInfo;
	}
	public boolean isValidated() {
		return validated;
	}
	public void setValidated(boolean validated) {
		this.validated = validated;
	}
	
}
