/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-08 15:10:09
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 15:54:37
 * @FilePath: /max_turbulence_size/src/lib.rs
 */

pub struct Solution;

impl Solution {
    pub fn max_turbulence_size(arr: Vec<i32>) -> i32 {
        if arr.len() <= 1 {
            return 1;
        }

        let mut pre_status: i32 = 0;
        let mut cur_status: i32 = 0;
        let mut cur_len = 0;
        let mut max_len = 0;
        for i in 1..arr.len() {
            pre_status = cur_status;
            if arr[i - 1] == arr[i] {
                cur_status = 0
            } else if arr[i - 1] < arr[i] {
                cur_status = -1
            } else {
                cur_status = 1
            }

            //遇到拐点
            if pre_status * cur_status == -1 {
                cur_len += 1;
            } else {
                max_len = std::cmp::max(max_len, cur_len);
                cur_len = cur_status.abs() + 1;
            }
        }

        std::cmp::max(max_len, cur_len)
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::max_turbulence_size(vec![9,4,2,10,7,8,8,1,9]),5);
        assert_eq!(Solution::max_turbulence_size(vec![4, 8, 12, 16]), 2);
        assert_eq!(Solution::max_turbulence_size(vec![100]), 1);
    
    }
}


/*

And I execute command:

$ rustup update && rustup component add rust-src

info: syncing channel updates for 'nightly-x86_64-unknown-linux-gnu'
info: latest update on 2020-10-30, rust version 1.49.0-nightly (6bdae9edd 2020-10-29)
info: skipping nightly which is missing installed components 'rls-preview', 'rustfmt-preview'
info: syncing channel updates for 'nightly-2020-10-29-x86_64-unknown-linux-gnu'
info: latest update on 2020-10-29, rust version 1.49.0-nightly (31ee872db 2020-10-28)
info: skipping nightly which is missing installed components 'rls-preview', 'rustfmt-preview'
info: syncing channel updates for 'nightly-2020-10-28-x86_64-unknown-linux-gnu'
info: latest update on 2020-10-28, rust version 1.49.0-nightly (07e968b64 2020-10-27)
info: skipping nightly which is missing installed components 'rls-preview', 'rustfmt-preview'
info: syncing channel updates for 'nightly-2020-10-27-x86_64-unknown-linux-gnu'
info: latest update on 2020-10-27, rust version 1.49.0-nightly (fd542592f 2020-10-26)
info: skipping nightly which is missing installed components 'rls-preview', 'rustfmt-preview'
info: syncing channel updates for 'nightly-2020-10-26-x86_64-unknown-linux-gnu'
info: latest update on 2020-10-26, rust version 1.49.0-nightly (4760b8fb8 2020-10-25)
info: skipping nightly which is missing installed components 'rls-preview', 'rustfmt-preview'
info: syncing channel updates for 'nightly-2020-10-25-x86_64-unknown-linux-gnu'

  nightly-x86_64-unknown-linux-gnu unchanged - rustc 1.49.0-nightly (ffa2e7ae8 2020-10-24)

info: cleaning up downloads & tmp directories
info: component 'rust-src' is up to date
I have following setting in ~/.zshrc:

export RUST_SRC_PATH="$(rustc --print sysroot)/lib/rustlib/src/rust/src"
*/