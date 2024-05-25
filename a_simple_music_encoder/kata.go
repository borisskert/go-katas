package kata

import (
	"fmt"
	"sort"
	"strings"
)

// Compress https://www.codewars.com/kata/58db9545facc51e3db00000a/train/go
func Compress(raw []int) string {
	return Music(raw).Stream().Encode().AsString()
}

/* ---------------------------------------------------------------------------------------------------------------------
 * Music Type
 * -------------------------------------------------------------------------------------------------------------------*/

type Music []int

func (ctx Music) Stream() MusicStream {
	return MusicStream(ListFrom(ctx))
}

/* ---------------------------------------------------------------------------------------------------------------------
 * MusicStream Type
 * -------------------------------------------------------------------------------------------------------------------*/

type MusicStream List[int]

func (ctx MusicStream) Encode() EncodedStream {
	encoded, remaining := ctx.encodeInitial()

	encodedStream := MusicStream(remaining).encodeRecursively(ListOf(encoded))

	return EncodedStream(encodedStream)
}

func (ctx MusicStream) encodeRecursively(encodedItems List[EncodedItem]) List[EncodedItem] {
	values := List[int](ctx)

	if values.IsEmpty() {
		return encodedItems
	}

	lastEncodedItem := encodedItems.Last()
	value := values.Head()

	if lastEncodedItem.accept(value) {
		extendedItem := lastEncodedItem.append(value)
		updatedItems := encodedItems.Init().Append(extendedItem)

		return MusicStream(values.Tail()).encodeRecursively(updatedItems)
	}

	encodedItem, remaining := ctx.encodeInitial()
	extendedItems := encodedItems.Append(encodedItem)

	return MusicStream(remaining).encodeRecursively(extendedItems)
}

func (ctx MusicStream) encodeInitial() (EncodedItem, List[int]) {
	values := List[int](ctx)

	encoded := ctx.tryToEncode3()

	if encoded != nil {
		return encoded, values.Drop(3)
	}

	encoded = ctx.tryToEncode2()
	if encoded != nil {
		return encoded, values.Drop(2)
	}

	encoded = Simple{value: values.Head()}

	return encoded, values.Tail()
}

func (ctx MusicStream) tryToEncode2() EncodedItem {
	values := List[int](ctx)

	if values.Length() < 2 {
		return nil
	}

	if MusicStream(values.Take(2)).areSameValue() {
		return SameValue{
			count: 2,
			value: values.Head(),
		}
	}

	return nil
}

func (ctx MusicStream) tryToEncode3() EncodedItem {
	values := List[int](ctx)

	if values.Length() < 3 {
		return nil
	}

	value := values.Take(3)

	if MusicStream(values.Take(3)).areAnInterval() {
		return Interval{
			start: value.Head(),
			end:   value.Last(),
			step:  value.Tail().Head() - values.Head(),
		}
	}

	return nil
}

func (ctx MusicStream) areSameValue() bool {
	values := List[int](ctx)

	sorted := values.SortedWith(func(a, b int) bool {
		return a < b
	})

	return sorted.Head() == sorted.Last()
}

func (ctx MusicStream) areAnInterval() bool {
	zipped := ctx.zipWithNeighbor()
	firstPair := zipped.Head()
	interval := firstPair.second - firstPair.first

	return abs(interval) > 0 && zipped.Tail().All(func(pair Pair) bool {
		return pair.second-pair.first == interval
	})
}

func (ctx MusicStream) zipWithNeighbor() List[Pair] {
	values := List[int](ctx)

	if len(values.items) < 2 {
		return EmptyList[Pair]()
	}

	first := values.Head()
	tail := values.Tail()
	second := tail.Head()

	pair := Pair{first: first, second: second}
	furtherPairs := MusicStream(tail).zipWithNeighbor()

	return ListOf(pair).AppendAll(furtherPairs)
}

/* ---------------------------------------------------------------------------------------------------------------------
 * EncodedItem Types
 * -------------------------------------------------------------------------------------------------------------------*/

type EncodedItem interface {
	accept(value int) bool
	append(value int) EncodedItem
	format() string
}

type EncodedStream List[EncodedItem]

