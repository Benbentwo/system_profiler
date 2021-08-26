package system_profiler

import (
	"encoding/json"
	"errors"
	"os/exec"
)

var DefaultArgs = []string{"-detailLevel", "mini"}

func New() (*SystemProfiler, error) {
	return GetInformation(DefaultArgs)
}

func GetInformation(opts []string) (*SystemProfiler, error) {
	args := append(opts, "-json")
	out, err := exec.Command("system_profiler", args...).Output()
	if err != nil {
		return nil, errors.Wrap(err, "system_profiler has failed")
	}

	var data SystemProfiler
	if err := json.Unmarshal(out, &data); err != nil {
		return nil, errors.Wrap(err, "system_profiler did not return valid json.")
	}

	return &data, nil
}
