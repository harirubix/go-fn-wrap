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

func add4(a, b, c, d int32) int32 {
	fmt.Println("Add4 from Go Wasmtime!")
	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)
	return a + b + c + d
}

func dummy(a int32) {
	fmt.Printf("Dummy from Go Wasmtime! %d\n", a)
}

func main() {
	// Create an engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Load the WebAssembly module
	module, err := wasmtime.NewModuleFromFile(store.Engine,
		"../target/wasm32-wasi/debug/callback.wasm",
	)
	if err != nil {
		fmt.Print("Error 0")
		panic(err)
	}

	// helloItem := wasmtime.WrapFunc(store, hello)
	// goodbyeItem := wasmtime.WrapFunc(store, bye)
	// addFN := wasmtime.WrapFunc(store, add)
	// instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{
	// 	helloItem, goodbyeItem, addFN})

	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{
		wasmtime.WrapFunc(store, add),
		wasmtime.WrapFunc(store, add4),
		wasmtime.WrapFunc(store, add),

		wasmtime.WrapFunc(store, add),

		wasmtime.WrapFunc(store, dummy),
	})

	if err != nil {
		fmt.Print("Error 1")
		panic(err)
	}

	addAndMultiply := instance.GetExport(store, "add_and_multiply")

	// Call the "add_and_multiply" function from Rust
	result, _ := addAndMultiply.Func().Call(store, 2, 3)
	fmt.Println("Result:", result.(int32))
}
