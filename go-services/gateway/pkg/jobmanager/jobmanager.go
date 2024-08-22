package jobmanager

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/patrickmn/go-cache"
)

type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
)

type JobResult struct {
	Status JobStatus
	Result any
}

type JobManager struct {
	cache *cache.Cache
}

func NewJobManager(defaultExpiration, cleanupInterval time.Duration) *JobManager {
	return &JobManager{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}
func (jm *JobManager) GenerateJobID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
func (jm *JobManager) SetJobResult(jobID string, status JobStatus, result any) {
	jm.cache.Set(jobID, &JobResult{
		Status: status,
		Result: result,
	}, cache.DefaultExpiration)
}

func (jm *JobManager) GetJobResult(jobID string) (*JobResult, bool) {
	res, found := jm.cache.Get(jobID)
	if !found {
		return nil, false
	}
	return res.(*JobResult), true
}
