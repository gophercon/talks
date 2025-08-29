package brain

type Interface interface {
	Think(prefix []string) (string, error)
}

func Think(br Interface, prompt string) (string, error) { panic("TODO") }
