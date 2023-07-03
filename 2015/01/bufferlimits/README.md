# Why this doesn't work?

Because the input is larger than 4096 bytes and the Linux kernel has a pipe buffer limit of 4096 bytes, the input is truncated.

The best approach to work around this limitation is to read the input from a file instead of stdin.

Read the complete history [here, in English](https://4s3ti.net/en/posts/advent0115/) or [here, in Portuguese](https://4s3ti.net/posts/advent0115/).