func (ctx EncodedStream) AsString() string {
	formatted := EmptyList[string]()
	encodedItems := List[EncodedItem](ctx)

	for _, item := range encodedItems.Items() {
		formatted = formatted.Append(item.format())
	}

	return strings.Join(formatted.Items(), ",")
}

type Simple struct {
	value int
}

func (ctx Simple) accept(int) bool {
	return false
}

func (ctx Simple) append(int) EncodedItem {
	panic("Not supported")
}

func (ctx Simple) format() string {
	return fmt.Sprintf("%d", ctx.value)
}

type SameValue struct {
	count int
	value int
}

func (ctx SameValue) accept(value int) bool {
	return ctx.value == value
}

func (ctx SameValue) append(int) EncodedItem {
	return SameValue{
		count: ctx.count + 1,
		value: ctx.value,
	}
}

func (ctx SameValue) format() string {
	return fmt.Sprintf("%d*%d", ctx.value, ctx.count)
}

type Interval struct {
	start int
	end   int
	step  int
}

func (ctx Interval) accept(value int) bool {
	return ctx.end+ctx.step == value
}

func (ctx Interval) append(value int) EncodedItem {
	return Interval{
		start: ctx.start,
		end:   value,
		step:  ctx.step,
	}
}

func (ctx Interval) format() string {
	if abs(ctx.step) == 1 {
		return fmt.Sprintf("%d-%d", ctx.start, ctx.end)
	}

	return fmt.Sprintf("%d-%d/%d", ctx.start, ctx.end, abs(ctx.step))
}

/* ---------------------------------------------------------------------------------------------------------------------
 * List Type
 * -------------------------------------------------------------------------------------------------------------------*/

type List[T interface{}] struct {
	items []T
}

func (ctx List[T]) Head() T {
	return ctx.items[0]
}

func (ctx List[T]) Last() T {
	lastIndex := len(ctx.items) - 1
	return ctx.items[lastIndex]
}

func (ctx List[T]) Init() List[T] {
	lastIndex := len(ctx.items) - 1
	items := ctx.items[:lastIndex]

	return List[T]{items}
}

func (ctx List[T]) Tail() List[T] {
	items := (ctx.items)[1:]
	return List[T]{items}
}

func (ctx List[T]) Take(n int) List[T] {
	items := (ctx.items)[0:n]
	return List[T]{items}
}

func (ctx List[T]) Drop(n int) List[T] {
	items := (ctx.items)[n:]
	return List[T]{items}
}

func (ctx List[T]) Append(item T) List[T] {
	if ctx.IsEmpty() {
		items := make([]T, 1)
		items[0] = item

		return List[T]{items}
	}

	items := append(ctx.items, item)

	return List[T]{items}
}

func (ctx List[T]) AppendAll(another List[T]) List[T] {
	if ctx.IsEmpty() {
		return another
	}

	items := append(ctx.items, another.items...)

	return List[T]{items}
}

func (ctx List[T]) IsEmpty() bool {
	return len(ctx.items) == 0
}

func (ctx List[T]) Get(index int) T {
	return ctx.items[index]
}

func (ctx List[T]) Length() int {
	return len(ctx.items)
}

func (ctx List[T]) SortedWith(less func(a, b T) bool) List[T] {
	sorted := createCopy(ctx.items)

	sort.SliceStable(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]

		return less(a, b)
	})

	return List[T]{sorted}
}

func (ctx List[T]) All(predicate func(T) bool) bool {
	for _, item := range ctx.items {
		if !predicate(item) {
			return false
		}
	}

	return true
}

func (ctx List[T]) Items() []T {
	items := make([]T, len(ctx.items))

	copy(items, ctx.items)

	return items
}

func ListFrom[T any](items []T) List[T] {
	copied := createCopy(items)
	return List[T]{copied}
}

func ListOf[T any](item T) List[T] {
	items := []T{item}
	return List[T]{items}
}

func EmptyList[T any]() List[T] {
	items := make([]T, 0)
	return List[T]{items}
}

func createCopy[T any](items []T) []T {
	size := len(items)

	copiedItems := make([]T, size)
	copy(copiedItems, items)

	return copiedItems
}

type Pair struct {
	first  int
	second int
}

func abs(value int) int {
	if value < 0 {
		return 0 - value
	}

	return value
}
