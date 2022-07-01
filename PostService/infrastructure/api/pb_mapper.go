package api

import (
	"PostService/domain"
	pb "github.com/dislinked/common/proto/post_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func mapPostFromDomainToPb(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		User:     post.User,
		Posttext: post.PostText,
		Date:     timestamppb.New(post.Date),
		Image:    post.Image,
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
			Id:       comment.Id.Hex(),
			Username: comment.User,
			Content:  comment.Content,
			Date:     timestamppb.New(comment.Date),
		})
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Id:        primitive.NewObjectID(),
		User:      postPb.User,
		PostText:  postPb.Posttext,
		Date:      time.Now(),
		IsDeleted: false,
		Image:     postPb.Image,
	}
	for _, link := range postPb.Links {
		post.Links = append(post.Links, link)
	}
	post.Comments = []domain.Comment{}
	for _, comment := range postPb.Comments {
		id, err := primitive.ObjectIDFromHex(comment.Id)
		if err != nil {
			continue
		}
		post.Comments = append(post.Comments, domain.Comment{
			Id:      id,
			Content: comment.Content,
			Date:    comment.Date.AsTime(),
			User:    comment.Username,
		})
	}

	return post
}

func mapReactionTypeToPb(reactionType domain.ReactionType) pb.Reaction_ReactionType {
	switch reactionType {
	case domain.LIKE:
		return pb.Reaction_LIKE
	case domain.DISLIKED:
		return pb.Reaction_DISLIKE
	}
	return pb.Reaction_LIKE
}
