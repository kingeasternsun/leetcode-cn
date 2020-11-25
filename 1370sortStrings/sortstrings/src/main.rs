fn main() {
    println!("Hello, world!");

    assert_eq!(Solution::sort_string("aaaabbbbcccc".to_string()),"abccbaabccba".to_string());
    assert_eq!(Solution::sort_string("aaaabbbbccc".to_string()),"abccbaabcba".to_string());
}


pub struct Solution;

impl Solution {
    //0ms 
    pub fn sort_string(s: String) -> String {

        let mut byte_map = vec![0;26];
        for b in s.as_bytes(){

            byte_map[ (b - b'a')as usize]+=1;

        }

        let mut res = String::new();
        let mut flat = true;

        while res.len()<s.len(){

            if flat{

                // for i in 0..26{
                //     if byte_map[i]>0{
                //         res.push( (i as u8 +97)as char);
                //         byte_map[i]-=1;
                //     }
                // }

               (0..26u8).for_each(|i|  if byte_map[i as usize] >0 {  res.push((i+97)as char); byte_map[i as usize]-=1;  } );

            }else{
                // for i in (0..26u8).rev(){
                //     if byte_map[i as usize]>0{
                //         res.push( (i +97)as char);
                //         byte_map[i as usize]-=1;
                //     }
                // } 

               (0..26u8).rev().for_each(|i|  if byte_map[i as usize] >0 {  res.push((i+97)as char); byte_map[i as usize]-=1;  } );

            }

            flat = !flat;

        }
        
        res 

    }
}