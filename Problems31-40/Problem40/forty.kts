fun champernowne() {
    val MAX_DIGIT = 1_000_000 
    var digit = 1
    var nextMultiplier = 1
    var product = 1
    var str = ""
    do {
        str = str + digit
        if (nextMultiplier == digit) {
            product = Integer.parseInt(str[digit - 1].toString()) * product
            nextMultiplier *= 10
        }
        digit++
    } while (str.length <= MAX_DIGIT)

    println("Product: $product")
}

champernowne()