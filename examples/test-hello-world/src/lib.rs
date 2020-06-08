#![no_std]

extern crate sgx_tstd as std;

#[cfg(feature = "with-testing")]
//#[cfg(test)]
pub mod tests {
    use std::prelude::v1::*;
    use testing::{generate_runner, test};

    generate_runner!();

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}