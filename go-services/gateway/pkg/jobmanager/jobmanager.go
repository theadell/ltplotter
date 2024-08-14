package jobmanager

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/patrickmn/go-cache"
)

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

func (jm *JobManager) Set(jobID string, result any) {
	jm.cache.Set(jobID, result, cache.DefaultExpiration)
}

func (jm *JobManager) Get(jobID string) (any, bool) {
	return jm.cache.Get(jobID)
}
