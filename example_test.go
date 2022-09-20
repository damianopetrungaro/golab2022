package workshop

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var (
	fakeTraceID = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fakeSpanID  = [8]byte{0, 1, 2, 3, 4, 5, 6, 7}
)

func ExampleLogger() {
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithIDGenerator(&fakeIDGenerator{}),
	)

	otel.SetTracerProvider(tracerProvider)

	ctx, _ := otel.Tracer("-").Start(context.Background(), "-")

	var logger Logger = New(os.Stdout, TraceDecorator)
	logger = logger.With(String("name", "golab"))
	logger = logger.With(Int("year", 2022))
	logger.Info(ctx, "Done!")
	// Output:
	// {"level":"INFO","message":"Done!","fields":[{"name":"golab"},{"year":2022},{"trace_id":"000102030405060708090a0b0c0d0e0f"},{"span_id":"0001020304050607"}]}
}

type fakeIDGenerator struct{}

func (f *fakeIDGenerator) NewIDs(_ context.Context) (trace.TraceID, trace.SpanID) {
	return fakeTraceID, fakeSpanID
}

func (f *fakeIDGenerator) NewSpanID(_ context.Context, _ trace.TraceID) trace.SpanID {
	return fakeSpanID
}
