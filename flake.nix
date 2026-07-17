{
  description = "CLI tool to manage tmux sessions with configuration file";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs, ... }:
    let
      systems = [
        "aarch64-darwin"
        "x86_64-darwin"
        "aarch64-linux"
        "x86_64-linux"
      ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
      # Prefer the commit hash over a hand-maintained version: it can never
      # go stale, and `nix run github:corrupt952/tmuxist` always builds main.
      version = self.shortRev or self.dirtyShortRev or "dev";
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "tmuxist";
            inherit version;
            src = pkgs.lib.cleanSource self;
            vendorHash = "sha256-M7nS+VeKXot/2ljQ30it2KYfwcIYmKPI13wwUWA1Omo=";
            # Keep in sync with .goreleaser.yml.
            ldflags = [ "-s" "-w" "-X" "tmuxist/command.Version=${version}" ];
            meta.mainProgram = "tmuxist";
          };
        });

      checks = forAllSystems (system: {
        default = self.packages.${system}.default;
      });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShellNoCC {
            packages = with pkgs; [
              go
              gopls
              gotools
              golangci-lint
              goreleaser
            ];
          };
        });
    };
}
