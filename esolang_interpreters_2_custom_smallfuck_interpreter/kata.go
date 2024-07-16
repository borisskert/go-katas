package kata

// Interpreter / https://www.codewars.com/kata/58678d29dbca9a68d80000d7/train/go
func Interpreter(code, tape string) string {
	machine := newMachine(code, tape)
	machine.run()

	return machine.tape.String()
}

type tape struct {
	read   string
	unread string
}

func newTapeFromString(s string) tape {
	return tape{unread: s}
}

func (t *tape) readBit() bool {
	if len(t.unread) == 0 {
		panic("cannot read bit")
	}

	return t.unread[:1] == "1"
}

func (t *tape) flipBit() {
	if len(t.unread) == 0 {
		panic("cannot flip bit")
	}

	if t.unread[:1] == "0" {
		t.unread = "1" + t.unread[1:]
	} else {
		t.unread = "0" + t.unread[1:]
	}
}

func (t *tape) moveRight() bool {
	if len(t.unread) == 0 {
		return false
	}

	t.read += t.unread[:1]
	t.unread = t.unread[1:]

	return true
}

func (t *tape) moveLeft() bool {
	if len(t.read) == 0 {
		return false
	}

	t.unread = t.read[len(t.read)-1:] + t.unread
	t.read = t.read[:len(t.read)-1]

	return true
}

func (t *tape) String() string {
	return t.read + t.unread
}

type code struct {
	tape
}

func newCodeFromString(s string) code {
	return code{tape: newTapeFromString(s)}
}

func (c *code) readByte() rune {
	if len(c.unread) == 0 {
		panic("cannot read byte")
	}

	return rune(c.unread[:1][0])
}

func (c *code) jumpBackward() bool {
	if len(c.read) == 0 {
		return false
	}

	c.moveLeft()
	current := c.readByte()
	var level = 0

	for level > 0 || current != '[' {
		if current == ']' {
			level++
		} else if current == '[' {
			level--
		}

		c.moveLeft()
		current = c.readByte()
	}

	return true
}

func (c *code) jumpForward() bool {
	if len(c.unread) == 0 {
		return false
	}

	c.moveRight()
	current := c.readByte()
	var level = 0

	for level > 0 || current != ']' {
		if current == '[' {
			level++
		} else if current == ']' {
			level--
		}

		c.moveRight()
		current = c.readByte()
	}

	return true
}

type machine struct {
	code code
	tape tape
}

func newMachine(code, tape string) machine {
	return machine{
		code: newCodeFromString(code),
		tape: newTapeFromString(tape),
	}
}

func (m *machine) run() {
	for len(m.code.unread) > 0 && len(m.tape.unread) > 0 {
		executedCorrectly := m.code.nextInstruction().execute(m)

		if !executedCorrectly {
			break
		}
	}
}

type instruction interface {
	execute(m *machine) bool
}

type flipBit struct{}

func (i *flipBit) execute(m *machine) bool {
	m.tape.flipBit()
	m.code.moveRight()

	return true
}

type moveLeft struct{}

func (i *moveLeft) execute(m *machine) bool {
	m.code.moveRight()
	return m.tape.moveLeft()
}

type moveRight struct{}

func (i *moveRight) execute(m *machine) bool {
	m.code.moveRight()
	return m.tape.moveRight()
}

type jumpForward struct{}

func (i *jumpForward) execute(m *machine) bool {
	if m.tape.readBit() {
		m.code.moveRight()
		return true
	}

	return m.code.jumpForward()
}

type jumpBackward struct{}

func (i *jumpBackward) execute(m *machine) bool {
	if !m.tape.readBit() {
		m.code.moveRight()
		return true
	}

	return m.code.jumpBackward()
}

type noop struct{}

func (i *noop) execute(m *machine) bool {
	m.code.moveRight()
	return true
}

func (c *code) nextInstruction() instruction {
	switch c.readByte() {
	case '*':
		return &flipBit{}
	case '>':
		return &moveRight{}
	case '<':
		return &moveLeft{}
	case '[':
		return &jumpForward{}
	case ']':
		return &jumpBackward{}
	default:
		return &noop{}
	}
}
