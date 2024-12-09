#!/bin/bash

output="/tmp/data_line_types.go"
rm -rf $output

find test_data/textfiles -name *.txt | while read f
do
    header=$(head -1 $f)
    read header_fields <<< $header
    for i in "${!header_fields[@]}"
    do
        if [[ "${header_fields[$i]}" = "LINE_TYPE" ]]
        then
            break
        fi
    done
    line_type_index=$i
    data_line=$(head -2 $f | tail -1)
    read data_line_fields <<< $data_line
    line_type="${data_line_fields[$line_type_index]}"
    echo "$header $line_type" | tr -s ' ' | awk -F' ALPHA LINE_TYPE ' '{print $2}'
done  | sort | uniq | while read data_header
do
    read data_fields <<< $data_header
    line_type="${data_fields[${#data_fields[@]} - 1]}"
    for i in "${!data_fields[@]}"
    do
        if [[ "${#data_fields[@]}" == "$i" ]]
        then
            unset data_fields[$i]
        fi
    done
    echo "" >> $output
    echo "type ${line_type} struct {" >> $output
    for term in "${data_fields[@]}"
    do
        lower_name="$(echo $term | tr '[:upper:]' '[:lower:]')"
        name="$(tr '[:lower:]' '[:upper:]' <<< ${lower_name:0:1})${lower_name:1}"
        echo "      ${name} string \`json:\"${term}\"\`" >> $output
    done
    echo "}" >> $output
    echo "" >> $output
    i=$((i+1))
done
