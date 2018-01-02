package pego

import "os/exec"

type ffmpeg struct {
	*exec.Cmd
}

func New() (*ffmpeg, error) {
	cmd, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	return &ffmpeg{exec.Command(
		cmd,
	)}, nil
}
