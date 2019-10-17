package app

import (
	"sort"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	ok, msg := createGroup("admins")
	if !ok || msg != "" {
		t.Fatalf("Unexpected error: " + msg + "\n")
	}

	ok, msg = createGroup("admins")
	if ok || msg == "" {
		t.Fatalf("Expected an error while creating existing group but none was returned" + "\n")
	}

	userList, ok := getGroup("admins")

	if !testEq(userList, []string{}) {
		t.Fatalf("Expected empty list of users for group admins." + "\n")
	}

	// cleanup state
	_, _ = deleteGroup("admins")
}

func TestCreateUser(t *testing.T) {
	user := userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	ok, msg := createUser(&user, []string{"admins"})
	if ok || msg == "" {
		t.Fatalf("Should not be able to create a user in a non-existent group." + "\n")
	}

	ok, msg = createGroup("admins")
	userList, ok := getGroup("admins")

	if !testEq(userList, []string{}) {
		t.Fatalf("Expected empty user list for group admins." + "\n")
	}

	ok, msg = createUser(&user, []string{"admins"})
	if !ok || msg != "" {
		t.Fatalf("Unexpected create user error." + "\n")
	}

	userList, ok = getGroup("admins")
	if !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected user jsmith in list of users for group admins." + "\n")
	}

	ok, msg = createUser(&user, []string{"admins"})
	if ok || msg == "" {
		t.Fatalf("Expected error trying to create existing user." + "\n")
	}

	// cleanup state
	_, _ = deleteGroup("admins")
	_, _ = deleteUser("jsmith")
}

func TestDeleteGroup(t *testing.T) {
	_, _ = createGroup("admins")
	_, _ = createGroup("users")

	_, ok := getGroup("admins")
	if !ok {
		t.Fatalf("Could not access group admins." + "\n")
	}

	_, ok = getGroup("users")
	if !ok {
		t.Fatalf("Could not access group users." + "\n")
	}

	user1 := userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	_, _ = createUser(&user1, []string{"admins", "users"})

	ok, msg := deleteGroup("admins")

	if !ok || msg != "" {
		t.Fatalf("Unexpected error while deleting group: " + msg + "\n")
	}

	_, ok = getGroup("admins")
	if ok {
		t.Fatalf("Should not be able to access deleted group." + "\n")
	}

	user, ok := getUser("jsmith")

	// The user should still be part of the group users but no longer part of the group admins
	if !ok || len(user.groups) != 1 || user.groups["admins"] || !user.groups["users"] {
		t.Fatalf("User data inconsistent after delte group." + "\n")
	}

	// cleanup
	_, _ = deleteUser("jsmith")
	_, _ = deleteGroup("users")
}

func TestDeleteUser(t *testing.T) {
	user1 := userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	user2 := userData{
		firstName: "Sally",
		lastName:  "Smith",
		userID:    "ssmith",
		groups:    make(map[string]bool),
	}

	_, _ = createGroup("admins")
	_, _ = createUser(&user1, []string{"admins"})
	_, _ = createUser(&user2, []string{"admins"})

	userList, ok := getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith"}) {
		t.Fatalf("Expected user jsmith and ssmith in list of users for group admins." + "\n")
	}

	_, ok = getUser("jsmith")
	if !ok {
		t.Fatalf("Expected user jsmith to exist." + "\n")
	}

	_, ok = getUser("ssmith")
	if !ok {
		t.Fatalf("Expected user ssmith to exist." + "\n")
	}

	ok, msg := deleteUser("ssmith")
	if !ok || msg != "" {
		t.Fatalf("Unexpected error deleting user ssmith. " + msg + "\n")
	}

	userList, ok = getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected user jsmith in list of users for group admins." + "\n")
	}

	_, ok = getUser("jsmith")
	if !ok {
		t.Fatalf("Expected user jsmith to exist." + "\n")
	}

	_, ok = getUser("ssmith")
	if ok {
		t.Fatalf("Expected user ssmith to not exist." + "\n")
	}

	// cleanup state
	_, _ = deleteGroup("admins")
	_, _ = deleteUser("jsmith")
}

