package main

import (
	"fmt"
	"kwx/pkg/awxclient"
	"log"
)

func main() {
	// Initialize the AWX client
	client := awxclient.NewClient("http://your-awx-instance.com", "your-token")

	// 1. Create a new credential
	newCredential := awxclient.Credential{
		Name:        "My New Credential",
		Description: "SSH Key for Production Server",
		Kind:        "ssh",
	}
	createdCredential, err := client.CreateCredential(newCredential)
	if err != nil {
		log.Fatalf("Error creating new credential: %v", err)
	}
	fmt.Printf("Created new credential with ID: %d\n", createdCredential.ID)

	// 2. Create a new project
	newProject := awxclient.Project{
		Name:        "My New Project",
		Description: "Repo containing deployment scripts",
		SCMType:     "git",
		SCMURL:      "https://github.com/your-repo.git",
	}
	createdProject, err := client.CreateProject(newProject)
	if err != nil {
		log.Fatalf("Error creating new project: %v", err)
	}
	fmt.Printf("Created new project with ID: %d\n", createdProject.ID)

	// 3. Fetch details of an inventory group by ID
	groupID := 123 // Replace with a group ID you're interested in
	group, err := client.GetGroup(groupID)
	if err != nil {
		log.Fatalf("Error fetching group details: %v", err)
	}
	fmt.Printf("Group Details - Name: %s, InventoryID: %d\n", group.Name, group.InventoryID)
}
