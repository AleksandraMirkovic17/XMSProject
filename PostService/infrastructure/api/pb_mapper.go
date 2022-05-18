package api

import (
	"PostService/domain"
	pb "github.com/dislinked/common/proto/post_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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

func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Id:        primitive.NewObjectID(),
		User:      postPb.User,
		PostText:  postPb.Posttext,
		Date:      time.Now(),
		IsDeleted: false,
	}
	for _, image := range postPb.Image {
		post.Images = append(post.Images, image)
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

/*

func mapPostToPb(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Username: post.Username,
		Content:  post.Content,
		Image:    post.Image,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Date:     timestamppb.New(post.Date),
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Id:       comment.Id.Hex(),
			Content:  comment.Content,
			Date:     timestamppb.New(comment.Date),
			Username: comment.Username,
		})
	}
	return postPb
}
*/

func inversePostMap(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		User:     post.User,
		Posttext: post.PostText,
		Image:    post.Images,
		Date:     timestamppb.New(post.Date),
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Id:       comment.Id.Hex(),
			Content:  comment.Content,
			Date:     timestamppb.New(comment.Date),
			Username: comment.User,
		})
	}
	/*for _, reaction := range post.Reactions {
		postPb.Reactions = append(postPb.Reactions, &pb.Reaction{
			Username:     reaction.User,
			ReactionType: reaction.Reaction,
		})
	}*/
	return postPb
}

/*func mapNewComment(commentPb *pb.Comment) *domain.Comment {
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
