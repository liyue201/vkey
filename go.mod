module github.com/liyue201/vkey

go 1.19

require (
	github.com/consensys/gnark v0.7.1
	github.com/consensys/gnark-crypto v0.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.26.0 // indirect
)

replace github.com/consensys/gnark v0.7.1 => github.com/stirlingx001/gnark v0.0.0-20240918053953-7e81a4765069

require github.com/stirlingx001/vksol v0.0.0-20241016084416-6057102ce9de

require (
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/sys v0.24.0 // indirect
)
