package cli

import "os"

type Args struct {
	Tech           string
	AdditionalArgs []string
}

func GetArgs() Args {

	args := Args{Tech: "", AdditionalArgs: []string{}}
	argsNumber := len(os.Args)

	switch {
	case argsNumber == 1:
	case argsNumber == 2:
		args.Tech = os.Args[1]
	default:
		args.Tech = os.Args[1]
		args.AdditionalArgs = os.Args[2:]
	}

	return args
}
