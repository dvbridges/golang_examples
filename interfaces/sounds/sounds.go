// A demo using interfaces
package sounds

import (
	"fmt"
	"strings"
)

// An audio driver
type EngineOne struct {
	Name string
}

// Another audio driver
type EngineTwo struct {
	Name string
}

// EngineOne play method
func (e EngineOne) Play() {
	fmt.Printf("Audio Player: %s playing\n", e.Name)
}

// EngineTwo play method
func (e EngineTwo) Play() {
	newName := strings.ToUpper(e.Name)
	fmt.Printf("Audio Player: %s playing\n", newName)
}

// EngineTwo also has a close method
func (e EngineTwo) Close() {
	fmt.Printf("Shutting down: %s\n", e.Name)
}