fn key_to_vec(key: &str) -> Vec<u8> {
    key.chars()
        .filter_map(|ch| {
            if ch.is_ascii_alphabetic() {
                Some(ch as u8 - if ch.is_ascii_uppercase() { b'A' } else { b'a' })
            } else {
                None
            }
        })
        .collect()
}

pub fn encrypt(plaintext: &str, key: &str) -> String {
    if key.len() == 0 {
        return plaintext.to_string();
    }
    let key = key_to_vec(key);
    let mut i = 0;
    plaintext
        .chars()
        .map(|mut ch| {
            if ch.is_ascii_alphabetic() {
                let f = if ch.is_ascii_uppercase() { b'A' } else { b'a' };
                ch = (f + (ch as u8 - f + key[i % key.len()]) % 26) as char;
                i += 1;
            }
            return ch;
        })
        .collect()
}

pub fn decrypt(plaintext: &str, key: &str) -> String {
    if key.len() == 0 {
        return plaintext.to_string();
    }
    let key = key_to_vec(key);
    let mut i = 0;
    plaintext
        .chars()
        .map(|mut ch| {
            if ch.is_ascii_alphabetic() {
                let f = if ch.is_ascii_uppercase() { b'A' } else { b'a' };
                ch = (f + (ch as u8 - f + 26 - key[i % key.len()]) % 26) as char;
                i += 1;
            }
            return ch;
        })
        .collect()
}

#[cfg(test)]
mod test {
    use super::{decrypt, encrypt};

    #[test]
    pub fn test_empty() {
        let test_cases = vec![
            ("", "", ""),
            ("plaintext with empty key", "", "plaintext with empty key"),
            ("", "key with no plaintext", ""),
        ];

        test_cases.iter().for_each(|&(input, key, expected)| {
            assert_eq!(encrypt(input, key), expected);
            assert_eq!(decrypt(expected, key), input);
        });
    }

    #[test]
    pub fn test_simple_text() {
        let input = "The quick brown fox jumps over 13 lazy dogs.";
        let expected = "Llg hybmo zjsye jhh nsetu fzxb 13 pyrc ffkl.";
        let key = "secretkey";
        assert_eq!(encrypt(input, key), expected);
        assert_eq!(decrypt(expected, key), input);
    }

    #[test]
    pub fn test_unicode() {
        let test_cases = vec![
            ("love: ❤️ vs brain: 🧠", "oops", "zckw: ❤️ jg qjowc: 🧠"),
            ("love: ❤️ vs brain: 🧠", "❤️oops🧠", "zckw: ❤️ jg qjowc: 🧠"),
        ];
        test_cases.iter().for_each(|&(input, key, expected)| {
            assert_eq!(encrypt(input, key), expected);
            assert_eq!(decrypt(expected, key), input);
        });
    }
}
