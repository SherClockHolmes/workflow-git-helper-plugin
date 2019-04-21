package git

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type Clone struct {
	args []string
	path string
}

func New(repository, branch string) (*Clone, error) {
	c := &Clone{}

	// git clone if the directory does not exist
	gitRepo := path.Base(repository)
	gitRepoPath, err := filepath.Abs(gitRepo[:len(gitRepo)-4])
	if err != nil {
		return nil, err
	}

	var gitArgs []string
	if _, err := os.Stat(gitRepoPath); os.IsNotExist(err) {
		gitArgs = []string{"clone", "--progress"}

		// Add branch to clone statement if it exists
		if branch != "" {
			gitArgs = append(gitArgs, []string{"-b", branch}...)
		}

		gitArgs = append(gitArgs, repository)
	} else {
		gitArgs = []string{"pull", "--progress"}
		c.path = gitRepoPath
	}

	c.args = gitArgs
	return c, nil
}

func (c *Clone) Run() ([]byte, error) {
	// Change directory if update
	if c.path != "" {
		if err := os.Chdir(c.path); err != nil {
			return nil, err
		}

		defer os.Chdir("..") // Go back to working directory after update
	}

	cmd := exec.Command("git", c.args...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
