"""
  Horner's method for polynomial evaluation.
  coefficients - vector of coefficients of polynomial sorted from the highest to the lowest degree.
  value - value for which polynomial should be evaluated.

  Time complexity (#FMA operations): Theta(n), where n is the polynomial degree.
"""
function horner(coefficients::Vector{T}, value::T)::T where {T<:Number}
  result = 0
  for val in coefficients
    result = result * value + val
  end
  return result
end

function main()
  # y = 2x^3 + 3x^2 + 4x + 5
  coefficients = [2, 3, 4, 5]
  value = 2
  println(horner(coefficients, value))
end

main()