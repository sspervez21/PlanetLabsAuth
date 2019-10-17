package app

// TODO: this data model will not hold up to a scalable FE service.
//       all data is in memory, we would need to find a way to cache this perhaps.

type userData struct {
	firstName string
	lastName  string
	userID    string
	groups    map[string]bool //using map for fast lookups
}

// userId -> userData
var userMap = make(map[string]*userData)

// groupName -> userID -> *userData OR groupMap[groupName][userID] -> *userData
var groupMap = make(map[string]map[string]*userData)

func createGroup(groupName string) (bool, string) {
	_, ok := groupMap[groupName]

	if ok {
		return false, "The group already exists."
	}

	groupMap[groupName] = make(map[string]*userData)
	return true, ""
}

func createUser(user *userData, groups []string) (bool, string) {
	// make suer user does not already exist
	_, ok := userMap[user.userID]
	if ok {
		return false, "This user already exists."
	}

	// check that all groups exist
	for _, groupName := range groups {
		if _, ok := groupMap[groupName]; !ok {
			// group not found
			return false, "Group " + groupName + " not found."
		}
	}

	// all groups exist, we should add the user
	// TODO: make the following two steps transactional
	for _, groupName := range groups {
		addUserToGroup(groupName, user)
		addGroupToUser(groupName, user)
	}

	userMap[user.userID] = user
	return true, ""
}

func deleteGroup(groupName string) (bool, string) {
	groupUserMap, ok := groupMap[groupName]

	if !ok {
		return false, "group does not exist."
	}

	// for each user that links to this group, remove the group
	for _, user := range groupUserMap {
		removeGroupFromUser(groupName, user)
	}

	delete(groupMap, groupName)
	return true, ""
}

func deleteUser(userID string) (bool, string) {
	user, ok := userMap[userID]

	if !ok {
		return false, "user does not exist"
	}

	// Remove the user from all groups they are a part of
	for groupName := range user.groups {
		removeUserFromGroup(groupName, userID)
	}

	delete(userMap, userID)
	return true, ""
}

func getGroup(groupName string) ([]string, bool) {
	userMap, ok := groupMap[groupName]

	if !ok {
		return nil, false
	}

	ret := []string{}

	for userID := range userMap {
		ret = append(ret, userID)
	}

	return ret, true
}

func getUser(userID string) (*userData, bool) {
	val, ok := userMap[userID]

	if !ok {
		return nil, false
	}

	return val, true
}

func updateGroup(groupName string, newUserList []string) (bool, string) {
	groupUserMap, ok := groupMap[groupName]

	if !ok {
		return false, "The specified group does not exist."
	}

	for _, newUserID := range newUserList {
		_, ok := userMap[newUserID]
		if !ok {
			return false, "The user " + newUserID + " does not exist, could not perform update."
		}
	}

	removedUsers, addedUsers := getUserDifference(newUserList, groupUserMap)

	for _, removedUserID := range removedUsers {
		removeGroupFromUser(groupName, userMap[removedUserID])
		removeUserFromGroup(groupName, removedUserID)
	}

	for _, addedUserID := range addedUsers {
		addGroupToUser(groupName, userMap[addedUserID])
		addUserToGroup(groupName, userMap[addedUserID])
	}

	return true, ""
}

func updateUser(newUser *userData) (bool, string) {
	existingUser, ok := userMap[newUser.userID]

	if !ok {
		return false, "This user does not exist."
	}

	// verify that all new groups actually exist
	for newGroup := range newUser.groups {
		_, ok := groupMap[newGroup]
		if !ok {
			return false, "The group " + newGroup + " does not exist, could not perform update."
		}
	}

	existingUser.firstName = newUser.firstName
	existingUser.lastName = newUser.lastName
	// existingUser.userID = newUser.userID

	removedGroups, addedGroups := getGroupDifference(newUser.groups, existingUser.groups)

	for _, removedGroup := range removedGroups {
		removeGroupFromUser(removedGroup, existingUser)
		removeUserFromGroup(removedGroup, existingUser.userID)
	}

	for _, addedGroup := range addedGroups {
		addGroupToUser(addedGroup, existingUser)
		addUserToGroup(addedGroup, existingUser)
	}

	return true, ""
}

// utility functions
func addUserToGroup(groupName string, user *userData) {
	groupMap[groupName][user.userID] = user
}

func addGroupToUser(groupName string, user *userData) {
	user.groups[groupName] = true
}

func removeGroupFromUser(groupName string, user *userData) {
	_, ok := user.groups[groupName]
	if !ok {
		// TODO: raise error
	}
	delete(user.groups, groupName)
}

func removeUserFromGroup(groupName string, userID string) {
	groupUserMap, ok := groupMap[groupName]
	if !ok {
		// TODO: raise error
	}
	delete(groupUserMap, userID)
}

func getGroupDifference(newGroupMap map[string]bool, existingGroupUserMap map[string]bool) ([]string, []string) {
	return getDifference(newGroupMap, existingGroupUserMap)
}

func getUserDifference(newUsers []string, existingUsers map[string]*userData) ([]string, []string) {
	newUserMap := make(map[string]bool)
	existingUserMap := make(map[string]bool)

	for _, newUser := range newUsers {
		newUserMap[newUser] = true
	}

	for existingUser := range existingUsers {
		existingUserMap[existingUser] = true
	}

	return getDifference(newUserMap, existingUserMap)
}

func getDifference(newMap map[string]bool, existingMap map[string]bool) ([]string, []string) {
	var removedElmnts []string
	var addedElmnts []string

	for newElmnt := range newMap {
		_, ok := existingMap[newElmnt]
		if !ok {
			// The user does not already belong to this group, we should add them
			addedElmnts = append(addedElmnts, newElmnt)
		}
	}

	for existingElmnt := range existingMap {
		_, ok := newMap[existingElmnt]
		if !ok {
			// The user should be removed from this group
			removedElmnts = append(removedElmnts, existingElmnt)
		}
	}

	return removedElmnts, addedElmnts
}
