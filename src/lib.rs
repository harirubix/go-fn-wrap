#[no_mangle]
pub extern "C" fn multiply(a: i32, b: i32) -> i32 {
    a * b
}

#[no_mangle]
pub extern "C" fn add_and_multiply(a: i32, b: i32) -> i32 {
    let go_add_result = unsafe { add(a, b) };
    let mul_result = multiply(go_add_result, b);
    mul_result
}

extern "C" {
    fn add(a: i32, b: i32) -> i32;
}
