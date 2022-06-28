// This is an simple implementation of ceaser cipher.
// It rotates the string left -> right for encription and right -> left for decription.
// The key can be any value from 0 to 26.
// note: Only Alphabets i.e A...Z and a..z are rotated, the special symbols remain unchanged.

pub fn encrypt(plaintext: &str, key: u8) -> String {
    let key = key % 26;
    plaintext
        .chars()
        .map(|ch| {
            if ch.is_ascii_alphabetic() {
                let f = if ch.is_ascii_uppercase() { b'A' } else { b'a' };
                (f + (ch as u8 - f + key) % 26) as char
            } else {
                ch
            }
        })
        .collect()
}

pub fn decrypt(plaintext: &str, key: u8) -> String {
    let key = key % 26;
    return encrypt(plaintext, 26 - key);
}

#[cfg(test)]
mod test {
    use super::{decrypt, encrypt};

    #[test]
    pub fn test_empty() {
        let input = "";
        let expected = "";
        let shift = 10;
        assert_eq!(encrypt(input,shift), expected);
        assert_eq!(encrypt(expected,26 - shift), input);
        assert_eq!(decrypt(expected,shift), input);
        assert_eq!(decrypt(input,26 - shift), expected);
    }

    #[test]
    pub fn test_hello_world() {
        let input = "\"Hello, World!\" using ceaser cipher.";
        let expected = "\"Rovvy, Gybvn!\" ecsxq mokcob mszrob.";
        let shift = 10;
        assert_eq!(encrypt(input,shift), expected);
        assert_eq!(encrypt(expected,26 - shift), input);
        assert_eq!(decrypt(expected,shift), input);
        assert_eq!(decrypt(input,26 - shift), expected);
    }

    #[test]
    pub fn test_unicode() {
        let input = "love: ❤️";
        let expected = "fipy: ❤️";
        let shift = 20;
        assert_eq!(encrypt(input,shift), expected);
        assert_eq!(encrypt(expected,26 - shift), input);
        assert_eq!(decrypt(expected,shift), input);
        assert_eq!(decrypt(input,26 - shift), expected);
    }

    #[test]
    pub fn test_overflow() {
        let input = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque nisl eros,
        pulvinar facilisis justo mollis, auctor consequat urna. Morbi a bibendum metus.
        Donec scelerisque sollicitudin enim eu venenatis. Duis tincidunt laoreet ex,
        in pretium orci vestibulum eget. Class aptent taciti sociosqu ad litora torquent
        per conubia nostra, per inceptos himenaeos. Duis pharetra luctus lacus ut
        vestibulum. Maecenas ipsum lacus, lacinia quis posuere ut, pulvinar vitae dolor.
        Integer eu nibh at nisi ullamcorper sagittis id vel leo. Integer feugiat
        faucibus libero, at maximus nisl suscipit posuere. Morbi nec enim nunc.
        Phasellus bibendum turpis ut ipsum egestas, sed sollicitudin elit convallis.
        Cras pharetra mi tristique sapien vestibulum lobortis. Nam eget bibendum metus,
        non dictum mauris. Nulla at tellus sagittis, viverra est a, bibendum metus.";
        for shift in 0..255 {
            let output = encrypt(input,shift);
            assert_eq!(encrypt(input,shift), output);
            assert_eq!(decrypt(&output ,shift), input);
        }
    }
}
