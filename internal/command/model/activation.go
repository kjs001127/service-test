package model

type Activation struct {
	ActivationID
	Enabled bool
}

type Activations []*Activation

func (activations Activations) ToMap() map[string]*Activation {
	ret := make(map[string]*Activation)
	for _, a := range activations {
		ret[a.CommandID] = a
	}
	return ret
}

type ActivationID struct {
	ChannelID string
	CommandID string
}
