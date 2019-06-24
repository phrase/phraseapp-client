package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func usedBranchName(useLocalBranchNameFlag bool, branchParam string) (string, error) {
	if useLocalBranchName(useLocalBranchNameFlag) && branchParam == "" {
		gitBranch, gitErr := checkedOutGitBranch()
		if gitErr != nil || gitBranch == "" {
			mercurialBranch, mercurialErr := checkedOutMercurialBranch()
			if mercurialErr != nil || mercurialBranch == "" {
				return "", errors.New("could not determine neither a git nor a mercurial branch")
			} else if mercurialErr != nil {
				return "", mercurialErr
			}

			return mercurialBranch, nil
		} else if gitErr != nil {
			return "", gitErr
		}

		return gitBranch, nil
	}

	return branchParam, nil
}

func useLocalBranchName(useLocalBranchNameFlag bool) bool {
	return os.Getenv("PHRASEAPP_USE_LOCAL_BRANCH_NAME") == "true" || useLocalBranchNameFlag
}

func checkedOutGitBranch() (string, error) {
	gitPath := os.Getenv("PHRASEAPP_GIT_BINARY")
	if gitPath == "" {
		systemGitPath, err := exec.LookPath("git")
		if err != nil {
			return "", errors.New("git is not installed")
		}
		gitPath = systemGitPath
	}

	gitCmd := exec.Command(gitPath, "rev-parse", "--abbrev-ref", "HEAD")
	out, err := gitCmd.Output()
	if err != nil {
		return "", err
	}

	gitBranch := strings.TrimSpace(string(out))
	if gitBranch == "HEAD" {
		return "", errors.New("no git branched checked out")
	}

	return gitBranch, nil
}

func checkedOutMercurialBranch() (string, error) {
	hgPath := os.Getenv("PHRASEAPP_MERCURIAL_BINARY")
	if hgPath == "" {
		systemHgPath, err := exec.LookPath("hg")
		if err != nil {
			return "", errors.New("hg is not installed")
		}
		hgPath = systemHgPath
	}

	hgCmd := exec.Command(hgPath, "identify", "-b")
	out, err := hgCmd.Output()
	if err != nil {
		return "", err
	}

	mercurialBranch := strings.TrimSpace(string(out))

	return mercurialBranch, nil
}