func TestGetGroup(t *testing.T) {
	_, _ = createGroup("admins")
	userList, ok := getGroup("admins")

	if !ok || !testEq(userList, []string{}) {
		t.Fatalf("Expected empty list of users for group admins." + "\n")
	}

	// cleanup state
	_, _ = deleteGroup("admins")
}

func TestGetUser(t *testing.T) {
	user, ok := getUser("jsmith")

	if user != nil || ok {
		t.Fatalf("Expected error while getting non-existent user." + "\n")
	}

	user = &userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	_, _ = createGroup("admins")
	_, _ = createUser(user, []string{"admins"})

	user, ok = getUser("jsmith")
	if !ok || user.firstName != "Joe" || user.lastName != "Smith" || user.userID != "jsmith" || len(user.groups) != 1 || !user.groups["admins"] {
		t.Fatalf("Unexpected error getting user.")
	}

	// cleanup state
	_, _ = deleteUser("jsmith")
}

func TestUpdateUser(t *testing.T) {
	user1 := userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	user2 := userData{
		firstName: "Sally",
		lastName:  "Smith",
		userID:    "ssmith",
		groups:    make(map[string]bool),
	}

	_, _ = createGroup("admins")
	_, _ = createGroup("users")
	_, _ = createGroup("frisbeePlayers")
	_, _ = createUser(&user1, []string{"admins", "frisbeePlayers"})
	_, _ = createUser(&user2, []string{"users", "frisbeePlayers"})

	user, ok := getUser("jsmith")
	if !ok || user.firstName != "Joe" || user.lastName != "Smith" || user.userID != "jsmith" || len(user.groups) != 2 || !user.groups["admins"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user jsmith.")
	}

	user, ok = getUser("ssmith")
	if !ok || user.firstName != "Sally" || user.lastName != "Smith" || user.userID != "ssmith" || len(user.groups) != 2 || !user.groups["users"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user ssmith.")
	}

	userList, ok := getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected jsmith in list of users for group admins." + "\n")
	}

	userList, ok = getGroup("users")
	if !ok || !testEq(userList, []string{"ssmith"}) {
		t.Fatalf("Expected ssmith in list of users for group users." + "\n")
	}

	userList, ok = getGroup("frisbeePlayers")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith"}) {
		t.Fatalf("Expected jsmith and ssmith in list of users for group frisbeePlayers." + "\n")
	}

	user3 := userData{
		firstName: "Jack",
		lastName:  "Reacher",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}
	user3.groups = map[string]bool{"admins": true, "users": true}

	ok, msg := updateUser(&user3)
	if !ok || msg != "" {
		t.Fatalf("Unexpected error while updating user jsmith. " + msg + "\n")
	}

	user, ok = getUser("jsmith")
	if !ok || user.firstName != "Jack" || user.lastName != "Reacher" || user.userID != "jsmith" || len(user.groups) != 2 || !user.groups["admins"] || !user.groups["users"] {
		t.Fatalf("Unexpected error validating user jsmith.")
	}

	user, ok = getUser("ssmith")
	if !ok || user.firstName != "Sally" || user.lastName != "Smith" || user.userID != "ssmith" || len(user.groups) != 2 || !user.groups["users"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user ssmith.")
	}

	userList, ok = getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected jsmith in list of users for group admins." + "\n")
	}

	userList, ok = getGroup("users")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith"}) {
		t.Fatalf("Expected jsmith and ssmith in list of users for group users." + "\n")
	}

	userList, ok = getGroup("frisbeePlayers")
	if !ok || !testEq(userList, []string{"ssmith"}) {
		t.Fatalf("Expected ssmith in list of users for group frisbeePlayers." + "\n")
	}

	//cleanup
	_, _ = deleteUser("jsmith")
	_, _ = deleteUser("ssmith")
	_, _ = deleteGroup("admins")
	_, _ = deleteGroup("users")
	_, _ = deleteGroup("frisbeePlayers")
}

