using SHA

function rabinKarp(str::String, pattern::String)::Bool
  n = length(str)
  m = length(pattern)
  h = sha256(pattern)
  for i in 1:n-m+1
    println(str[i:i+m-1])
    if sha256(str[i:i+m-1]) == h && str[i:i+m-1] == pattern
      return true
    end #if
  end #for

  return false
end #rabinKarp

function main()
  pattern = ARGS[1]
  text = ARGS[2]

  result = rabinKarp(text, pattern)
  if result
    println("Wzorzec występuje")
  else
    println("Wzorzec nie występuje")
  end #if
end #main

main()