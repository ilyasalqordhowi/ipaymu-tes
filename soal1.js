// 1. Validate Parentheses (20 points)
// Write a function that takes a string containing just the characters '(', ')', '{', '}', ‘<’,’>’, '[' and ']' and determines if the input string is valid. The string is valid if:
// Open brackets are closed by the same type of brackets.
// Open brackets are closed in the correct order.
// Input: "([{}])"
// Output: True
// Input: "(]"
// Output: False

function validate(s) {
  const results = [];

  for (let i = 0; i < s.length; i++) {
    const characters = s[i];

    if (
      characters === "(" ||
      characters === "{" ||
      characters === "[" ||
      characters === "<"
    ) {
      results.push(characters);
    } else if (characters === ")") {
      if (results.length === 0 || results.pop() !== "(") return false;
    } else if (characters === "}") {
      if (results.length === 0 || results.pop() !== "{") return false;
    } else if (characters === "]") {
      if (results.length === 0 || results.pop() !== "[") return false;
    } else if (characters === ">") {
      if (results.length === 0 || results.pop() !== "<") return false;
    }
  }

  return results.length === 0;
}

console.log(validate("([{}])"));
console.log(validate("(]"));
