#!/usr/bin/env sh
usage() {
    echo "Usage: $0 -h <host> -u <username> -p <password> -b <bucket> -s <scope> -c <collection> -j <json_file>"
    echo "Example: $0 -h localhost:8091 -u Administrator -p 'your_password' -b met -s _default -c MET -j /tmp/sample_output.json"
    exit 1
}
while getopts ":h:u:p:b:s:c:f:" opt; do
    case $opt in
        h) cb_host=$OPTARG;;
        u) cb_user=$OPTARG;;
        p) cb_pwd=$OPTARG;;
        b) bucket=$OPTARG;;
        s) scope=$OPTARG;;
        c) collection=$OPTARG;;
        f) json_f=$OPTARG;;
        \?) echo "Invalid option -$OPTARG" >&2
        ;;
    esac
done

[ -z "${cb_host}" ] && usage
[ -z "${cb_user}" ] && usage
[ -z "${cb_pwd}" ] && usage
[ -z "${bucket}" ] && usage
[ -z "${scope}" ] && usage
[ -z "${collection}" ] && usage
[ -z "${json_f}" ] && usage

file ${json_f} | grep "gzip compressed data"
ret=$?
if [ $ret -eq 0 ]; then
    echo "File is compressed - uncompressing it"
    gunzip ${json_f}
    export json=$(echo ${json_f} | sed 's/.gz//')
else
    export json="${json_f}"
fi

export number_of_cpus=$(nproc)
cbimport json --threads ${number_of_cpus} --cluster ${cb_host} --bucket ${bucket} --scope-collection-exp ${scope}.${collection} --username ${cb_user} --password ${cb_pwd} --format list --generate-key %id% --dataset file:///${json}
