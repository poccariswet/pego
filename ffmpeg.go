package pego

import (
	"os/exec"

	"github.com/pkg/errors"
)

type ffmpeg struct {
	*exec.Cmd
}

func New() (*ffmpeg, error) {
	cmd, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, errors.New("cannot find ffmpeg")
	}

	return &ffmpeg{exec.Command(
		cmd,
	)}, nil
}

func (f *ffmpeg) Dir(dir string) {
	f.Dir = dir
}

func (f *ffmeg) ArgsSet(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) Run(out string) {
	f.Args = append(f.Args, out)
	return f.Run()
}
