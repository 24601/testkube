/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

type TestMetrics struct {
	Executions []TestMetricsExecutions `json:"executions,omitempty"`
	// Percentage pass to fail ratio
	PassFailRatio float64 `json:"pass_fail_ratio,omitempty"`
	// 50th percentile of all durations
	ExecutionDurationP50 string `json:"execution_duration_p50,omitempty"`
	// 90th percentile of all durations
	ExecutionDurationP90 string `json:"execution_duration_p90,omitempty"`
	// 99th percentile of all durations
	ExecutionDurationP99 string `json:"execution_duration_p99,omitempty"`
	// total executions number
	TotalExecutions int `json:"total_executions,omitempty"`
	// failed executions number
	FailedExecutions int `json:"failed_executions,omitempty"`
}
