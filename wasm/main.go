package main

import (
	"errors"
	"syscall/js"
	"unicode/utf8"
)

const (
	dec_ptr = iota
	inc_ptr

	decrease
	increase

	output
	input

	loop_start
	loop_end
)

const mem_size int = 30000

func main() {
	js.Global().Set("executeBf", Run())
	<-make(chan struct{})
}

func compile(code string) (compiled [][2]int, err error) {
	stack := []int{}
	i := 0

	for n := 0; n < utf8.RuneCountInString(code); n++ {
		switch string([]rune(code)[n]) {
		case "<":
			compiled = append(compiled, [2]int{dec_ptr, 0})
		case ">":
			compiled = append(compiled, [2]int{inc_ptr, 0})

		case "-":
			compiled = append(compiled, [2]int{decrease, 0})
		case "+":
			compiled = append(compiled, [2]int{increase, 0})

		case ".":
			compiled = append(compiled, [2]int{output, 0})
		case ",":
			compiled = append(compiled, [2]int{input, 0})

		case "[":
			compiled = append(compiled, [2]int{loop_start, 0})
			stack = append(stack, i)
		case "]":
			if len(stack) == 0 {
				return nil, errors.New("brainf**king syntax error")
			}

			compiled = append(compiled, [2]int{loop_end, stack[len(stack)-1]})
			compiled[stack[len(stack)-1]][1] = i
			stack = stack[:len(stack)-1]
		default:
			i--
		}

		i++
	}

	if len(stack) != 0 {
		return nil, errors.New("brainf**king syntax error")
	}

	return
}

func execute(jsdoc js.Value, program [][2]int, stdinput string) {
	ptr := 0
	memory := [mem_size]uint8{}
	input_index := 0
	// ptr = 0
	// memory = [mem_size]uint8{}

	rect_size := 20
	pixel_size := 16

	canvas := jsdoc.Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")

	ctx.Call("clearRect", 0, 0, canvas.Get("width"), canvas.Get("height"))

	colors := []string{
		"white",
		"red",
		"cyan",
		"blue",
		"darkblue",
		"lightblue",
		"purple",
		"yellow",
		"lime",
		"magenta",
		"pink",
		"silver",
		"gray",
		"black",
		"orange",
		"brown",
		"maroon",
		"green",
		"olive",
		"aquamarine",
	}

	for i := 0; i < len(program); i++ {
		e := program[i]

		switch e[0] {
		case dec_ptr:
			if ptr <= 0 {
				ptr = mem_size - 1
				break
			}

			ptr--

		case inc_ptr:
			if ptr >= mem_size-1 {
				ptr = 0
				break
			}

			ptr++

		case decrease:
			memory[ptr]--

		case increase:
			memory[ptr]++

		case output:
			ctx.Set("fillStyle", colors[int(memory[ptr])%len(colors)])
			ctx.Call("fillRect", ptr%pixel_size*rect_size, ptr/pixel_size*rect_size, rect_size, rect_size)

		case input:
			if input_index < len(stdinput) {
				memory[ptr] = byte(stdinput[input_index])
				input_index++
			}

		case loop_start:
			if memory[ptr] == 0 {
				i = e[1] - 1
			}

		case loop_end:
			if memory[ptr] != 0 {
				i = e[1] - 1
			}
		}
	}
}

func Run() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "invalid arguments"
		}

		doc := js.Global().Get("document")

		program, err := compile(args[0].String())

		if program == nil {
			panic(err)
		}

		execute(doc, program, args[1].String())

		return nil
	})

	return jsfunc
}
