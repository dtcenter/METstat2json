for ds in G2G_v12 MODE_compref statfiles tc_data tcstfiles textfiles; do
    case "$ds" in
    "G2G_v12")
        bin/darwin/arm64/sample_parser -dataset G2Gv12 -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    "MODE_compref")
        bin/darwin/arm64/sample_parser -dataset ModeCmpRef -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    "statfiles")
        bin/darwin/arm64/sample_parser -dataset StatFiles -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    "tc_data")
        bin/darwin/arm64/sample_parser -dataset TcPairsPst -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    "tcstfiles")
        bin/darwin/arm64/sample_parser -dataset TcstFiles -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    "textfiles")
        bin/darwin/arm64/sample_parser -dataset TextFiles -outdir ~/metplusdata -path /tmp/testdata/testdata/${ds}
        ;;
    esac
done
