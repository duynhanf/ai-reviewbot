package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")
	pullRequestID := os.Getenv("GITHUB_PULL_REQUEST_ID")
	requirement := getRequirementFromGithubPullRequest(owner, repo, pullRequestID)
	reviewContent := getDiffFromGithubPullRequest(owner, repo, pullRequestID)

	getReviewFromOpenAI(requirement, reviewContent)
}
