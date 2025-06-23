{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.nodejs_20
    pkgs.tailwindcss
    pkgs.air
    pkgs.makeWrapper
  ];

  shellHook = ''
    echo "✅ nix-shell is ready."
    echo "Run: make dev"
  '';
}
