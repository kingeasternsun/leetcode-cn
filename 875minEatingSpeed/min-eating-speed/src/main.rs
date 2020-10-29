fn main() {
    println!("Hello, world!");
    // print!("{}\n",Solution::min_eating_speed(vec![3,6,7,11], 8));
    // print!("{}\n",Solution::min_eating_speed(vec![30,11,23,4,20],5));
    // print!("{}\n",Solution::min_eating_speed(vec![30,11,23,4,20],6));
    // print!(
    //     "{}\n",
    //     Solution::min_eating_speed(
    //         vec![
    //             332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673,
    //             679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097,
    //             692137887, 718203285, 629455728,
    //              941802184
    //         ],
    //         823855818
    //     )
    // );

    print!("{}\n",Solution::min_eating_speed(vec![312884470],968709470));

}
pub struct Solution;

impl Solution {
    pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
        if piles.len() == 0 {
            return 0;
        }
        // print!("{}\n", std::i32::MAX);

        let mut left = 1;
        let mut right = *piles.iter().max().unwrap();
        // print!("max {}\n", right);

        let mut mid;

        // print!("max {}\n", right);
        while left < right {
            mid = left + (right - left) / 2;
            // print!("mid {}\n", mid);
            // let mut piles2 = piles.clone();
            if Self::finish(&piles, h, mid) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        left
    }

    //timeout  关键点
    pub fn finish_timeout(piles: &mut Vec<i32>, h: i32, k: i32) -> bool {
        let mut t = 0;
        let mut i = 0;

        while i < piles.len() {
            //已经过去h个小时了 ，还没吃完
            if t == h {
                return false;
            }

            //当前数量大于k 就吃完等下一个小时
            if piles[i] > k {
                piles[i] -= k;
            } else {
                //吃完到下一堆等
                i += 1;
            }

            t += 1;
        }

        true
    }

    pub fn finish(piles: & Vec<i32>, h: i32, k: i32) -> bool {

        // 计算需要多少小时
        let takes = piles.iter().fold(0,|acc,x| acc+ x/k  + if x%k ==0 {0} else {1});


        takes <= h

    }
}
