import "dart:io";

main() {
  File namesFile = new File('names.txt');
  List<String> namesList = namesFile.readAsStringSync().split(',');
  namesList.sort((a,b) => a.compareTo(b));
  
  int total = 0;

  for (var x = 0; x < namesList.length; x++) {
    String name = namesList[x];
    total += (name.trim().codeUnits.map((val) => val - 64).reduce((val, element) => element + val)) * (x+1);
  }

  print('Total names score: $total');
}