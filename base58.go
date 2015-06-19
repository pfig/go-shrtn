package shrtn

/*
base58 - encode/decode a numeric value using base 58

Base 58 is a modification to the "standard" base 62 encoding generally used by
URL-shortening services - some characters were removed for easier reading.

Flickr give the rationale for this modification here:

    https://www.flickr.com/groups/api/discuss/72157616713786392/

(and in that thread you can find code for a multitude of languages, including
Objective-C).

This code is loosely based on Tatsuhiko Miyagawa's Encode::Base58 Perl module
(and the test data a straight rip off of his).
*/

import (
  "errors"
  "strings"
)

const (
  // Chars list of valid characters
  Chars string = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
  // Base is the numeric base used
  Base uint = uint(len(Chars))
)

/*
Encode a given base 10 number into base 58.
Returns a string.
*/
func Encode(number uint) string {
  if number == 0 {
    return string(Chars[0])
  }

  value := number
  var result []byte
  for value > 0 {
    div := value / Base
    mod := value % Base
    result = append([]byte{Chars[mod]}, result...)
    value = div
  }

  return string(result)
}

/*
Decode a given base 58 value into a base 10 number.
Returns an int64.
*/
func Decode(short string) (decoded uint, err error) {
  n := len(short)
  reversed := []byte(short)
  for i := 0; i < n/2; i++ {
    reversed[i], reversed[n-1-i] = reversed[n-1-i], reversed[i]
  }

  var result uint
  multi := uint(1)

  for _, c := range reversed {
    idx := strings.Index(Chars, string(c))
    if idx == -1 {
      return 0, errors.New("Invalid character")
    }
    result += multi * uint(idx)
    multi *= Base
  }

  return result, nil
}
