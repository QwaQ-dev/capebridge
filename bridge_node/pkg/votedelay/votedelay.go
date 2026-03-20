package votedelay

import "time"

func VoteDelay(nodeID string) time.Duration {
	switch nodeID {
	case "node-1":
		return 15 * time.Second
	case "node-2":
		return 0
	case "node-3":
		return 30 * time.Second
	default:
		return 0
	}
}
