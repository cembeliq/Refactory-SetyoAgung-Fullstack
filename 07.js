const numbers = [9, 4, 2, 4, 1, 5, 3, 0];
const sortOdd = sort_odd(numbers);
console.log(sortOdd);


function sort_odd(array) {
    const odd = array.filter((x) => x % 2).sort((a, b) => a - b);
    return array.map((x) => x % 2 ? odd.shift() : x);
}