package notion

import "time"

// GetMetrics returns current source metrics (copy without mutex)
func (ns *NotionSource) GetMetrics() NotionSourceMetrics {
	ns.metrics.mu.RLock()
	defer ns.metrics.mu.RUnlock()

	// Return a copy without the mutex to avoid lock copying
	return NotionSourceMetrics{
		ObjectsRead:       ns.metrics.ObjectsRead,
		PagesRead:         ns.metrics.PagesRead,
		BlocksRead:        ns.metrics.BlocksRead,
		CommentsRead:      ns.metrics.CommentsRead,
		DatabasesRead:     ns.metrics.DatabasesRead,
		UsersRead:         ns.metrics.UsersRead,
		RequestsMade:      ns.metrics.RequestsMade,
		ErrorsEncountered: ns.metrics.ErrorsEncountered,
		TotalDuration:     ns.metrics.TotalDuration,
		StartTime:         ns.metrics.StartTime,
		EndTime:           ns.metrics.EndTime,
	}
}

// Helper methods for metrics

func (ns *NotionSource) incrementPageCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.PagesRead++
	ns.metrics.ObjectsRead++
}

func (ns *NotionSource) incrementBlockCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.BlocksRead++
	ns.metrics.ObjectsRead++
}

func (ns *NotionSource) incrementCommentCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.CommentsRead++
	ns.metrics.ObjectsRead++
}

func (ns *NotionSource) incrementDatabaseCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.DatabasesRead++
	ns.metrics.ObjectsRead++
}

func (ns *NotionSource) incrementUserCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.UsersRead++
	ns.metrics.ObjectsRead++
}

func (ns *NotionSource) incrementRequestCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.RequestsMade++
}

func (ns *NotionSource) incrementErrorCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.ErrorsEncountered++
}

func (ns *NotionSource) updateEndTime() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	now := time.Now()
	ns.metrics.EndTime = &now
	ns.metrics.TotalDuration = now.Sub(ns.metrics.StartTime)
}
