package api

import (
	"UserService/application"
	"UserService/infrastructure/mappers"
	orchestrators "UserService/infrastructure/orchestrators"
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service                *application.UserService
	updateUserOrchestrator *orchestrators.UpdateUserOrchestrator
}

func NewUserHandler(service *application.UserService, updateUserOrchestrator *orchestrators.UpdateUserOrchestrator) *UserHandler {
	return &UserHandler{service: service, updateUserOrchestrator: updateUserOrchestrator}
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetUserBySearchParamsRequest) (*pb.GetAllUserResponse, error) {
	println("getting all users")
	users, err := handler.service.GetAllByUsernameAndNameAndSurname(request.Username, request.Username, request.Username)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{Users: []*pb.User{}}
	for _, user := range users {
		current := mappers.MapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	println(request.Id)
	userId, _ := primitive.ObjectIDFromHex(request.Id)
	user, err := handler.service.GetOne(userId)
	if err != nil || user == nil {
		return nil, err
	}
	response := &pb.GetUserResponse{
		User: mappers.MapUser(user),
	}
	return response, nil
}

func (handler *UserHandler) Insert(ctx context.Context, request *pb.RegisterUserRequest) (*pb.User, error) {
	user := mappers.MapUserPbToDomain(request.User.User)
	fmt.Println("mapper zavrsio")

	err := handler.service.Insert(user)
	fmt.Println("Register zavrsio")
	if err != nil {
		return nil, err
	}

	return mappers.MapUser(user), nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	println("Starting to update")
	println("Before mapping", request.User.OldUser.User)
	user := mappers.MapUserPbToDomain(request.User.OldUser.User)
	foundUser, findErr := handler.service.FindByUsername((*user).Username)
	if findErr != nil {
		return nil, findErr
	}
	if foundUser == nil {
		println("found user is null")
		return nil, findErr
	}
	user.Id = foundUser.Id
	println("foun user id is:", user.Id.Hex())

	updateErr := handler.service.Update(user.Id, mappers.MapUserPbToDomain(request.User.User.User))
	if updateErr != nil {
		return nil, updateErr
	}

	//pozivanje sage prilikom update
	handler.updateUserOrchestrator.Start(mappers.MapPbUserToEventUpdateUser((*request.User.User.User)), mappers.MapPbUserToEventUpdateUser(*request.User.OldUser.User))

	return mappers.MapUser(user), nil
}

/*func (handler *UserHandler) SearchProfile(ctx context.Context, request *pb.SearchProfileRequest) (*pb.GetAllUserResponse, error) {

	users, err := handler.service.Search(request.GetParam())
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUserResponse{
		Users: []*pb.User{},
	}
	println("Name comming")
	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}*/

func (handler *UserHandler) FindByUsername(ctx context.Context, request *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	print("Searching fo user by username")
	println("Searching for: " + request.Username)
	user, err := handler.service.FindByUsername(request.Username)
	if err != nil || user == nil {
		return nil, err
	}
	response := &pb.GetUserResponse{
		User: mappers.MapUser(user),
	}
	return response, nil
}

func (handler *UserHandler) GenerateAPIToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.ApiToken, error) {
	user, err := handler.service.FindByUsername(request.Username.Username)
	if err != nil {
		return nil, err
	}
	token, err := handler.service.GenerateApiToken(user)
	if err != nil {
		println("Error is in generate token")
		println(err.Error())
		return nil, err
	}
	return &pb.ApiToken{
		ApiToken: token,
	}, nil

}

func (handler *UserHandler) ShareJobOffer(ctx context.Context, request *pb.ShareJobRequest) (*pb.ShareJobResponse, error) {
	token := request.ShareJob.ApiToken
	hasAccess, username, err := handler.service.CheckAccess(token)
	if err != nil {
		println("Impossible to check the access of the token:", err.Error())
		return &pb.ShareJobResponse{
			Job:   nil,
			Valid: false,
		}, err
	}
	if !hasAccess {
		return &pb.ShareJobResponse{
			Job:   nil,
			Valid: false,
		}, nil
	}
	user, err := handler.service.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	println("User id is:", user.Id.Hex())

	return &pb.ShareJobResponse{
		Job: &pb.Job{
			JobID:          "",
			PublisherId:    user.Id.Hex(),
			RequiredSkills: request.ShareJob.Job.RequiredSkills,
			DatePosted:     request.ShareJob.Job.DatePosted,
			DateValid:      request.ShareJob.Job.DateValid,
			CompanyName:    request.ShareJob.Job.CompanyName,
			Position:       request.ShareJob.Job.Position,
			JobDescription: request.ShareJob.Job.JobDescription,
		},
		Valid: true,
	}, nil

	/*postRequestBody, _ := json.Marshal(map[string]any{
		"jobID":          "",
		"publisherId":    user.Id.Hex(),
		"requiredSkills": request.ShareJob.Job.RequiredSkills,
		"datePosted":     request.ShareJob.Job.DatePosted,
		"dateValid":      request.ShareJob.Job.DateValid,
		"companyName":    request.ShareJob.Job.CompanyName,
		"position":       request.ShareJob.Job.Position,
		"jobDescription": request.ShareJob.Job.JobDescription,
	})

	jsonBody := []byte(postRequestBody)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%s/job")
	println("request url", requestURL)

	println("company name:", request.ShareJob.Job.CompanyName)

	/requestBodt := bytes.NewBuffer(postRequestBody)
	resp, err := http.Post("http://localhost:4200/job", "application/json", requestBodt)

	resp, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	print("response status code")
	println(resp)
	if err != nil {
		return &pb.ActionResult{
			Status: 400,
			Msg:    "Error while sending post request:" + err.Error(),
		}, nil
	}
	defer resp.Body.Close()
	return &pb.ActionResult{
		Status: 200,
		Msg:    "Request sent to post service! Status:" + resp.Response.Status,
	}, nil*/
}
