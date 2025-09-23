package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// HasChanges checks if there are remote changes available
func HasChanges(repo Repository) (bool, error) {
	// First, fetch the latest remote information
	fetchCmd := exec.Command("git", "fetch", "origin", repo.Branch)
	fetchCmd.Dir = repo.Path

	if output, err := fetchCmd.CombinedOutput(); err != nil {
		return false, fmt.Errorf("git fetch failed: %w\nOutput: %s", err, string(output))
	}

	// Get local commit hash
	localCmd := exec.Command("git", "rev-parse", "HEAD")
	localCmd.Dir = repo.Path
	localOutput, err := localCmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to get local commit: %w", err)
	}
	localCommit := strings.TrimSpace(string(localOutput))

	// Get remote commit hash
	remoteCmd := exec.Command("git", "rev-parse", fmt.Sprintf("origin/%s", repo.Branch))
	remoteCmd.Dir = repo.Path
	remoteOutput, err := remoteCmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to get remote commit: %w", err)
	}
	remoteCommit := strings.TrimSpace(string(remoteOutput))

	// Compare commits
	hasChanges := localCommit != remoteCommit

	if hasChanges {
		logger := NewLogger()
		logger.Debug("Local: %s, Remote: %s", localCommit[:8], remoteCommit[:8])
	}

	return hasChanges, nil
}

// PullChanges pulls the latest changes from the remote repository
func PullChanges(repo Repository) error {
	// Execute git pull
	pullCmd := exec.Command("git", "pull", "origin", repo.Branch)
	pullCmd.Dir = repo.Path

	output, err := pullCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git pull failed: %w\nOutput: %s", err, string(output))
	}

	// Check if the output indicates success
	outputStr := string(output)
	if strings.Contains(outputStr, "Already up to date") ||
	   strings.Contains(outputStr, "Already up-to-date") {
		// This shouldn't happen if HasChanges returned true, but handle it gracefully
		return nil
	}

	return nil
}

// Helper functions

func ensureDirectoryExists(path string) error {
	// Expand home directory if needed
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = strings.Replace(path, "~", homeDir, 1)
	}

	// Check if directory exists
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create the directory
		return os.MkdirAll(path, 0755)
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", path)
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}