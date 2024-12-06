package day6

type State string

const (
	Up    State = "Up"
	Down  State = "Down"
	Left  State = "Left"
	Right State = "Right"
)

type FSM struct {
	currentState State
	transitions  map[State]State
}

func NewFSM(initial State, transitions map[State]State) *FSM {
	return &FSM{
		currentState: initial,
		transitions:  transitions,
	}
}

func (f *FSM) CurrentState() State {
	return f.currentState
}

func (f *FSM) NextState() {
	if nextState, ok := f.transitions[f.currentState]; ok {
		f.currentState = nextState
	}
}

func (f *FSM) GetStateChar() string {
	if f.currentState == Left {
		return "<"
	}
	if f.currentState == Up {
		return "^"
	}

	if f.currentState == Right {
		return ">"
	}

	if f.currentState == Down {
		return "V"
	}

	return "Thefuck?"
}
