const secret_text = "23dn3ir30fd2eddd";

const masking = (str) => {
    const strLength = str.length;
    const last3Letter = str.substr(strLength - 3); // => "Tabs1"

    let strMasking = "";
    for (let i = 0; i < strLength - 3; i++) {
        strMasking += "*";
    }
    strMasking = strMasking.concat(last3Letter);
    return strMasking;
}

console.log(masking(secret_text));