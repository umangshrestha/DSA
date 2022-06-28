mod vigenere_cipher;
use vigenere_cipher::*;

fn main() {
    let input = "The quick brown fox jumps over 13 lazy dogs.";
    let expected = "Llg hybmo zjsye jhh nsetu fzxb 13 pyrc ffkl.";
    let key = "secretkey";

    let output = encrypt(input, key);
    println!("input:{}\nkey:{}\noutput:{}\nexpected:{}", input, output, decrypt(&output, key), expected);
    
}