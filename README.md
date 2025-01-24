# Canvasfuck
Paint the canvas using basic Brainfuck. The Brainfuck implementation of Canvasfuck is written in Go and compiled to WebAssembly.

## Get started
The size of the memory is 30,000, and 0 ~ 324 is for the canvas pixel. The size of canvas is 32 × 32 pixel.

Use `.` command to paint. It paint the pixel of `(pointer % 32, pointer / 32)`. For example, if the pointer is 40 and `.` command is went, (8, 1) pixel will be painted.

## Run in local
```bash
# You need to install Go.
make
python3 -m http.server
```