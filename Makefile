gen-abis:
	mkdir -p ./abis/srcstatemanager
	abigen --abi=./abis/json/src_state_manager.json \
		--pkg=srcstatemanager \
		--type=srcstatemanager \
		--out=./abis/srcstatemanager/binding.go
	mkdir -p ./abis/erc20
	abigen --abi=./abis/json/erc20.json \
		--pkg=erc20 \
		--type=erc20 \
		--out=./abis/erc20/binding.go
	mkdir -p ./abis/destentrypoint
	abigen --abi=./abis/json/dest_entrypoint.json \
		--pkg=destentrypoint \
		--type=destentrypoint \
		--out=./abis/destentrypoint/binding.go
	mkdir -p ./abis/deststatemanager
	abigen --abi=./abis/json/dest_state_manager.json \
		--pkg=deststatemanager \
		--type=deststatemanager \
		--out=./abis/deststatemanager/binding.go