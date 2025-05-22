nixpkgs/update:
	@nix flake lock --override-input nixpkgs github:NixOS/nixpkgs/$(rev)

.PHONY: build test run clean

test:
	@go test -short -cover ./...
