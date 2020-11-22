fn main() {
    println!("Hello, world!");
    assert_eq!(Solution::is_anagram("anagram".to_string(), "nagaram".to_string()),true);
    assert_eq!(Solution::is_anagram("rat".to_string(), "car".to_string()),false);
    assert_eq!(Solution::is_anagram("rat".to_string(), "ra".to_string()),false);
    assert_eq!(Solution::is_anagram("ra".to_string(), "rat".to_string()),false);
    assert_eq!(Solution::is_anagram("abc".to_string(), "acb".to_string()),true);
}

pub struct Solution;

impl Solution {

    //8ms 2.3MB
    pub fn is_anagram1(s: String, t: String) -> bool {

        if s.len()!=t.len(){
            return false
        }


        let mut s_map = std::collections::HashMap::new();
        let mut t_map = std::collections::HashMap::new();
        for r in s.chars(){
            let entry = s_map.entry(r).or_insert(0);
            *entry +=1;
        }


        for r in t.chars(){
            let entry = t_map.entry(r).or_insert(0);
            *entry +=1;
        }

        if s_map.len()!=t_map.len(){
            return false
        }

        for r in s.chars(){

            let s_cnt = s_map.get(&r).unwrap();
            let t_entry = t_map.get(&r);
            
            if t_entry.is_none(){
                return false
            }

            if t_entry.unwrap()!=s_cnt{
                return false
            }


        }

        true
    }


    //0ms 2.3MB
    pub fn is_anagram2(s: String, t: String) -> bool {

        if s.len()!=t.len(){
            return false
        }

        let mut s_array = [0;26];
        let mut t_array = [0;26];

        s.bytes().for_each(|r| s_array[(r - b'a') as usize]+=1 );
        t.bytes().for_each(|r| t_array[(r - b'a' )as usize]+=1 );

        s_array == t_array
    }

    
    //0ms 2.1MB
    pub fn is_anagram(s: String, t: String) -> bool {

        if s.len()!=t.len(){
            return false
        }

        let mut s_array = [0;26];
        // let mut t_array = [0;26];

        s.bytes().for_each(|r| s_array[(r - b'a') as usize]+=1 );
        
        for r in t.bytes(){

            if s_array[(r-b'a') as usize]==0{
                return false
            }

            s_array[(r-b'a') as usize] -=1;

        }

        true
    }


}