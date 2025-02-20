//go:build integration

package testresult

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/kubeshop/testkube/internal/pkg/api/repository/storage"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/rand"
)

const (
	mongoDns    = "mongodb://localhost:27017"
	mongoDbName = "testkube-test"
)

func TestTestExecutionsMetrics(t *testing.T) {
	assert := require.New(t)

	repository, err := getRepository()
	assert.NoError(err)

	err = repository.Coll.Drop(context.TODO())
	assert.NoError(err)

	testName := "example-test"

	err = repository.insertExecutionResult(testName, testkube.FAILED_TestSuiteExecutionStatus, time.Now().Add(2*-time.Hour), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Hour), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(10*-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(10*-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.FAILED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.FAILED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.FAILED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key1": "value1", "key2": "value2"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.FAILED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)
	err = repository.insertExecutionResult(testName, testkube.PASSED_TestSuiteExecutionStatus, time.Now().Add(-time.Minute), map[string]string{"key3": "value3", "key4": "value4"})
	assert.NoError(err)

	metrics, err := repository.GetTestSuiteMetrics(context.Background(), testName)
	assert.NoError(err)

	t.Run("getting execution metrics for test data", func(t *testing.T) {
		assert.NoError(err)
		assert.Equal(20, metrics.TotalExecutions)
		assert.Equal(5, metrics.FailedExecutions)
		assert.Len(metrics.Executions, 20)
	})

	t.Run("getting pass/fail ratio", func(t *testing.T) {
		assert.Equal(float64(75), metrics.PassFailRatio)
	})

	t.Run("getting percentiles of execution duration", func(t *testing.T) {
		assert.Contains(metrics.ExecutionDurationP50, "1m0.00")
		assert.Contains(metrics.ExecutionDurationP90, "10m0.00")
		assert.Contains(metrics.ExecutionDurationP99, "1h0m0.00")
	})
}

func getRepository() (*MongoRepository, error) {
	db, err := storage.GetMongoDataBase(mongoDns, mongoDbName)
	repository := NewMongoRespository(db)
	return repository, err
}

func (repository *MongoRepository) insertExecutionResult(testSuiteName string, execStatus testkube.TestSuiteExecutionStatus, startTime time.Time, labels map[string]string) error {
	return repository.Insert(context.Background(),
		testkube.TestSuiteExecution{
			Id:        rand.Name(),
			TestSuite: &testkube.ObjectRef{Namespace: "testkube", Name: testSuiteName},
			Name:      "dummyName",
			StartTime: startTime,
			EndTime:   time.Now(),
			Duration:  time.Since(startTime).String(),
			Labels:    labels,
			Status:    &execStatus,
		})
}
