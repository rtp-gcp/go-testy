two types of floats:

* float32
* float64

IEEE 754 standard

# max numbers

Max

* math.MaxFloat32 is 2.4e38
* matchMaxFloat64 is 1.8e308

Smallest positive

* float32
  - 1.4e-45
* float64
  - 4.9e-324


# regarding floating point numbers

```
const Avogadro = 6.02214129e23  // 6.02214129 x 10 ^ 23
```

This number is max float e23 < e38  

* one portion of the 32 bit value is the mantissa, so n-bits hold Avogadro
number portion 6
* another portion holds the Avogadro number .02214129 as 2214129
* another portion records this as a positive number
* another portion holds the exponenent as positive
* the last portion holds the exponent as 23 

Looking at the number, I see

* 1 024
* 64 000
* 9680 0000
* 1M

1024 + 65536 + 131072 + 1048576 = 1,246,208 that is close to 2,214,129



