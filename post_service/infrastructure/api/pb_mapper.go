package api

import (
	"PostService/domain"
	"time"

	pb "github.com/dislinked/common/proto/post_service"
	events "github.com/dislinked/common/saga/create_notification"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func MapDomainNotificationToEventNotification(postDomain *domain.Post) *events.NotificationDetails {
	notification := &events.NotificationDetails{
		User:    postDomain.User,
		Content: "NOTIFICATION_CONTENT_NOT_GENERATED_ERR",
		Url:     "profile/" + postDomain.User,
		Seen:    false,
	}
	return notification
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
	post.Reactions = []*domain.Reaction{}
	for _, reaction := range postPb.Reactions {
		id, err := primitive.ObjectIDFromHex(reaction.PostId)
		if err != nil {
			continue
		}
		post.Reactions = append(post.Reactions, &domain.Reaction{
			Id:       id,
			User:     reaction.Username,
			Reaction: mapPbReactionTypeToDomain(reaction.ReactionType),
		})
	}

	post.Comments = []*domain.Comment{}
	for _, comment := range postPb.Comments {
		id, err := primitive.ObjectIDFromHex(comment.Id)
		if err != nil {
			continue
		}
		post.Comments = append(post.Comments, &domain.Comment{
			Id:      id,
			Content: comment.Content,
			Date:    time.Now(),
			User:    comment.Username,
		})
	}

	return post
}

func mapDomainCommentsToPbComments(comments []*domain.Comment) *pb.MultipleCommentsResponse {
	response := &pb.MultipleCommentsResponse{}
	for _, comment := range comments {
		response.Comment = append(response.Comment, &pb.Comment{
			Id:       comment.Id.Hex(),
			Content:  comment.Content,
			Date:     timestamppb.New(comment.Date),
			Username: comment.User,
		})
	}
	return response
}

func mapNewPbCommentToDomainComment(comment *pb.Comment) *domain.Comment {
	commentD := &domain.Comment{
		Id:      primitive.NewObjectID(),
		Content: comment.Content,
		Date:    comment.Date.AsTime(),
		User:    comment.Username,
	}
	return commentD
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

func mapPbReactionTypeToDomain(reactionType pb.Reaction_ReactionType) domain.ReactionType {
	switch reactionType {
	case pb.Reaction_LIKE:
		return domain.LIKE
	case pb.Reaction_DISLIKE:
		return domain.DISLIKED
	}
	return domain.LIKE
}
