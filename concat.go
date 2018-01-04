package pego

func (f *ffmpeg) Concat(str []string) {
	var res []byte
	for _, f := range str {
		res = append(res, f.Name()...)
		res = append(res, '|')
	}

	files := string(res[:len(res)-1])
	f.Args = append(f.Args, "-i")
	f.Args = append(f.Args, files)
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
