#!/bin/bash
output="/tmp/header_type.go"
rm -rf $output
header_line=$(find test_data/statfiles -name *.stat | while read f; do head -1 $f; done | tr -s ' ' | awk -F' LINE_TYPE ' '{print $1}' | sort | uniq | grep LINE_TYPE)
echo "type statHeader struct {" > $output

for term in $header_line
do
    lower_name="$(echo $term | tr '[:upper:]' '[:lower:]')"
    name="$(tr '[:lower:]' '[:upper:]' <<< ${lower_name:0:1})${lower_name:1}"
    echo "      ${name} string \`json:\"${term}\"\`" >> $output
done
echo "      ID string \`json:\"ID\"\`" >> $output
echo "}" >> $output

