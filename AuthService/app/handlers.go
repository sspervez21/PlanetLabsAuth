package app

import (
	"PlanetLabs/AuthService/models"
	"PlanetLabs/AuthService/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// CreateGroup function
func CreateGroup(params operations.CreateGroupParams) middleware.Responder {
	// TODO: Create admins group by default?
	ok, msg := createGroup(*params.CreateGroupInput.Name)

	if !ok {
		return operations.NewCreateGroupConflict().WithPayload(&models.BadRequest{
			Code:    int64(operations.CreateGroupConflictCode),
			Message: msg,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewCreateGroupOK()
}

// CreateUser function
func CreateUser(params operations.CreateUserParams) middleware.Responder {
	user := userData{
		firstName: *params.CreateUserInput.FirstName,
		lastName:  *params.CreateUserInput.LastName,
		userID:    *params.CreateUserInput.UserID,
		groups:    make(map[string]bool),
	}

	ok, errorMessage := createUser(&user, params.CreateUserInput.Groups)

	if !ok {
		return operations.NewCreateUserConflict().WithPayload(&models.BadRequest{
			Code:    int64(operations.CreateUserConflictCode),
			Message: errorMessage,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewCreateUserOK()
}

// DeleteGroup function
func DeleteGroup(params operations.DeleteGroupParams) middleware.Responder {
	//TODO: Disallow deletion of admins group?
	ok, msg := deleteGroup(params.GroupName)

	if !ok {
		return operations.NewDeleteGroupNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.DeleteGroupNotFoundCode),
			Message: msg,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewDeleteGroupOK()
}

// DeleteUser function
func DeleteUser(params operations.DeleteUserParams) middleware.Responder {
	ok, msg := deleteUser(params.UserID)

	if !ok {
		return operations.NewDeleteUserNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.DeleteUserNotFoundCode),
			Message: msg,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewDeleteUserOK()
}

// GetGroup function
func GetGroup(params operations.GetGroupParams) middleware.Responder {
	userList, ok := getGroup(params.GroupName)

	if !ok {
		return operations.NewGetGroupNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.GetGroupNotFoundCode),
			Message: "This group does not exist.",
		})
	}

	return operations.NewGetGroupOK().WithPayload(userList)
}

// GetUser function
func GetUser(params operations.GetUserParams) middleware.Responder {
	user, ok := getUser(params.UserID)

	if !ok {
		return operations.NewGetUserNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.GetUserNotFoundCode),
			Message: "This user does not exist.",
		})
	}

	return operations.NewGetUserOK().WithPayload(makeUserRecord(user))
}

// UpdateGroup function
func UpdateGroup(params operations.UpdateGroupParams) middleware.Responder {
	//TODO: Disallow updates to admins group?
	ok, msg := updateGroup(params.GroupName, params.UpdateGroupInput.List)

	if !ok {
		return operations.NewUpdateGroupNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.UpdateGroupNotFoundCode),
			Message: msg,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewUpdateGroupOK()
}

// UpdateUser function
func UpdateUser(params operations.UpdateUserParams) middleware.Responder {
	if params.UserID != *params.UpdateUserInput.UserID {
		return operations.NewUpdateUserNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.UpdateUserNotFoundCode),
			Message: "Expected UserID to match the one in user body.",
		})
	}

	user := userData{
		firstName: *params.UpdateUserInput.FirstName,
		lastName:  *params.UpdateUserInput.LastName,
		userID:    *params.UpdateUserInput.UserID,
		groups:    make(map[string]bool),
	}

	for _, group := range params.UpdateUserInput.Groups {
		user.groups[group] = true
	}

	ok, msg := updateUser(&user)

	if !ok {
		return operations.NewUpdateUserNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.UpdateUserNotFoundCode),
			Message: msg,
		})
	}

	err := PersistData()
	if err != nil {
		// TODO: log and panic...
	}

	return operations.NewUpdateUserOK()
}

// Helper functions
func makeUserRecord(user *userData) *models.UserRecord {
	appObject := &models.UserRecord{
		FirstName: &user.firstName,
		LastName:  &user.lastName,
		UserID:    &user.userID,
	}

	for element := range user.groups {
		// TODO: make append more performant
		appObject.Groups = append(appObject.Groups, element)
	}

	return appObject
}
