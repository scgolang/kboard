// Package kboard makes it easy to use MIDI input from the Keith McMillen instruments K-Board.
package kboard

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/scgolang/midi"
)

// KBoard represents a MIDI connection to a K-Board.
type KBoard struct {
	*midi.Device
}

// New initializes a MIDI connection to the K-Board.
// If there doesn't appear to be a K-Board connected to your computer, an error will be returned.
func New() (*KBoard, error) {
	devices, err := midi.Devices()
	if err != nil {
		return nil, errors.Wrap(err, "listing midi devices")
	}
	var kb *midi.Device
	for _, d := range devices {
		if strings.Contains(d.Name, "K-Board") {
			kb = d
		}
	}
	if kb == nil {
		return nil, errors.New("K-Board not found")
	}
	return &KBoard{Device: kb}, nil
}