func TestUpdateGroup(t *testing.T) {
	user1 := userData{
		firstName: "Joe",
		lastName:  "Smith",
		userID:    "jsmith",
		groups:    make(map[string]bool),
	}

	user2 := userData{
		firstName: "Sally",
		lastName:  "Smith",
		userID:    "ssmith",
		groups:    make(map[string]bool),
	}

	user3 := userData{
		firstName: "Salman",
		lastName:  "Pervez",
		userID:    "spervez",
		groups:    make(map[string]bool),
	}

	_, _ = createGroup("admins")
	_, _ = createGroup("users")
	_, _ = createGroup("frisbeePlayers")
	_, _ = createUser(&user1, []string{"admins", "frisbeePlayers"})
	_, _ = createUser(&user2, []string{"users", "frisbeePlayers"})
	_, _ = createUser(&user3, []string{"users", "frisbeePlayers"})

	user, ok := getUser("jsmith")
	if !ok || user.firstName != "Joe" || user.lastName != "Smith" || user.userID != "jsmith" || len(user.groups) != 2 || !user.groups["admins"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user jsmith.")
	}

	user, ok = getUser("ssmith")
	if !ok || user.firstName != "Sally" || user.lastName != "Smith" || user.userID != "ssmith" || len(user.groups) != 2 || !user.groups["users"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user ssmith.")
	}

	user, ok = getUser("spervez")
	if !ok || user.firstName != "Salman" || user.lastName != "Pervez" || user.userID != "spervez" || len(user.groups) != 2 || !user.groups["users"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user spervez.")
	}

	userList, ok := getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected jsmith in list of users for group admins." + "\n")
	}

	userList, ok = getGroup("users")
	if !ok || !testEq(userList, []string{"ssmith", "spervez"}) {
		t.Fatalf("Expected ssmith and spervez in list of users for group users." + "\n")
	}

	userList, ok = getGroup("frisbeePlayers")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith", "spervez"}) {
		t.Fatalf("Expected jsmith, ssmith and spervez in list of users for group frisbeePlayers." + "\n")
	}

	// {"ssmith", "spervez"} -> "jsmith", "ssmith"}
	ok, msg := updateGroup("users", []string{"jsmith", "ssmith"})
	if !ok || msg != "" {
		t.Fatalf("Unexpected error while updating group users. " + msg + "\n")
	}

	user, ok = getUser("jsmith")
	if !ok || user.firstName != "Joe" || user.lastName != "Smith" || user.userID != "jsmith" || len(user.groups) != 3 || !user.groups["admins"] || !user.groups["frisbeePlayers"] || !user.groups["users"] {
		t.Fatalf("Unexpected error validating user jsmith.")
	}

	user, ok = getUser("ssmith")
	if !ok || user.firstName != "Sally" || user.lastName != "Smith" || user.userID != "ssmith" || len(user.groups) != 2 || !user.groups["users"] || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user ssmith.")
	}

	user, ok = getUser("spervez")
	if !ok || user.firstName != "Salman" || user.lastName != "Pervez" || user.userID != "spervez" || len(user.groups) != 1 || !user.groups["frisbeePlayers"] {
		t.Fatalf("Unexpected error validating user spervez.")
	}

	userList, ok = getGroup("admins")
	if !ok || !testEq(userList, []string{"jsmith"}) {
		t.Fatalf("Expected jsmith in list of users for group admins." + "\n")
	}

	userList, ok = getGroup("users")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith"}) {
		t.Fatalf("Expected jsmith and ssmith in list of users for group users." + "\n")
	}

	userList, ok = getGroup("frisbeePlayers")
	if !ok || !testEq(userList, []string{"jsmith", "ssmith", "spervez"}) {
		t.Fatalf("Expected jsmith, ssmith and spervez in list of users for group frisbeePlayers." + "\n")
	}

	//cleanup
	_, _ = deleteUser("jsmith")
	_, _ = deleteUser("ssmith")
	_, _ = deleteUser("spervez")
	_, _ = deleteGroup("admins")
	_, _ = deleteGroup("users")
	_, _ = deleteGroup("frisbeePlayers")
}

func testEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
