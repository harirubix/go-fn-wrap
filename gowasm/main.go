package main

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go"
)

func hello() {
	fmt.Println("Hello from GO wasmtime!")
}

func bye() {
	fmt.Println("Goodbye from Go Wasmtime!")
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

	helloItem := wasmtime.WrapFunc(store, hello)
	goodbyeItem := wasmtime.WrapFunc(store, bye)
	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{helloItem, goodbyeItem})

	if err != nil {
		fmt.Print("Error 1")
		panic(err)
	}

	fireEventFunc := instance.GetExport(store, "run")

	_, err = fireEventFunc.Func().Call(store)
	if err != nil {
		fmt.Print("Error 2")
		panic(err)
	}
}
