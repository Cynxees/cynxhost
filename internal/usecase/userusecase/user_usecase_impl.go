package userusecase

import (
	"context"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"strconv"
)

type UserUseCaseImpl struct {
	tblUser    database.TblUser
	jwtManager *dependencies.JWTManager
}

func New(tblUser database.TblUser, jwtManager *dependencies.JWTManager) usecase.UserUseCase {
	return &UserUseCaseImpl{
		tblUser:    tblUser,
		jwtManager: jwtManager,
	}
}

func (usecase *UserUseCaseImpl) PaginateUser(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse) {
	_, users, err := usecase.tblUser.PaginateUser(ctx, req)
	if err != nil {
		resp.Code = responsecode.CodeTblUserError
		resp.Error = err.Error()
		return
	}

	var convertedUsers []responsedata.User
	for _, user := range users {
		convertedUsers = append(convertedUsers, responsedata.User{
			Id:          user.Id,
			Username:    user.Username,
			Coin:        user.Coin,
			CreatedDate: user.CreatedDate,
			UpdatedDate: user.UpdatedDate,
		})
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginateUserResponseData{
		Users: convertedUsers,
	}
}

func (usecase *UserUseCaseImpl) CheckUsername(ctx context.Context, req request.CheckUsernameRequest, resp *response.APIResponse) {
	_, exists, err := usecase.tblUser.CheckUserExists(ctx, "username", req.Username)
	if err != nil {
		resp.Code = responsecode.CodeTblUserError
		resp.Error = err.Error()
		return
	}

	if exists {
		resp.Code = responsecode.CodeNotAllowed
		return
	}

	resp.Code = responsecode.CodeSuccess
}

func (usecase *UserUseCaseImpl) RegisterUser(ctx context.Context, req request.RegisterUserRequest, resp *response.APIResponse) {
	_, id, err := usecase.tblUser.InsertUser(ctx, entity.TblUser{
		Username: req.Username,
		Password: req.Password,
		Coin:     0,
	})
	if err != nil {
		resp.Code = responsecode.CodeTblUserError
		resp.Error = err.Error()
		return
	}

	token, err := usecase.jwtManager.GenerateToken(id)
	if err != nil {
		resp.Code = responsecode.CodeInternalError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.AuthResponseData{
		AccessToken: token.AccessToken,
		TokenType:   "Bearer",
	}

}

func (usecase *UserUseCaseImpl) LoginUser(ctx context.Context, req request.LoginUserRequest, resp *response.APIResponse) {
	_, user, err := usecase.tblUser.GetUser(ctx, "username", req.Username)
	if err != nil {
		resp.Code = responsecode.CodeTblUserError
		resp.Error = err.Error()
		return
	}

	if user == nil {
		resp.Code = responsecode.CodeNotFound
		return
	}

	if user.Password != req.Password {
		resp.Code = responsecode.CodeAuthenticationError
		return
	}

	token, err := usecase.jwtManager.GenerateToken(user.Id)
	if err != nil {
		resp.Code = responsecode.CodeInternalError
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.AuthResponseData{
		AccessToken: token.AccessToken,
		TokenType:   "Bearer",
	}
}

func (usecase *UserUseCaseImpl) GetProfile(ctx context.Context, resp *response.APIResponse) {

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return
	}

	_, user, err := usecase.tblUser.GetUser(ctx, "id", strconv.Itoa(contextUser.Id))
	if err != nil {
		resp.Code = responsecode.CodeTblUserError
		resp.Error = err.Error()
		return
	}

	if user == nil {
		resp.Code = responsecode.CodeNotFound
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.User{
		Id:          user.Id,
		Username:    user.Username,
		Coin:        user.Coin,
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
	}
}
