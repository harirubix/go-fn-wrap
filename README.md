Demonstrating Go function call from wasm.

cd gowasm
go run main.go



Rust code

extern "C" {
    fn hello();
    fn goodbye();
}

#[no_mangle]
pub extern "C" fn run() {
    unsafe {
        hello();
        goodbye();
    }
}

Here it expects two functions passed from the execution environment.
GO SCRIPT has this
helloItem := wasmtime.WrapFunc(store, func () {})
goodbyeItem := wasmtime.WrapFunc(store, func () {})
instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{helloItem, goodbyeItem})


To build cargo / rust with updates
cargo build --target wasm32-unknown-unknown