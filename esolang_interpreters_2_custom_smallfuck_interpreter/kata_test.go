package kata_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/esolang_interpreters_2_custom_smallfuck_interpreter"
	"testing"
)

var examples = []testCase{
	{
		input{"*", "00101100"},
		"10101100",
		"Flips the leftmost cell of the tape",
	},
	{
		input{">*>*", "00101100"},
		"01001100",
		"Flips the second and third cell of the tape",
	},
	{
		input{"*>*>*>*>*>*>*>*", "00101100"},
		"11010011",
		"Flips all the bits in the tape",
	},
	{
		input{"*>*>>*>>>*>*", "00101100"},
		"11111111",
		"Flips all the bits that are initialized to 0",
	},
	{
		input{">>>>>*<*<<*", "00101100"},
		"00000000",
		"Goes somewhere to the right of the tape and then flips all bits that are initialized to 1, progressing leftwards through the tape",
	},
	{
		input{"iwmlis *!BOSS 333 ^v$#@", "00101100"},
		"10101100",
		"Ignore non-command characters",
	},
	{
		input{"*>>>*>*>>*>>>>>>>*>*>*>*>>>**>>**", "0000000000000000"},
		"1001101000000111",
		"Return when pointer goes out of bounds",
	},
	{
		input{"<<<*>*>*>*>*>>>*>>>>>*>*", "0000000000000000"},
		"0000000000000000",
		"Return when pointer goes out of bounds",
	},
	{
		input{"*[>*]", "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		"1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
		"Simple and nested loops",
	},
	{
		input{"[>*]", "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"Simple and nested loops",
	},
	{
		input{"[[]*>*>*>]", "000"},
		"000",
		"Simple and nested loops",
	},
	{
		input{"[*>[>*>]>]", "11001"},
		"01100",
		"Simple and nested loops",
	},
}

var _ = Describe("Smallfuck Interpreter", func() {
	Describe("Examples", func() {
		for i, p := range examples {
			in, out, msg := p.in, p.out, p.msg
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := Interpreter(in.code, in.tape)
				if ans != out {
					fmt.Printf("<LOG::Input>(%#v, %#v)\n", in.code, in.tape)
					fmt.Printf("<LOG::Returned>%#v\n", ans)
					fmt.Printf("<LOG::Expected>%#v\n", out)
					fmt.Println(msg)
				}
				Expect(ans).To(Equal(out))
			})
		}
	})
})

type input struct {
	code string
	tape string
}

type testCase struct {
	in  input
	out string
	msg string
}

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Esolang Interpreters #2 - Custom Smallfuck Interpreter")
}
