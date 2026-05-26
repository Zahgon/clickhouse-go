package clickhouse_api

func QBitSubcolumns() error { _ = "STUB: not implemented"; return nil }

// QBit is an experimental feature in ClickHouse

// Create table with QBit column

// Insert vectors with special float values to demonstrate bit patterns
// Float32 has 32 bits: 1 sign bit + 8 exponent bits + 23 mantissa bits
// Bit numbering in QBit: .1 is MSB (sign bit), .32 is LSB

// In Go, compiler doesn't let you use the constant -0.0

// Insert vectors with positive and negative zeros to show sign bit difference

// Access subcolumns to examine bit patterns
// .1 is the sign bit (MSB of Float32)
// .2-.9 are exponent bits
// .10-.32 are mantissa bits
