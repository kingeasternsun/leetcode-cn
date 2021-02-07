use std::mem::swap;

/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 14:37:20
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 15:10:50
 * @FilePath: \1528restoreString\restore-string\src\lib.rs
 */
pub struct Solution;
impl Solution {
    pub fn restore_string(s: String, indices: Vec<i32>) -> String {

        let mut s = s.into_bytes();
        let mut indices :Vec<usize>= indices.iter().map(|x| *x as usize).collect();

        for i in 0..s.len(){

            if indices[i]==i{
                continue;
            }

            let mut ch = s[i];
            let mut next_indice =indices[i];
            
            while next_indice!=i {
                std::mem::swap(& mut s[next_indice],& mut ch); //ch放到要调整的位置next_indice上，并且把next_indice上原来的数值移出来放在ch里

                //这里最巧妙，1. 保存 next_indice 要移动的新位置  indices[next_indice] 到 next_indice 2. 把 indices[next_indice]置为 next_indice，避免死循环
                std::mem::swap(& mut indices[next_indice],&mut next_indice);  
            }
            //这里关键，因为终止条件是 next_indice!=i ，所以最后next_indice回到i的时候就跳出循环了，s[i]上还没有最后的赋值
            s[i]=ch
        }

        unsafe {String::from_utf8_unchecked(s)}

    }

}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::restore_string(String::from("codeleet"), vec![4,5,6,7,0,2,1,3]),String::from("leetcode"));
        assert_eq!(Solution::restore_string(String::from("abc"), vec![0,1,2]),String::from("abc"));
        assert_eq!(Solution::restore_string(String::from("aiohn"), vec![3,1,4,2,0]),String::from("nihao"));
        assert_eq!(Solution::restore_string(String::from("aaiougrt"), vec![4,0,2,6,7,3,1,5]),String::from("arigatou"));
        assert_eq!(Solution::restore_string(String::from("art"), vec![1,0,2]),String::from("rat"));
    }
}
