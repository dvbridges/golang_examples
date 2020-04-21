// Demo of interfaces
package main

import (
	"interfaces/sounds"
)

var Sound interface {
	Play()
}

func main() {
	// Assign the interface to a variable
	engine := Sound

	// Choose which compatible struct to apply to the interface
	for i, _ := range []int{1,2} {
		if i == 0 {
			engine = sounds.EngineOne{Name: "Engine one"}
        } else {
			engine = sounds.EngineTwo{Name: "Engine two"}
		}
		engine.Play()
	}		

}