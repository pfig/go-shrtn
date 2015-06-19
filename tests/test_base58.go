package shrtn

import(
  "github.com/pfig/shrtn"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
)

func testBase58() {
  assert(shrtn.Base == 58)

  /* These characters shouldn't be included: 0, O, I, and l */
  invalid := "0OIl"
  assert(strings.ContainsAny(shrtn.Chars, invalid) == false)

  /* Test encoding and decoding */

  /* Encoding 0 should return the first of our valid characters */
  assert(shrtn.Encode(0) == "1")

  /* We should get an error when we find an invalid character */
  _, err := shrtn.Decode("0")
  assert(err != nil)
  assert(err.Error() == "Invalid character")

  /* Read in the test data */
  buf, _ := ioutil.ReadFile("test_encodings.yaml")
  var encodings []struct {
    Short string
    Number uint
  }
  err = yaml.Unmarshal([]byte(buf), &encodings)
  /* Some very minimal sanity checking of the test data */
  assert(err == nil)
  assert(len(encodings) == 500)

  /* Test data */
  for _, pair := range encodings {
    assert(shrtn.Encode(pair.Number) == pair.Short)
    res, err := shrtn.Decode(pair.Short)
    assert(res == pair.Number)
    assert(err == nil)
  }
}
