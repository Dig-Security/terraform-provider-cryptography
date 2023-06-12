# sha512 (Data Source)

The sha512 computes the SHA512 hash of a given string and encodes it with hexadecimal digits.
The given string is first encoded with provided encoding
and then the SHA512 algorithm is applied. The raw hash is then encoded to lowercase hexadecimal
digits before returning.

## Example Usage


```terraform

data "cryptography_sha" "sha512" {
  input = "hello world"
  encoding = "ISO-8859-1"
}

```

## Schema

### Required

- `input` (String) String to calculate hash.
- `encoding` (String) Encoding to encode.

### Read-Only

- `sha` (String) SHA512 of `input`.
- `id` (String) String to calculat