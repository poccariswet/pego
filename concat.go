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
