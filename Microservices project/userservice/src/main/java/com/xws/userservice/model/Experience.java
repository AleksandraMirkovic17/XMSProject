package com.xws.userservice.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import com.xws.userservice.enums.EExperienceType;

@Entity
@Table(name="education_experiences")
public class Experience {
	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Integer id;
	
	EExperienceType type;
	String institutionName;
	DateRange dateRange;
	public Integer getId() {
		return id;
	}
	public EExperienceType getType() {
		return type;
	}
	public DateRange getDateRange() {
		return dateRange;
	}
	public void setId(Integer id) {
		this.id = id;
	}
	public void setType(EExperienceType type) {
		this.type = type;
	}
	public void setDateRange(DateRange dateRange) {
		this.dateRange = dateRange;
	}
	public String getInstitutionName() {
		return institutionName;
	}
	public void setInstitutionName(String institutionName) {
		this.institutionName = institutionName;
	}
	
}
