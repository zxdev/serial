# Serial Enumerator and Number

Serial Enumerator and Number types

## Serial Enumerator

Serial provides an Enumerator type generates a reuseable pool of numbers and returns a set of functions to ```request()``` and ```release()``` a SID (serialized numeric ID) from an available limited numeric id pool range which is given by the user as 0...{n} range.

A SID can be released in any order, but requests are always forward looking and will be the next sequential numeric ID available for use or re-use. 

Size the Enumerator properly to avoid exhaustion and blocking while waiting for a SID to become available because with exhaustion of all SID the Enumerator will block waiting for a SID to be released for reuse.

For example, if you consume 1000 sid/sec and hold a SID for 3 seconds on average, an appropraite size might be found by taking your consumtion times average duration and adding in an additional padding factor to prevent blocking. (eg. 1000x3 x2 = 6000)

```golang 

  request,release := serial.Enumerator(6000)
  
  for !exit {
    go func(id int) {
      defer release(id)
      // ... do stuff
      
    }(request())
  }

```

## Number 

The Serial Number type provides a sequential enumerator that generates serial numbers sequentially utilizing the first 7 bytes of a uint64 to generate a number in the range N00000000000001 to NFFFFFFFFFFFFFF. The type provides method to persist the current Number to disk.

```golang

  var persist = "./number.persist"
  func Example() {
    var n Number
    var exit bool
    n.Load(persist)
    defer n.Save(persist)
    for !exit {
      sn := n.Next()
      // ... do stuff

      fmt.Println(sn)
      // N00000000000001
    }
  }