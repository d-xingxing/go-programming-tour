package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHost string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{ //jaeger client的配置项
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //（固定采样、对所有数据都进行采样）
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ // 是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHost,
		},
	}
	// 根据配置项初始化Tracer对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局的Tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
