// 2. Merge Intervals (15 points)
// Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals.
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]

function mergeIntervals(intervalsInput) {
  intervalsInput.sort((a, b) => a[0] - b[0]);

  const merged = [];
  for (const interval of intervalsInput) {
    if (merged.length === 0 || merged[merged.length - 1][1] < interval[0]) {
      merged.push(interval);
    } else {
      merged[merged.length - 1][1] = interval[1];
    }
  }

  for (let i = 1; i < merged.length; i++) {
    if (merged[i][0] <= merged[i - 1][1]) {
      merged[i - 1][1] = merged[i][1];
      merged.splice(i, 1);
      i--;
    }
  }

  return merged;
}

const intervalsInput = [
  [1, 3],
  [2, 6],
  [8, 10],
  [15, 18],
];

const result = mergeIntervals(intervalsInput);
console.log(result);
