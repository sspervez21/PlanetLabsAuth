package app

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const groupDataFile = "groupData"
const userDataFile = "userData"

// InitializeDataStore function
func InitializeDataStore() error {
	dataDir := os.Getenv("DATA_DIR")

	groupDataFileName := filepath.Join(dataDir, groupDataFile)
	userDataFileName := filepath.Join(dataDir, userDataFile)

	_, err := os.Stat(groupDataFileName)
	if err != nil {
		// We ignore missing files, the server should just initialize to the empty state.
		return nil
	}

	_, err = os.Stat(userDataFileName)
	if err != nil {
		// We ignore missing files, the server should just initialize to the empty state.
		return nil
	}

	allGroups, err := readFile(groupDataFileName)
	if err != nil {
		return err
	}

	allUsers, err := readFile(userDataFileName)
	if err != nil {
		return err
	}

	// create all groups
	groupMap = make(map[string]map[string]*userData)
	for _, groupStr := range allGroups {
		ok, msg := createGroup(groupStr)
		if !ok {
			return errors.New(msg)
		}
	}

	// create all users
	userMap = make(map[string]*userData)
	for _, userStr := range allUsers {
		tokens := strings.Fields(userStr)
		if len(tokens) < 3 {
			return errors.New("malformed user string. " + userStr)
		}

		user := &userData{firstName: tokens[0], lastName: tokens[1], userID: tokens[2], groups: make(map[string]bool)}

		ok, msg := createUser(user, tokens[3:])
		if !ok {
			return errors.New(msg)
		}
	}

	return nil
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ret := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret, nil
}

// PersistData function
func PersistData() error {

	dataDir := os.Getenv("DATA_DIR")

	allGroups := make([]string, 0)
	allUsers := make([]string, 0)

	for group := range groupMap {
		allGroups = append(allGroups, group)
	}

	for _, userData := range userMap {
		user := make([]string, 0)
		user = append(user, userData.firstName)
		user = append(user, userData.lastName)
		user = append(user, userData.userID)
		for group := range userData.groups {
			user = append(user, group)
		}
		allUsers = append(allUsers, strings.Join(user, " "))
	}

	groupDataFileName := filepath.Join(dataDir, groupDataFile)
	userDataFileName := filepath.Join(dataDir, userDataFile)

	err := writeToFile(groupDataFileName, allGroups)
	if err != nil {
		return err
	}

	err = writeToFile(userDataFileName, allUsers)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(fileName string, lines []string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
