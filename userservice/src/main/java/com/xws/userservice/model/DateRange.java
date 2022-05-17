package com.xws.userservice.model;

import java.time.LocalDateTime;

import javax.persistence.Embeddable;

@Embeddable
public class DateRange {
	private final LocalDateTime fromDate;
	private final LocalDateTime toDate;
	
	public DateRange() {
		//No validation for now
		fromDate = null;
		toDate = null;
	}
	
	public DateRange(LocalDateTime from, LocalDateTime to) throws Exception {
		this.fromDate = from;
		this.toDate = to;
		
		validate();
	}
	
	void validate() throws Exception {
		if (fromDate.isAfter(toDate)) {
			throw new Exception("Invalid DateRange");
		}
	}

	public LocalDateTime getFromDate() {
		return fromDate;
	}

	public LocalDateTime getToDate() {
		return toDate;
	}
	
	public static boolean dateRangesOverlap(DateRange dateRange1, DateRange dateRange2) {
		if(dateRange1.getToDate().isBefore(dateRange2.getFromDate()) || dateRange1.getFromDate().isAfter(dateRange2.getToDate()))
			return false;
		return true;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((fromDate == null) ? 0 : fromDate.hashCode());
		result = prime * result + ((toDate == null) ? 0 : toDate.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		DateRange other = (DateRange) obj;
		if (fromDate == null) {
			if (other.fromDate != null)
				return false;
		} else if (!fromDate.equals(other.fromDate))
			return false;
		if (toDate == null) {
			if (other.toDate != null)
				return false;
		} else if (!toDate.equals(other.toDate))
			return false;
		return true;
	}
}
