fn main() {
    println!("Hello, world!");


    let mut dict = WordDictionary::new();
    dict.add_word("bad".to_owned());
    dict.add_word("dad".to_owned());
    dict.add_word("mad".to_owned());
    // dict.search("pad".to_owned());
    // dict.search("bad".to_owned());
    dict.search(".ad".to_owned());
    // dict.search("da.".to_owned());
}


use std::collections::HashMap;


#[derive(Default)]
struct Trie {
    is_end: bool,
    child: HashMap<usize,Box<Trie>>,
}

impl Trie {

    /** Initialize your data structure here. */
    fn new() -> Self {
         Default::default()
    }


        /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    fn search(&self, chars:& [char]) -> bool {

        let mut cur = self;
        // let mut chs:Vec<char> = word.chars().collect();
        for id in 0..chars.len(){
            // let next = cur.child.get(&i);

            if chars[id]!='.'{

                let i = chars[id] as usize - 'a' as usize;

                //如果存在，就继续往下找
                if  let Some(next_d) = cur.child.get(&i){
                    cur = next_d;
                    continue;
                }
    
                // 找不到 ，而且 不是 通配符 . ,就可以直接返回了
                if 0 <= i && i<= 25{
                    return false;
                }
            }else{

                for(_,v) in cur.child.iter(){
                    if  v.search(& chars[id+1..]){
                        return true;
                    }
                 }
            }

            return false;


        }

        return cur.is_end;

    }

}

struct WordDictionary {
    root:Box<Trie>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl WordDictionary {

    /** Initialize your data structure here. */
    fn new() -> Self {
         WordDictionary{root: Box::new(Trie::new())}
    }
    
    /** Adds a word into the data structure. */
    fn add_word(&mut self, word: String) {

        let mut cur = self.root.as_mut();
        for i in word.chars().map(|ch| (ch as u8 - 'a' as u8) as usize){
            cur = cur.child.entry(i).or_insert_with(||Box::new(Trie::new()));
        }

        cur.is_end = true;

    }
    
    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    fn search(&self, word: String) -> bool {

        let chars :Vec<char>= word.chars().collect();
        self.root.search(&chars[..])

    }
}



#[cfg(test)]


    #[test]
    fn test_211() {
        let mut dict = WordDictionary::new();
        dict.add_word("bad".to_owned());
        dict.add_word("dad".to_owned());
        dict.add_word("mad".to_owned());
        assert_eq!(dict.search("pad".to_owned()), false);
        assert_eq!(dict.search("bad".to_owned()), true);
        assert_eq!(dict.search(".ad".to_owned()), true);
        assert_eq!(dict.search("da.".to_owned()), true);
    }
