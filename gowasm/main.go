package main

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go"
)

func add(a, b int32) int32 {
	fmt.Println("Add from Go Wasmtime!")
	fmt.Printf("a: %d, b: %d\n", a, b)
	return a + b
}

func main() {
	// Create an engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Load the WebAssembly module
	module, err := wasmtime.NewModuleFromFile(store.Engine,
		"../target/wasm32-unknown-unknown/debug/callback.wasm",
	)
	if err != nil {
		fmt.Print("Error 0")
		panic(err)
	}

	linker := wasmtime.NewLinker(store.Engine)
	linker.DefineFunc(store, "env", "add", add)

	instance, err := linker.Instantiate(store, module)

	if err != nil {
		fmt.Print("Error 1")
		panic(err)
	}

	addAndMultiply := instance.GetExport(store, "add_and_multiply")

	// Call the "add_and_multiply" function from Rust
	result, _ := addAndMultiply.Func().Call(store, 2, 3)
	fmt.Println("Result:", result.(int32))
}
