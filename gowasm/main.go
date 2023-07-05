// package main

// import (
// 	"fmt"

// 	"github.com/bytecodealliance/wasmtime-go"
// )

// func add(a, b int32) int32 {
// 	fmt.Println("Add from Go Wasmtime!")
// 	fmt.Printf("a: %d, b: %d\n", a, b)
// 	return a + b
// }

// func main() {
// 	// Create an engine and store
// 	engine := wasmtime.NewEngine()
// 	store := wasmtime.NewStore(engine)

// 	// Load the WebAssembly module
// 	module, err := wasmtime.NewModuleFromFile(store.Engine,
// 		"../target/wasm32-unknown-unknown/debug/callback.wasm",
// 	)
// 	if err != nil {
// 		fmt.Print("Error 0")
// 		panic(err)
// 	}

// 	linker := wasmtime.NewLinker(store.Engine)
// 	linker.DefineFunc(store, "env", "add", add)

// 	instance, err := linker.Instantiate(store, module)

// 	if err != nil {
// 		fmt.Print("Error 1")
// 		panic(err)
// 	}

// 	addAndMultiply := instance.GetExport(store, "add_and_multiply")

// 	// Call the "add_and_multiply" function from Rust
// 	result, _ := addAndMultiply.Func().Call(store, 2, 3)
// 	fmt.Println("Result:", result.(int32))
// }

package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/bytecodealliance/wasmtime-go"
)

type WasmtimeRuntime struct {
	store   *wasmtime.Store
	memory  *wasmtime.Memory
	handler *wasmtime.Func

	input  []byte
	output []byte
}

func (r *WasmtimeRuntime) Init(wasmFile string) {
	engine := wasmtime.NewEngine()
	linker := wasmtime.NewLinker(engine)
	linker.DefineWasi()
	wasiConfig := wasmtime.NewWasiConfig()
	r.store = wasmtime.NewStore(engine)
	r.store.SetWasi(wasiConfig)
	linker.FuncWrap("env", "load_input", r.loadInput)
	linker.FuncWrap("env", "dump_output", r.dumpOutput)
	wasmBytes, _ := os.ReadFile(wasmFile)
	module, _ := wasmtime.NewModule(r.store.Engine, wasmBytes)
	instance, _ := linker.Instantiate(r.store, module)
	r.memory = instance.GetExport(r.store, "memory").Memory()
	r.handler = instance.GetFunc(r.store, "handler")
}

func (r *WasmtimeRuntime) loadInput(pointer int32) {
	copy(r.memory.UnsafeData(r.store)[pointer:pointer+int32(len(r.input))], r.input)
}

func (r *WasmtimeRuntime) dumpOutput(pointer int32, muliply int32, length int32) {
	fmt.Println("multiply :", muliply)
	r.output = make([]byte, length)
	copy(r.output, r.memory.UnsafeData(r.store)[pointer:pointer+length])
}

func (r *WasmtimeRuntime) RunHandler(data []byte, buf int32, b1 int32, b2 int32) []byte {
	r.input = data
	r.handler.Call(r.store, buf, b1, b2)
	fmt.Println("Result:", r.output)
	return r.output
}

func main() {
	size := 16
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = 'a'
	}
	b1 := make([]byte, 4)
	binary.LittleEndian.PutUint32(b1, 12)
	fmt.Println(b1)

	b2 := make([]byte, 4)
	binary.LittleEndian.PutUint32(b2, 96)
	fmt.Println(b2)

	merge1 := append(buf, b1...)
	merge := append(merge1, b2...)
	fmt.Println(merge)

	runtime := &WasmtimeRuntime{}
	runtime.Init("../target/wasm32-unknown-unknown/debug/callback.wasm")
	runtime.RunHandler(merge, int32(len(buf)), int32(len(b1)), int32(len(b2)))
}
