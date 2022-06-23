package api

/*func mapUser(user *domain.User) *pb.User {
	postPb := &pb.Post{
		Id:       user.Id.Hex(),
		Date:     timestamppb.New(user.Date),
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
*/
