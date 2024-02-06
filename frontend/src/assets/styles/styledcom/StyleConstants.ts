export const SPACE_UNIT = 4
export const SPACE_TIMES = (multiple?: number) =>
  `${SPACE_UNIT * (multiple || 1)}px`
export const SPACE = SPACE_TIMES(1) // 4
export const SPACE_XS = SPACE_TIMES(2) // 8
export const SPACE_SM = SPACE_TIMES(3) // 12
export const SPACE_MD = SPACE_TIMES(4) // 16
export const SPACE_LG = SPACE_TIMES(5) // 20
export const SPACE_XL = SPACE_TIMES(9) // 36 = 16 + 20
export const SPACE_XXL = SPACE_TIMES(14) // 56 = 20 + 36
export const BLUE = '#1946b9'
