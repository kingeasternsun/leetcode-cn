fn main() {
    println!("Hello, world!");
    assert_eq!(5, Solution::num_tilings(3));
    assert_eq!(1255, Solution::num_tilings(10));
    assert_eq!(979232805, Solution::num_tilings(1000));
}

pub struct Solution;
impl Solution {
    pub fn num_tilings(n: i32) -> i32 {

        let m  = 1000000007;
        if n == 0{
            return 0
        }

        if n == 1{
            return 1
        }

        if n == 2{
            return 2
        }

        let mut dp:Vec<u32> = vec![0;4];
        dp[0]=1;
        dp[1]=1;

        for _ in 2..n+1{
            dp[3] = dp[1]%m; //保存之前的dp[1]
            // dp[i] = (dp[i-1] + dp[i-2] + 2*dp1[i-1]) % mod // dp[i-1]:表示横着铺一个X，dp[i-2]表示两个竖着的X， dp1[i-1]表示 两种L
            dp[1] = dp[1]%m + dp[0]%m + 2*dp[2]%m;
            // dp1[i] = (dp[i-2] + dp1[i-1]) % m            // dp[i-2]:表示一个L 放在整齐的前i-2上， dp[i-1]表示X竖着插在之前的L上
            dp[2] = dp[0]%m + dp[2] % m;
            dp[0] = dp[3];
        }

        (dp[1] %m) as i32
        
    }
}