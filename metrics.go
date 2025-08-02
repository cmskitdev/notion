package notion

import "time"

// GetMetrics returns current source metrics (copy without mutex)
func (ns *Plugin) GetMetrics() NotionSourceMetrics {
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

func (ns *Plugin) incrementPageCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.PagesRead++
	ns.metrics.ObjectsRead++
}

func (ns *Plugin) incrementBlockCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.BlocksRead++
	ns.metrics.ObjectsRead++
}

func (ns *Plugin) incrementCommentCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.CommentsRead++
	ns.metrics.ObjectsRead++
}

func (ns *Plugin) incrementDatabaseCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.DatabasesRead++
	ns.metrics.ObjectsRead++
}

func (ns *Plugin) incrementUserCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.UsersRead++
	ns.metrics.ObjectsRead++
}

func (ns *Plugin) incrementRequestCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.RequestsMade++
}

func (ns *Plugin) incrementErrorCount() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	ns.metrics.ErrorsEncountered++
}

func (ns *Plugin) updateEndTime() {
	ns.metrics.mu.Lock()
	defer ns.metrics.mu.Unlock()
	now := time.Now()
	ns.metrics.EndTime = &now
	ns.metrics.TotalDuration = now.Sub(ns.metrics.StartTime)
}
