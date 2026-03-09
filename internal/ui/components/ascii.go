package components

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// LoadFrames reads numbered .txt files from the given directory,
// sorted numerically (1.txt, 2.txt, ... 10.txt).
func LoadFrames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading frames dir: %w", err)
	}

	type numberedFrame struct {
		num     int
		content string
	}

	var frames []numberedFrame
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".txt") {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".txt")
		num, err := strconv.Atoi(name)
		if err != nil {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", e.Name(), err)
		}
		frames = append(frames, numberedFrame{num: num, content: string(data)})
	}

	sort.Slice(frames, func(i, j int) bool {
		return frames[i].num < frames[j].num
	})

	result := make([]string, len(frames))
	for i, f := range frames {
		result[i] = f.content
	}
	return result, nil
}
