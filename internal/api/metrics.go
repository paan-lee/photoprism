package api

import (
	"io"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/common/expfmt"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/get"
)

// GET /api/v1/metrics
func GetMetrics(router *gin.RouterGroup) {
	router.GET("/metrics", func(c *gin.Context) {
		conf := get.Config()
		counts := conf.ClientPublic().Count

		c.Stream(func(w io.Writer) bool {
			reg := prometheus.NewRegistry()
			reg.MustRegister(collectors.NewGoCollector())

			factory := promauto.With(reg)

			registerCountMetrics(factory, counts)
			registerBuildInfoMetric(factory, conf.ClientPublic())

			metrics, err := reg.Gather()
			if err != nil {
				logError("metrics", err)
				return false
			}

			for _, metric := range metrics {
				if _, err := expfmt.MetricFamilyToText(w, metric); err != nil {
					logError("metrics", err)
					return false
				}
			}

			return false
		})
	})
}

// Register metrics that exposes various statistics for this instance.
func registerCountMetrics(factory promauto.Factory, counts config.ClientCounts) {
	metric := factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "photoprism",
			Subsystem: "statistics",
			Name:      "media_count",
			Help:      "media statistics for this photoprism instance",
		}, []string{"stat"},
	)

	metric.With(prometheus.Labels{"stat": "all"}).Set(float64(counts.All))
	metric.With(prometheus.Labels{"stat": "photos"}).Set(float64(counts.Photos))
	metric.With(prometheus.Labels{"stat": "videos"}).Set(float64(counts.Videos))
	metric.With(prometheus.Labels{"stat": "albums"}).Set(float64(counts.Albums))
	metric.With(prometheus.Labels{"stat": "folders"}).Set(float64(counts.Folders))
	metric.With(prometheus.Labels{"stat": "files"}).Set(float64(counts.Files))
}

// Register a metric that exposes build information for this instance.
func registerBuildInfoMetric(factory promauto.Factory, conf config.ClientConfig) {
	factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "photoprism",
			Name:      "build_info",
			Help:      "information about the photoprism instance",
		}, []string{"edition", "goversion", "version"},
	).With(prometheus.Labels{
		"edition":   conf.Edition,
		"goversion": runtime.Version(),
		"version":   conf.Version,
	}).Set(1.0)
}
