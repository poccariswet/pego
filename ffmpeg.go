package pego

import (
	"os/exec"
)

type ffmpeg struct {
	*exec.Cmd
}

func New() (*ffmpeg, error) {
	cmd, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	return &ffmpeg{
		exec.Command(cmd),
	}, nil
}

func (f *ffmpeg) Dir(dir string) {
	f.Dir = dir
}

func (f *ffmpeg) ArgsSet(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) run(out string) {
	f.Args = append(f.Args, out)
	return f.Run()
}
