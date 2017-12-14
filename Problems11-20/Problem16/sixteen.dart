/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/
import 'dart:math' as math;

main(List<String> args) {
  int sum = 0;

  math.pow(2, 1000).toString().split('').forEach((c) {
    sum += int.parse(c);
  });

  print("Sum of all digits in 2^1000: $sum");
}