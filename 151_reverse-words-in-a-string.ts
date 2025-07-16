function reverseWords(s: string): string {
    const words: Array<string> = [];
    let word = "";

    for (let i = 0; i < s.length; i++) {

        if (s[i] === " ") {
            if (word !== "") {
                words.unshift(word);
                word = "";
            }
            continue;
        } else {
            word = word + s[i];
            if(i===s.length-1) {
                words.unshift(word);
            }
        }
    }

    return words.join(" ")
};
