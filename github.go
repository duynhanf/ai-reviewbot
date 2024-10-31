package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func getRequirementFromGithubPullRequest(owner string, repo string, pullRequestNumber string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pullRequestNumber, nil)

	if err != nil {
		panic(err)
	}

	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")
	req.Header.Set("Authorization", "Bearer "+GITHUB_TOKEN)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// Parse the response body
	var pullRequest map[string]interface{}
	err = json.Unmarshal(body, &pullRequest)
	if err != nil {
		return ""
	}

	if pullRequest["body"] == nil {
		return ""
	}

	return pullRequest["body"].(string)
}

func getDiffFromGithubPullRequest(owner string, repo string, pullRequestNumber string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pullRequestNumber+"/files", nil)

	if err != nil {
		panic(err)
	}

	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")
	req.Header.Set("Authorization", "Bearer "+GITHUB_TOKEN)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the response body
	var pullRequest []map[string]interface{}
	err = json.Unmarshal(body, &pullRequest)
	if err != nil {
		panic(err)
	}

	return pullRequest[0]["patch"].(string)

}
