echo -n >  analysis/5bfg.log
echo -n >  analysis/583bfg.log
echo -n >  analysis/1024bfg.log
rm test/*.meta
rm test/*.data


go build qfilePerftest.go
echo "testing 5 bytes workload"
./qfilePerftest workload/5.byte > analysis/5bfg.log
php analyser.php analysis/5bfg.log

rm test/*.meta
rm test/*.data
echo "testing 583 bytes workload"
./qfilePerftest workload/583.byte > analysis/583bfg.log
php analyser.php analysis/583bfg.log

rm test/*.meta
rm test/*.data
echo "testing 1024 bytes workload"
./qfilePerftest workload/1024.byte > analysis/1024bfg.log
php analyser.php analysis/1024bfg.log
