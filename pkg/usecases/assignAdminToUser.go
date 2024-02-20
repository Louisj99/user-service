package usecases

import "context"

type AssignAdminToUserRequest struct {
	UserId string `json:"userId"`
}
type AssignAdminToUserResponse struct {
	Message string `json:"message"`
}

type AssignAdminToUserInterface interface {
	AssignAdminToUser(ctx context.Context, userId string) error
}

func AssignAdminToUser(AssignAdminToUserInterface AssignAdminToUserInterface, userId string) error {
	err := AssignAdminToUserInterface.AssignAdminToUser(context.Background(), userId)
	if err != nil {
		return err
	}
	return nil
}
