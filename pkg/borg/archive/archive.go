package archive

type OptionFunc func(*Options)

type Options struct {
	Compression      string
	WorkingDirectory string
	Excludes         []string
}

func WithCompression(compression string) OptionFunc {
	return func(o *Options) {
		o.Compression = compression
	}
}

func WithWorkingDirectory(workingDirectory string) OptionFunc {
	return func(o *Options) {
		o.WorkingDirectory = workingDirectory
	}
}

func WithExcludes(excludes []string) OptionFunc {
	return func(o *Options) {
		o.Excludes = excludes
	}
}
