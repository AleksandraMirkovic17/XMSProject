package com.agentska.dto;

public class CommentCreationDTO {
	private String text;

	public CommentCreationDTO() {}

	public CommentCreationDTO(String text) {
		this.text = text;
	}

	public String getText() {
		return text;
	}

	public void setText(String text) {
		this.text = text;
	}
	
}