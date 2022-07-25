package com.agentska.dto;

public class CompanyRegisterDTO {
	private String name;
	private String description;
	private String contactInfo;
	
	public CompanyRegisterDTO() {}
	
	public CompanyRegisterDTO(String name, String description, String contactInfo) {
		this.name = name;
		this.description = description;
		this.contactInfo = contactInfo;
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
