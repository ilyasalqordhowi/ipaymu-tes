// 3. Longest Consecutive Subsequence (15 points)
// You are given an unsorted array of integers. Write a function that finds the length of the longest sequence of consecutive integers.
// Input: [100, 4, 200, 1, 3, 2]

function longestConsecutive(nums) {
  const numSet = new Set(nums);
  let longestStreak = 0;

  for (const num of numSet) {
    if (!numSet.has(num - 1)) {
      let currentNum = num;
      let currentStreak = 1;

      while (numSet.has(currentNum + 1)) {
        currentNum += 1;
        currentStreak += 1;
      }

      longestStreak = Math.max(longestStreak, currentStreak);
    }
  }

  return longestStreak;
}

const input = [100, 4, 200, 1, 3, 2];
const result = longestConsecutive(input);
console.log(result);
