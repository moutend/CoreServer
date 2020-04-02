ARGF.each do |line|
  name = line.gsub("\n", "")
puts "func acc#{name}(v *IAccessible) error {"
puts <<EOF
	return ole.NewError(ole.E_NOTIMPL)
}

EOF
end
