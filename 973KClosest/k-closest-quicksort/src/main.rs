fn main() {
    println!("Hello, world!");
    assert_eq!(Solution::k_closest(vec![vec![1,3],vec![-2,2]], 1),vec![vec![-2,2]]);
    assert_eq!(Solution::k_closest(vec![vec![-2,2]], 1),vec![vec![-2,2]]);
    assert_eq!(
        Solution::k_closest(vec![vec![3, 3], vec![5, -1], vec![-2, 4]], 2),
        vec![vec![-2, 4], vec![3, 3]]
    );


}
pub struct Solution;

impl Solution {

    //利用快排
    pub fn k_closest(mut points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let k = k as usize;
        if points.len() == k {
            return points;
        }

    
        Self::quick_sort(&mut points[0..], k);



       points.into_iter().take(k).collect()
    }

    //利用快排
    pub fn quick_sort( points: &mut [Vec<i32>], k: usize)-> usize  {

        let mut beg = 0;//[0 beg]里面的都是比points[0]小于或相等的
        let mut end = 1; //(beg ~]里面的都是比points[0]大的

        while end < points.len(){

            if points[end][0]*points[end][0] + points[end][1]*points[end][1] > points[0][0]*points[0][0] +points[0][1]*points[0][1]{
                end+=1;
                continue;
            }

            points.swap(beg+1, end);
            end +=1;
            beg +=1;

        }

        // print!("{:?}\n",points);
        points.swap(0, beg);


        if beg == k-1{
            return  beg
        }

        if beg > k-1{
            return Self::quick_sort(& mut points[0..beg], k);
        }

        return Self::quick_sort(& mut points[beg+1..], k-(beg+1));
        
    }
}
