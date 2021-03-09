/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-09 14:12:10
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-09 14:13:14
 * @FilePath: /452findMinArrowShots/find_min_arrow_shots/src/lib.rs
 */

pub struct Solution;
//16ms 2.9MB
impl Solution {
    pub fn find_min_arrow_shots(points: Vec<Vec<i32>>) -> i32 {
        if points.len() == 0 {
            return 0;
        }

        let mut points = points;
        points.sort_by(|a, b| {
            if a[1] == b[1] {
                return a[0].cmp(&a[0]);
            }

            return a[1].cmp(&b[1]);
        });

        let mut mark = points[0][1];
        let mut res = 1;

        for v in points.iter() {
            if v[0] > mark {
                mark = v[1];
                res += 1;
            }
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        let a = vec![
            vec![10,16],
            vec![2,8],
            vec![1,6],
            vec![7,12],
        ];
        assert_eq!(Solution::find_min_arrow_shots(a),2);
    
        let b = vec![
            vec![1,2],
            vec![3,4],
            vec![5,6],
            vec![7,8],
        ];
        assert_eq!(Solution::find_min_arrow_shots(b),4);
    
        let c = vec![
            vec![1,2],
            vec![2,3],
            vec![3,4],
            vec![4,5],
        ];
        assert_eq!(Solution::find_min_arrow_shots(c),2);
    }
}
