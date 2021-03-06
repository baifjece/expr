package expr

// Option expr option
type Option func(*options)

type options struct {
	ops         map[string]CalcFunc
	typs        map[string]VarType
	defaultType VarType
}

func newOptions() *options {
	return &options{
		ops:         make(map[string]CalcFunc),
		typs:        make(map[string]VarType),
		defaultType: Str,
	}
}

// WithOp add a op
func WithOp(symbol string, opFunc CalcFunc) Option {
	return func(opts *options) {
		opts.ops[symbol] = opFunc
	}
}

// WithVarType with var type
func WithVarType(symbol string, value VarType) Option {
	return func(opts *options) {
		opts.typs[symbol] = value
	}
}

// WithDefaultVarType set default var type
func WithDefaultVarType(value VarType) Option {
	return func(opts *options) {
		opts.defaultType = value
	}
}
