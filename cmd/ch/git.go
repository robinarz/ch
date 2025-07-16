package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func commit() error {
	// Run the interactive mode to get the commit message.
	commitMsg, err := runInteractiveMode()
	if err != nil {
		// This error is handled in main.go, just pass it up.
		return err
	}

	// Open the git repository in the current directory.
	repo, err := git.PlainOpen(".")
	if err != nil {
		return fmt.Errorf("error opening repository: %w", err)
	}

	// 3. Get the working tree.
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("error getting worktree: %w", err)
	}

	// 4. Stage all changes (equivalent to `git add .`).
	fmt.Fprintln(os.Stderr, faintStyle.Render("\nStaging files..."))
	_, err = worktree.Add(".")
	if err != nil {
		return fmt.Errorf("error staging files: %w", err)
	}

	// Get the current status.
	status, err := worktree.Status()
	if err != nil {
		return fmt.Errorf("error getting status: %w", err)
	}
	if status.IsClean() {
		fmt.Fprintln(os.Stderr, errorStyle.Render("No changes to commit."))
		return nil
	}

	// Get user signature from git config
	author, err := getAuthorSignature()
	if err != nil {
		// Fallback if no user is configured locally or globally
		fmt.Fprintf(os.Stderr, "Warning: Could not get git author info: %v. Using default.\n", err)
		author = &object.Signature{
			Name:  "Unknown Author",
			Email: "author@example.com",
			When:  time.Now(),
		}
	}

	// Create the commit.
	fmt.Fprintln(os.Stderr, faintStyle.Render("Creating commit..."))
	commit, err := worktree.Commit(commitMsg, &git.CommitOptions{
		Author: author,
	})
	if err != nil {
		return fmt.Errorf("error creating commit: %w", err)
	}

	// Print the commit hash.
	obj, err := repo.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("error getting commit object: %w", err)
	}

	fmt.Fprintln(os.Stderr, successStyle.Render("\n✅ Commit successful!"), faintStyle.Render(obj.Hash.String()[:7]))
	return nil
}

// getAuthorSignature retrieves the user's name and email from the git config,
// by loading and merging all standard git config files (system, global, local).
func getAuthorSignature() (*object.Signature, error) {
	// Get the user's name from git config.
	nameCmd := exec.Command("git", "config", "user.name")
	var nameOut bytes.Buffer
	nameCmd.Stdout = &nameOut
	if err := nameCmd.Run(); err != nil {
		return nil, fmt.Errorf("could not run 'git config user.name': %w", err)
	}
	name := strings.TrimSpace(nameOut.String())

	// Get the user's email from git config.
	emailCmd := exec.Command("git", "config", "user.email")
	var emailOut bytes.Buffer
	emailCmd.Stdout = &emailOut
	if err := emailCmd.Run(); err != nil {
		return nil, fmt.Errorf("could not run 'git config user.email': %w", err)
	}
	email := strings.TrimSpace(emailOut.String())

	// Check if the values are empty.
	if name == "" || email == "" {
		return nil, fmt.Errorf("user.name or user.email not set in git config")
	}

	return &object.Signature{
		Name:  name,
		Email: email,
		When:  time.Now(),
	}, nil
}
