// This file has been automatically generated. Don't edit it.

package events

/*
InputVolumeChanged represents the event body for the "InputVolumeChanged" event.
Since v5.0.0.
*/
type InputVolumeChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// New volume level in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// New volume level in multimap
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}