#[no_mangle]
pub extern "C" fn multiply(a: i32, b: i32) -> i32 {
    a * b
}

// #[no_mangle]
// pub extern "C" fn add_and_multiply(a: i32, b: i32) -> i32 {
//     let go_add_result = unsafe { add(a, b) };
//     multiply(go_add_result, b)
// }

// extern "C" {
//     fn add(a: i32, b: i32) -> i32;
// }

extern "C" {
    fn load_input(pointer: *mut u8);
    fn dump_output(pointer: *const u8, multiply: u32, length: usize);
}

#[no_mangle]
pub extern "C" fn handler(input_length: usize , b1_length: usize , b2_length: usize) {
    // load input data
    let mut input = Vec::with_capacity(input_length + b1_length + b2_length);
    // let mut buf: Vec<u8> = Vec::with_capacity(input_length);
    // let mut b1: Vec<u8> = Vec::with_capacity(b1_length);
    // let mut b2: Vec<u8> = Vec::with_capacity(b2_length);
    
    unsafe {
        load_input(input.as_mut_ptr());
        input.set_len(input_length + b1_length + b2_length);
    }


    let (buf, b1_b2) = input.split_at(input_length);
    let (b1, b2) = b1_b2.split_at(b1_length);
    // process app data
    let output = buf.to_ascii_uppercase();
    // let b1_out: Vec<u8> = b1.to_ascii_uppercase();
    // let b2_out: Vec<u8> = b2.to_ascii_uppercase();

    let first = u32::from_ne_bytes(b1[0..b1_length].try_into().unwrap());
    let second = u32::from_ne_bytes(b2[0..b2_length].try_into().unwrap());

    let multiply = first * second;

    // dump output data
    unsafe {
        dump_output(output.as_ptr() , multiply , output.len());
        // dump_output(output.as_ptr(), output.len());
        // dump_output(b1_out.as_ptr(), b1.len());

    }
}
