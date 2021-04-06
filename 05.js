const text_1 = "Mammals";
const text_2 = "Bruiser build";

const chars = (_text) => {
    _text = _text.replace(/\s/g, '');
    const objCount = {};

    _split = _text.toLowerCase().split('');
    const uniqueChars = new Set(_split);
    for (char of uniqueChars) {
        let n = _text.toLowerCase().split(char).length - 1;
        let asterix = ''
        for (let i = 0; i < n; i++) {
            asterix += '*';
        }
        objCount[char] = asterix;
    }

    return objCount;
}

console.log(chars(text_1));
console.log(chars(text_2));