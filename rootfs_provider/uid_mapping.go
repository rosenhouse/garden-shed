package rootfs_provider

import (
	"fmt"
	"strings"

	"github.com/opencontainers/runtime-spec/specs-go"
)

type MappingList []specs.IDMapping

func (m MappingList) Map(id int) int {
	for _, m := range m {
		if delta := id - int(m.ContainerID); delta < int(m.Size) {
			return int(m.HostID) + delta
		}
	}

	return id
}

func (m MappingList) String() string {
	if len(m) == 0 {
		return "empty"
	}

	var parts []string
	for _, entry := range m {
		parts = append(parts, fmt.Sprintf("%d-%d-%d", entry.ContainerID, entry.HostID, entry.Size))
	}

	return strings.Join(parts, ",")
}
