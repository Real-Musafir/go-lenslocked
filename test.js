const input_array = [1, -1, 2];
let left_sum = input_array[0];
let right_sum = input_array[input_array.length - 1];

let leftPosition = 0;
let rightPosition = input_array.length - 1;
let findedMiden = null;
while (true) {
  console.log(leftPosition, rightPosition, "Check the position");
  if (leftPosition >= rightPosition) break;
  if (
    right_sum === 0 &&
    leftPosition === 0 &&
    rightPosition - leftPosition == 1
  ) {
    findedMiden = 0;
    break;
  }
  if (
    left_sum === 0 &&
    rightPosition === input_array.length - 1 &&
    rightPosition - leftPosition == 1
  ) {
    findedMiden = rightPosition;
    break;
  }

  if (left_sum > right_sum) {
    rightPosition = rightPosition - 1;
    right_sum = right_sum + input_array[rightPosition];
  } else if (left_sum < right_sum) {
    leftPosition = leftPosition + 1;
    left_sum = left_sum + input_array[leftPosition];
  } else if (left_sum === right_sum && rightPosition - leftPosition === 2) {
    findedMiden = leftPosition + 1;
    break;
  } else {
    leftPosition = leftPosition + 1;
    rightPosition = rightPosition - 1;
    right_sum = right_sum + input_array[rightPosition];
    left_sum = left_sum + input_array[leftPosition];
  }
}
console.log(findedMiden, "OKKK");
