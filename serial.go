// MIT License

// Copyright (c) 2020 zxdev

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serial

import "sync"

// Enumerator returns a set of functions to request() and release() a SID
// (serialized numeric ID) from an available limited numeric pool range.
// The SID can be released in any order, but requests are always forward
// looking and will be the next sequential one available.
//
// Size the pool properly to avoid exhaustion and blocking while waiting for
// a SID to become available. For example, if you consume 1000 sid/sec and
// hold a SID for 3 seconds on average, an appropraite size might be found
// like this: (1000*3) *2 = 6000
func Enumerator(size int) (request func() (sid int), release func(sid int)) {

	var p int                            // positional index
	var mutex sync.Mutex                 // mutex control
	var bits = make([]uint8, size/8%8+1) // bit array

	return func() (sid int) { // request() function
			mutex.Lock()
			for bits[p/8]&(1<<(p%8)) != 0 { // check sid bit
				p++          // advance
				p = p % size // constrain
			}
			bits[p/8] |= 1 << (p % 8) // set sid bit
			sid = p                   // store, then advance
			p++                       // advance
			p = p % size              // constrain
			mutex.Unlock()
			return
		},
		func(sid int) { // release(sid) function
			mutex.Lock()
			bits[sid/8] &^= 1 << (sid % 8) // clear sid bit
			mutex.Unlock()
		}
}
