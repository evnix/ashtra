rm analysis/*bfg.log
rm test/*.meta
rm test/*.data

./scripts/bfq/100000.sh

./scripts/bfq/500000.sh

./scripts/bfq/1000000.sh

