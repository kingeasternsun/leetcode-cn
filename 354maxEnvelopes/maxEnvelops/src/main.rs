fn main() {
    println!("Hello, world!");
    let  envelopes = vec![vec![5, 4], vec![6, 4], vec![6, 7], vec![2, 3]];
    let a = Solution::max_envelopes(envelopes);
    println!("{}",a);
}

pub struct Solution {}

impl Solution {
    pub fn max_envelopes(mut envelopes: Vec<Vec<i32>>) -> i32 {
        if envelopes.len() <= 1 {
            return envelopes.len() as i32;
        }

        // 经过排序后，后面的信封肯定不可能装入到前面的信封中
        envelopes.sort_by(|a, b| a[0].cmp(&b[0]));

        let mut dp = vec![0; envelopes.len()];
        dp[0] = 1;

        let mut max_ens = 1;
        for i in 1..envelopes.len() {
            if i == 0 {
                continue;
            }

            let mut tmp_ens = 0;
            for j in (0..i).rev() {
                if envelopes[i][0] > envelopes[j][0]
                    && envelopes[i][1] > envelopes[j][1]
                    && dp[j] > tmp_ens
                {
                    tmp_ens = dp[j]
                }
            }

            dp[i] = tmp_ens + 1;

            if dp[i] > max_ens {
                max_ens = dp[i]
            }
        }

        max_ens
    }
}
