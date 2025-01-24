# Canvasfuck
Paint the canvas using basic Brainfuck. The Brainfuck implementation of Canvasfuck is written in Go and compiled to WebAssembly.

## Get started
Everything sames to [basic Brainfuck](https://en.wikipedia.org/wiki/Brainfuck). But:

The size of the memory is 30,000, and values from 0 to 255 is for the canvas pixels. The size of canvas is 16 × 16 pixels.

Use `.` command to paint. It paints the pixel of `(pointer % 16, pointer / 16)`. For example, if the pointer is 40 and `.` command is executed, `(8, 2)` pixel will be painted. If the pointer is more than 256, `.` command won't do anything.

The memory locations from 0 to 255 memory represent color types. The color types are as follows:

1. "white"
1. "red"
1. "cyan"
1. "blue"
1. "darkblue"
1. "lightblue"
1. "purple"
1. "yellow"
1. "lime"
1. "magenta"
1. "pink"
1. "silver"
1. "gray"
1. "black"
1. "orange"
1. "brown"
1. "maroon"
1. "green"
1. "olive"
1. "aquamarine"

## Yes, it sucks
It's so inefficient and unnecessary. It's so inefficient and unnecessary. You need an extremely long code just to paint a small rectangle in the center of the canvas:
```
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >>>>> >>>>> >
>>>>> >> +. > +. >> >>>>> >
>>>>> >> +. > +.
```

But there are no additional commands like `^` and `v`. Only the basic Brainfuck commands are needed. I simply wanted to create the canvas extension that uses only the basic Brainfuck commands.


## Run in local
```bash
# You need to install Go.
make
python3 -m http.server
```

# License
MIT License.