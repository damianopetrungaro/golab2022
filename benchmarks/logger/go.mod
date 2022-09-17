module github.com/damianopetrungaro/workshop/benchmarks/logger

require (
	github.com/damianopetrungaro/workshop v0.0.0
	github.com/sirupsen/logrus v1.9.0
	go.uber.org/zap v1.23.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace github.com/damianopetrungaro/workshop v0.0.0 => ./../../

go 1.18
