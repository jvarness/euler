main() {
  var amicableNumbers = [];
  for (var number = 1; number < 10000; number++) {
    var numberDivisors1 = _divisors(number);
    var sumDivisors1 = numberDivisors1.reduce((x,y)=>x+y);
    var numberDivisors2 = _divisors(sumDivisors1);
    var sumDivisors2 = numberDivisors2.reduce((x,y)=>x+y);

    if (number == sumDivisors2 && number != sumDivisors1) {
      print('$number : $sumDivisors1');
      if (!amicableNumbers.contains(number)) {
        amicableNumbers.add(number);
      }
      if (!amicableNumbers.contains(sumDivisors1)) {
        amicableNumbers.add(sumDivisors1);
      }
    }
  }

  var sum = amicableNumbers.reduce((x,y)=>x+y);

  print('Sum of amicable numbers: $sum');
}

List<int> _divisors(int n) {
  var divisors = [0];
  for (var x = 1; x < n; x++) {
    if (n % x == 0) {
      divisors.add(x);
    }
  }
  return divisors;
}