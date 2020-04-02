ARGF.each do |line|
  name = line.gsub("\n", "")
puts "func (v *IAccessible) #{name}() error {"
puts "return acc#{name}(v)"
puts <<EOF
}

EOF
end
