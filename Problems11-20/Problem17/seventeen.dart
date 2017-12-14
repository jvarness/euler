main(List<String> args) {
  var totalWordCount = 0;
  for (var i = 1; i <= 1000; i++) {
    totalWordCount += getWordCount(i);
  }
  print('Word count: $totalWordCount');
}

int getWordCount(int i) {
  const _numbers = const ['', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine', 'ten', 'eleven','twelve','thirteen','fourteen','fifteen','sixteen','seventeen','eighteen','nineteen'];
  const _tens = const ['','','twenty','thirty','forty','fifty','sixty','seventy','eighty','ninety'];
  const _hundred = 'hundred';
  const _thousand = 'thousand';
  const _and = 'and';

  var thousandPlace = i ~/ 1000;
  var hundredPlace = (i % 1000) ~/ 100;
  var tensPlace = (i % 100) ~/ 10;
  var onesPlace = i % 10;

  var wordCount = 0;

  if (thousandPlace >= 1) {
    wordCount += (_thousand.length + _numbers[thousandPlace].length);
  }
  if (hundredPlace >= 1) {
    wordCount += (_hundred.length + _numbers[hundredPlace].length);
    if (onesPlace != 0 || tensPlace != 0) {
      wordCount += _and.length;
    }
  }
  if (tensPlace >= 1){
    if ((tensPlace * 10) + onesPlace < _numbers.length) {
      wordCount += _numbers[(tensPlace * 10) + onesPlace].length;
    } else {
      wordCount += _tens[tensPlace].length + _numbers[onesPlace].length;
    }
  } else if (onesPlace >= 1) {
    wordCount += _numbers[onesPlace].length;
  }

  return wordCount;
}