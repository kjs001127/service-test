package datadog

import (
	"strconv"

	"github.com/go-resty/resty/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func DecorateResty(r *resty.Client) *resty.Client {
	r.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		opts := []ddtrace.StartSpanOption{
			tracer.ResourceName(request.URL),
			tracer.SpanType(ext.SpanTypeWeb),
			tracer.Tag(ext.HTTPMethod, request.Method),
			tracer.Tag(ext.HTTPURL, request.URL),
			tracer.ServiceName(ddServiceName),
		}

		_, ctx := tracer.StartSpanFromContext(request.Context(), "resty.request", opts...)
		request.SetContext(ctx)

		return nil
	})

	r.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		ctx := resp.Request.Context()

		span, ok := tracer.SpanFromContext(ctx)
		if ok {
			span.SetTag(ext.HTTPURL, resp.Request.URL)
			span.SetTag(ext.HTTPUserAgent, resp.Request.Header.Get("User-Agent"))
			span.SetTag("http.host", resp.Request.RawRequest.Host)

			span.SetTag(ext.HTTPCode, strconv.Itoa(resp.StatusCode()))
			span.SetTag(ext.Error, resp.Error())
			span.Finish()
		}

		return nil
	})

	return r
}
