# fawchain

## Overview

**Fawchain** is a rudimentary blockchain implementation written in Go. At its core, it's designed to provide a clear and concise demonstration of how blockchains operate. Each block in this blockchain has a data field representing the base price and a calculated price based on a unique algorithm.


## What It Does

- **Fawchain** initializes with a genesis block.
- New blocks can be added, with each block containing data (base price).
- Every new block is linked to the previous block through its hash, ensuring data integrity.
- The blockchain can be viewed as a whole, displaying all the blocks and their content.

## Price Algorithm

The price for each block is determined based on the following factors:

1. **Seed Complexity**: A hash is generated from a random seed.
2. **Time Factor**: The current time is hashed, ensuring that the price varies over time.
3. **Price Multiplier**: Combining the two hashes, a final hash is generated. The first 8 bytes of this hash are then converted to a float number between 0 and 1. This multiplier is adjusted to a range between 0.001 (lucky) and 11 (unlucky) (see below).
4. The base price is multiplied by this multiplier to determine the final price for the block.

![fawchain example](img/blockchain.png)
