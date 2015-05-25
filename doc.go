/*
Package bitset represents a fixed size sequence of N bits.

Lets create a new bitset:

    package main

    import (
        "fmt"
        "github.com/umahmood/bitset"
    )

    func main() {
        bits := bitset.New(64)
        ...
    }

This will allocate 64 bits all initialized to zero:

    0000000000000000000000000000000000000000000000000000000000000000

Note: The desired size is rounded upto be a muliple of 64. i.e:

    bitset.New(50)  // Size is rounded upto 64.
    bitset.New(100) // Size is rounded upto 128.
    bitset.New(150) // Size is rounded upto 192.

To set a bit:

    bits.Set(42)

Note: The Most Significant Bit (MSB) is at index zero, so performing:

    bits.Set(0)

Will set the bit at index zero and will result in the bit pattern:

    1000000000000000000000000000000000000000000000000000000000000000
    ↑ msb

Reset a bit:

    bits.Reset(42)

Flip a bit:

    bits.Flip(28)

Test a bit:

    if bits.Test(10) {
        fmt.Println("bit a position 10 set to 1 (true).")
    } else {
        fmt.Println("bit a position 10 set to 0 (false).")
    }

Size of the bitset:

    n := bits.Size()

Check if all the bits in the bitset are set to 1.

    if bits.All() {...}

Check if any of the bits in the bitset are set to 1.

    if bits.Any() {...}

Check if none of the bits in the bitset are set to 1.

    if bits.None() {...}

Get the number of bits in the bitset set to 1.

    n := bits.TrueCount()

Make a copy of a bitset:

    a := bitset.New(200)
    a.Set(42)
    b := bitset.New(100)
    b.Copy(a)
    b.Reset(42)
    a.Test(42) // true
    b.Test(42) // false

Printing a bitset:

    fmt.Println(bits.String())
    // or
    fmt.Println(bits)

*/
package bitset
