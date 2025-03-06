package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GithubApiClient struct {
	Endpoint string
}

var (
	realGithubApiClient = GithubApiClient{
		Endpoint: "https://api.github.com",
	}
)

func CheckForUpdate(currentVersion string) (string, error) {
	return realGithubApiClient.CheckForUpdate(currentVersion)
}

func GetLatestVersion() (string, error) {
	return realGithubApiClient.GetLatestVersion()
}

func (apiClient *GithubApiClient) GetLatestVersion() (string, error) {
	resp, err := http.Get(apiClient.Endpoint + "/repos/KingJorjai/BONK/releases/latest")
	if err != nil {
		return "", fmt.Errorf("error getting the latest version: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response: %v", err)
	}

	var release struct {
		TagName string `json:"tag_name"` // The version tag (e.g., "v1.0.5")
	}
	if err := json.Unmarshal(body, &release); err != nil {
		return "", fmt.Errorf("error decoding the response: %v", err)
	}

	// Remove the "v" prefix if present
	version := strings.TrimPrefix(release.TagName, "v")
	return version, nil
}

func (apiClient *GithubApiClient) CheckForUpdate(currentVersion string) (string, error) {
	latestVersion, err := apiClient.GetLatestVersion()
	if err != nil {
		return "", err
	}
	if latestVersion != currentVersion {
		return latestVersion, nil
	}
	return "", fmt.Errorf("no update available")
}
