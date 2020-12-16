#!/bin/sh

lines=$(grep -o -e 'new[A-Z].*' ../define_*)

out=$(cat << EOS
# Error Codes

| Category | Http Status | Code | Info Message |
----|----|----|----
EOS
)

# ../define_payment.go:newBadRequest("invalid_product_id", "商品情報が不正です")

IFS=$')\n'
for l in $lines
do
	l=$(echo $l|sed -e 's/\n/<br>/g')
	category=$(echo $l|sed -e 's/.*define_\(.*\)\.go:.*/\1/g')
	status=$(echo $l|sed -e 's/.*define_.*\.go:new\(.*\)(.*/\1/g')
	code=$(echo $l|sed -e 's/.*define_.*\.go:new.*("\(.*\)",.*/\1/g')
	message=$(echo $l|sed -e 's/.*define_.*\.go:new.*(".*", \(.*\)/\1/g'|sed -e 's/"//g')
	out=$(cat << EOS
${out}
| ${category} | ${status} | ${code} | ${message} |
EOS
)
done

echo "$out" > README.md

