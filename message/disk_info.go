package message

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/tychoish/grip/level"
)

type DiskStats struct {
	Message   string                 `json:"message,omitempty" bson:"message,omitempty"`
	Paritions []disk.PartitionStat   `json:"partitions,omitempty" bson:"partitions,omitempty"`
	Usage     []*disk.UsageStat      `json:"usage,omitempty" bson:"usage,omitempty"`
	IOStat    []*disk.IOCountersStat `json:"iostat,omitempty" bson:"iostat,omitempty"`
	Errors    []string               `json:"errors,omitempty" bson:"errors,omitempty"`
	Base      `json:"metadata,omitempty" bson:"metadata,omitempty"`
	loggable  bool
	rendered  string
}

func CollectDiskStats() Composer {
	return NewDiskStats(level.Trace, "")
}

func MakeDiskStats(message string) Composer {
	return NewDiskStats(level.Info, message)
}

func NewDiskStats(priority level.Priority, message string) Composer {
	s := &DiskStats{
		Message: message,
	}

	if err := s.SetPriority(priority); err != nil {
		s.Errors = append(s.Errors, err.Error())
		return s
	}
	s.loggable = true

	partitions, err := disk.Partitions(true)
	s.saveError(err)
	if err != nil {
		for _, p := range partitions {
			u, err := disk.Usage(p.Mountpoint)
			s.saveError(err)
			if err != nil {
				continue
			}

			s.Usage = append(s.Usage, u)
		}

		s.Paritions = partitions
	}

	iostatMap, err := disk.IOCounters()
	s.saveError(err)
	for _, stat := range iostatMap {
		s.IOStat = append(s.IOStat, &stat)
	}

	return s

}

func (s *DiskStats) saveError(err error) {
	if shouldSaveError(err) {
		s.Errors = append(s.Errors, err.Error())
	}
}

// Loggable returns true when the Processinfo structure has been
// populated.
func (s *DiskStats) Loggable() bool { return s.loggable }

// Raw always returns the DiskStats object, however it will call the
// Collect method of the base operation first.
func (s *DiskStats) Raw() interface{} { _ = s.Collect(); return s }

// String returns a string representation of the message, lazily
// rendering the message, and caching it privately.
func (s *DiskStats) String() string {
	if s.rendered == "" {
		s.rendered = renderStatsString(s.Message, s)
	}

	return s.rendered
}
