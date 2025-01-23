build:
	cd wasm && GOOS=js GOARCH=wasm go build -o ./interpreter.wasm

# get:
# 	cd wasm && go get $(PKG)
