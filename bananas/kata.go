package kata

// Bananas / https://www.codewars.com/kata/5917fbed9f4056205a00001e/train/go
func Bananas(text string) []string {
	return search("banana", text).toSlice()
}

func search(banana string, text string) stringSet {
	if len(banana) == 0 {
		return setOf(masked(text))
	}

	if len(text) == 0 {
		return emptySet()
	}

	if banana[0] == text[0] {
		taken := takeOne(banana, text)
		omitted := omitOne(banana, text)

		return taken.merge(omitted)
	}

	return omitOne(banana, text)
}

func takeOne(banana string, text string) stringSet {
	results := search(banana[1:], text[1:])

	return results.mapTo(func(s string) string {
		return string(banana[0]) + s
	})
}

func omitOne(banana string, text string) stringSet {
	results := search(banana, text[1:])

	return results.mapTo(func(s string) string {
		return "-" + s
	})
}

func masked(text string) string {
	maskedString := ""

	for range text {
		maskedString += "-"
	}

	return maskedString
}

type stringSet map[string]struct{}

func setOf(s string) stringSet {
	result := stringSet{}
	result.add(s)

	return result
}

func emptySet() stringSet {
	return stringSet{}
}

func (s stringSet) merge(other stringSet) stringSet {
	result := stringSet{}

	for k := range s {
		result.add(k)
	}

	for k := range other {
		result.add(k)
	}

	return result
}

func (s stringSet) mapTo(f func(string) string) stringSet {
	result := stringSet{}

	for k := range s {
		result.add(f(k))
	}

	return result
}

func (s stringSet) toSlice() []string {
	result := make([]string, 0, len(s))

	for k := range s {
		result = append(result, k)
	}

	return result
}

func (s stringSet) add(str string) {
	s[str] = struct{}{}
}
