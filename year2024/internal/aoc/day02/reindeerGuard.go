package day02

type State int

const (
	Unknown    State = 0
	Increasing State = 1
	Decreasing State = 2
)

type ReindeerGuard struct {
	currentState State
	isValid      bool
}

func NewReindeerGuard() *ReindeerGuard {
	return &ReindeerGuard{
		currentState: Unknown,
		isValid:      true,
	}
}

func (sm *ReindeerGuard) Reset() {
	sm.currentState = Unknown
	sm.isValid = true
}

func (sm *ReindeerGuard) Next(prev, curr int) {
	difference := curr - prev

	if difference == 0 || difference < -3 || difference > 3 {
		sm.isValid = false
		return
	}

	if sm.currentState == Unknown {
		if difference > 0 {
			sm.currentState = Increasing
		} else {
			sm.currentState = Decreasing
		}
	} else if sm.currentState == Increasing && difference < 0 {
		sm.isValid = false
	} else if sm.currentState == Decreasing && difference > 0 {
		sm.isValid = false
	}
}
