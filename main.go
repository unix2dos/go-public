package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const (
	initialExperimentRatio = 0.5
	updatedExperimentRatio = 0.7
)

func main() {
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9", "user10"}

	for _, user := range users {
		initialGroup := assignGroup(user, initialExperimentRatio)
		updatedGroup := assignGroup(user, updatedExperimentRatio)

		fmt.Printf("User: %s, Initial Group: %s, Updated Group: %s\n", user, initialGroup, updatedGroup)
	}
}

func assignGroup(userID string, experimentRatio float64) string {
	hash := sha256.Sum256([]byte(userID))
	hashValue := binary.BigEndian.Uint64(hash[:8])

	// Normalize the hash value to a [0, 1) range
	normalizedValue := float64(hashValue) / float64(^uint64(0))

	if normalizedValue < experimentRatio {
		return "Experiment"
	} else {
		return "Control"
	}
}
