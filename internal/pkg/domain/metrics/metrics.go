package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	Native     *NativeMetrics
	ThirdParty *ThirdPartyMetrics
}

type NativeMetrics struct {
	Hits    *prometheus.CounterVec
	Timings *prometheus.SummaryVec
}

type ThirdPartyMetrics struct {
	Hits    *prometheus.CounterVec
	Timings *prometheus.SummaryVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		Native:     NewNativeMetrics(),
		ThirdParty: NewThirdPartyMetrics(),
	}
}

func NewThirdPartyMetrics() *ThirdPartyMetrics {
	newThirdPartyMetrics := &ThirdPartyMetrics{
		Timings: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "my_service",
			Name:       "timings_third_party",
			Help:       "Latency by each endpoint in third-party API",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"path"}),

		Hits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "my_service",
			Name:      "hits_third_party",
			Help:      "Number of hits per endpoint in third-party API with distribution by response status",
		}, []string{"status", "path"}),
	}

	prometheus.MustRegister(
		newThirdPartyMetrics.Hits,
		newThirdPartyMetrics.Timings,
	)

	return newThirdPartyMetrics
}

func NewNativeMetrics() *NativeMetrics {
	newNativeMetrics := &NativeMetrics{
		Timings: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "my_service",
			Name:       "timings",
			Help:       "Latency by each endpoint",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"path"}),

		Hits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "my_service",
			Name:      "hits",
			Help:      "Number of hits per endpoint with distribution by response status",
		}, []string{"status", "path"}),
	}

	prometheus.MustRegister(
		newNativeMetrics.Hits,
		newNativeMetrics.Timings,
	)

	return newNativeMetrics
}
