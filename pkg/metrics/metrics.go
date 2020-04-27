package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Metrics includes metrics used in theia pods controller
type Metrics struct {
	cli                   client.Client
	runningTheias         *prometheus.GaugeVec
	TheiaCreation         *prometheus.CounterVec
	TheiaFailCreation     *prometheus.CounterVec
	TheiaCullingCount     *prometheus.CounterVec
	TheiaCullingTimestamp *prometheus.GaugeVec
}

func NewMetrics(cli client.Client) *Metrics {
	m := &Metrics{
		cli: cli,
		runningTheias: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "theia_running",
				Help: "Current running theia pods in the cluster",
			},
			[]string{"namespace"},
		),
		TheiaCreation: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "theia_create_total",
				Help: "Total times of creating theia pods",
			},
			[]string{"namespace"},
		),
		TheiaFailCreation: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "theia_create_failed_total",
				Help: "Total failure times of creating theia pods",
			},
			[]string{"namespace"},
		),
		TheiaCullingCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "theia_culling_total",
				Help: "Total times of culling theia pods",
			},
			[]string{"namespace", "name"},
		),
		TheiaCullingTimestamp: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "last_theia_culling_timestamp_seconds",
				Help: "Timestamp of the last theia pod culling in seconds",
			},
			[]string{"namespace", "name"},
		),
	}

	metrics.Registry.MustRegister(m)
	return m
}

// Describe implements the prometheus.Collector interface.
func (m *Metrics) Describe(ch chan<- *prometheus.Desc) {
	m.runningTheias.Describe(ch)
	m.TheiaCreation.Describe(ch)
	m.TheiaFailCreation.Describe(ch)
}

// Collect implements the prometheus.Collector interface.
func (m *Metrics) Collect(ch chan<- prometheus.Metric) {
	m.scrape()
	m.runningTheias.Collect(ch)
	m.TheiaCreation.Collect(ch)
	m.TheiaFailCreation.Collect(ch)
}

// scrape gets current running theia statefulsets.
func (m *Metrics) scrape() {
	stsList := &appsv1.StatefulSetList{}
	err := m.cli.List(context.TODO(), stsList)
	if err != nil {
		return
	}
	stsCache := make(map[string]float64)
	for _, v := range stsList.Items {
		name, ok := v.Spec.Template.GetLabels()["theia-pod-name"]
		if ok && name == v.Name {
			stsCache[v.Namespace] += 1
		}
	}

	for ns, v := range stsCache {
		m.runningTheias.WithLabelValues(ns).Set(v)
	}
}
