make install-darwin

rm -rf ~/.simapp/

simd init mynode --chain-id=simnet

simd keys add validator --keyring-backend="test"
simd add-genesis-account $(simd keys show validator -a --keyring-backend="test") 2000000stake
simd gentx validator 2000000stake --keyring-backend="test" --chain-id=simnet
simd collect-gentxs

sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.simapp/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.simapp/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.simapp/config/config.toml
sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.simapp/config/config.toml

simd start --pruning=nothing