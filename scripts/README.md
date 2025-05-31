# Example usage of this tool

## Building the parser

```text
> cd cloned_directory
>./scripts/build.sh
formatting primary packages
linting primary packages
building package generator
building for linux/arm64
building for darwin/arm64
building for linux/amd64
building for darwin/amd64
saving current metLineTypeDefinitions.go to /tmp/metLineTypeDefinitions.go.bak
building new metLineTypeDefinitions.go
format and lint metLineTypeDefinitions
formatting ./pkg/metLineTypeDefinitions
linting ./pkg/metLineTypeDefinitions
building sample program
building for linux/arm64
building for darwin/arm64
building for linux/amd64
building for darwin/amd64
finished
>
```

After the above step you should have a complete, generated, formatted, and linted GO package in pkg/metLineTypeDefinitions and supporting
pks pkg/util, pkg/structColumnDefs and pkg/ample_parser.

## running the parser

```text
bin/darwin/arm64/sample_parser -outdir /tmp -path <MET-parser-testdata>
environment:darwin_arm64
Recovered: runtime error: slice bounds out of range [24:10] for fileName grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_420000L_20241101_180000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_540000L_20241101_180000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_600000L_20241101_180000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_660000L_20241101_180000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_900000L_20241101_180000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241102-00z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_180000L_20241102_000000V.stat - skipping
2025/02/24 14:35:05 empty file <MET-parser-testdata>G2G_v12/20241102-00z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_480000L_20241102_000000V.stat - skipping
2025/02/24 14:35:07 empty file <MET-parser-testdata>G2G_v12/20241103-12z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1020000L_20241103_120000V.stat - skipping
2025/02/24 14:35:09 missing VERSION at start of header line - bad header line? for file <MET-parser-testdata>__MACOSX/._G2G_v12 - skipping rest of file
2025/02/24 14:35:11 empty file <MET-parser-testdata>tc_data/CMC/2023062606/tc_pairs_al92.dat.tcst - skipping
2025/02/24 14:35:11 empty file <MET-parser-testdata>tc_data/CMC/2023062606/tc_pairs_al92.dat_PROBRIRW.tcst - skipping
2025/02/24 14:35:14 empty file <MET-parser-testdata>tc_data/GFSO/2024081218/tc_pairs_al05.dat_PROBRIRW.tcst - skipping
2025/02/24 14:35:17 empty file <MET-parser-testdata>tc_data/HMON/2023062606/tc_pairs_al03.dat.tcst - skipping
2025/02/24 14:35:17 empty file <MET-parser-testdata>tc_data/HMON/2023062606/tc_pairs_al03.dat_PROBRIRW.tcst - skipping
2025/02/24 14:35:17 empty file <MET-parser-testdata>tc_data/HMON/2023062606/tc_pairs_al92.dat.tcst - skipping
2025/02/24 14:35:18 empty file <MET-parser-testdata>tc_data/HMON/2024081306/tc_pairs_al05.dat_PROBRIRW.tcst - skipping
2025/02/24 14:35:18 empty file <MET-parser-testdata>tc_data/HWRF/2023062606/tc_pairs_al92.dat.tcst - skipping
2025/02/24 14:35:20 missing VERSION at start of header line - bad header line? for file <MET-parser-testdata>tcstfiles/PROBRIRW_filter_ee.tcst - skipping rest of file
2025/02/24 14:35:20 missing VERSION at start of header line - bad header line? for file <MET-parser-testdata>tcstfiles/PROBRIRW_summary_tk_err.tcst - skipping rest of file
```

These messages are expected because of poorly formatted data files, empty files, or data files with missing header lines.
After this runs you will have a compressed json file in the output path (/tmp) with the name sample_output.json.gz

## importing the compressed json file

```text
scripts/import_compressed_file.sh -h localhost:8091 -u Administrator -p 'pwd_av!d' -b metplusdata -s _default -c MET -f /tmp/sample_output.json.gz
/tmp/sample_output.json.gz: gzip compressed data, original size modulo 2^32 486133947
File is compressed - uncompressing it
/tmp/sample_output.json already exists -- do you wish to overwrite (y or n)? y
JSON `//tmp/sample_output.json` imported to `localhost:8091` successfully
Documents imported: 320952 Documents failed: 0
```

In this case there was an existing uncompressed file from a previous run. Whether or not the uncompressed file exists, the compressed file will be imported into the database.
