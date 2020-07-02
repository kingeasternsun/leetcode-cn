fn main() {
    println!("Hello, world!");
}

use std::collections::HashMap;

#[derive(Default)]
struct Trie {
    is_end: bool,
    // nodes: HashMap<usize,Box<Trie>>,
    nodes: HashMap<usize,Trie>, //36ms 11.6MB
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Trie {
    /** Initialize your data structure here. */
    fn new() -> Self {
        Default::default()
    }

    /** insert a word into the trie. */
    fn insert(&mut self, word: String) {
        let mut curr = self;

        for i in word.chars().map(|ch| (ch as u8 - 'a' as u8) as usize) {

            let next_node = curr.nodes.entry(i).or_insert_with(||Trie::new());

            curr = next_node
        }
        curr.is_end = true;
    }

    /** Returns if the word is in the trie. */
    fn search(&self, word: String) -> bool {
        self.find(word).map_or(false, |t| t.is_end)
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    fn starts_with(&self, prefix: String) -> bool {
        self.find(prefix).is_some()
    }

    fn find(&self, word: String) -> Option<&Trie> {
        let mut curr = self;
        for i in word.chars().map(|ch| (ch as u8 - 'a' as u8) as usize) {
           curr =  curr.nodes.get(&i)?
         
        }
        Some(curr)
    }
}

// submission codes end

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_208() {
        let mut trie = Trie::new();
        trie.insert("apple".to_owned());
        assert_eq!(trie.search("apple".to_owned()), true); // returns true
        assert_eq!(trie.search("app".to_owned()), false);
        assert_eq!(trie.starts_with("app".to_owned()), true); // returns true
        trie.insert("app".to_owned());
        assert_eq!(trie.search("app".to_owned()), true); // returns true
    }
}