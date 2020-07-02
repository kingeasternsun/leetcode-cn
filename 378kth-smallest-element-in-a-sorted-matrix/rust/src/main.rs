fn main() {
    println!("Hello, world!");
}
pub struct Solution {}

impl Solution {
    pub fn kth_smallest(matrix: Vec<Vec<i32>>, k: i32) -> i32 {

        if k ==0{
            return matrix[0][0];
        }

        let h = matrix.len();
        if k as usize == h*h{
            return matrix[h-1][h-1];
        }

        let mut left = matrix[0][0];
        let mut right = matrix[h-1][h-1];
        let mut mid  = 0;
        let mut cnt = 0;
        while left < right {
            mid = (left + right)/2;

            cnt = Self::less_count(&matrix, mid);
            if cnt <k{
                left = mid+1;
            }else{
                right = mid;
            }

        }

        left

        
    }

    pub fn less_count(matrix: & Vec<Vec<i32>>, mid_value: i32)->i32 {

        let mut count = 0;
        let h = matrix.len();
        let mut row = h-1 ;
        let mut col = 0 ;

        while  col < h {

            if matrix[row][col]<= mid_value{
                count = count+ row as i32 +1;
                col = col + 1;
            }else{

                if row == 0{
                    break
                }
                row  = row - 1;
            }

        }


        count as i32
    }


}