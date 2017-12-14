main() {
  var factorial = _factorial(100);
  var sum = factorial.toString().split('').map((x) => int.parse(x)).reduce((x,y)=>x+y);
  print('Factorial sum: $sum');
}

int _factorial(int n) => n == 1 ? n : n * _factorial(n-1);