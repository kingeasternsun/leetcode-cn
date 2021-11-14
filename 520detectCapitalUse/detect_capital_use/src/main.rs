fn main() {
    println!("Hello, world!");
}

pub fn detect_capital_use(word: String) -> bool {

    let mut word = word.chars();
    let first = word.next();
    if first.is_none(){
        return true
    }
    let first = first.unwrap();

    if let Some(second) = word.next(){
        let res = word.try_for_each(move |x|{
            if second.is_lowercase() && x.is_lowercase(){
                return Ok(())
            }
    
            if second.is_uppercase() && x.is_uppercase(){
                return Ok(())
            }
    
            Err(())
        });

        if res.is_err(){
            return false
        }
        if first.is_uppercase(){
            return true
        }

        if first.is_lowercase() && second.is_lowercase(){
            return true
        }
    
        false
    }else{
        true
    }

}
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_name() {
        assert!(detect_capital_use("U".to_string()));
        assert!(detect_capital_use("US".to_string()));
        assert!(detect_capital_use("USA".to_string()));
        assert!(detect_capital_use("a".to_string()));
        assert!(detect_capital_use("ab".to_string()));
        assert!(detect_capital_use("abc".to_string()));
        assert!(detect_capital_use("Abc".to_string()));
        assert!(detect_capital_use("Ab".to_string()));
        assert!(!detect_capital_use("bAc".to_string()));
        assert!(!detect_capital_use("FlaG".to_string()));
    }
}
