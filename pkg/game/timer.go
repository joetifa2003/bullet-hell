package game

type Timer struct {
	passed float32
	target float32
}

func (t *Timer) Update(dt float32) {
	if t.passed < t.target {
		t.passed += dt
	}
}

func (t *Timer) Done() bool {
	return t.passed >= t.target
}

func (t *Timer) Reset() {
	t.passed -= t.target
}

func NewTimer(target float32) Timer {
	return Timer{
		passed: target,
		target: target,
	}
}
