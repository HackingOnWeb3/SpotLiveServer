solc-select install 0.8.13


```shell
brew install solc-select
solc-select install 0.8.13
solc-select use 0.8.13

solc --abi contracts/SpotLive/SpotLive.sol -o contracts/SpotLive/SpotLive.abi
mkdir -p SpotLive && abigen --abi=contracts/SpotLive/SpotLive.abi --pkg=SpotLive --out=SpotLive/SpotLive.go


```

