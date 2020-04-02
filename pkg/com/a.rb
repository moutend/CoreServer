found = false
is_put = false
is_get = false
ARGF.each do |line|
  line.gsub(/  */, " ").split(" ").each do |token|

    if found
      t = token.gsub(/\(.*/, "")

      t[0] = t[0].upcase
      t = "Put#{t}" if is_put
      t = "Get#{t}" if is_get

      puts "#{t} uintptr"

      found = false
      is_put = false
      is_get = false

      next
    end
    is_put = true if token == "propput,"
    is_get = true if token == "propget,"
    found = true if token == "HRESULT"
  end
end
