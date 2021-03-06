import { BigNumber } from 'bignumber.js'

export function parseBigNumber (number) {
  return BigNumber(number.toString()).toString(10)
}

export function ceil (bigNumber) {
  return bigNumber.integerValue(BigNumber.ROUND_CEIL)
}
