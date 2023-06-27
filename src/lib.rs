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