# Serial Pool

Serial provides an enumerator that returns a set of functions to ```request()``` and ```release()``` a SID (serialized numeric ID) from an available limited numeric id pool range.

The SID can be released in any order, but requests are always forward looking and will be the next sequential numeric ID available for use or re-use. 

```golang 

  request,release := serial.Enumerator(6000)
  
  for !exit {
    go func(id int) {
      defer release(id)
      // ... do stuff
    }(request())
  }

```

## Note

Size the pool properly to avoid exhaustion and blocking while waiting for a SID to become available. 

For example, if you consume 1000 sid/sec and hold a SID for 3 seconds on average, an appropraite size might be found by taking your consumtion times average duration and adding in an additional padding factor to prevent blocking. (eg. 1000x3 x2 = 6000)

