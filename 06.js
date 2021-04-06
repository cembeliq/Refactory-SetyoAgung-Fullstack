const list_letters_first = ["c", "d", "e", "g", "h"]
const list_letters_second = ["X", "Z"]

function findMissLetter(arrs) {
    var i, j = 0,
        m = 122;
    let str = "";
    for (arr of arrs) {
        str += arr;
    }
    if (str) {
        i = str.charCodeAt(0);
        while (i <= m && j < str.length) {
            if (String.fromCharCode(i) !== str.charAt(j)) {
                return String.fromCharCode(i);
            }
            i++;
            j++;
        }
    }
    return undefined;
}

console.log("list_letters_first = " + findMissLetter(list_letters_first));
console.log("list_letters_second = " + findMissLetter(list_letters_second));