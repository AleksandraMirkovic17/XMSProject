package com.agentska.dto;

public class CompanyDTO {
	private Integer id;
	private Integer ownerId;
	private String name;
	private String description;
	private String contactInfo;
	
	public CompanyDTO() {}
	
	public CompanyDTO(Integer id, String name, String description, String contactInfo, Integer ownerId) {
		this.ownerId = ownerId;
		this.id = id;
		this.name = name;
		this.description = description;
		this.contactInfo = contactInfo;
	}
	
	public Integer getOwnerId() {
		return ownerId;
	}
	public void setOwnerId(Integer ownerId) {
		this.ownerId = ownerId;
	}
	public Integer getId() {
		return id;
	}
	public void setId(Integer id) {
		this.id = id;
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
	
	
}
