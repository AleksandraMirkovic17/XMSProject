package com.agentska.dto;

public class InterviewCommentCreationDTO {
	private String text;
	private int rating;
	private String difficulty;

	public InterviewCommentCreationDTO() {}

	public InterviewCommentCreationDTO(String text, int rating, String difficulty) {
		this.text = text;
		this.rating = rating;
		this.difficulty = difficulty;
	}

	public String getText() {
		return text;
	}

	public void setText(String text) {
		this.text = text;
	}

	public int getRating() {
		return rating;
	}

	public void setRating(int rating) {
		this.rating = rating;
	}

	public String getDifficulty() {
		return difficulty;
	}

	public void setDifficulty(String difficulty) {
		this.difficulty = difficulty;
	}

}
