package rpc1

type Args struct {
	A, B float32
}

type Result struct {
	C float32
}

type MathService struct {
}

func (s *MathService) AddService(args *Args, result *Result) error {
	result.C = args.A + args.B
	return nil
}

func (s *MathService) DivideService(args *Args, result *Result) error {
	result.C = args.A / args.B
	return nil
}
