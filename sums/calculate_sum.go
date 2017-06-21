package calculate_sum

import (
      "crypto/md5"
      )

func GetMD5Hash(s string) [16]byte {
  data := []byte(s)
  return [16]byte(md5.Sum(data))
}
