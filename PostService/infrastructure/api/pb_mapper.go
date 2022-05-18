package api

import (
	"PostService/domain"
	pb "github.com/dislinked/common/proto/post_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		User:     post.User,
		Posttext: post.PostText,
		Date:     timestamppb.New(post.Date),
	}
	for _, image := range post.Images {
		postPb.Image = append(postPb.Image, image)
	}
	for _, link := range post.Links {
		postPb.Links = append(postPb.Links, link)
	}
	for _, reaction := range post.Reactions {
		postPb.Reactions = append(postPb.Reactions, &pb.Reaction{
			Username:     reaction.User,
			ReactionType: mapReactionTypeToPb(reaction.Reaction),
		})
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Username: comment.User,
			Content:  comment.Content,
		})
	}
	return postPb
}

/*func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Id:         primitive.NewObjectID(),
		UserId:     postPb.UserId,
		PostText:   postPb.PostText,
		DatePosted: time.Now(),
	}
	for _, image := range postPb.ImagePaths {
		post.ImagePaths = append(post.ImagePaths, image)
	}
	for _, link := range postPb.Links {
		post.Links = append(post.Links, link)
	}

	return post
}

func mapNewComment(commentPb *pb.Comment) *domain.Comment {
	comment := &domain.Comment{
		Username:    commentPb.Username,
		CommentText: commentPb.CommentText,
	}

	return comment
}*/

func mapReactionTypeToPb(reactionType domain.ReactionType) pb.Reaction_ReactionType {
	switch reactionType {
	case domain.LIKE:
		return pb.Reaction_LIKE
	case domain.DISLIKED:
		return pb.Reaction_DISLIKE
	}
	return pb.Reaction_LIKE
}
